package bench

import "testing"

func TestDefaultModelsIncludeArgon(t *testing.T) {
	seen := map[string]bool{}
	for _, m := range DefaultModels {
		if m.ID == "" || m.Family == "" {
			t.Errorf("model with empty id or family: %+v", m)
		}
		if seen[m.ID] {
			t.Errorf("duplicate model id %q", m.ID)
		}
		seen[m.ID] = true
	}
	found := false
	for _, m := range DefaultModels {
		if m.ID == "google/gemini-4-argon" {
			found = true
			if m.Family != "Gemini Argon" {
				t.Errorf("argon family = %q, want %q", m.Family, "Gemini Argon")
			}
		}
	}
	if !found {
		t.Errorf("google/gemini-4-argon is missing from DefaultModels")
	}
}
