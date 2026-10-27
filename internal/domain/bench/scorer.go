package bench

import "github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"

// NewScorer builds a DAT model with the cleaning rules of the language.
func NewScorer(l Language, v dat.Vectors) *dat.Model {
	if l == LangJA {
		return dat.NewWithCandidates(v, dat.JapaneseCandidates)
	}
	return dat.New(v)
}
