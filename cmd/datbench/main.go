// Command datbench measures LLM creativity with the Divergent Association
// Task over OpenRouter. It is the composition root: it wires adapters into
// the use cases and parses the command line.
//
//	datbench prepare --lang en|ja        build data/<lang>.vectors.bin from the raw embeddings
//	datbench plan                        show the reasoning effort each model gets per arm
//	datbench run [flags]                 query models and append trials to results/trials.jsonl
//	datbench report                      write results/ranking.md and results/summary.json
//	datbench score [--lang en] [--min 7] w1 w2 ...   score an ad-hoc word list
package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"time"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/adapter/jsonl"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/adapter/markdown"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/adapter/openrouter"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/adapter/vectors"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

const (
	dataDir     = "data"
	resultsDir  = "results"
	gloveZip    = "data/glove/glove.840B.300d.zip"
	fastTextGz  = "data/fasttext/cc.ja.300.vec.gz"
	englishDict = "assets/words.txt"
	trialsFile  = "results/trials.jsonl"
	// jaVocabLimit keeps the most frequent Japanese tokens of fastText's
	// 2M-word vocabulary (the file is frequency-ordered).
	jaVocabLimit = 300000
)

func main() {
	log.SetFlags(log.Ltime)
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	var err error
	switch os.Args[1] {
	case "prepare":
		err = cmdPrepare(os.Args[2:])
	case "plan":
		err = cmdPlan(ctx, os.Args[2:])
	case "run":
		err = cmdRun(ctx, os.Args[2:])
	case "report":
		err = cmdReport(os.Args[2:])
	case "score":
		err = cmdScore(os.Args[2:])
	default:
		usage()
		os.Exit(2)
	}
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: datbench <prepare|plan|run|report|score> [flags]")
}

func cmdPrepare(args []string) error {
	fs := flag.NewFlagSet("prepare", flag.ExitOnError)
	lang := fs.String("lang", "en", "language to prepare: en or ja")
	fs.Parse(args)
	store := vectors.NewFileStore(dataDir)
	var (
		src  usecase.EmbeddingSource
		keep func(string) bool
		lim  int
		err  error
	)
	switch bench.ParseLanguage(*lang) {
	case bench.LangEN:
		var n int
		keep, n, err = vectors.EnglishDictionary(englishDict)
		if err != nil {
			return err
		}
		log.Printf("dictionary: %d words match %s", n, vectors.DictionaryPattern)
		src, err = vectors.NewZipTextSource(gloveZip)
	case bench.LangJA:
		keep, lim = vectors.JapaneseKeep, jaVocabLimit
		src, err = vectors.NewGzipTextSource(fastTextGz)
	default:
		return fmt.Errorf("unknown language %q", *lang)
	}
	if err != nil {
		return err
	}
	start := time.Now()
	n, err := usecase.PrepareEmbeddings(src, keep, lim, bench.Language(*lang), store)
	if err != nil {
		return err
	}
	log.Printf("wrote %d vectors to %s in %s", n, store.Path(bench.Language(*lang)), time.Since(start).Round(time.Second))
	return nil
}

func modelFlags(fs *flag.FlagSet) (*string, *string) {
	return fs.String("models", "", "comma-separated OpenRouter model IDs (default: built-in list)"),
		fs.String("family", "", "only models whose family contains this text")
}

func selectModels(models, family string) []bench.ModelSpec {
	out := bench.DefaultModels
	if models != "" {
		out = nil
		for _, id := range strings.Split(models, ",") {
			id = strings.TrimSpace(id)
			fam := id
			for _, m := range bench.DefaultModels {
				if m.ID == id {
					fam = m.Family
				}
			}
			out = append(out, bench.ModelSpec{ID: id, Family: fam})
		}
	}
	if family != "" {
		var f []bench.ModelSpec
		for _, m := range out {
			if strings.Contains(strings.ToLower(m.Family), strings.ToLower(family)) {
				f = append(f, m)
			}
		}
		out = f
	}
	return out
}

func cmdPlan(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("plan", flag.ExitOnError)
	models, family := modelFlags(fs)
	fs.Parse(args)
	client := openrouter.New(os.Getenv("OPENROUTER_API_KEY"), "")
	b := &usecase.Benchmark{Catalog: client}
	plans, err := b.Plan(ctx, selectModels(*models, *family))
	if err != nil {
		return err
	}
	fmt.Printf("%-32s %-14s %-9s %-10s %-12s %s\n", "MODEL", "FAMILY", "MANDATORY", "THINKING", "NO-THINKING", "SUPPORTED")
	for _, p := range plans {
		fmt.Printf("%-32s %-14s %-9v %-10s %-12s %s\n", p.ID, p.Family, p.ReasoningMandatory, p.ThinkingEffort, p.NoThinkingEffort, strings.Join(p.SupportedEfforts, ","))
	}
	return nil
}

