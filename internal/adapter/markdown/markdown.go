// Package markdown renders a usecase.Report as a Markdown document.
package markdown

import (
	"fmt"
	"strings"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

// Meta is extra information printed in the header.
type Meta struct {
	Runs      int
	Generated string
	Notes     []string
	// Trials, when given, are listed in an appendix (words and scores).
	Trials []bench.Trial
}

var langTitle = map[bench.Language]string{bench.LangEN: "English", bench.LangJA: "日本語"}

// Render produces the ranking document.
func Render(r usecase.Report, m Meta) string {
	var b strings.Builder
	fmt.Fprintf(&b, "# LLM Creativity Benchmark (DAT) — Ranking\n\n")
	fmt.Fprintf(&b, "Generated: %s · Trials: %d · Total cost: $%.4f\n\n", m.Generated, r.TotalTrials, r.TotalCostUSD)
	b.WriteString("Score = Divergent Association Task (Olson et al., 2021): mean cosine distance of the first 7 valid words × 100. Higher is more divergent (humans average ≈ 78, SD ≈ 6).\n\n")
	for _, n := range m.Notes {
		fmt.Fprintf(&b, "- %s\n", n)
	}
	if len(m.Notes) > 0 {
		b.WriteString("\n")
	}
	for _, l := range r.Languages() {
		fmt.Fprintf(&b, "## %s\n\n", langTitle[l])
		writeSideBySide(&b, r, l)
		for _, mode := range bench.AllModes {
			rows := r.Ranking(l, mode)
			if len(rows) == 0 {
				continue
			}
			fmt.Fprintf(&b, "### %s\n\n", mode)
			b.WriteString("| Rank | Model | Family | Effort | Mean | SD | Min | Max | Scored/N | Errors | Unscorable | Avg reasoning tokens |\n")
			b.WriteString("|---:|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|\n")
			for _, s := range rows {
				fmt.Fprintf(&b, "| %d | %s | %s | %s | %s | %s | %s | %s | %d/%d | %d | %d | %.0f |\n",
					s.Rank, s.ModelID, s.Family, s.Effort, num(s.Mean, s.Scored), num(s.SD, s.Scored), num(s.Min, s.Scored), num(s.Max, s.Scored),
					s.Scored, s.N, s.Errors, s.Unscorable, s.MeanReasoningTokens)
			}
			b.WriteString("\n")
		}
	}
	if len(m.Trials) > 0 {
		b.WriteString(RenderTrials(m.Trials))
	}
	return b.String()
}

func writeSideBySide(b *strings.Builder, r usecase.Report, l bench.Language) {
	rows := r.SideBySide(l)
	if len(rows) == 0 {
		return
	}
	b.WriteString("| Model | Family | thinking (mean ± SD) | no-thinking (mean ± SD) | Δ (thinking − no-thinking) |\n")
	b.WriteString("|---|---|---:|---:|---:|\n")
	for _, row := range rows {
		th, hasTh := row.ByMode[bench.ModeThinking]
		nt, hasNt := row.ByMode[bench.ModeNoThinking]
		delta := "—"
		if hasTh && hasNt && th.Scored > 0 && nt.Scored > 0 {
			delta = fmt.Sprintf("%+.2f", th.Mean-nt.Mean)
		}
		fmt.Fprintf(b, "| %s | %s | %s | %s | %s |\n", row.ModelID, row.Family, cell(th, hasTh), cell(nt, hasNt), delta)
	}
	b.WriteString("\n")
}

func cell(s bench.Summary, has bool) string {
	if !has {
		return "—"
	}
	if s.Scored == 0 {
		return fmt.Sprintf("— (%d errors)", s.Errors)
	}
	return fmt.Sprintf("%.2f ± %.2f (%s)", s.Mean, s.SD, s.Effort)
}

func num(v float64, scored int) string {
	if scored == 0 {
		return "—"
	}
	return fmt.Sprintf("%.2f", v)
}

// RenderTrials lists every trial's words and score, grouped by language,
// model and mode, so readers can inspect what each model actually said.
func RenderTrials(trials []bench.Trial) string {
	var b strings.Builder
	b.WriteString("## Appendix: words per trial\n")
	last := ""
	for _, t := range trials {
		head := fmt.Sprintf("### %s · %s · %s (effort=%s)", langTitle[t.Language], t.ModelID, t.Mode, t.Effort)
		if head != last {
			fmt.Fprintf(&b, "\n%s\n\n", head)
			last = head
		}
		switch {
		case t.Error != "":
			fmt.Fprintf(&b, "- run %d: ERROR %s\n", t.Run, t.Error)
		case t.Score == nil:
			fmt.Fprintf(&b, "- run %d: unscorable (%d valid) — %s\n", t.Run, len(t.Valid), strings.Join(t.Words, ", "))
		default:
			fmt.Fprintf(&b, "- run %d: **%.2f** — %s\n", t.Run, *t.Score, strings.Join(t.Used, ", "))
		}
	}
	return b.String()
}
