package usecase

import (
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

// PrepareEmbeddings streams a source, keeps the words accepted by keep
// (at most limit distinct words when limit > 0; sources are expected to be
// frequency-ordered so this keeps the most common ones) and stores them.
// A repeated word overwrites the earlier vector, like the reference
// implementation's dict assignment.
func PrepareEmbeddings(src EmbeddingSource, keep func(string) bool, limit int, l bench.Language, store VectorStore) (int, error) {
	v := dat.Vectors{}
	err := src.Each(func(word string, vec []float32) error {
		if !keep(word) {
			return nil
		}
		if _, seen := v[word]; !seen && limit > 0 && len(v) >= limit {
			return ErrStop
		}
		v[word] = vec
		return nil
	})
	if err != nil {
		return 0, err
	}
	return len(v), store.Save(l, v)
}
