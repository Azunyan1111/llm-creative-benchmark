package bench

// PromptEN reproduces the DAT instructions from Olson et al. (2021), PNAS,
// with an output-format line so replies can be parsed mechanically.
// The 4-minute time limit from the human study is omitted (not applicable).
const PromptEN = `Please enter 10 nouns that are as different from each other as possible, in all meanings and uses of the words.

Rules:
- Use only single words in English.
- Use only nouns (e.g., things, objects, concepts).
- Avoid proper nouns (e.g., no specific people or places).
- Avoid specialized vocabulary (e.g., no technical terms).
- Think of the words on your own (e.g., do not just look at objects in your surroundings).

Output format: reply with exactly 10 words, one word per line, with no numbering, punctuation, or explanation.`

// PromptJA is a Japanese rendering of the same instructions.
const PromptJA = `互いにできるだけ意味が離れた名詞を10個挙げてください。単語のあらゆる意味・用法を考慮したうえで、互いにできるだけ異なる単語を選んでください。

ルール:
- 日本語の単語を1語ずつ（句や文、説明は不可）。
- 名詞のみ（例: 物、物体、概念）。
- 固有名詞は不可（例: 特定の人名や地名）。
- 専門用語は不可（例: 学術用語や業界用語）。
- 自分で考えてください（例: 周囲にある物をそのまま挙げない）。

出力形式: 単語だけを1行に1つ、合計10行で出力してください。番号、句読点、説明は付けないでください。`

// PromptFor returns the task prompt for a language.
func PromptFor(l Language) (string, bool) {
	switch l {
	case LangEN:
		return PromptEN, true
	case LangJA:
		return PromptJA, true
	}
	return "", false
}
