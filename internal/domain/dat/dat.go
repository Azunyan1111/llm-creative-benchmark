// Package dat implements the Divergent Association Task (DAT) score.
//
// The algorithm is a faithful port of the reference implementation by
// Jay Olson (https://github.com/jayolson/divergent-association-task, MIT),
// described in Olson et al. (2021), "Naming unrelated words predicts
// creativity", PNAS 118(25), e2022340118, doi:10.1073/pnas.2022340118.
//
// Score = mean cosine distance over all pairs of the first `minimum` (7)
// valid unique words, multiplied by 100.
package dat

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strings"
)

// Vectors maps a dictionary word to its GloVe embedding.
type Vectors map[string][]float32

var (
	stripRe   = regexp.MustCompile(`[^a-zA-Z- ]+`)
	spacesRe  = regexp.MustCompile(` +`)
	hyphensRe = regexp.MustCompile(`-+`)
)

// DefaultMinimum is the number of valid words used for scoring (7 in the paper).
const DefaultMinimum = 7

// Candidates turns a raw answer into lookup candidates, most preferred first.
// It encodes the language-specific cleaning rules of Validate.
type Candidates func(word string) []string

// Model scores word lists against a fixed set of embeddings.
type Model struct {
	vectors    Vectors
	candidates Candidates
}

// New creates an English scoring model (reference dat.py rules).
func New(v Vectors) *Model {
	return &Model{vectors: v, candidates: EnglishCandidates}
}

// NewWithCandidates creates a model with custom cleaning rules (e.g. Japanese).
func NewWithCandidates(v Vectors, c Candidates) *Model {
	return &Model{vectors: v, candidates: c}
}

// EnglishCandidates mirrors Model.validate in dat.py: strips non-letters,
// lowercases, then tries compound forms ("cul de sac" -> "cul-de-sac",
// "culdesac"; "top-hat" -> "top-hat", "tophat").
func EnglishCandidates(word string) []string {
	clean := strings.ToLower(strings.TrimSpace(stripRe.ReplaceAllString(word, "")))
	if len(clean) <= 1 {
		return nil
	}
	if strings.Contains(clean, " ") {
		return []string{
			spacesRe.ReplaceAllString(clean, "-"),
			spacesRe.ReplaceAllString(clean, ""),
		}
	}
	c := []string{clean}
	if strings.Contains(clean, "-") {
		c = append(c, hyphensRe.ReplaceAllString(clean, ""))
	}
	return c
}

// Size returns the number of words in the model.
func (m *Model) Size() int { return len(m.vectors) }

// Validate cleans a word and returns the first candidate found in the model.
func (m *Model) Validate(word string) (string, bool) {
	for _, c := range m.candidates(word) {
		if _, ok := m.vectors[c]; ok {
			return c, true
		}
	}
	return "", false
}

// Distance returns the cosine distance (0..2) between two validated words.
func (m *Model) Distance(a, b string) (float64, error) {
	va, ok := m.vectors[a]
	if !ok {
		return 0, fmt.Errorf("word %q not in model", a)
	}
	vb, ok := m.vectors[b]
	if !ok {
		return 0, fmt.Errorf("word %q not in model", b)
	}
	return cosineDistance(va, vb), nil
}

// cosineDistance matches scipy.spatial.distance.cosine: 1 - u.v/(|u||v|),
// computed in float64.
func cosineDistance(u, v []float32) float64 {
	var dot, nu, nv float64
	for i := range u {
		a, b := float64(u[i]), float64(v[i])
		dot += a * b
		nu += a * a
		nv += b * b
	}
	if nu == 0 || nv == 0 {
		return math.NaN()
	}
	return 1 - dot/(math.Sqrt(nu)*math.Sqrt(nv))
}

// Result is the outcome of scoring one word list.
type Result struct {
	// Valid holds every unique valid word, in input order.
	Valid []string `json:"valid"`
	// Used holds the first `minimum` valid words that were actually scored.
	Used []string `json:"used"`
	// Distances holds the pairwise cosine distances of Used (in combination order).
	Distances []float64 `json:"distances,omitempty"`
	// Score is the DAT score (mean distance * 100). Only meaningful if OK.
	Score float64 `json:"score"`
	// OK is false when fewer than `minimum` valid words were given.
	OK bool `json:"ok"`
}

// ErrNotEnoughWords is returned when fewer than `minimum` valid words exist.
var ErrNotEnoughWords = errors.New("not enough valid words")

// Score computes the DAT score of a word list, using the first `minimum`
// unique valid words. Mirrors Model.dat in dat.py.
func (m *Model) Score(words []string, minimum int) (Result, error) {
	if minimum < 2 {
		minimum = DefaultMinimum
	}
	res := Result{}
	seen := map[string]bool{}
	for _, w := range words {
		v, ok := m.Validate(w)
		if ok && !seen[v] {
			seen[v] = true
			res.Valid = append(res.Valid, v)
		}
	}
	if len(res.Valid) < minimum {
		return res, ErrNotEnoughWords
	}
	res.Used = res.Valid[:minimum]
	var sum float64
	for i := 0; i < len(res.Used); i++ {
		for j := i + 1; j < len(res.Used); j++ {
			d, err := m.Distance(res.Used[i], res.Used[j])
			if err != nil {
				return res, err
			}
			res.Distances = append(res.Distances, d)
			sum += d
		}
	}
	res.Score = sum / float64(len(res.Distances)) * 100
	res.OK = true
	return res, nil
}
