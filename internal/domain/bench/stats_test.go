package bench

import (
	"math"
	"testing"
)

func f(v float64) *float64 { return &v }

func TestSummarize(t *testing.T) {
	trials := []Trial{
		{Language: LangEN, ModelID: "a", Family: "A", Mode: ModeThinking, Score: f(80), ReasoningTokens: 100},
		{Language: LangEN, ModelID: "a", Family: "A", Mode: ModeThinking, Score: f(90), ReasoningTokens: 300},
		{Language: LangEN, ModelID: "a", Family: "A", Mode: ModeThinking, Error: "boom"},
		{Language: LangEN, ModelID: "a", Family: "A", Mode: ModeThinking, Score: nil}, // < 7 valid words
		{Language: LangEN, ModelID: "b", Family: "B", Mode: ModeThinking, Score: f(70)},
		{Language: LangEN, ModelID: "b", Family: "B", Mode: ModeNoThinking, Score: f(95)},
		{Language: LangJA, ModelID: "b", Family: "B", Mode: ModeThinking, Score: f(60)},
	}
	sums := Summarize(trials)
	find := func(l Language, id string, m Mode) Summary {
		for _, s := range sums {
			if s.Language == l && s.ModelID == id && s.Mode == m {
				return s
			}
		}
		t.Fatalf("missing summary %s %s %s", l, id, m)
		return Summary{}
	}
	a := find(LangEN, "a", ModeThinking)
	if a.N != 4 || a.Scored != 2 || a.Errors != 1 || a.Unscorable != 1 {
		t.Errorf("counts: %+v", a)
	}
	if math.Abs(a.Mean-85) > 1e-9 || math.Abs(a.SD-math.Sqrt(50)) > 1e-9 || a.Min != 80 || a.Max != 90 {
		t.Errorf("stats: %+v", a)
	}
	if a.MeanReasoningTokens != 200 {
		t.Errorf("reasoning tokens: %+v", a)
	}
	if len(sums) != 4 {
		t.Errorf("expected 4 groups, got %d", len(sums))
	}

	r := Rank(sums, LangEN, ModeThinking)
	if len(r) != 2 || r[0].ModelID != "a" || r[1].ModelID != "b" {
		t.Errorf("ranking: %+v", r)
	}
	if r[0].Rank != 1 || r[1].Rank != 2 {
		t.Errorf("rank numbers: %+v", r)
	}
}

func TestRankPutsUnscoredLast(t *testing.T) {
	sums := []Summary{
		{Language: LangEN, Mode: ModeThinking, ModelID: "none", N: 2},
		{Language: LangEN, Mode: ModeThinking, ModelID: "x", N: 2, Scored: 2, Mean: 50},
	}
	r := Rank(sums, LangEN, ModeThinking)
	if r[0].ModelID != "x" || r[1].ModelID != "none" {
		t.Errorf("ranking: %+v", r)
	}
}
