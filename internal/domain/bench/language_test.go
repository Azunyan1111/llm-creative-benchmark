package bench

import (
	"strings"
	"testing"
)

func TestLanguages(t *testing.T) {
	for _, l := range []Language{LangEN, LangJA} {
		p, ok := PromptFor(l)
		if !ok || !strings.Contains(p, "10") {
			t.Errorf("prompt for %s: %q", l, p)
		}
	}
	if _, ok := PromptFor(Language("xx")); ok {
		t.Error("unknown language should not have a prompt")
	}
	if ParseLanguage("en") != LangEN || ParseLanguage("ja") != LangJA || ParseLanguage("zz") != "" {
		t.Error("ParseLanguage")
	}
}
