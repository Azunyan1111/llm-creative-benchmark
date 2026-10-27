package usecase

import (
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
)

func fp(v float64) *float64 { return &v }

func TestBuildReport(t *testing.T) {
	trials := []bench.Trial{
		{Language: bench.LangEN, ModelID: "a", Mode: bench.ModeThinking, Score: fp(80), CostUSD: 1},
		{Language: bench.LangEN, ModelID: "b", Mode: bench.ModeThinking, Score: fp(90), CostUSD: 1},
		{Language: bench.LangEN, ModelID: "a", Mode: bench.ModeNoThinking, Score: fp(85)},
		{Language: bench.LangJA, ModelID: "a", Mode: bench.ModeThinking, Score: fp(70)},
	}
	r := BuildReport(trials)
	if r.TotalTrials != 4 || r.TotalCostUSD != 2 {
		t.Errorf("totals: %+v", r)
	}
	en := r.Ranking(bench.LangEN, bench.ModeThinking)
	if len(en) != 2 || en[0].ModelID != "b" || en[0].Rank != 1 {
		t.Errorf("en thinking ranking: %+v", en)
	}
	if got := r.Ranking(bench.LangJA, bench.ModeNoThinking); len(got) != 0 {
		t.Errorf("expected empty ranking, got %+v", got)
	}
	// Side-by-side rows: one per (language, model) with both modes,
	// ordered by the best mean across modes (b: 90 beats a: 85).
	rows := r.SideBySide(bench.LangEN)
	if len(rows) != 2 {
		t.Fatalf("rows: %+v", rows)
	}
	if rows[0].ModelID != "b" {
		t.Errorf("row order: %+v", rows)
	}
	if _, ok := rows[0].ByMode[bench.ModeNoThinking]; ok {
		t.Errorf("row b should lack no-thinking: %+v", rows[0])
	}
	if rows[1].ModelID != "a" || rows[1].ByMode[bench.ModeThinking].Mean != 80 || rows[1].ByMode[bench.ModeNoThinking].Mean != 85 {
		t.Errorf("row a: %+v", rows[1])
	}
	if r.Languages()[0] != bench.LangEN || r.Languages()[1] != bench.LangJA {
		t.Errorf("languages: %v", r.Languages())
	}
}

func TestBuildReportKeepsLatestTrialPerJob(t *testing.T) {
	trials := []bench.Trial{
		{Language: bench.LangEN, ModelID: "a", Mode: bench.ModeThinking, Run: 0, Error: "http 502"},
		{Language: bench.LangEN, ModelID: "a", Mode: bench.ModeThinking, Run: 0, Score: fp(80)}, // retry of the same job
		{Language: bench.LangEN, ModelID: "a", Mode: bench.ModeThinking, Run: 1, Score: fp(90)},
	}
	r := BuildReport(trials)
	if r.TotalTrials != 2 {
		t.Errorf("duplicates should collapse: %+v", r)
	}
	s := r.Ranking(bench.LangEN, bench.ModeThinking)[0]
	if s.N != 2 || s.Errors != 0 || s.Scored != 2 || s.Mean != 85 {
		t.Errorf("summary: %+v", s)
	}
}
