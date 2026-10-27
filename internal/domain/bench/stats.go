package bench

import (
	"math"
	"sort"
)

// Summary aggregates the trials of one (language, model, mode) group.
type Summary struct {
	Language            Language `json:"language"`
	ModelID             string   `json:"model_id"`
	Family              string   `json:"family"`
	Mode                Mode     `json:"mode"`
	Effort              string   `json:"effort"`
	N                   int      `json:"n"`          // trials attempted
	Scored              int      `json:"scored"`     // trials with a score
	Unscorable          int      `json:"unscorable"` // replied but < 7 valid words
	Errors              int      `json:"errors"`     // request failed
	Mean                float64  `json:"mean"`
	SD                  float64  `json:"sd"` // sample standard deviation
	Min                 float64  `json:"min"`
	Max                 float64  `json:"max"`
	MeanReasoningTokens float64  `json:"mean_reasoning_tokens"`
	TotalCostUSD        float64  `json:"total_cost_usd"`
	Rank                int      `json:"rank,omitempty"`
}

// Summarize groups trials and computes per-group statistics.
func Summarize(trials []Trial) []Summary {
	type key struct {
		l Language
		m string
		d Mode
	}
	groups := map[key]*Summary{}
	scores := map[key][]float64{}
	var order []key
	for _, t := range trials {
		k := key{t.Language, t.ModelID, t.Mode}
		s, ok := groups[k]
		if !ok {
			s = &Summary{Language: t.Language, ModelID: t.ModelID, Family: t.Family, Mode: t.Mode, Effort: t.Effort}
			groups[k] = s
			order = append(order, k)
		}
		s.N++
		s.TotalCostUSD += t.CostUSD
		switch {
		case t.Error != "":
			s.Errors++
		case t.Score == nil:
			s.Unscorable++
		default:
			s.Scored++
			scores[k] = append(scores[k], *t.Score)
			s.MeanReasoningTokens += float64(t.ReasoningTokens)
		}
	}
	out := make([]Summary, 0, len(order))
	for _, k := range order {
		s := groups[k]
		if xs := scores[k]; len(xs) > 0 {
			s.Mean, s.SD, s.Min, s.Max = describe(xs)
			s.MeanReasoningTokens /= float64(len(xs))
		}
		out = append(out, *s)
	}
	return out
}

func describe(xs []float64) (mean, sd, min, max float64) {
	min, max = math.Inf(1), math.Inf(-1)
	for _, x := range xs {
		mean += x
		min = math.Min(min, x)
		max = math.Max(max, x)
	}
	mean /= float64(len(xs))
	if len(xs) > 1 {
		var ss float64
		for _, x := range xs {
			ss += (x - mean) * (x - mean)
		}
		sd = math.Sqrt(ss / float64(len(xs)-1))
	}
	return
}

// Rank filters summaries to one language and mode and orders them by mean
// score, best first. Groups without any score go last. Rank numbers are set.
func Rank(sums []Summary, l Language, m Mode) []Summary {
	var out []Summary
	for _, s := range sums {
		if s.Language == l && s.Mode == m {
			out = append(out, s)
		}
	}
	sort.SliceStable(out, func(i, j int) bool {
		a, b := out[i], out[j]
		if (a.Scored > 0) != (b.Scored > 0) {
			return a.Scored > 0
		}
		return a.Mean > b.Mean
	})
	for i := range out {
		out[i].Rank = i + 1
	}
	return out
}
