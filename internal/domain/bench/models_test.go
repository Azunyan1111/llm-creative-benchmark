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

func TestDefaultModelsIncludeGemini4Lineup(t *testing.T) {
	family := map[string]string{}
	for _, m := range DefaultModels {
		family[m.ID] = m.Family
	}
	for _, id := range []string{"google/gemini-4-carbon", "google/gemini-4-hydrogen", "google/gemini-4-neon"} {
		f, ok := family[id]
		if !ok {
			t.Errorf("%s is missing from DefaultModels", id)
			continue
		}
		if f != "Gemini 4" {
			t.Errorf("%s family = %q, want %q", id, f, "Gemini 4")
		}
	}
}