func cmdRun(ctx context.Context, args []string) error {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	models, family := modelFlags(fs)
	langs := fs.String("lang", "en,ja", "comma-separated languages")
	modes := fs.String("modes", "thinking,no-thinking", "comma-separated modes")
	runs := fs.Int("runs", 10, "trials per (language, model, mode)")
	conc := fs.Int("concurrency", 6, "parallel requests")
	maxTokens := fs.Int("max-tokens", 16000, "completion budget per request (includes reasoning)")
	out := fs.String("out", trialsFile, "JSONL file to append trials to (resumes if present)")
	fs.Parse(args)

	client, err := openrouter.NewFromEnv()
	if err != nil {
		return err
	}
	cfg := usecase.RunConfig{Runs: *runs, Concurrency: *conc, MaxTokens: *maxTokens, Models: selectModels(*models, *family)}
	for _, l := range strings.Split(*langs, ",") {
		if pl := bench.ParseLanguage(strings.TrimSpace(l)); pl != "" {
			cfg.Languages = append(cfg.Languages, pl)
		} else {
			return fmt.Errorf("unknown language %q", l)
		}
	}
	for _, m := range strings.Split(*modes, ",") {
		cfg.Modes = append(cfg.Modes, bench.Mode(strings.TrimSpace(m)))
	}
	repo := jsonl.New(*out)
	b := &usecase.Benchmark{
		Gen:     client,
		Catalog: client,
		Vectors: vectors.NewFileStore(dataDir),
		Trials:  repo,
		Log:     log.Printf,
	}
	plans, err := b.Plan(ctx, cfg.Models)
	if err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(filepath.Dir(*out), "plan.json"), plans); err != nil {
		return err
	}
	stats, err := b.Run(ctx, cfg)
	log.Printf("planned=%d skipped=%d done=%d errors=%d unscorable=%d", stats.Planned, stats.Skipped, stats.Done, stats.Errors, stats.Unscorable)
	return err
}

func cmdReport(args []string) error {
	fs := flag.NewFlagSet("report", flag.ExitOnError)
	in := fs.String("in", trialsFile, "JSONL trials file")
	fs.Parse(args)
	trials, err := jsonl.New(*in).All()
	if err != nil {
		return err
	}
	if len(trials) == 0 {
		return fmt.Errorf("no trials in %s", *in)
	}
	report := usecase.BuildReport(trials)
	trials = usecase.LatestPerJob(trials)
	dir := filepath.Dir(*in)
	notes := []string{
		"thinking = reasoning effort `high`; no-thinking = effort `none` where the provider allows switching reasoning off, otherwise the weakest effort the provider accepts (`minimal` or `low`) — see the Effort column and the average reasoning tokens.",
		"English uses GloVe 840B-300d (the paper's model) with the reference dictionary; Japanese uses fastText cc.ja.300 (Common Crawl + Wikipedia). Scores are comparable within a language, not across languages.",
		"Unscorable = the reply had fewer than 7 valid dictionary words; such trials are excluded from the mean.",
	}
	md := markdown.Render(report, markdown.Meta{Generated: time.Now().Format("2006-01-02 15:04 MST"), Notes: notes, Trials: trials})
	if err := os.WriteFile(filepath.Join(dir, "ranking.md"), []byte(md), 0o644); err != nil {
		return err
	}
	if err := writeJSON(filepath.Join(dir, "summary.json"), report.Summaries); err != nil {
		return err
	}
	fmt.Print(md[:strings.Index(md, "## Appendix")])
	log.Printf("wrote %s and %s", filepath.Join(dir, "ranking.md"), filepath.Join(dir, "summary.json"))
	return nil
}

func cmdScore(args []string) error {
	fs := flag.NewFlagSet("score", flag.ExitOnError)
	lang := fs.String("lang", "en", "language: en or ja")
	min := fs.Int("min", 0, "valid words to score (default 7, as in the paper)")
	fs.Parse(args)
	l := bench.ParseLanguage(*lang)
	if l == "" {
		return fmt.Errorf("unknown language %q", *lang)
	}
	v, err := vectors.NewFileStore(dataDir).Load(l)
	if err != nil {
		return err
	}
	res, err := bench.NewScorer(l, v).Score(fs.Args(), *min)
	if err != nil {
		return fmt.Errorf("%w (valid: %v)", err, res.Valid)
	}
	fmt.Printf("score=%.2f used=%v\n", res.Score, res.Used)
	return nil
}

func writeJSON(path string, v any) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(b, '\n'), 0o644)
}
