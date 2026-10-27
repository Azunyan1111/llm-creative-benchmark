// Package bench holds the benchmark's domain entities and pure rules:
// languages, modes, models, effort planning, trial records and statistics.
// It depends on nothing outside the standard library and domain/dat.
package bench

import "time"

// Language of the task (prompt language and embedding space).
type Language string

const (
	LangEN Language = "en"
	LangJA Language = "ja"
)

// AllLanguages lists the supported languages in display order.
var AllLanguages = []Language{LangEN, LangJA}

// ParseLanguage returns the Language for a code, or "" if unknown.
func ParseLanguage(s string) Language {
	for _, l := range AllLanguages {
		if string(l) == s {
			return l
		}
	}
	return ""
}

// Mode is one of the two arms of the benchmark.
type Mode string

const (
	ModeThinking   Mode = "thinking"
	ModeNoThinking Mode = "no-thinking"
)

// AllModes lists the arms in display order.
var AllModes = []Mode{ModeThinking, ModeNoThinking}

// ModelSpec is a model under test.
type ModelSpec struct {
	ID     string `json:"id"`
	Family string `json:"family"`
}

// ModelInfo is the domain view of a catalogue entry (what a provider reports
// about a model's reasoning controls).
type ModelInfo struct {
	ID               string
	Mandatory        bool     // reasoning cannot be switched off
	SupportedEfforts []string // e.g. ["max","high","medium","low","none"]
	// ReasoningToggle is true when the provider accepts a reasoning on/off
	// switch for this model even though it lists no effort levels.
	ReasoningToggle bool
}

// Trial is one request/response and its score.
type Trial struct {
	Language         Language  `json:"language"`
	ModelID          string    `json:"model_id"`
	Family           string    `json:"family"`
	Mode             Mode      `json:"mode"`
	Effort           string    `json:"effort"`
	Run              int       `json:"run"`
	Timestamp        time.Time `json:"timestamp"`
	Provider         string    `json:"provider,omitempty"`
	RawContent       string    `json:"raw_content"`
	Words            []string  `json:"words"`
	Valid            []string  `json:"valid"`
	Used             []string  `json:"used"`
	Score            *float64  `json:"score"` // nil when unscorable or errored
	Error            string    `json:"error,omitempty"`
	FinishReason     string    `json:"finish_reason,omitempty"`
	PromptTokens     int       `json:"prompt_tokens"`
	CompletionTokens int       `json:"completion_tokens"`
	ReasoningTokens  int       `json:"reasoning_tokens"`
	CostUSD          float64   `json:"cost_usd"`
	LatencyMs        int64     `json:"latency_ms"`
}

// Key identifies the job a trial belongs to (used for resume).
func (t Trial) Key() string {
	return string(t.Language) + "|" + t.ModelID + "|" + string(t.Mode) + "|" + itoa(t.Run)
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}
	neg := i < 0
	if neg {
		i = -i
	}
	var b [20]byte
	p := len(b)
	for i > 0 {
		p--
		b[p] = byte('0' + i%10)
		i /= 10
	}
	if neg {
		p--
		b[p] = '-'
	}
	return string(b[p:])
}
