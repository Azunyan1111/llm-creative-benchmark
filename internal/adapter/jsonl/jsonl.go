// Package jsonl persists trials as one JSON object per line (append-only).
package jsonl

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
)

// Repository implements usecase.TrialRepository on a JSONL file.
type Repository struct {
	path string
	mu   sync.Mutex
}

// New returns a repository for path (created on first Append).
func New(path string) *Repository { return &Repository{path: path} }

// Path returns the backing file.
func (r *Repository) Path() string { return r.path }

// Append writes one trial as a line.
func (r *Repository) Append(t bench.Trial) error {
	b, err := json.Marshal(t)
	if err != nil {
		return err
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if err := os.MkdirAll(filepath.Dir(r.path), 0o755); err != nil {
		return err
	}
	f, err := os.OpenFile(r.path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(append(b, '\n'))
	return err
}

// All reads every trial; a missing file yields an empty list.
func (r *Repository) All() ([]bench.Trial, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	f, err := os.Open(r.path)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var out []bench.Trial
	sc := bufio.NewScanner(f)
	sc.Buffer(make([]byte, 1<<20), 16<<20)
	for sc.Scan() {
		if len(sc.Bytes()) == 0 {
			continue
		}
		var t bench.Trial
		if err := json.Unmarshal(sc.Bytes(), &t); err != nil {
			return nil, err
		}
		out = append(out, t)
	}
	return out, sc.Err()
}
