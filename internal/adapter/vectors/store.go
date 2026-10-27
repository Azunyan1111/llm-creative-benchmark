package vectors

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
	"path/filepath"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

const magic = "DATVEC01"

// FileStore keeps one binary vectors file per language in a directory.
// It implements usecase.VectorStore and usecase.VectorRepository.
type FileStore struct{ dir string }

// NewFileStore creates a store rooted at dir.
func NewFileStore(dir string) *FileStore { return &FileStore{dir: dir} }

// Path is the file for a language.
func (s *FileStore) Path(l bench.Language) string {
	return filepath.Join(s.dir, string(l)+".vectors.bin")
}

// Save writes vectors: magic, dim, count, then (len, word, floats) entries.
func (s *FileStore) Save(l bench.Language, vectors dat.Vectors) error {
	if err := os.MkdirAll(s.dir, 0o755); err != nil {
		return err
	}
	dim := 0
	for _, v := range vectors {
		dim = len(v)
		break
	}
	tmp := s.Path(l) + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}
	w := bufio.NewWriterSize(f, 4<<20)
	hdr := make([]byte, len(magic)+8)
	copy(hdr, magic)
	binary.LittleEndian.PutUint32(hdr[len(magic):], uint32(dim))
	binary.LittleEndian.PutUint32(hdr[len(magic)+4:], uint32(len(vectors)))
	w.Write(hdr)
	buf := make([]byte, 4*dim)
	for word, vec := range vectors {
		if len(vec) != dim || len(word) > math.MaxUint16 {
			f.Close()
			os.Remove(tmp)
			return fmt.Errorf("inconsistent entry %q (dim %d, want %d)", word, len(vec), dim)
		}
		var l [2]byte
		binary.LittleEndian.PutUint16(l[:], uint16(len(word)))
		w.Write(l[:])
		w.WriteString(word)
		for i, v := range vec {
			binary.LittleEndian.PutUint32(buf[4*i:], math.Float32bits(v))
		}
		w.Write(buf)
	}
	if err := w.Flush(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, s.Path(l))
}

// Load reads the vectors of a language.
func (s *FileStore) Load(l bench.Language) (dat.Vectors, error) {
	f, err := os.Open(s.Path(l))
	if err != nil {
		return nil, fmt.Errorf("%w (run `datbench prepare --lang %s` first)", err, l)
	}
	defer f.Close()
	r := bufio.NewReaderSize(f, 4<<20)
	hdr := make([]byte, len(magic)+8)
	if _, err := io.ReadFull(r, hdr); err != nil {
		return nil, err
	}
	if string(hdr[:len(magic)]) != magic {
		return nil, fmt.Errorf("%s: bad magic", s.Path(l))
	}
	dim := int(binary.LittleEndian.Uint32(hdr[len(magic):]))
	count := int(binary.LittleEndian.Uint32(hdr[len(magic)+4:]))
	vectors := make(dat.Vectors, count)
	buf := make([]byte, 4*dim)
	for i := 0; i < count; i++ {
		var l [2]byte
		if _, err := io.ReadFull(r, l[:]); err != nil {
			return nil, err
		}
		wb := make([]byte, binary.LittleEndian.Uint16(l[:]))
		if _, err := io.ReadFull(r, wb); err != nil {
			return nil, err
		}
		if _, err := io.ReadFull(r, buf); err != nil {
			return nil, err
		}
		vec := make([]float32, dim)
		for j := range vec {
			vec[j] = math.Float32frombits(binary.LittleEndian.Uint32(buf[4*j:]))
		}
		vectors[string(wb)] = vec
	}
	return vectors, nil
}
