package usecase

import (
	"errors"
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

type sliceSource [][2]any // [word, vec]

func (s sliceSource) Each(fn func(word string, vec []float32) error) error {
	for _, e := range s {
		if err := fn(e[0].(string), e[1].([]float32)); err != nil {
			if errors.Is(err, ErrStop) {
				return nil
			}
			return err
		}
	}
	return nil
}

type memStore struct {
	lang bench.Language
	v    dat.Vectors
}

func (m *memStore) Save(l bench.Language, v dat.Vectors) error { m.lang, m.v = l, v; return nil }

func TestPrepareEmbeddingsFiltersAndLimits(t *testing.T) {
	src := sliceSource{
		{"cat", []float32{1, 0}},
		{"Cat", []float32{0, 1}}, // rejected by keep
		{"dog", []float32{0, 1}},
		{"cat", []float32{2, 0}},  // duplicate: later wins, does not count twice
		{"bird", []float32{1, 1}}, // beyond limit
	}
	store := &memStore{}
	keep := func(w string) bool { return w == "cat" || w == "dog" || w == "bird" }
	n, err := PrepareEmbeddings(src, keep, 2, bench.LangEN, store)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 || store.lang != bench.LangEN || len(store.v) != 2 {
		t.Errorf("n=%d store=%+v", n, store)
	}
	if store.v["cat"][0] != 2 {
		t.Errorf("later duplicate should overwrite: %v", store.v["cat"])
	}
	if _, ok := store.v["bird"]; ok {
		t.Error("limit should stop before bird")
	}
}

func TestPrepareEmbeddingsNoLimit(t *testing.T) {
	src := sliceSource{{"a", []float32{1}}, {"b", []float32{1}}}
	store := &memStore{}
	n, err := PrepareEmbeddings(src, func(string) bool { return true }, 0, bench.LangJA, store)
	if err != nil || n != 2 {
		t.Errorf("n=%d err=%v", n, err)
	}
}
