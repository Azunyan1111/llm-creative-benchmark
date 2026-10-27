package dat

import (
	"math"
	"testing"
)

func unit(x, y float32) []float32 { return []float32{x, y} }

func TestValidateAndScore(t *testing.T) {
	m := New(Vectors{
		"cat":        unit(1, 0),
		"dog":        unit(0, 1),
		"cul-de-sac": unit(1, 1),
		"tophat":     unit(-1, 0),
	})
	cases := map[string]string{
		"Cat":        "cat",
		" dog. ":     "dog",
		"cul de sac": "cul-de-sac",
		"top-hat":    "tophat",
		"x":          "",
		"zebra":      "",
		"1. cat":     "cat",
	}
	for in, want := range cases {
		got, ok := m.Validate(in)
		if got != want || ok != (want != "") {
			t.Errorf("Validate(%q) = %q,%v want %q", in, got, ok, want)
		}
	}
	r, err := m.Score([]string{"cat", "Cat", "dog"}, 2)
	if err != nil || !r.OK {
		t.Fatalf("score failed: %v", err)
	}
	if math.Abs(r.Score-100) > 1e-9 { // orthogonal vectors: distance 1 -> 100
		t.Errorf("score = %v want 100", r.Score)
	}
	if len(r.Used) != 2 || r.Used[0] != "cat" || r.Used[1] != "dog" {
		t.Errorf("used = %v", r.Used)
	}
	if _, err := m.Score([]string{"cat"}, 2); err != ErrNotEnoughWords {
		t.Errorf("expected ErrNotEnoughWords, got %v", err)
	}
}

func TestJapaneseCandidates(t *testing.T) {
	cases := map[string][]string{
		"猫":        {"猫"},
		"1. 犬。":    {"犬"},
		"林檎（りんご）":  {"林檎"},
		"ぬいぐるみ・人形": {"ぬいぐるみ人形", "ぬいぐるみ"},
		"ラーメン":     {"ラーメン"},
		"cat":      nil,
		"**自転車**":  {"自転車"},
		"電子 レンジ":   {"電子レンジ", "電子"},
	}
	for in, want := range cases {
		got := JapaneseCandidates(in)
		if len(got) != len(want) {
			t.Errorf("JapaneseCandidates(%q) = %v want %v", in, got, want)
			continue
		}
		for i := range got {
			if got[i] != want[i] {
				t.Errorf("JapaneseCandidates(%q) = %v want %v", in, got, want)
			}
		}
	}
}
