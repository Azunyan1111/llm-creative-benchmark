package vectors

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

// DictionaryPattern mirrors the default `pattern` of the reference dat.Model.
var DictionaryPattern = regexp.MustCompile(`^[a-z][a-z-]*[a-z]$`)

// EnglishDictionary reads words.txt and returns a keep function accepting
// the words matching DictionaryPattern, plus how many there are.
func EnglishDictionary(path string) (func(string) bool, int, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()
	words := map[string]bool{}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		w := strings.TrimRight(sc.Text(), "\r\n")
		if DictionaryPattern.MatchString(w) {
			words[w] = true
		}
	}
	if err := sc.Err(); err != nil {
		return nil, 0, err
	}
	return func(w string) bool { return words[w] }, len(words), nil
}

// JapaneseKeep accepts tokens made only of kana/kanji (no digits, Latin or
// punctuation), i.e. plausible single Japanese words from fastText's vocabulary.
func JapaneseKeep(w string) bool { return dat.JapanesePattern.MatchString(w) }
