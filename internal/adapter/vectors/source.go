// Package vectors adapts embedding files (GloVe zip, fastText gzip) to the
// usecase ports and stores the filtered subset in a compact binary file.
package vectors

import (
	"archive/zip"
	"bufio"
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

// TextSource streams "word v1 v2 ... vN" lines (GloVe / fastText .vec).
// A fastText header line ("<count> <dim>") is skipped. Lines whose key
// contains spaces (e.g. ". . .") or whose values do not parse are skipped,
// matching how the reference implementation never sees them.
type TextSource struct {
	r     io.Reader
	close func() error
}

// NewTextSource wraps a reader of embedding lines.
func NewTextSource(r io.Reader) *TextSource { return &TextSource{r: r} }

// NewZipTextSource opens the first .txt entry of a zip (GloVe distribution).
func NewZipTextSource(path string) (*TextSource, error) {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, err
	}
	for _, f := range zr.File {
		if strings.HasSuffix(f.Name, ".txt") {
			rc, err := f.Open()
			if err != nil {
				zr.Close()
				return nil, err
			}
			return &TextSource{r: rc, close: func() error { rc.Close(); return zr.Close() }}, nil
		}
	}
	zr.Close()
	return nil, fmt.Errorf("no .txt entry in %s", path)
}

// NewGzipTextSource opens a gzip-compressed text file (fastText .vec.gz).
func NewGzipTextSource(path string) (*TextSource, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	gz, err := gzip.NewReader(f)
	if err != nil {
		f.Close()
		return nil, err
	}
	return &TextSource{r: gz, close: func() error { gz.Close(); return f.Close() }}, nil
}

// Each implements usecase.EmbeddingSource.
func (s *TextSource) Each(fn func(word string, vec []float32) error) (err error) {
	if s.close != nil {
		defer func() {
			if cerr := s.close(); err == nil {
				err = cerr
			}
		}()
	}
	sc := bufio.NewScanner(bufio.NewReaderSize(s.r, 8<<20))
	sc.Buffer(make([]byte, 1<<20), 1<<20)
	dim := 0
	first := true
	for sc.Scan() {
		line := sc.Text()
		if first {
			first = false
			if isHeader(line) {
				continue
			}
		}
		sp := strings.IndexByte(line, ' ')
		if sp <= 0 {
			continue
		}
		word := line[:sp]
		fields := strings.Fields(line[sp+1:])
		if dim == 0 {
			dim = len(fields)
		}
		if len(fields) != dim || dim == 0 {
			continue
		}
		vec := make([]float32, dim)
		ok := true
		for i, f := range fields {
			v, perr := strconv.ParseFloat(f, 32)
			if perr != nil {
				ok = false
				break
			}
			vec[i] = float32(v)
		}
		if !ok {
			continue
		}
		if err := fn(word, vec); err != nil {
			if errors.Is(err, usecase.ErrStop) {
				return nil
			}
			return err
		}
	}
	return sc.Err()
}

func isHeader(line string) bool {
	f := strings.Fields(line)
	if len(f) != 2 {
		return false
	}
	_, e1 := strconv.Atoi(f[0])
	_, e2 := strconv.Atoi(f[1])
	return e1 == nil && e2 == nil
}
