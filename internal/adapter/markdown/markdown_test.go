package markdown

import (
	"strings"
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

func fp(v float64) *float64 { return &v }

func TestRender(t *testing.T) {
	trials := []bench.Trial{
		{Language: bench.LangEN, ModelID: "openai/x", Family: "GPT", Mode: bench.ModeThinking, Effort: "high", Score: fp(80), ReasoningTokens: 100, Used: []string{"cat", "dog"}},
		{Language: bench.LangEN, ModelID: "openai/x", Family: "GPT", Mode: bench.ModeNoThinking, Effort: "none", Score: fp(85.5)},
		{Language: bench.LangEN, ModelID: "y/y", Family: "Y", Mode: bench.ModeThinking, Effort: "low", Error: "boom"},
		{Language: bench.LangJA, ModelID: "openai/x", Family: "GPT", Mode: bench.ModeThinking, Effort: "high", Score: fp(60), Used: []string{"猫"}},
	}
	out := Render(usecase.BuildReport(trials), Meta{Runs: 1, Generated: "2026-10-27", Trials: trials})
	for _, want := range []string{
		"# ", "## English", "## 日本語", "### thinking", "### no-thinking",
		"| 1 | openai/x |", "85.50", "80.00", "y/y", "boom", "cat, dog", "猫", "high", "none",
	} {
		if !strings.Contains(out, want) {
			t.Errorf("rendered report lacks %q\n%s", want, out)
		}
	}
	if !strings.Contains(out, "| Rank | Model |") {
		t.Errorf("table header missing\n%s", out)
	}
}
