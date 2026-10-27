package bench

import (
	"regexp"
	"strings"
)

var (
	listMarkerRe = regexp.MustCompile(`^\s*(?:\d+[\.\):]|[-*•]|\(\d+\))\s*`)
	mdRe         = regexp.MustCompile("[`*_~\"'“”‘’]")
)

// ParseWords extracts candidate words from a model reply.
// One entry per line is expected; a single comma-separated line is also
// accepted. Lines with more than three tokens are treated as prose and dropped.
func ParseWords(content string) []string {
	content = strings.ReplaceAll(content, "\r\n", "\n")
	lines := strings.Split(content, "\n")
	var items []string
	for _, l := range lines {
		l = strings.TrimSpace(mdRe.ReplaceAllString(listMarkerRe.ReplaceAllString(l, ""), ""))
		if l == "" {
			continue
		}
		items = append(items, l)
	}
	// Fallback: everything on one line, separated by commas or semicolons.
	if len(items) == 1 && strings.ContainsAny(items[0], ",;") {
		items = regexp.MustCompile(`[,;]`).Split(items[0], -1)
	}
	var words []string
	for _, it := range items {
		it = strings.TrimSpace(strings.Trim(it, ".,;:!?"))
		if it == "" {
			continue
		}
		if len(strings.Fields(it)) > 3 {
			continue // prose, not a word
		}
		words = append(words, it)
	}
	return words
}
