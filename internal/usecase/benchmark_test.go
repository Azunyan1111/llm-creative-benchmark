package usecase

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

// --- fakes -------------------------------------------------------------

type fakeGen struct {
	mu    sync.Mutex
	calls []GenerateRequest
	reply func(GenerateRequest) (Generation, error)
}

func (g *fakeGen) Generate(_ context.Context, r GenerateRequest) (Generation, error) {
	g.mu.Lock()
	g.calls = append(g.calls, r)
	g.mu.Unlock()
	return g.reply(r)
}

type fakeCatalog map[string]bench.ModelInfo

func (c fakeCatalog) Lookup(_ context.Context, id string) (bench.ModelInfo, bool, error) {
	mi, ok := c[id]
	return mi, ok, nil
}

type fakeVectors map[bench.Language]dat.Vectors

func (v fakeVectors) Load(l bench.Language) (dat.Vectors, error) {
	vec, ok := v[l]
	if !ok {
		return nil, fmt.Errorf("no vectors for %s", l)
	}
	return vec, nil
}

type memTrials struct {
	mu sync.Mutex
	ts []bench.Trial
}

func (m *memTrials) Append(t bench.Trial) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ts = append(m.ts, t)
	return nil
}
func (m *memTrials) All() ([]bench.Trial, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return append([]bench.Trial(nil), m.ts...), nil
}

func oneHot(words ...string) dat.Vectors {
	v := dat.Vectors{}
	for i, w := range words {
		vec := make([]float32, len(words))
		vec[i] = 1
		v[w] = vec
	}
	return v
}

var enWords = []string{"cat", "dog", "house", "river", "violin", "tomato", "ticket", "hippo"}
var jaWords = []string{"猫", "犬", "家", "川", "音楽", "トマト", "切符", "河馬"}

func newBench(gen *fakeGen) (*Benchmark, *memTrials) {
	trials := &memTrials{}
	b := &Benchmark{
		Gen: gen,
		Catalog: fakeCatalog{
			"m1": {ID: "m1", Mandatory: false, SupportedEfforts: []string{"high", "low", "none"}},
			"m2": {ID: "m2", Mandatory: true, SupportedEfforts: []string{"high", "medium", "low"}},
		},
		Vectors: fakeVectors{bench.LangEN: oneHot(enWords...), bench.LangJA: oneHot(jaWords...)},
		Trials:  trials,
		Now:     func() time.Time { return time.Unix(0, 0) },
	}
	return b, trials
}

// --- tests -------------------------------------------------------------

