package usecase

import (
	"sort"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
)

// Report is the aggregated view of all trials.
type Report struct {
	Summaries    []bench.Summary
	TotalTrials  int
	TotalCostUSD float64
}

// ModelRow shows one model's summaries for every mode side by side.
type ModelRow struct {
	ModelID string
	Family  string
	ByMode  map[bench.Mode]bench.Summary
}

// BuildReport aggregates trials. When a job was run more than once (a retry
// after an error), only its latest trial counts.
func BuildReport(trials []bench.Trial) Report {
	trials = LatestPerJob(trials)
	r := Report{Summaries: bench.Summarize(trials), TotalTrials: len(trials)}
	for _, t := range trials {
		r.TotalCostUSD += t.CostUSD
	}
	return r
}

// LatestPerJob keeps the last trial recorded for each job, in first-seen order.
func LatestPerJob(trials []bench.Trial) []bench.Trial {
	idx := map[string]int{}
	var out []bench.Trial
	for _, t := range trials {
		if i, ok := idx[t.Key()]; ok {
			out[i] = t
			continue
		}
		idx[t.Key()] = len(out)
		out = append(out, t)
	}
	return out
}

// Languages lists the languages present, in bench.AllLanguages order.
func (r Report) Languages() []bench.Language {
	present := map[bench.Language]bool{}
	for _, s := range r.Summaries {
		present[s.Language] = true
	}
	var out []bench.Language
	for _, l := range bench.AllLanguages {
		if present[l] {
			out = append(out, l)
		}
	}
	return out
}

// Ranking orders models for one language and mode, best first.
func (r Report) Ranking(l bench.Language, m bench.Mode) []bench.Summary {
	return bench.Rank(r.Summaries, l, m)
}

// SideBySide returns one row per model for a language, ordered by the best
// mean across modes.
func (r Report) SideBySide(l bench.Language) []ModelRow {
	rows := map[string]*ModelRow{}
	var order []string
	for _, s := range r.Summaries {
		if s.Language != l {
			continue
		}
		row, ok := rows[s.ModelID]
		if !ok {
			row = &ModelRow{ModelID: s.ModelID, Family: s.Family, ByMode: map[bench.Mode]bench.Summary{}}
			rows[s.ModelID] = row
			order = append(order, s.ModelID)
		}
		row.ByMode[s.Mode] = s
	}
	best := func(row *ModelRow) float64 {
		b := -1.0
		for _, s := range row.ByMode {
			if s.Scored > 0 && s.Mean > b {
				b = s.Mean
			}
		}
		return b
	}
	sort.SliceStable(order, func(i, j int) bool { return best(rows[order[i]]) > best(rows[order[j]]) })
	out := make([]ModelRow, 0, len(order))
	for _, id := range order {
		out = append(out, *rows[id])
	}
	return out
}
