package bench

import (
	"reflect"
	"testing"
)

func TestParseWords(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"cat\ndog\n", []string{"cat", "dog"}},
		{"1. cat\n2) dog\n- bird\n* fish", []string{"cat", "dog", "bird", "fish"}},
		{"**cat**\n`dog`", []string{"cat", "dog"}},
		{"cat, dog, cul de sac", []string{"cat", "dog", "cul de sac"}},
		{"Here are ten words that are very different:\ncat\ndog", []string{"cat", "dog"}},
	}
	for _, c := range cases {
		if got := ParseWords(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("ParseWords(%q) = %v want %v", c.in, got, c.want)
		}
	}
}
