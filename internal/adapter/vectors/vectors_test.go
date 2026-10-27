package vectors

import (
	"archive/zip"
	"compress/gzip"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

func collect(t *testing.T, src usecase.EmbeddingSource, stopAfter int) map[string][]float32 {
	t.Helper()
	out := map[string][]float32{}
	err := src.Each(func(w string, v []float32) error {
		out[w] = v
		if stopAfter > 0 && len(out) >= stopAfter {
			return usecase.ErrStop
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	return out
}

func TestTextSourceParsesGloveAndFastText(t *testing.T) {
	glove := "cat 1.0 2.0\n. . . 0.5 0.5\ndog -1 0.25\nbad x y\n"
	got := collect(t, NewTextSource(strings.NewReader(glove)), 0)
	if len(got) != 2 || got["cat"][1] != 2 || got["dog"][0] != -1 {
		t.Errorf("glove parse: %v", got)
	}
	ft := "3 2\n猫 1 0\n犬 0 1\n鳥 1 1\n"
	got = collect(t, NewTextSource(strings.NewReader(ft)), 2)
	if len(got) != 2 || got["猫"][0] != 1 || got["犬"][1] != 1 {
		t.Errorf("fasttext parse / early stop: %v", got)
	}
}

func TestZipAndGzipSources(t *testing.T) {
	dir := t.TempDir()
	zp := filepath.Join(dir, "g.zip")
	zf, _ := os.Create(zp)
	zw := zip.NewWriter(zf)
	w, _ := zw.Create("glove.txt")
	w.Write([]byte("cat 1 2\n"))
	zw.Close()
	zf.Close()
	src, err := NewZipTextSource(zp)
	if err != nil {
		t.Fatal(err)
	}
	if got := collect(t, src, 0); got["cat"][1] != 2 {
		t.Errorf("zip: %v", got)
	}

	gp := filepath.Join(dir, "ft.vec.gz")
	gf, _ := os.Create(gp)
	gw := gzip.NewWriter(gf)
	gw.Write([]byte("1 2\n猫 3 4\n"))
	gw.Close()
	gf.Close()
	src, err = NewGzipTextSource(gp)
	if err != nil {
		t.Fatal(err)
	}
	if got := collect(t, src, 0); got["猫"][0] != 3 {
		t.Errorf("gzip: %v", got)
	}
}

func TestFileStoreRoundTrip(t *testing.T) {
	dir := t.TempDir()
	st := NewFileStore(dir)
	v := dat.Vectors{"cat": {1, 2, 3}, "猫": {-1, 0.5, 0}}
	if err := st.Save(bench.LangJA, v); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(st.Path(bench.LangJA)); err != nil {
		t.Fatal(err)
	}
	got, err := st.Load(bench.LangJA)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) != 2 || got["猫"][1] != 0.5 || got["cat"][2] != 3 {
		t.Errorf("roundtrip: %v", got)
	}
	if _, err := st.Load(bench.LangEN); err == nil {
		t.Error("missing language should error")
	}
}

func TestEnglishDictionaryKeep(t *testing.T) {
	p := filepath.Join(t.TempDir(), "words.txt")
	os.WriteFile(p, []byte("cat\nCat\ncul-de-sac\n-bad\nx\nokay's\n"), 0o644)
	keep, n, err := EnglishDictionary(p)
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 || !keep("cat") || !keep("cul-de-sac") || keep("Cat") || keep("x") || keep("okay's") || keep("dog") {
		t.Errorf("dictionary keep wrong (n=%d)", n)
	}
}

func TestJapaneseKeep(t *testing.T) {
	if !JapaneseKeep("猫") || !JapaneseKeep("ラーメン") || !JapaneseKeep("佐々木") || JapaneseKeep("cat") || JapaneseKeep("猫。") || JapaneseKeep("１２") {
		t.Error("JapaneseKeep")
	}
}
