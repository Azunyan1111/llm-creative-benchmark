package jsonl

import (
	"path/filepath"
	"sync"
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
)

func TestAppendAllRoundTrip(t *testing.T) {
	p := filepath.Join(t.TempDir(), "sub", "trials.jsonl")
	repo := New(p)
	all, err := repo.All()
	if err != nil || len(all) != 0 {
		t.Fatalf("empty repo: %v %v", all, err)
	}
	s := 80.5
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			if err := repo.Append(bench.Trial{Language: bench.LangEN, ModelID: "m", Mode: bench.ModeThinking, Run: i, Score: &s, Words: []string{"a", "b"}}); err != nil {
				t.Error(err)
			}
		}(i)
	}
	wg.Wait()
	all, err = repo.All()
	if err != nil || len(all) != 20 {
		t.Fatalf("got %d trials, err=%v", len(all), err)
	}
	seen := map[int]bool{}
	for _, tr := range all {
		seen[tr.Run] = true
		if tr.Score == nil || *tr.Score != 80.5 || len(tr.Words) != 2 {
			t.Errorf("bad trial: %+v", tr)
		}
	}
	if len(seen) != 20 {
		t.Errorf("runs lost: %v", seen)
	}
}
