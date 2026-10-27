package dat

import (
	"regexp"
	"strings"
	"unicode"
)

// Japanese rules. The reference DAT is English-only; this adaptation keeps the
// same scoring and applies analogous cleaning for Japanese answers:
//   - keep only kana, kanji, the long-vowel mark and iteration mark
//   - drop everything else (ASCII, punctuation, spaces, parentheses)
//   - for an answer with an interpunct or space ("ぬいぐるみ・人形") also try the
//     joined form, as dat.py does for "cul de sac"
var (
	jaSepRe = regexp.MustCompile(`[\s・/／]+`)
	// JapanesePattern matches a single Japanese token in the fastText vocabulary.
	JapanesePattern = regexp.MustCompile(`^[\p{Hiragana}\p{Katakana}\p{Han}ー々〆ヶ]+$`)
)

func isJa(r rune) bool {
	return unicode.Is(unicode.Hiragana, r) || unicode.Is(unicode.Katakana, r) ||
		unicode.Is(unicode.Han, r) || r == 'ー' || r == '々' || r == '〆' || r == 'ヶ'
}

// JapaneseCandidates returns lookup candidates for a Japanese answer.
func JapaneseCandidates(word string) []string {
	word = strings.TrimSpace(word)
	// Drop a parenthesised reading/explanation such as "林檎（りんご）".
	if i := strings.IndexAny(word, "（("); i > 0 {
		word = word[:i]
	}
	parts := jaSepRe.Split(word, -1)
	var kept []string
	for _, p := range parts {
		var b strings.Builder
		for _, r := range p {
			if isJa(r) {
				b.WriteRune(r)
			}
		}
		if b.Len() > 0 {
			kept = append(kept, b.String())
		}
	}
	if len(kept) == 0 {
		return nil
	}
	joined := strings.Join(kept, "")
	if len([]rune(joined)) == 0 {
		return nil
	}
	if len(kept) == 1 {
		return []string{joined}
	}
	// Multi-part answer: prefer the joined compound, then the first part.
	return []string{joined, kept[0]}
}
