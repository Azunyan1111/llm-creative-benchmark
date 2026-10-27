package bench

import "sort"

// EffortPlan says which reasoning effort each arm sends for a model.
type EffortPlan struct {
	ID                 string   `json:"id"`
	Family             string   `json:"family"`
	ReasoningMandatory bool     `json:"reasoning_mandatory"`
	SupportedEfforts   []string `json:"supported_efforts"` // weak -> strong
	ThinkingEffort     string   `json:"thinking_effort"`
	NoThinkingEffort   string   `json:"no_thinking_effort"`
	// NoThinkingIsTrueOff is true when the no-thinking arm really disables
	// reasoning (effort "none" or a model without reasoning); false means
	// "the weakest effort the provider allows".
	NoThinkingIsTrueOff bool `json:"no_thinking_is_true_off"`
}

// EffortFor returns the effort to send for a mode ("" = send no reasoning field).
func (p EffortPlan) EffortFor(m Mode) string {
	if m == ModeThinking {
		return p.ThinkingEffort
	}
	return p.NoThinkingEffort
}

// ThinkingEffort is what the thinking arm asks for when supported.
const ThinkingEffort = "high"

// Pseudo-efforts for models that only expose a reasoning on/off switch.
// Adapters translate them into the provider's toggle.
const (
	EffortEnabled  = "enabled"
	EffortDisabled = "disabled"
)

// effortRank orders efforts from least to most reasoning.
var effortRank = map[string]int{"none": 0, "minimal": 1, "low": 2, "medium": 3, "high": 4, "xhigh": 5, "max": 6}

// ResolveEfforts decides efforts from catalogue metadata:
//   - thinking:    "high" if supported, else the strongest supported effort
//   - no-thinking: "none" if supported, else "minimal", else the weakest supported
//   - toggle-only models (no effort levels): "enabled" / "disabled"
func ResolveEfforts(spec ModelSpec, info ModelInfo) EffortPlan {
	p := EffortPlan{ID: spec.ID, Family: spec.Family, ReasoningMandatory: info.Mandatory}
	if len(info.SupportedEfforts) == 0 {
		p.NoThinkingIsTrueOff = true // nothing to switch off, or a plain toggle
		if info.ReasoningToggle {
			p.ThinkingEffort, p.NoThinkingEffort = EffortEnabled, EffortDisabled
		}
		return p
	}
	eff := append([]string(nil), info.SupportedEfforts...)
	sort.SliceStable(eff, func(i, j int) bool { return effortRank[eff[i]] < effortRank[eff[j]] })
	p.SupportedEfforts = eff
	has := func(e string) bool {
		for _, x := range eff {
			if x == e {
				return true
			}
		}
		return false
	}
	if has(ThinkingEffort) {
		p.ThinkingEffort = ThinkingEffort
	} else {
		p.ThinkingEffort = eff[len(eff)-1]
	}
	switch {
	case has("none"):
		p.NoThinkingEffort, p.NoThinkingIsTrueOff = "none", true
	case has("minimal"):
		p.NoThinkingEffort = "minimal"
	default:
		p.NoThinkingEffort = eff[0]
	}
	return p
}
