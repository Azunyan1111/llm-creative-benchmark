package bench

import (
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/dat"
)

func TestNewScorerUsesLanguageRules(t *testing.T) {
	v := dat.Vectors{"cat": {1, 0}, "猫": {0, 1}}
	en := NewScorer(LangEN, v)
	if w, ok := en.Validate("Cat!"); !ok || w != "cat" {
		t.Errorf("en validate: %q %v", w, ok)
	}
	if _, ok := en.Validate("猫"); ok {
		t.Error("en scorer must not accept Japanese")
	}
	ja := NewScorer(LangJA, v)
	if w, ok := ja.Validate("猫。"); !ok || w != "猫" {
		t.Errorf("ja validate: %q %v", w, ok)
	}
}