func TestRunCoversAllJobsAndScores(t *testing.T) {
	gen := &fakeGen{reply: func(r GenerateRequest) (Generation, error) {
		words := enWords
		if strings.Contains(r.Prompt, "名詞") {
			words = jaWords
		}
		return Generation{Content: strings.Join(words, "\n"), ReasoningTokens: 7, CostUSD: 0.01}, nil
	}}
	b, trials := newBench(gen)
	cfg := RunConfig{
		Languages:   []bench.Language{bench.LangEN, bench.LangJA},
		Models:      []bench.ModelSpec{{ID: "m1", Family: "F1"}, {ID: "m2", Family: "F2"}},
		Modes:       bench.AllModes,
		Runs:        3,
		Concurrency: 4,
		MaxTokens:   123,
	}
	stats, err := b.Run(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	all, _ := trials.All()
	if want := 2 * 2 * 2 * 3; len(all) != want || stats.Done != want || stats.Skipped != 0 {
		t.Fatalf("trials=%d stats=%+v want %d", len(all), stats, want)
	}
	seen := map[string]bool{}
	for _, tr := range all {
		if seen[tr.Key()] {
			t.Errorf("duplicate job %s", tr.Key())
		}
		seen[tr.Key()] = true
		if tr.Score == nil || *tr.Score != 100 {
			t.Errorf("%s: score %v want 100 (orthogonal one-hot vectors)", tr.Key(), tr.Score)
		}
		if len(tr.Used) != 7 || len(tr.Words) != 8 {
			t.Errorf("%s: used=%v words=%v", tr.Key(), tr.Used, tr.Words)
		}
		if tr.ReasoningTokens != 7 || tr.CostUSD != 0.01 || tr.Timestamp.IsZero() {
			t.Errorf("%s: metadata not copied: %+v", tr.Key(), tr)
		}
		wantEffort := map[string]map[bench.Mode]string{
			"m1": {bench.ModeThinking: "high", bench.ModeNoThinking: "none"},
			"m2": {bench.ModeThinking: "high", bench.ModeNoThinking: "low"},
		}[tr.ModelID][tr.Mode]
		if tr.Effort != wantEffort {
			t.Errorf("%s: effort %q want %q", tr.Key(), tr.Effort, wantEffort)
		}
		if tr.Family == "" {
			t.Errorf("%s: family missing", tr.Key())
		}
	}
	// Requests carry the resolved effort, the language prompt and max tokens.
	for _, c := range gen.calls {
		if c.MaxTokens != 123 || c.Prompt == "" || c.ModelID == "" {
			t.Errorf("bad request: %+v", c)
		}
	}
}

func TestRunResumesAndRecordsErrorsAndUnscorable(t *testing.T) {
	gen := &fakeGen{reply: func(r GenerateRequest) (Generation, error) {
		switch r.ModelID {
		case "m2":
			return Generation{}, errors.New("provider down")
		}
		if r.Effort == "none" {
			return Generation{Content: "cat\ndog\nhouse"}, nil // too few valid words
		}
		return Generation{Content: strings.Join(enWords, "\n")}, nil
	}}
	b, trials := newBench(gen)
	// A finished job from a previous invocation must be skipped.
	trials.Append(bench.Trial{Language: bench.LangEN, ModelID: "m1", Mode: bench.ModeThinking, Run: 0, Score: new(float64)})
	cfg := RunConfig{
		Languages: []bench.Language{bench.LangEN},
		Models:    []bench.ModelSpec{{ID: "m1"}, {ID: "m2"}},
		Modes:     bench.AllModes,
		Runs:      2,
	}
	stats, err := b.Run(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	if stats.Skipped != 1 || stats.Done != 7 || stats.Errors != 4 || stats.Unscorable != 2 {
		t.Errorf("stats %+v", stats)
	}
	if len(gen.calls) != 7 {
		t.Errorf("generator called %d times, want 7", len(gen.calls))
	}
	all, _ := trials.All()
	var errs, unscorable int
	for _, tr := range all {
		if tr.Error != "" {
			errs++
			if tr.Score != nil {
				t.Error("errored trial must have nil score")
			}
		} else if tr.Score == nil {
			unscorable++
			if len(tr.Valid) != 3 {
				t.Errorf("unscorable trial should keep valid words: %+v", tr.Valid)
			}
		}
	}
	if errs != 4 || unscorable != 2 {
		t.Errorf("errs=%d unscorable=%d", errs, unscorable)
	}
}

func TestRunRejectsUnknownModelOrLanguage(t *testing.T) {
	gen := &fakeGen{reply: func(GenerateRequest) (Generation, error) { return Generation{}, nil }}
	b, _ := newBench(gen)
	_, err := b.Run(context.Background(), RunConfig{Languages: []bench.Language{bench.LangEN}, Models: []bench.ModelSpec{{ID: "nope"}}, Modes: bench.AllModes, Runs: 1})
	if err == nil {
		t.Error("unknown model should fail before any request")
	}
	_, err = b.Run(context.Background(), RunConfig{Languages: []bench.Language{"xx"}, Models: []bench.ModelSpec{{ID: "m1"}}, Modes: bench.AllModes, Runs: 1})
	if err == nil {
		t.Error("unknown language should fail")
	}
	if len(gen.calls) != 0 {
		t.Error("no requests expected")
	}
}

func TestPlan(t *testing.T) {
	b, _ := newBench(&fakeGen{})
	plans, err := b.Plan(context.Background(), []bench.ModelSpec{{ID: "m1", Family: "F"}, {ID: "m2"}})
	if err != nil || len(plans) != 2 || plans[0].NoThinkingEffort != "none" || plans[1].NoThinkingEffort != "low" {
		t.Errorf("plans=%+v err=%v", plans, err)
	}
}

func TestRunRetriesErroredJobsOnResumeAndDropsCancelled(t *testing.T) {
	gen := &fakeGen{reply: func(r GenerateRequest) (Generation, error) {
		return Generation{Content: strings.Join(enWords, "\n")}, nil
	}}
	b, trials := newBench(gen)
	// An errored trial from a previous invocation must be re-run; a scored one must not.
	trials.Append(bench.Trial{Language: bench.LangEN, ModelID: "m1", Mode: bench.ModeThinking, Run: 0, Error: "http 502"})
	trials.Append(bench.Trial{Language: bench.LangEN, ModelID: "m1", Mode: bench.ModeThinking, Run: 1, Score: new(float64)})
	cfg := RunConfig{Languages: []bench.Language{bench.LangEN}, Models: []bench.ModelSpec{{ID: "m1"}}, Modes: []bench.Mode{bench.ModeThinking}, Runs: 2}
	stats, err := b.Run(context.Background(), cfg)
	if err != nil {
		t.Fatal(err)
	}
	if stats.Skipped != 1 || stats.Done != 1 || len(gen.calls) != 1 || gen.calls[0].ModelID != "m1" {
		t.Errorf("resume: stats=%+v calls=%d", stats, len(gen.calls))
	}

	// A job interrupted by cancellation must not be written as an errored trial.
	ctx, cancel := context.WithCancel(context.Background())
	gen2 := &fakeGen{reply: func(r GenerateRequest) (Generation, error) {
		cancel()
		<-ctx.Done()
		return Generation{}, ctx.Err()
	}}
	b2, trials2 := newBench(gen2)
	_, err = b2.Run(ctx, RunConfig{Languages: []bench.Language{bench.LangEN}, Models: []bench.ModelSpec{{ID: "m1"}}, Modes: []bench.Mode{bench.ModeThinking}, Runs: 3})
	if err == nil {
		t.Error("cancelled run should report the context error")
	}
	all, _ := trials2.All()
	for _, tr := range all {
		if strings.Contains(tr.Error, "context canceled") {
			t.Errorf("cancelled trial must not be recorded: %+v", tr)
		}
	}
}
