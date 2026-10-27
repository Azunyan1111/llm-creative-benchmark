// Package usecase contains the application logic of the benchmark. It
// depends only on the domain and on the interfaces (ports) declared here;
// adapters implement them.
package usecase

import (
	"context"
	"errors"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

// GenerateRequest asks a model to answer the DAT prompt.
type GenerateRequest struct {
	ModelID   string
	Prompt    string
	Effort    string // reasoning effort; "" sends no reasoning control
	MaxTokens int
}

// Generation is a model's reply plus usage metadata.
type Generation struct {
	Content          string
	Reasoning        string
	Provider         string
	FinishReason     string
	PromptTokens     int
	CompletionTokens int
	ReasoningTokens  int
	CostUSD          float64
}

// Generator talks to an LLM provider.
type Generator interface {
	Generate(ctx context.Context, req GenerateRequest) (Generation, error)
}

// Catalog reports provider metadata for a model.
type Catalog interface {
	Lookup(ctx context.Context, id string) (info bench.ModelInfo, found bool, err error)
}

// VectorRepository loads prepared embeddings for a language.
type VectorRepository interface {
	Load(l bench.Language) (dat.Vectors, error)
}

// TrialRepository persists trials. Append must be safe for concurrent use.
type TrialRepository interface {
	Append(t bench.Trial) error
	All() ([]bench.Trial, error)
}

// EmbeddingSource streams raw word vectors (e.g. a GloVe text file).
// Each stops early, returning nil, when fn returns ErrStop.
type EmbeddingSource interface {
	Each(fn func(word string, vec []float32) error) error
}

// ErrStop tells an EmbeddingSource to stop iterating without error.
var ErrStop = errors.New("stop")

// VectorStore persists prepared embeddings for a language.
type VectorStore interface {
	Save(l bench.Language, v dat.Vectors) error
}

// Logger receives progress messages; nil loggers are ignored.
type Logger func(format string, args ...any)

func (l Logger) printf(format string, args ...any) {
	if l != nil {
		l(format, args...)
	}
}
