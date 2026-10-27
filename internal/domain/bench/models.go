package bench

// DefaultModels is the set requested for this benchmark (OpenRouter IDs).
var DefaultModels = []ModelSpec{
	// GPT-5.6 family
	{"openai/gpt-5.6-luna", "GPT-5.6"},
	{"openai/gpt-5.6-sol", "GPT-5.6"},
	{"openai/gpt-5.6-terra", "GPT-5.6"},
	// GPT-6 / 6.1
	{"openai/gpt-6-astra", "GPT-6"},
	{"openai/gpt-6.1-sol", "GPT-6.1"},
	// Gemini Flash family
	{"google/gemini-3.8-flash", "Gemini Flash"},
	{"google/gemini-3.7-flash", "Gemini Flash"},
	{"google/gemini-3.6-flash", "Gemini Flash"},
	{"google/gemini-3.5-flash", "Gemini Flash"},
	{"google/gemini-3.5-flash-lite", "Gemini Flash"},
	// Gemini Argon
	{"google/gemini-4-argon", "Gemini Argon"},
	// Grok >= 4.2 (incl. multi-agent)
	{"x-ai/grok-4.20", "Grok"},
	{"x-ai/grok-4.20-multi-agent", "Grok"},
	{"x-ai/grok-4.3", "Grok"},
	{"x-ai/grok-4.5", "Grok"},
	{"x-ai/grok-4.6", "Grok"},
	{"x-ai/grok-4.7", "Grok"},
	// Meta Muse
	{"meta/muse-wonder", "Meta Muse"},
	{"meta/muse-spark-1.3", "Meta Muse"},
	{"meta/muse-spark-open", "Meta Muse"},
	{"meta/muse-glimmer-30b", "Meta Muse"},
	// Claude
	{"anthropic/claude-fable-5", "Claude Fable"},
	{"anthropic/claude-fable-5.1", "Claude Fable"},
	{"anthropic/claude-opus-5.5", "Claude 5.5"},
	{"anthropic/claude-sonnet-5.5", "Claude 5.5"},
	{"anthropic/claude-haiku-5.5", "Claude 5.5"},
}
