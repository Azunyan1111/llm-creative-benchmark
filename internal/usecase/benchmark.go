package usecase

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

// RunConfig describes one benchmark invocation.
type RunConfig struct {
	Languages   []bench.Language
	Models      []bench.ModelSpec
	Modes       []bench.Mode
	Runs        int // trials per (language, model, mode)
	Concurrency int // parallel requests (default 1)
	MaxTokens   int // completion budget incl. reasoning
	Minimum     int // valid words used for scoring (default 7)
}

// RunStats summarises an invocation.
type RunStats struct {
	Planned    int
	Skipped    int // already present in the repository
	Done       int
	Errors     int
	Unscorable int
}

// Benchmark is the "run the DAT against many models" use case.
type Benchmark struct {
	Gen     Generator
	Catalog Catalog
	Vectors VectorRepository
	Trials  TrialRepository
	Log     Logger
	Now     func() time.Time
}

// Plan resolves the reasoning efforts for each model from the catalogue.
func (b *Benchmark) Plan(ctx context.Context, models []bench.ModelSpec) ([]bench.EffortPlan, error) {
	plans := make([]bench.EffortPlan, 0, len(models))
	for _, m := range models {
		info, found, err := b.Catalog.Lookup(ctx, m.ID)
		if err != nil {
			return nil, fmt.Errorf("catalogue lookup %s: %w", m.ID, err)
		}
		if !found {
			return nil, fmt.Errorf("model %q not found in catalogue", m.ID)
		}
		plans = append(plans, bench.ResolveEfforts(m, info))
	}
	return plans, nil
}

type job struct {
	lang   bench.Language
	plan   bench.EffortPlan
	mode   bench.Mode
	run    int
	prompt string
	scorer *dat.Model
}

// Run executes every (language, model, mode, run) job not already stored.
func (b *Benchmark) Run(ctx context.Context, cfg RunConfig) (RunStats, error) {
	var stats RunStats
	if cfg.Runs <= 0 {
		return stats, errors.New("runs must be positive")
	}
	if cfg.Concurrency <= 0 {
		cfg.Concurrency = 1
	}
	if cfg.Minimum <= 0 {
		cfg.Minimum = dat.DefaultMinimum
	}
	now := b.Now
	if now == nil {
		now = time.Now
	}

	// Validate everything before the first request.
	plans, err := b.Plan(ctx, cfg.Models)
	if err != nil {
		return stats, err
	}
	scorers := map[bench.Language]*dat.Model{}
	prompts := map[bench.Language]string{}
	for _, l := range cfg.Languages {
		p, ok := bench.PromptFor(l)
		if !ok {
			return stats, fmt.Errorf("unsupported language %q", l)
		}
		v, err := b.Vectors.Load(l)
		if err != nil {
			return stats, fmt.Errorf("load vectors for %s: %w", l, err)
		}
		prompts[l] = p
		scorers[l] = bench.NewScorer(l, v)
		b.Log.printf("[%s] %d vectors loaded", l, len(v))
	}
	existing, err := b.Trials.All()
	if err != nil {
		return stats, err
	}
	// Finished jobs are skipped on resume; errored ones are retried.
	done := map[string]bool{}
	for _, t := range existing {
		if t.Error == "" {
			done[t.Key()] = true
		}
	}

	// Jobs are ordered run-major so concurrent workers spread across models
	// and providers instead of hammering one model at a time.
	var jobs []job
	for r := 0; r < cfg.Runs; r++ {
		for _, l := range cfg.Languages {
			for _, p := range plans {
				for _, m := range cfg.Modes {
					stats.Planned++
					key := bench.Trial{Language: l, ModelID: p.ID, Mode: m, Run: r}.Key()
					if done[key] {
						stats.Skipped++
						continue
					}
					jobs = append(jobs, job{lang: l, plan: p, mode: m, run: r, prompt: prompts[l], scorer: scorers[l]})
				}
			}
		}
	}
	b.Log.printf("planned %d jobs, %d already done, %d to run", stats.Planned, stats.Skipped, len(jobs))

	var (
		mu    sync.Mutex
		wg    sync.WaitGroup
		queue = make(chan job)
	)
	worker := func() {
		defer wg.Done()
		for j := range queue {
			t := b.runOne(ctx, cfg, j, now)
			if t.Error != "" && ctx.Err() != nil {
				continue // interrupted, not a model failure: leave the job for the next run
			}
			mu.Lock()
			stats.Done++
			switch {
			case t.Error != "":
				stats.Errors++
			case t.Score == nil:
				stats.Unscorable++
			}
			if err := b.Trials.Append(t); err != nil {
				b.Log.printf("append: %v", err)
			}
			mu.Unlock()
			b.Log.printf("[%s] %-32s %-11s run=%d %s", t.Language, t.ModelID, t.Mode, t.Run, describe(t))
		}
	}
	for i := 0; i < cfg.Concurrency; i++ {
		wg.Add(1)
		go worker()
	}
	for _, j := range jobs {
		select {
		case <-ctx.Done():
			close(queue)
			wg.Wait()
			return stats, ctx.Err()
		case queue <- j:
		}
	}
	close(queue)
	wg.Wait()
	return stats, ctx.Err()
}

func (b *Benchmark) runOne(ctx context.Context, cfg RunConfig, j job, now func() time.Time) bench.Trial {
	t := bench.Trial{
		Language:  j.lang,
		ModelID:   j.plan.ID,
		Family:    j.plan.Family,
		Mode:      j.mode,
		Effort:    j.plan.EffortFor(j.mode),
		Run:       j.run,
		Timestamp: now(),
	}
	start := now()
	g, err := b.Gen.Generate(ctx, GenerateRequest{ModelID: j.plan.ID, Prompt: j.prompt, Effort: t.Effort, MaxTokens: cfg.MaxTokens})
	t.LatencyMs = now().Sub(start).Milliseconds()
	if err != nil {
		t.Error = err.Error()
		return t
	}
	t.RawContent = g.Content
	t.Provider = g.Provider
	t.FinishReason = g.FinishReason
	t.PromptTokens = g.PromptTokens
	t.CompletionTokens = g.CompletionTokens
	t.ReasoningTokens = g.ReasoningTokens
	t.CostUSD = g.CostUSD
	t.Words = bench.ParseWords(g.Content)
	res, err := j.scorer.Score(t.Words, cfg.Minimum)
	t.Valid = res.Valid
	t.Used = res.Used
	if err == nil {
		s := res.Score
		t.Score = &s
	}
	return t
}

func describe(t bench.Trial) string {
	switch {
	case t.Error != "":
		return "ERROR " + t.Error
	case t.Score == nil:
		return fmt.Sprintf("unscorable (%d valid: %v)", len(t.Valid), t.Valid)
	default:
		return fmt.Sprintf("score=%.2f reasoning_tokens=%d %v", *t.Score, t.ReasoningTokens, t.Used)
	}
}
