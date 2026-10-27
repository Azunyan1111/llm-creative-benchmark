package bench

import "testing"

func TestResolveEfforts(t *testing.T) {
	cases := []struct {
		name        string
		info        ModelInfo
		wantThink   string
		wantNoThink string
		wantTrueOff bool
	}{
		{"optional reasoning with none",
			ModelInfo{ID: "m", Mandatory: false, SupportedEfforts: []string{"max", "xhigh", "high", "medium", "low", "none"}},
			"high", "none", true},
		{"mandatory with minimal",
			ModelInfo{ID: "m", Mandatory: true, SupportedEfforts: []string{"high", "medium", "low", "minimal"}},
			"high", "minimal", false},
		{"mandatory low only",
			ModelInfo{ID: "m", Mandatory: true, SupportedEfforts: []string{"xhigh", "high", "medium", "low"}},
			"high", "low", false},
		{"no high supported",
			ModelInfo{ID: "m", Mandatory: true, SupportedEfforts: []string{"medium", "low"}},
			"medium", "low", false},
		{"no reasoning metadata",
			ModelInfo{ID: "m"},
			"", "", true},
	}
	for _, c := range cases {
		p := ResolveEfforts(ModelSpec{ID: "m", Family: "F"}, c.info)
		if p.ThinkingEffort != c.wantThink || p.NoThinkingEffort != c.wantNoThink || p.NoThinkingIsTrueOff != c.wantTrueOff {
			t.Errorf("%s: got %+v", c.name, p)
		}
		if p.ID != "m" || p.Family != "F" {
			t.Errorf("%s: spec not carried: %+v", c.name, p)
		}
	}
	p := ResolveEfforts(ModelSpec{ID: "m"}, ModelInfo{SupportedEfforts: []string{"low", "high", "none", "medium"}})
	if got := p.SupportedEfforts; len(got) != 4 || got[0] != "none" || got[3] != "high" {
		t.Errorf("supported efforts not sorted weak->strong: %v", got)
	}
	if p.EffortFor(ModeThinking) != "high" || p.EffortFor(ModeNoThinking) != "none" {
		t.Errorf("EffortFor mismatch: %+v", p)
	}
}

func TestResolveEffortsToggleOnlyModel(t *testing.T) {
	// grok-4.20 style: reasoning can be switched on/off but has no effort levels.
	p := ResolveEfforts(ModelSpec{ID: "m"}, ModelInfo{ID: "m", ReasoningToggle: true})
	if p.ThinkingEffort != EffortEnabled || p.NoThinkingEffort != EffortDisabled || !p.NoThinkingIsTrueOff {
		t.Errorf("toggle-only plan: %+v", p)
	}
}
