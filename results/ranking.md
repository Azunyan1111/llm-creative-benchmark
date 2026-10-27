# LLM Creativity Benchmark (DAT) — Ranking

Generated: 2026-10-27 11:42 JST · Trials: 1000 · Total cost: $14.2572

Score = Divergent Association Task (Olson et al., 2021): mean cosine distance of the first 7 valid words × 100. Higher is more divergent (humans average ≈ 78, SD ≈ 6).

- thinking = reasoning effort `high`; no-thinking = effort `none` where the provider allows switching reasoning off, otherwise the weakest effort the provider accepts (`minimal` or `low`) — see the Effort column and the average reasoning tokens.
- English uses GloVe 840B-300d (the paper's model) with the reference dictionary; Japanese uses fastText cc.ja.300 (Common Crawl + Wikipedia). Scores are comparable within a language, not across languages.
- Unscorable = the reply had fewer than 7 valid dictionary words; such trials are excluded from the mean.

## English

| Model | Family | thinking (mean ± SD) | no-thinking (mean ± SD) | Δ (thinking − no-thinking) |
|---|---|---:|---:|---:|
| google/gemini-3.8-flash | Gemini Flash | 89.17 ± 2.84 (high) | 87.64 ± 2.97 (low) | +1.53 |
| anthropic/claude-opus-5.5 | Claude 5.5 | 88.07 ± 1.44 (high) | 87.34 ± 0.34 (low) | +0.73 |
| google/gemini-3.5-flash | Gemini Flash | 84.05 ± 2.74 (high) | 88.04 ± 2.00 (minimal) | -3.99 |
| google/gemini-3.6-flash | Gemini Flash | 86.61 ± 2.13 (high) | 87.76 ± 4.14 (minimal) | -1.15 |
| google/gemini-3.7-flash | Gemini Flash | 87.68 ± 1.20 (high) | 87.10 ± 2.90 (low) | +0.58 |
| openai/gpt-6.1-sol | GPT-6.1 | 87.33 ± 2.03 (high) | 85.24 ± 2.40 (low) | +2.09 |
| anthropic/claude-sonnet-5.5 | Claude 5.5 | 86.62 ± 2.26 (high) | 86.13 ± 2.95 (low) | +0.49 |
| openai/gpt-6-astra | GPT-6 | 86.41 ± 1.47 (high) | 86.31 ± 2.51 (low) | +0.10 |
| anthropic/claude-fable-5.1 | Claude Fable | 86.08 ± 1.55 (high) | 85.89 ± 2.32 (low) | +0.19 |
| meta/muse-wonder | Meta Muse | 85.31 ± 3.43 (high) | 79.94 ± 3.86 (minimal) | +5.37 |
| anthropic/claude-haiku-5.5 | Claude 5.5 | 82.73 ± 2.03 (high) | 84.91 ± 2.50 (none) | -2.18 |
| anthropic/claude-fable-5 | Claude Fable | 84.79 ± 1.66 (high) | 84.84 ± 0.11 (low) | -0.05 |
| google/gemini-3.5-flash-lite | Gemini Flash | 84.15 ± 3.57 (high) | 80.63 ± 2.86 (minimal) | +3.51 |
| openai/gpt-5.6-sol | GPT-5.6 | 83.59 ± 3.34 (high) | 82.67 ± 2.88 (none) | +0.92 |
| openai/gpt-5.6-luna | GPT-5.6 | 80.86 ± 2.80 (high) | 82.18 ± 2.48 (none) | -1.32 |
| x-ai/grok-4.6 | Grok | 80.47 ± 2.71 (high) | 79.73 ± 2.56 (low) | +0.74 |
| meta/muse-spark-open | Meta Muse | 79.65 ± 3.00 (high) | 76.45 ± 2.71 (low) | +3.20 |
| x-ai/grok-4.7 | Grok | 78.77 ± 3.41 (high) | 77.99 ± 3.19 (low) | +0.78 |
| openai/gpt-5.6-terra | GPT-5.6 | 77.31 ± 2.60 (high) | 78.12 ± 3.17 (none) | -0.81 |
| x-ai/grok-4.5 | Grok | 74.85 ± 3.76 (high) | 77.12 ± 4.40 (low) | -2.27 |
| x-ai/grok-4.20-multi-agent | Grok | 75.69 ± 2.29 (high) | 76.86 ± 4.08 (low) | -1.18 |
| meta/muse-spark-1.3 | Meta Muse | 76.16 ± 2.60 (high) | 75.52 ± 4.09 (minimal) | +0.63 |
| x-ai/grok-4.20 | Grok | 74.93 ± 2.74 (enabled) | 73.53 ± 1.87 (disabled) | +1.41 |
| x-ai/grok-4.3 | Grok | 74.57 ± 2.11 (high) | 72.25 ± 3.70 (none) | +2.32 |
| meta/muse-glimmer-30b | Meta Muse | 72.56 ± 2.91 (high) | 73.66 ± 2.10 (low) | -1.10 |

### thinking

| Rank | Model | Family | Effort | Mean | SD | Min | Max | Scored/N | Errors | Unscorable | Avg reasoning tokens |
|---:|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| 1 | google/gemini-3.8-flash | Gemini Flash | high | 89.17 | 2.84 | 85.40 | 94.12 | 10/10 | 0 | 0 | 9499 |
| 2 | anthropic/claude-opus-5.5 | Claude 5.5 | high | 88.07 | 1.44 | 86.33 | 90.65 | 10/10 | 0 | 0 | 287 |
| 3 | google/gemini-3.7-flash | Gemini Flash | high | 87.68 | 1.20 | 85.69 | 89.39 | 10/10 | 0 | 0 | 2108 |
| 4 | openai/gpt-6.1-sol | GPT-6.1 | high | 87.33 | 2.03 | 84.50 | 89.84 | 10/10 | 0 | 0 | 398 |
| 5 | anthropic/claude-sonnet-5.5 | Claude 5.5 | high | 86.62 | 2.26 | 83.96 | 91.82 | 10/10 | 0 | 0 | 0 |
| 6 | google/gemini-3.6-flash | Gemini Flash | high | 86.61 | 2.13 | 83.92 | 90.11 | 10/10 | 0 | 0 | 2660 |
| 7 | openai/gpt-6-astra | GPT-6 | high | 86.41 | 1.47 | 84.66 | 88.86 | 10/10 | 0 | 0 | 444 |
| 8 | anthropic/claude-fable-5.1 | Claude Fable | high | 86.08 | 1.55 | 83.22 | 87.47 | 10/10 | 0 | 0 | 335 |
| 9 | meta/muse-wonder | Meta Muse | high | 85.31 | 3.43 | 78.39 | 89.01 | 10/10 | 0 | 0 | 3068 |
| 10 | anthropic/claude-fable-5 | Claude Fable | high | 84.79 | 1.66 | 82.62 | 88.15 | 10/10 | 0 | 0 | 67 |
| 11 | google/gemini-3.5-flash-lite | Gemini Flash | high | 84.15 | 3.57 | 79.67 | 91.51 | 10/10 | 0 | 0 | 1818 |
| 12 | google/gemini-3.5-flash | Gemini Flash | high | 84.05 | 2.74 | 80.08 | 89.22 | 10/10 | 0 | 0 | 2153 |
| 13 | openai/gpt-5.6-sol | GPT-5.6 | high | 83.59 | 3.34 | 77.72 | 87.60 | 10/10 | 0 | 0 | 205 |
| 14 | anthropic/claude-haiku-5.5 | Claude 5.5 | high | 82.73 | 2.03 | 79.77 | 85.63 | 10/10 | 0 | 0 | 564 |
| 15 | openai/gpt-5.6-luna | GPT-5.6 | high | 80.86 | 2.80 | 75.73 | 85.36 | 10/10 | 0 | 0 | 286 |
| 16 | x-ai/grok-4.6 | Grok | high | 80.47 | 2.71 | 75.13 | 83.94 | 10/10 | 0 | 0 | 1981 |
| 17 | meta/muse-spark-open | Meta Muse | high | 79.65 | 3.00 | 74.99 | 84.67 | 10/10 | 0 | 0 | 2038 |
| 18 | x-ai/grok-4.7 | Grok | high | 78.77 | 3.41 | 72.10 | 83.28 | 10/10 | 0 | 0 | 1859 |
| 19 | openai/gpt-5.6-terra | GPT-5.6 | high | 77.31 | 2.60 | 74.06 | 80.94 | 10/10 | 0 | 0 | 20 |
| 20 | meta/muse-spark-1.3 | Meta Muse | high | 76.16 | 2.60 | 72.23 | 80.35 | 10/10 | 0 | 0 | 4899 |
| 21 | x-ai/grok-4.20-multi-agent | Grok | high | 75.69 | 2.29 | 72.17 | 79.35 | 10/10 | 0 | 0 | 33358 |
| 22 | x-ai/grok-4.20 | Grok | enabled | 74.93 | 2.74 | 71.55 | 79.90 | 10/10 | 0 | 0 | 1882 |
| 23 | x-ai/grok-4.5 | Grok | high | 74.85 | 3.76 | 69.24 | 81.00 | 10/10 | 0 | 0 | 749 |
| 24 | x-ai/grok-4.3 | Grok | high | 74.57 | 2.11 | 70.16 | 77.03 | 10/10 | 0 | 0 | 972 |
| 25 | meta/muse-glimmer-30b | Meta Muse | high | 72.56 | 2.91 | 67.62 | 76.78 | 10/10 | 0 | 0 | 938 |

### no-thinking

| Rank | Model | Family | Effort | Mean | SD | Min | Max | Scored/N | Errors | Unscorable | Avg reasoning tokens |
|---:|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| 1 | google/gemini-3.5-flash | Gemini Flash | minimal | 88.04 | 2.00 | 85.12 | 91.16 | 10/10 | 0 | 0 | 0 |
| 2 | google/gemini-3.6-flash | Gemini Flash | minimal | 87.76 | 4.14 | 79.60 | 94.43 | 10/10 | 0 | 0 | 0 |
| 3 | google/gemini-3.8-flash | Gemini Flash | low | 87.64 | 2.97 | 84.57 | 92.75 | 10/10 | 0 | 0 | 962 |
| 4 | anthropic/claude-opus-5.5 | Claude 5.5 | low | 87.34 | 0.34 | 87.14 | 88.07 | 10/10 | 0 | 0 | 0 |
| 5 | google/gemini-3.7-flash | Gemini Flash | low | 87.10 | 2.90 | 82.27 | 91.88 | 10/10 | 0 | 0 | 757 |
| 6 | openai/gpt-6-astra | GPT-6 | low | 86.31 | 2.51 | 83.18 | 90.23 | 10/10 | 0 | 0 | 101 |
| 7 | anthropic/claude-sonnet-5.5 | Claude 5.5 | low | 86.13 | 2.95 | 82.35 | 91.19 | 10/10 | 0 | 0 | 0 |
| 8 | anthropic/claude-fable-5.1 | Claude Fable | low | 85.89 | 2.32 | 81.19 | 88.23 | 10/10 | 0 | 0 | 0 |
| 9 | openai/gpt-6.1-sol | GPT-6.1 | low | 85.24 | 2.40 | 80.87 | 88.48 | 10/10 | 0 | 0 | 102 |
| 10 | anthropic/claude-haiku-5.5 | Claude 5.5 | none | 84.91 | 2.50 | 79.98 | 87.46 | 10/10 | 0 | 0 | 0 |
| 11 | anthropic/claude-fable-5 | Claude Fable | low | 84.84 | 0.11 | 84.78 | 85.08 | 10/10 | 0 | 0 | 13 |
| 12 | openai/gpt-5.6-sol | GPT-5.6 | none | 82.67 | 2.88 | 78.23 | 86.97 | 10/10 | 0 | 0 | 0 |
| 13 | openai/gpt-5.6-luna | GPT-5.6 | none | 82.18 | 2.48 | 78.06 | 85.22 | 10/10 | 0 | 0 | 0 |
| 14 | google/gemini-3.5-flash-lite | Gemini Flash | minimal | 80.63 | 2.86 | 75.15 | 84.88 | 10/10 | 0 | 0 | 0 |
| 15 | meta/muse-wonder | Meta Muse | minimal | 79.94 | 3.86 | 72.64 | 84.31 | 10/10 | 0 | 0 | 162 |
| 16 | x-ai/grok-4.6 | Grok | low | 79.73 | 2.56 | 76.86 | 85.77 | 10/10 | 0 | 0 | 697 |
| 17 | openai/gpt-5.6-terra | GPT-5.6 | none | 78.12 | 3.17 | 71.44 | 81.59 | 10/10 | 0 | 0 | 0 |
| 18 | x-ai/grok-4.7 | Grok | low | 77.99 | 3.19 | 73.43 | 84.01 | 10/10 | 0 | 0 | 426 |
| 19 | x-ai/grok-4.5 | Grok | low | 77.12 | 4.40 | 73.21 | 85.50 | 10/10 | 0 | 0 | 378 |
| 20 | x-ai/grok-4.20-multi-agent | Grok | low | 76.86 | 4.08 | 71.15 | 83.82 | 10/10 | 0 | 0 | 8738 |
| 21 | meta/muse-spark-open | Meta Muse | low | 76.45 | 2.71 | 72.57 | 82.71 | 10/10 | 0 | 0 | 290 |
| 22 | meta/muse-spark-1.3 | Meta Muse | minimal | 75.52 | 4.09 | 70.85 | 82.77 | 10/10 | 0 | 0 | 136 |
| 23 | meta/muse-glimmer-30b | Meta Muse | low | 73.66 | 2.10 | 69.50 | 77.17 | 10/10 | 0 | 0 | 177 |
| 24 | x-ai/grok-4.20 | Grok | disabled | 73.53 | 1.87 | 69.96 | 75.12 | 10/10 | 0 | 0 | 0 |
| 25 | x-ai/grok-4.3 | Grok | none | 72.25 | 3.70 | 67.32 | 76.10 | 10/10 | 0 | 0 | 0 |

## 日本語

| Model | Family | thinking (mean ± SD) | no-thinking (mean ± SD) | Δ (thinking − no-thinking) |
|---|---|---:|---:|---:|
| google/gemini-3.8-flash | Gemini Flash | 81.53 ± 1.89 (high) | 82.82 ± 2.62 (low) | -1.30 |
| openai/gpt-6-astra | GPT-6 | 82.00 ± 1.50 (high) | 81.60 ± 1.74 (low) | +0.40 |
| openai/gpt-6.1-sol | GPT-6.1 | 80.74 ± 1.94 (high) | 81.61 ± 0.89 (low) | -0.86 |
| google/gemini-3.5-flash | Gemini Flash | 81.28 ± 1.46 (high) | 76.56 ± 2.57 (minimal) | +4.72 |
| google/gemini-3.7-flash | Gemini Flash | 81.14 ± 1.87 (high) | 80.57 ± 1.68 (low) | +0.57 |
| google/gemini-3.6-flash | Gemini Flash | 80.74 ± 2.26 (high) | 78.27 ± 1.60 (minimal) | +2.47 |
| openai/gpt-5.6-sol | GPT-5.6 | 80.56 ± 1.95 (high) | 78.00 ± 2.32 (none) | +2.56 |
| anthropic/claude-sonnet-5.5 | Claude 5.5 | 79.64 ± 1.98 (high) | 80.50 ± 1.64 (low) | -0.85 |
| meta/muse-wonder | Meta Muse | 80.46 ± 1.42 (high) | 77.50 ± 3.62 (minimal) | +2.96 |
| anthropic/claude-fable-5.1 | Claude Fable | 80.46 ± 2.15 (high) | 78.24 ± 1.67 (low) | +2.22 |
| x-ai/grok-4.7 | Grok | 80.08 ± 2.25 (high) | 77.02 ± 2.31 (low) | +3.05 |
| anthropic/claude-opus-5.5 | Claude 5.5 | 79.59 ± 1.31 (high) | 79.73 ± 1.44 (low) | -0.14 |
| anthropic/claude-haiku-5.5 | Claude 5.5 | 79.70 ± 1.24 (high) | 78.17 ± 2.06 (none) | +1.53 |
| openai/gpt-5.6-luna | GPT-5.6 | 77.68 ± 1.45 (high) | 79.61 ± 0.97 (none) | -1.93 |
| anthropic/claude-fable-5 | Claude Fable | 79.45 ± 2.28 (high) | 79.53 ± 2.49 (low) | -0.08 |
| meta/muse-glimmer-30b | Meta Muse | 78.35 ± 1.85 (high) | 74.82 ± 1.71 (low) | +3.53 |
| openai/gpt-5.6-terra | GPT-5.6 | 78.07 ± 2.08 (high) | 75.57 ± 1.60 (none) | +2.50 |
| meta/muse-spark-open | Meta Muse | 78.05 ± 2.32 (high) | 77.92 ± 1.66 (low) | +0.13 |
| google/gemini-3.5-flash-lite | Gemini Flash | 77.86 ± 2.64 (high) | 77.72 ± 2.65 (minimal) | +0.14 |
| x-ai/grok-4.6 | Grok | 77.02 ± 2.13 (high) | 77.79 ± 1.54 (low) | -0.77 |
| meta/muse-spark-1.3 | Meta Muse | 77.72 ± 2.20 (high) | 75.07 ± 3.77 (minimal) | +2.65 |
| x-ai/grok-4.20-multi-agent | Grok | 77.18 ± 1.37 (high) | 76.68 ± 2.23 (low) | +0.50 |
| x-ai/grok-4.5 | Grok | 76.69 ± 1.34 (high) | 77.02 ± 2.38 (low) | -0.34 |
| x-ai/grok-4.20 | Grok | 76.74 ± 1.81 (enabled) | 75.58 ± 2.03 (disabled) | +1.16 |
| x-ai/grok-4.3 | Grok | 74.23 ± 2.14 (high) | 74.64 ± 3.52 (none) | -0.40 |

### thinking

| Rank | Model | Family | Effort | Mean | SD | Min | Max | Scored/N | Errors | Unscorable | Avg reasoning tokens |
|---:|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| 1 | openai/gpt-6-astra | GPT-6 | high | 82.00 | 1.50 | 79.83 | 84.28 | 10/10 | 0 | 0 | 1127 |
| 2 | google/gemini-3.8-flash | Gemini Flash | high | 81.53 | 1.89 | 78.67 | 85.47 | 10/10 | 0 | 0 | 8087 |
| 3 | google/gemini-3.5-flash | Gemini Flash | high | 81.28 | 1.46 | 78.56 | 83.36 | 10/10 | 0 | 0 | 2518 |
| 4 | google/gemini-3.7-flash | Gemini Flash | high | 81.14 | 1.87 | 77.78 | 83.75 | 10/10 | 0 | 0 | 2708 |
| 5 | openai/gpt-6.1-sol | GPT-6.1 | high | 80.74 | 1.94 | 77.51 | 83.61 | 10/10 | 0 | 0 | 1064 |
| 6 | google/gemini-3.6-flash | Gemini Flash | high | 80.74 | 2.26 | 76.92 | 85.33 | 10/10 | 0 | 0 | 2989 |
| 7 | openai/gpt-5.6-sol | GPT-5.6 | high | 80.56 | 1.95 | 77.63 | 83.98 | 10/10 | 0 | 0 | 924 |
| 8 | meta/muse-wonder | Meta Muse | high | 80.46 | 1.42 | 78.57 | 82.40 | 10/10 | 0 | 0 | 4179 |
| 9 | anthropic/claude-fable-5.1 | Claude Fable | high | 80.46 | 2.15 | 77.55 | 84.21 | 10/10 | 0 | 0 | 524 |
| 10 | x-ai/grok-4.7 | Grok | high | 80.08 | 2.25 | 74.62 | 83.16 | 10/10 | 0 | 0 | 3262 |
| 11 | anthropic/claude-haiku-5.5 | Claude 5.5 | high | 79.70 | 1.24 | 76.83 | 81.00 | 10/10 | 0 | 0 | 801 |
| 12 | anthropic/claude-sonnet-5.5 | Claude 5.5 | high | 79.64 | 1.98 | 76.00 | 83.15 | 10/10 | 0 | 0 | 124 |
| 13 | anthropic/claude-opus-5.5 | Claude 5.5 | high | 79.59 | 1.31 | 78.07 | 81.98 | 10/10 | 0 | 0 | 354 |
| 14 | anthropic/claude-fable-5 | Claude Fable | high | 79.45 | 2.28 | 76.62 | 82.80 | 10/10 | 0 | 0 | 256 |
| 15 | meta/muse-glimmer-30b | Meta Muse | high | 78.35 | 1.85 | 74.36 | 80.98 | 10/10 | 0 | 0 | 2150 |
| 16 | openai/gpt-5.6-terra | GPT-5.6 | high | 78.07 | 2.08 | 74.88 | 80.40 | 10/10 | 0 | 0 | 500 |
| 17 | meta/muse-spark-open | Meta Muse | high | 78.05 | 2.32 | 73.65 | 80.55 | 10/10 | 0 | 0 | 2103 |
| 18 | google/gemini-3.5-flash-lite | Gemini Flash | high | 77.86 | 2.64 | 73.19 | 81.03 | 10/10 | 0 | 0 | 1311 |
| 19 | meta/muse-spark-1.3 | Meta Muse | high | 77.72 | 2.20 | 75.60 | 81.52 | 9/10 | 0 | 1 | 6155 |
| 20 | openai/gpt-5.6-luna | GPT-5.6 | high | 77.68 | 1.45 | 75.24 | 79.70 | 10/10 | 0 | 0 | 908 |
| 21 | x-ai/grok-4.20-multi-agent | Grok | high | 77.18 | 1.37 | 74.18 | 78.27 | 10/10 | 0 | 0 | 49833 |
| 22 | x-ai/grok-4.6 | Grok | high | 77.02 | 2.13 | 74.24 | 81.61 | 10/10 | 0 | 0 | 3044 |
| 23 | x-ai/grok-4.20 | Grok | enabled | 76.74 | 1.81 | 74.57 | 79.73 | 10/10 | 0 | 0 | 2367 |
| 24 | x-ai/grok-4.5 | Grok | high | 76.69 | 1.34 | 73.77 | 78.68 | 10/10 | 0 | 0 | 1069 |
| 25 | x-ai/grok-4.3 | Grok | high | 74.23 | 2.14 | 70.69 | 76.92 | 10/10 | 0 | 0 | 1143 |

### no-thinking

| Rank | Model | Family | Effort | Mean | SD | Min | Max | Scored/N | Errors | Unscorable | Avg reasoning tokens |
|---:|---|---|---|---:|---:|---:|---:|---:|---:|---:|---:|
| 1 | google/gemini-3.8-flash | Gemini Flash | low | 82.82 | 2.62 | 79.65 | 87.49 | 10/10 | 0 | 0 | 474 |
| 2 | openai/gpt-6.1-sol | GPT-6.1 | low | 81.61 | 0.89 | 80.35 | 83.06 | 10/10 | 0 | 0 | 189 |
| 3 | openai/gpt-6-astra | GPT-6 | low | 81.60 | 1.74 | 79.23 | 84.68 | 10/10 | 0 | 0 | 181 |
| 4 | google/gemini-3.7-flash | Gemini Flash | low | 80.57 | 1.68 | 78.02 | 83.52 | 10/10 | 0 | 0 | 479 |
| 5 | anthropic/claude-sonnet-5.5 | Claude 5.5 | low | 80.50 | 1.64 | 77.73 | 83.21 | 10/10 | 0 | 0 | 0 |
| 6 | anthropic/claude-opus-5.5 | Claude 5.5 | low | 79.73 | 1.44 | 77.72 | 82.06 | 10/10 | 0 | 0 | 0 |
| 7 | openai/gpt-5.6-luna | GPT-5.6 | none | 79.61 | 0.97 | 77.83 | 80.74 | 9/10 | 0 | 1 | 0 |
| 8 | anthropic/claude-fable-5 | Claude Fable | low | 79.53 | 2.49 | 76.37 | 83.67 | 10/10 | 0 | 0 | 10 |
| 9 | google/gemini-3.6-flash | Gemini Flash | minimal | 78.27 | 1.60 | 75.83 | 81.27 | 10/10 | 0 | 0 | 0 |
| 10 | anthropic/claude-fable-5.1 | Claude Fable | low | 78.24 | 1.67 | 75.93 | 80.96 | 10/10 | 0 | 0 | 0 |
| 11 | anthropic/claude-haiku-5.5 | Claude 5.5 | none | 78.17 | 2.06 | 75.16 | 81.08 | 10/10 | 0 | 0 | 0 |
| 12 | openai/gpt-5.6-sol | GPT-5.6 | none | 78.00 | 2.32 | 74.88 | 81.21 | 10/10 | 0 | 0 | 0 |
| 13 | meta/muse-spark-open | Meta Muse | low | 77.92 | 1.66 | 75.21 | 81.04 | 10/10 | 0 | 0 | 383 |
| 14 | x-ai/grok-4.6 | Grok | low | 77.79 | 1.54 | 74.57 | 79.94 | 10/10 | 0 | 0 | 811 |
| 15 | google/gemini-3.5-flash-lite | Gemini Flash | minimal | 77.72 | 2.65 | 73.84 | 82.41 | 10/10 | 0 | 0 | 0 |
| 16 | meta/muse-wonder | Meta Muse | minimal | 77.50 | 3.62 | 70.70 | 81.87 | 10/10 | 0 | 0 | 237 |
| 17 | x-ai/grok-4.5 | Grok | low | 77.02 | 2.38 | 73.64 | 81.07 | 10/10 | 0 | 0 | 704 |
| 18 | x-ai/grok-4.7 | Grok | low | 77.02 | 2.31 | 73.56 | 81.33 | 10/10 | 0 | 0 | 753 |
| 19 | x-ai/grok-4.20-multi-agent | Grok | low | 76.68 | 2.23 | 73.56 | 79.82 | 10/10 | 0 | 0 | 12316 |
| 20 | google/gemini-3.5-flash | Gemini Flash | minimal | 76.56 | 2.57 | 71.71 | 80.13 | 10/10 | 0 | 0 | 0 |
| 21 | x-ai/grok-4.20 | Grok | disabled | 75.58 | 2.03 | 71.53 | 78.71 | 10/10 | 0 | 0 | 0 |
| 22 | openai/gpt-5.6-terra | GPT-5.6 | none | 75.57 | 1.60 | 72.94 | 78.75 | 10/10 | 0 | 0 | 0 |
| 23 | meta/muse-spark-1.3 | Meta Muse | minimal | 75.07 | 3.77 | 68.61 | 80.67 | 10/10 | 0 | 0 | 279 |
| 24 | meta/muse-glimmer-30b | Meta Muse | low | 74.82 | 1.71 | 72.37 | 77.68 | 10/10 | 0 | 0 | 290 |
| 25 | x-ai/grok-4.3 | Grok | none | 74.64 | 3.52 | 69.51 | 79.96 | 10/10 | 0 | 0 | 0 |

## Appendix: words per trial

### English · openai/gpt-5.6-luna · thinking (effort=high)

- run 2: **80.54** — freedom, elephant, spoon, thunder, friendship, mountain, calendar
- run 3: **81.25** — freedom, elephant, spoon, thunder, childhood, glacier, melody
- run 4: **77.49** — freedom, elephant, rainbow, spoon, mushroom, justice, whisper
- run 7: **75.73** — happiness, elephant, mountain, justice, spoon, thunder, child
- run 5: **81.51** — freedom, elephant, needle, thunder, memory, garden, justice
- run 6: **80.36** — elephant, democracy, volcano, happiness, needle, whisper, ocean

### English · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 0: **81.29** — elephant, democracy, spoon, thunder, friendship, mountain, pencil
- run 1: **85.07** — thunder, democracy, spoon, cactus, memory, whale, pillow
- run 2: **82.45** — teapot, democracy, cactus, whisper, mountain, jealousy, pigeon
- run 3: **79.13** — mountain, justice, lantern, whisper, elephant, democracy, pillow
- run 4: **85.22** — volcano, democracy, butterfly, suitcase, happiness, needle, orchestra
- run 5: **78.06** — river, justice, lantern, butterfly, silence, hammer, dream
- run 6: **80.47** — mountain, democracy, spoon, whisper, elephant, kindness, volcano

### English · openai/gpt-5.6-luna · thinking (effort=high)

- run 1: **85.36** — elephant, hammer, violin, potato, ocean, calendar, democracy

### English · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 7: **84.78** — umbrella, justice, volcano, whisper, democracy, caterpillar, nostalgia
- run 9: **83.14** — mountain, whisper, democracy, lantern, elephant, jealousy, pancake

### English · openai/gpt-5.6-luna · thinking (effort=high)

- run 0: **84.12** — umbrella, penguin, hurricane, poetry, jealousy, chocolate, village
- run 9: **80.36** — elephant, bread, mirror, thunder, democracy, friendship, hunger
- run 8: **81.84** — democracy, elephant, volcano, symphony, hunger, needle, ocean

### English · openai/gpt-5.6-sol · thinking (effort=high)

- run 4: **83.32** — justice, banana, thunder, carpet, childhood, engine, ocean
- run 0: **81.51** — penguin, justice, thunder, velvet, century, village, hunger
- run 3: **84.94** — justice, penguin, hammer, volcano, silk, laughter, teacher
- run 5: **87.60** — lantern, sorrow, democracy, glacier, melody, infancy, vinegar
- run 1: **86.88** — democracy, octopus, avalanche, lullaby, hammer, jealousy, bread
- run 2: **86.39** — justice, cabbage, lullaby, volcano, hammer, dolphin, velvet
- run 6: **86.26** — justice, volcano, melody, childhood, furniture, galaxy, treaty
- run 7: **81.00** — mountain, justice, violin, insect, ocean, childhood, brick

### English · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 0: **82.25** — apple, justice, thunder, child, silk, memory, island
- run 1: **86.32** — freedom, walrus, lantern, jealousy, canyon, rhythm, blanket
- run 2: **84.82** — happiness, volcano, needle, democracy, whale, music, cabbage
- run 3: **78.23** — justice, banana, thunder, child, velvet, island, laughter
- run 4: **81.95** — lantern, justice, whale, laughter, desert, needle, kingdom
- run 5: **82.12** — justice, volcano, lullaby, spoon, childhood, penguin, thunder

### English · openai/gpt-5.6-sol · thinking (effort=high)

- run 8: **77.72** — mountain, mercy, teapot, wolf, thunder, music, bread

### English · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 6: **84.17** — justice, volcano, lullaby, spoon, insect, winter, laughter
- run 7: **79.91** — lantern, justice, whale, melody, desert, child, thunder
- run 8: **79.91** — lantern, justice, whale, melody, desert, child, thunder

### English · openai/gpt-5.6-terra · thinking (effort=high)

- run 0: **79.21** — mountain, whisper, justice, spoon, child, rain, memory
- run 1: **74.99** — stone, laughter, river, shadow, justice, infant, forest

### English · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 9: **86.97** — justice, volcano, lullaby, spoon, jealousy, otter, galaxy

### English · openai/gpt-5.6-sol · thinking (effort=high)

- run 9: **80.28** — apple, freedom, mountain, teacher, music, insect, pillow

### English · openai/gpt-5.6-terra · thinking (effort=high)

- run 3: **79.03** — mountain, whisper, justice, spoon, infant, thunder, garden
- run 6: **74.06** — rock, idea, child, river, music, shadow, bread
- run 5: **75.58** — mountain, whisper, justice, spoon, child, rain, shadow
- run 8: **80.82** — mountain, whisper, justice, spoon, infant, thunder, memory
- run 7: **74.24** — rock, joy, shadow, river, bread, thunder, child
- run 2: **77.48** — basket, freedom, tooth, thunder, honey, mountain, laughter
- run 9: **76.74** — river, thought, butterfly, hammer, justice, laughter, mountain

### English · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 2: **81.59** — mountain, whisper, justice, spoon, puppy, rain, memory
- run 0: **78.11** — mountain, whisper, justice, spoon, butterfly, childhood, thunder

### English · openai/gpt-5.6-terra · thinking (effort=high)

- run 4: **80.94** — pebble, laughter, justice, forest, hunger, clock, blanket

### English · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 3: **79.21** — mountain, whisper, justice, spoon, child, rain, memory
- run 4: **71.44** — stone, joy, river, whisper, garden, shadow, bread
- run 1: **81.59** — mountain, whisper, justice, spoon, puppy, rain, memory
- run 7: **75.58** — mountain, whisper, justice, spoon, rainbow, child, forest
- run 8: **81.54** — mountain, whisper, justice, spoon, kitten, rainbow, memory
- run 9: **78.11** — butterfly, mountain, whisper, justice, spoon, childhood, thunder
- run 6: **76.87** — mountain, whisper, bread, justice, kitten, shadow, river
- run 5: **77.16** — mountain, whisper, spoon, justice, kitten, rainbow, bread

### English · openai/gpt-6-astra · thinking (effort=high)

- run 5: **87.39** — glacier, jealousy, violin, democracy, cabbage, elbow, thunder
- run 3: **85.37** — democracy, jellyfish, velvet, thunder, hunger, galaxy, staircase
- run 1: **84.79** — justice, kangaroo, perfume, galaxy, hammer, hunger, velvet
- run 2: **85.34** — galaxy, velvet, justice, cactus, laughter, anchor, hunger
- run 4: **85.63** — democracy, mushroom, velvet, thunder, ankle, riddle, galaxy
- run 0: **84.66** — justice, cactus, lullaby, galaxy, velvet, hunger, suitcase
- run 6: **87.71** — democracy, cabbage, galaxy, pillow, melody, drought, jealousy

### English · openai/gpt-6-astra · no-thinking (effort=low)

- run 0: **84.98** — justice, cactus, thunder, velvet, infancy, galaxy, spoon

### English · openai/gpt-6-astra · thinking (effort=high)

- run 8: **86.48** — justice, cactus, velvet, galaxy, hunger, melody, teaspoon

### English · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 8: **82.17** — elephant, democracy, spoon, thunder, friendship, volcano, puzzle

### English · openai/gpt-6-astra · thinking (effort=high)

- run 9: **87.90** — democracy, jellyfish, volcano, lullaby, velvet, ankle, jealousy
- run 7: **88.86** — democracy, cinnamon, galaxy, grief, velvet, kangaroo, arithmetic

### English · openai/gpt-6-astra · no-thinking (effort=low)

- run 2: **88.75** — elephant, democracy, vinegar, galaxy, violin, nostalgia, ladder
- run 1: **83.76** — justice, galaxy, cabbage, violin, grief, velvet, thunder
- run 3: **89.75** — glacier, justice, teaspoon, laughter, fungus, galaxy, debt
- run 5: **85.98** — glacier, democracy, perfume, sorrow, violin, fungus, eternity
- run 7: **86.62** — elephant, democracy, perfume, glacier, violin, hunger, alphabet
- run 6: **90.23** — kangaroo, democracy, perfume, glacier, violin, jealousy, alphabet
- run 4: **83.18** — elephant, justice, kettle, thunder, velvet, boredom, galaxy
- run 8: **84.02** — elephant, democracy, perfume, glacier, laughter, needle, eternity
- run 9: **85.84** — elephant, democracy, perfume, glacier, violin, envy, soup

### English · openai/gpt-6.1-sol · thinking (effort=high)

- run 1: **85.81** — galaxy, cabbage, justice, velvet, ankle, whisper, nostalgia
- run 0: **86.78** — justice, cabbage, galaxy, laughter, hammer, perfume, octopus
- run 2: **85.62** — octopus, justice, galaxy, velvet, hunger, ladder, lullaby
- run 5: **89.84** — justice, galaxy, velvet, mushroom, elbow, arithmetic, lullaby
- run 4: **84.50** — justice, cactus, thunder, velvet, galaxy, laughter, engine
- run 3: **85.68** — justice, cabbage, galaxy, velvet, laughter, ankle, pyramid
- run 7: **86.57** — justice, cabbage, galaxy, velvet, laughter, hammer, drought
- run 9: **89.22** — freedom, mushroom, thunder, velvet, arithmetic, jealousy, satellite

### English · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 0: **85.27** — glacier, jealousy, spoon, thunder, freedom, mushroom, ankle
- run 1: **86.03** — justice, galaxy, cabbage, violin, hunger, velvet, staircase

### English · openai/gpt-6.1-sol · thinking (effort=high)

- run 6: **89.68** — justice, octopus, velvet, galaxy, arithmetic, perfume, hinge

### English · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 2: **84.09** — thunder, jealousy, spoon, freedom, mushroom, elbow, alphabet
- run 3: **87.47** — glacier, jealousy, violin, democracy, mushroom, elbow, thunder

### English · openai/gpt-6.1-sol · thinking (effort=high)

- run 8: **89.57** — galaxy, garlic, velvet, jealousy, passport, octopus, arithmetic

### English · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 5: **87.42** — glacier, justice, violin, mosquito, laughter, velvet, galaxy
- run 4: **80.87** — elephant, justice, thunder, spoon, galaxy, velvet, hunger
- run 6: **83.31** — anchor, butterfly, justice, velvet, thunder, hunger, galaxy
- run 9: **82.93** — elephant, democracy, thunder, spoon, jealousy, galaxy, velvet
- run 8: **88.48** — elephant, democracy, perfume, glacier, arithmetic, lullaby, cabbage
- run 7: **86.52** — elephant, democracy, velvet, galaxy, vinegar, laughter, avalanche

### English · google/gemini-3.8-flash · thinking (effort=high)

- run 1: **86.84** — jellyfish, irony, anvil, vinegar, nausea, canyon, melody
- run 6: **89.17** — octopus, anvil, jealousy, lava, comet, melody, citizenship
- run 5: **88.68** — jealousy, lightning, walrus, custard, bankruptcy, alphabet, century
- run 3: **88.70** — octopus, scissors, jealousy, century, odor, magma, geometry

### English · google/gemini-3.8-flash · no-thinking (effort=low)

- run 0: **90.81** — jealousy, glacier, stapler, octopus, cathedral, mustard, bankruptcy
- run 2: **84.83** — justice, whale, brick, dentist, swamp, century, elbow
- run 1: **90.99** — remorse, monsoon, squid, zipper, nepotism, honey, fortnight
- run 3: **84.57** — galaxy, octopus, grief, needle, democracy, tornado, shadow

### English · google/gemini-3.8-flash · thinking (effort=high)

- run 4: **85.40** — giraffe, canyon, zipper, grief, bleach, alphabet, lightning

### English · google/gemini-3.8-flash · no-thinking (effort=low)

- run 4: **92.75** — toad, glacier, galaxy, stapler, grief, treaty, perfume

### English · google/gemini-3.8-flash · thinking (effort=high)

- run 8: **89.17** — grief, mosquito, zipper, fraction, lava, shadow, citizenship

### English · google/gemini-3.8-flash · no-thinking (effort=low)

- run 6: **86.51** — whale, jealousy, granite, symphony, elbow, kerosene, ghost
- run 5: **84.98** — glacier, spider, nostalgia, custard, skull, anvil, monarchy
- run 8: **88.85** — elbow, glacier, democracy, cinnamon, ladder, grief, mosquito

### English · google/gemini-3.8-flash · thinking (effort=high)

- run 7: **90.24** — octopus, galaxy, grief, century, mustard, stapler, debt
- run 2: **86.06** — jealousy, planet, parliament, century, blizzard, giraffe, triangle

### English · google/gemini-3.8-flash · no-thinking (effort=low)

- run 9: **85.96** — jealousy, volcano, needle, spleen, eternity, cabbage, echo
- run 7: **86.16** — moss, needle, grief, volcano, debt, galaxy, vinegar

### English · google/gemini-3.7-flash · thinking (effort=high)

- run 3: **87.05** — grief, comet, mushroom, cathedral, fraction, vinegar, treason
- run 0: **88.63** — galaxy, jellyfish, needle, jealousy, kidney, melody, bankruptcy
- run 1: **86.20** — octopus, jealousy, glacier, statue, orchestra, gravity, custard
- run 2: **89.39** — octopus, galaxy, guilt, bankruptcy, flute, vinegar, dandruff
- run 5: **88.92** — mosquito, grief, glacier, zipper, alphabet, elbow, echo
- run 4: **87.89** — grief, octopus, galaxy, anvil, fraction, kidney, glacier

### English · google/gemini-3.7-flash · no-thinking (effort=low)

- run 0: **89.68** — moss, nostalgia, zipper, eclipse, kidney, echo, debt

### English · google/gemini-3.7-flash · thinking (effort=high)

- run 7: **87.99** — octopus, needle, glacier, echo, bankruptcy, custard, irony

### English · google/gemini-3.7-flash · no-thinking (effort=low)

- run 1: **88.35** — galaxy, itch, bankruptcy, swamp, eyelash, brick, nostalgia
- run 2: **86.29** — sorrow, zipper, thunder, yeast, orphan, mustard, comet

### English · google/gemini-3.7-flash · thinking (effort=high)

- run 9: **85.69** — jellyfish, galaxy, glacier, mirage, nostalgia, symmetry, auction
- run 8: **86.84** — octopus, jealousy, anvil, glacier, lightning, debt, honey

### English · google/gemini-3.7-flash · no-thinking (effort=low)

- run 3: **83.81** — whale, needle, justice, moss, vinegar, canyon, elbow
- run 4: **91.88** — nostalgia, padlock, jellyfish, glacier, cinnamon, treaty, spleen

### English · google/gemini-3.7-flash · thinking (effort=high)

- run 6: **88.19** — volcano, nostalgia, needle, jellyfish, treason, echo, marrow

### English · google/gemini-3.7-flash · no-thinking (effort=low)

- run 5: **82.27** — shark, justice, mud, spoon, galaxy, laughter, garlic
- run 8: **84.62** — courage, comet, yeast, chimney, laughter, apron, wax

### English · google/gemini-3.8-flash · thinking (effort=high)

- run 0: **93.33** — galaxy, iguana, zipper, grief, century, salary, vinegar

### English · google/gemini-3.7-flash · no-thinking (effort=low)

- run 7: **88.61** — whale, envy, lava, trombone, midnight, garlic, ankle
- run 6: **87.23** — anvil, moss, irony, avalanche, liver, eclipse, mustard
- run 9: **88.25** — nostalgia, volcano, squid, anvil, democracy, vinegar, parade

### English · google/gemini-3.6-flash · thinking (effort=high)

- run 0: **84.08** — glacier, grief, anvil, fungus, carnival, lightning, skeleton
- run 5: **88.38** — grief, volcano, mosquito, treason, lantern, echo, cinnamon
- run 3: **86.36** — nostalgia, granite, whale, spoon, lightning, democracy, saliva
- run 6: **83.92** — whale, grief, needle, volcano, century, bankruptcy, rainbow
- run 2: **88.64** — octopus, spleen, anvil, vinegar, lava, grief, statute

### English · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 1: **90.46** — galaxy, nostalgia, spoon, bacterium, justice, avalanche, syllable
- run 0: **89.33** — hyperbole, microbe, justice, volcano, tango, nostalgia, pawn

### English · google/gemini-3.6-flash · thinking (effort=high)

- run 4: **84.14** — grief, anvil, octopus, glacier, velvet, abyss, decade

### English · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 2: **86.29** — galaxy, nostalgia, spoon, courage, fungus, silence, verdict
- run 3: **79.60** — grief, molecule, justice, banana, thunder, shadow, ghost
- run 5: **85.54** — grief, volcano, justice, syllable, sponge, gossip, horizon
- run 4: **84.90** — thunder, grief, triangle, oxygen, ancestor, velvet, whistle
- run 6: **91.80** — sympathy, glacier, tariff, nostril, algorithm, sieve, origami
- run 7: **86.67** — justice, mushroom, thunder, nostalgia, whisk, glacier, syllable
- run 8: **88.59** — thunder, dignity, elbow, algorithm, soup, nostalgia, pawn
- run 9: **94.43** — grief, elbow, algorithm, velvet, archipelago, syllogism, leukocyte

### English · google/gemini-3.6-flash · thinking (effort=high)

- run 1: **86.36** — alphabet, bicycle, curiosity, drought, epoch, fungus, gravity
- run 8: **90.11** — galaxy, nostalgia, thimble, octopus, spleen, bankruptcy, earthquake
- run 9: **87.81** — nostalgia, volcano, octopus, ladder, debt, galaxy, symmetry
- run 7: **86.31** — octopus, grief, symmetry, porcelain, democracy, lightning, kidney

### English · google/gemini-3.5-flash · thinking (effort=high)

- run 1: **80.41** — nostalgia, anvil, mushroom, thief, lightning, honey, shadow
- run 2: **86.10** — justice, volcano, octopus, gravity, needle, debt, nostalgia
- run 0: **89.22** — justice, galaxy, lava, fungus, zipper, echo, kidney
- run 3: **86.33** — justice, volcano, whale, umbrella, gravity, yeast, alphabet
- run 4: **82.94** — grief, mud, star, octopus, spoon, shadow, law

### English · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 1: **90.76** — grief, silicon, caterpillar, galaxy, lawsuit, melody, ointment
- run 0: **89.27** — grief, granite, concept, gasp, microbe, hovercraft, glint

### English · google/gemini-3.5-flash · thinking (effort=high)

- run 7: **80.08** — infinity, moss, grief, spoon, sand, feather, century

### English · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 3: **87.73** — justice, granite, symphony, microbe, nostalgia, oxygen, glacier
- run 2: **87.82** — nostalgia, iron, cloud, gossip, elbow, republic, kelp

### English · google/gemini-3.5-flash · thinking (effort=high)

- run 5: **83.33** — nostalgia, galaxy, gravity, umbrella, oak, decade, granite

### English · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 4: **86.27** — gravity, granite, grief, gargoyle, glimmer, gut, gossip
- run 5: **85.12** — thunder, grief, granite, concept, saliva, parachute, whisper
- run 6: **87.05** — grief, granite, concept, gasp, microbe, velocity, puddle
- run 7: **86.14** — thunder, syllable, shampoo, gravity, tissue, greed, volcano
- run 8: **89.06** — grief, granite, velocity, enzyme, plug, republic, puddle
- run 9: **91.16** — justice, granite, whisper, grommet, nostalgia, amoeba, galaxy

### English · google/gemini-3.5-flash · thinking (effort=high)

- run 6: **84.73** — gravity, octopus, clay, ladder, fern, tornado, lullaby
- run 9: **83.35** — gossip, glacier, octopus, paradox, umbrella, twilight, soup
- run 8: **84.02** — nostalgia, granite, octopus, gravity, triangle, cabinet, symphony

### English · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 0: **81.64** — octopus, cathedral, infinity, volcano, turnip, echo, labyrinth
- run 2: **79.67** — grief, brick, volcano, democracy, fungus, shadow, horizon
- run 1: **83.68** — grief, glacier, thimble, jellyfish, shadow, melody, egg
- run 5: **91.51** — grief, magma, tariff, thimble, fungus, vacuum, rhythm
- run 3: **87.06** — grief, worm, canyon, monarchy, echo, justice, thimble
- run 4: **87.69** — serendipity, avalanche, tooth, thimble, persimmon, midnight, melody

### English · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 0: **78.21** — whisper, brick, democracy, shadow, horizon, spoon, melody

### English · google/gemini-3.8-flash · thinking (effort=high)

- run 9: **94.12** — asteroid, fern, nostalgia, fraction, curfew, forehead, stapler

### English · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 2: **81.89** — whisper, mountain, democracy, spoon, gravity, melancholy, umbrella
- run 1: **78.72** — whisper, brick, democracy, shadow, spoon, horizon, rhythm
- run 3: **83.58** — whisper, anchor, melancholy, turnip, horizon, velocity, velvet
- run 5: **75.15** — whisper, brick, thought, ocean, feather, clock, shadow
- run 4: **80.36** — whisper, brick, melody, horizon, sponge, justice, feather
- run 6: **82.88** — melody, whisper, glacier, anchor, pebble, thistle, velvet

### English · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 6: **81.16** — justice, swamp, pencil, spider, shadow, lung, soup
- run 8: **82.79** — despair, blizzard, lever, root, eternity, threshold, salt

### English · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 7: **80.89** — whisper, brick, democracy, shadow, spoon, horizon, entropy
- run 8: **79.78** — whisper, brick, democracy, shadow, spoon, horizon, memory
- run 9: **84.88** — soup, mountain, whisper, democracy, cactus, tragedy, magnet

### English · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 7: **82.99** — hunger, lightning, octopus, anvil, song, thought, clay
- run 9: **83.30** — galaxy, sorrow, basket, night, cotton, memory, turtle

### English · x-ai/grok-4.20 · thinking (effort=enabled)

- run 1: **72.19** — joy, mountain, butterfly, justice, melody, thunder, wisdom
- run 0: **75.00** — freedom, mountain, elephant, justice, bicycle, thunder, melody
- run 2: **73.73** — love, mountain, justice, bicycle, war, book, elephant

### English · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 0: **75.02** — apple, justice, whisper, mountain, freedom, echo, shadow
- run 1: **75.12** — apple, freedom, river, laughter, mountain, justice, whisper
- run 2: **73.94** — tree, freedom, echo, river, laughter, shadow, courage

### English · x-ai/grok-4.20 · thinking (effort=enabled)

- run 4: **71.55** — hope, stone, laughter, forest, wisdom, sword, rainbow

### English · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 3: **74.85** — apple, freedom, whisper, gravity, mountain, laughter, ocean
- run 4: **69.96** — tree, love, ocean, justice, fire, silence, mountain
- run 5: **70.92** — tree, freedom, whisper, mountain, laughter, ocean, shadow
- run 6: **74.02** — apple, justice, ocean, whisper, mountain, freedom, shadow
- run 7: **74.71** — apple, freedom, ocean, laughter, mountain, justice, whisper

### English · x-ai/grok-4.20 · thinking (effort=enabled)

- run 5: **79.90** — justice, volcano, butterfly, melody, eternity, dragon, poverty

### English · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 8: **74.71** — apple, justice, mountain, whisper, freedom, ocean, laughter
- run 9: **72.05** — love, mountain, whisper, freedom, apple, shadow, laughter

### English · x-ai/grok-4.20 · thinking (effort=enabled)

- run 3: **75.67** — happiness, justice, mountain, elephant, piano, chocolate, silence
- run 8: **74.60** — love, mountain, justice, bicycle, wisdom, thunder, diamond
- run 7: **75.62** — courage, ocean, butterfly, eternity, whisper, castle, chocolate
- run 9: **78.75** — happiness, elephant, volcano, bicycle, justice, soup, courage
- run 6: **72.34** — friendship, ocean, elephant, castle, laughter, fire, dream

### English · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 0: **73.52** — love, mountain, elephant, bread, freedom, thunder, castle
- run 2: **75.48** — joy, elephant, ocean, justice, bicycle, thunder, poetry
- run 1: **74.61** — love, elephant, mountain, bicycle, justice, thunder, diamond
- run 3: **79.35** — joy, elephant, justice, bicycle, thunder, book, galaxy
- run 4: **77.12** — happiness, elephant, justice, banana, hammer, castle, eternity
- run 6: **76.03** — elephant, mountain, justice, bicycle, chocolate, hammer, ocean
- run 5: **74.08** — joy, elephant, mountain, justice, bicycle, star, fire
- run 7: **78.99** — joy, octopus, cactus, diamond, castle, sword, poem
- run 9: **75.52** — hope, hammer, eagle, pizza, castle, laughter, thunder

### English · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 2: **76.95** — elephant, ocean, galaxy, bicycle, pyramid, thunder, chocolate
- run 1: **73.27** — joy, mountain, hammer, butterfly, thunder, diamond, melody
- run 0: **75.43** — apple, elephant, mountain, thunder, galaxy, bicycle, guitar

### English · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 8: **72.17** — love, mountain, elephant, justice, ocean, thunder, bicycle

### English · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 3: **81.11** — rainbow, elephant, bicycle, courage, volcano, book, galaxy
- run 4: **79.95** — joy, hammer, octopus, volcano, bread, eternity, justice
- run 5: **73.10** — joy, elephant, ocean, justice, fire, wisdom, bicycle
- run 6: **74.45** — love, table, elephant, ocean, mountain, justice, bicycle
- run 7: **79.43** — elephant, justice, volcano, bicycle, chocolate, eternity, shadow

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 1: **70.16** — mountain, freedom, bicycle, justice, elephant, happiness, river

### English · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 9: **71.15** — love, elephant, mountain, justice, time, fire, music

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 0: **74.59** — mountain, happiness, bicycle, justice, ocean, pencil, freedom

### English · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 8: **83.82** — joy, canyon, octopus, hammer, truth, blizzard, violin

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 2: **75.63** — happiness, mountain, justice, sandwich, shadow, freedom, pencil

### English · x-ai/grok-4.3 · no-thinking (effort=none)

- run 0: **75.91** — apple, river, justice, mountain, whisper, candle, shadow
- run 1: **67.57** — mountain, whisper, friendship, horizon, shadow, echo, silence
- run 2: **67.32** — river, mountain, silence, shadow, freedom, laughter, horizon
- run 3: **73.09** — apple, river, justice, cloud, mountain, silence, freedom
- run 4: **68.93** — river, mountain, dream, silence, freedom, echo, shadow

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 3: **72.68** — mountain, freedom, elephant, harmony, river, justice, chair

### English · x-ai/grok-4.3 · no-thinking (effort=none)

- run 5: **75.94** — apple, ocean, justice, mountain, friendship, bicycle, shadow

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 8: **75.89** — mountain, happiness, bicycle, justice, cloud, courage, table

### English · x-ai/grok-4.3 · no-thinking (effort=none)

- run 6: **68.90** — mountain, justice, whisper, shadow, thunder, dream, river
- run 7: **76.10** — apple, justice, mountain, whisper, bicycle, freedom, thunder
- run 8: **75.67** — apple, river, justice, mountain, whisper, blanket, horizon
- run 9: **73.04** — cloud, whisper, anchor, freedom, echo, shadow, horizon

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 4: **76.80** — elephant, freedom, ocean, chair, mountain, pencil, justice
- run 9: **73.01** — happiness, elephant, justice, mountain, dream, pencil, courage
- run 6: **77.03** — elephant, happiness, freedom, bread, chair, music, shadow

### English · x-ai/grok-4.5 · thinking (effort=high)

- run 2: **77.52** — brick, truth, tiger, anger, cheese, lightning, finger
- run 1: **75.59** — tree, book, happiness, river, knife, dream, dust

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 7: **75.29** — happiness, mountain, justice, bicycle, cloud, courage, diamond

### English · x-ai/grok-4.5 · thinking (effort=high)

- run 3: **73.59** — cat, idea, stone, happiness, ocean, guitar, freedom
- run 0: **69.54** — dog, rock, love, table, sky, bread, idea

### English · x-ai/grok-4.3 · thinking (effort=high)

- run 5: **74.63** — mountain, happiness, freedom, chair, apple, justice, cloud

### English · x-ai/grok-4.5 · thinking (effort=high)

- run 4: **72.94** — dog, freedom, table, rain, joy, bread, mountain

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 0: **78.19** — bread, anger, knife, tiger, city, democracy, finger

### English · x-ai/grok-4.5 · thinking (effort=high)

- run 9: **78.47** — elephant, democracy, hammer, cloud, justice, bread, mountain
- run 5: **81.00** — elephant, freedom, scissors, ocean, laughter, brick, music
- run 8: **76.64** — elephant, freedom, mountain, laughter, bread, river, guitar

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 2: **74.05** — tree, justice, hammer, joy, ocean, idea, dog
- run 1: **85.50** — banana, justice, volcano, laughter, screwdriver, octopus, infinity
- run 3: **74.83** — apple, freedom, mountain, hammer, love, ocean, dog

### English · x-ai/grok-4.5 · thinking (effort=high)

- run 7: **69.24** — dog, love, mountain, book, fire, idea, ocean

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 7: **83.92** — banana, liberty, volcano, violin, philosophy, tiger, equality
- run 4: **75.16** — apple, freedom, mountain, guitar, joy, ocean, book

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 2: **83.94** — sadness, submarine, bamboo, marble, liberty, earthquake, bread

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 0: **84.31** — volcano, symphony, freedom, sandwich, penguin, library, thunderstorm

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 0: **76.48** — freedom, thunder, bread, ocean, child, music, stone

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 0: **72.25** — light, silence, hunger, stone, music, journey, mother

### English · meta/muse-wonder · thinking (effort=high)

- run 0: **78.39** — elephant, bread, hammer, grief, moon, desert, democracy

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 0: **77.25** — 空, 理念, 痛み, 旋律, 砂利, 鏡, 影

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 0: **79.08** — 未来, 岩石, 感情, 整数, 沈黙, 毛皮, 影

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 0: **77.18** — 愛, 影, 時間, 夢, 塩, 記憶, 沈黙

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 0: **75.42** — 重力, 感情, 液体, 憲法, 微生物, 冗談, 金属

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 0: **75.41** — mountain, childhood, silence, fire, ocean, justice, paper

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 0: **88.60** — volcano, whisper, spreadsheet, grandmother, umbrella, justice, cactus

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 0: **79.96** — 愛, 銀河, 微生物, 香水, 地震, 記憶, 砂漠

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 0: **83.96** — volcano, whisper, spoon, justice, elephant, nostalgia, bicycle

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 0: **87.46** — thunder, blanket, democracy, glacier, violin, mustard, orbit

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 0: **74.88** — 石, 夢, 匂い, 約束, 海, 沈黙, 祭り

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 0: **72.94** — 愛, 沈黙, 影, 匂い, 砂, 雷, 夢

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 0: **79.51** — 嫉妬, 火山, 醤油, 鉛筆, 幽霊, 河馬, 裁判

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 0: **77.52** — 石, 幸福, 約束, 雨, 犬, 歌, 椅子

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 0: **81.52** — 未来, 嫉妬, 泡, 洞窟, 蝶, 松, 味噌

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 0: **72.74** — 愛, 石, 時間, 夢, 音, 風, 影

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 0: **74.52** — ocean, freedom, thunder, mirror, garden, music, shadow

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 0: **73.02** — 雲, 針, 影, 夢, 鍵, 風, 声

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 0: **87.72** — volcano, spoon, jealousy, giraffe, algebra, sneeze, cathedral

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 0: **78.71** — 宇宙, 針, 契約, 旋律, 匂い, 憎悪, 墓

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 0: **84.78** — cloud, justice, spoon, bacteria, melody, volcano, grief

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 0: **85.85** — volcano, spoon, justice, elephant, whisper, carpet, bacteria

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 0: **69.02** — mountain, love, child, river, dream, stone, bread

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 0: **79.00** — 嫉妬, 稲妻, 法律, キノコ, 永遠, 歯車, 隙間

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 0: **82.62** — cloud, justice, spoon, elephant, whisper, volcano, nostalgia

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 0: **83.83** — freedom, bread, elbow, walrus, sorrow, hammer, century

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 0: **75.16** — 風, 愛, 鉄, 音, 数, 影, 鍵

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 0: **79.47** — 祖母, 霧, 怒り, 正義, 靴, 音, 休日

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 0: **76.21** — 山, 海, 涙, 椅子, 風, 卵, 鏡

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 0: **86.33** — volcano, spoon, jealousy, giraffe, algebra, sneeze, velvet

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 0: **81.38** — 鯨, 郵便, 正義, 砂糖, 雷, 膝, 退屈

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 0: **77.73** — 鯨, 正義, ピアノ, 雨, 祖母, 砂糖, 夢

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 0: **76.24** — 時間, 感情, 石鹸, 国家, 海藻, リズム, 借金

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 0: **78.90** — 火山, 靴下, 約束, 蜂蜜, 沈黙, 算数, 祖母

### English · meta/muse-wonder · thinking (effort=high)

- run 1: **88.68** — volcano, democracy, ankle, jellyfish, echo, grief, quilt

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 1: **72.64** — sand, freedom, elephant, music, mountain, ocean, anger

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 1: **74.53** — sorrow, bread, wolf, bridge, ocean, shadow, eternity

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 1: **72.92** — thunder, silence, mother, iron, hunger, journey, shadow

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 0: **76.60** — 石, 愛, 時間, 猫, 音楽, 鉛筆, 海

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 0: **83.22** — volcano, spoon, democracy, whisper, cactus, velvet, pirate

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 0: **82.80** — 石, 犬, 自由, 悲しみ, 昨日, 雷, 指

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 0: **80.30** — 山, 怒り, 靴, 昨日, 犬, 塩, 歌

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 0: **74.93** — stone, laughter, ocean, mother, winter, apple, shadow

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 0: **81.24** — 回答, 石, 愛, 雷, 魚, 約束, 椅子

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 0: **78.49** — 宇宙, 憎悪, 旋律, 砂漠, 法律, 墓, 針

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 1: **73.89** — mountain, river, child, silence, paper, light, dream

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 0: **73.64** — 猫, 山, 愛, 机, 風, 夢, 石

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 0: **81.17** — 虹, 良心, 砂糖, 靴下, 噂, 税金, 眠り

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 0: **81.23** — 虹, 税金, くしゃみ, 歯車, 郷愁, 豆腐, 子孫

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 0: **76.83** — 海, 時間, 味噌, 恋, 靴, 背中, 祭り

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 0: **78.07** — 火山, 嫉妬, 靴下, 蜂蜜, 葬式, 歯車, 冗談

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 1: **87.13** — volcano, spoon, justice, whisper, basement, giraffe, nostalgia

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 1: **83.69** — volcano, whisper, spoon, justice, basement, cloud, banana

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 1: **75.97** — ocean, thunder, childhood, freedom, stone, mirror, music

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 1: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 0: **85.77** — volcano, scissors, happiness, octopus, clock, justice, thunder

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 1: **84.79** — thunder, blanket, democracy, glacier, harmony, pepper, satellite

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 1: **79.81** — 愛情, 彗星, 鋏, 沈黙, 珊瑚, 蜃気楼, 糞

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 1: **84.81** — cloud, justice, spoon, bacteria, symphony, volcano, grief

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 1: **84.68** — justice, mushroom, hammer, galaxy, whisper, elbow, tornado

### English · meta/muse-spark-open · thinking (effort=high)

- run 1: **81.57** — galaxy, hammer, ocean, laughter, feather, justice, volcano

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 1: **86.76** — volcano, spoon, justice, whisper, giraffe, algebra, pillow

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 1: **78.75** — 沈黙, 渇き, 祖母, 雲, 約束, 穴, 笑顔

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 1: **75.20** — 石, 夢, 祭り, 孤独, 匂い, 約束, 海

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 0: **80.57** — 石, 自由, 魚, 針, 昨日, 怒り, 雲

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 0: **78.34** — 愛, 石, 時間, 猫, 音楽, 水, 戦争

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 0: **77.78** — 宇宙, 鯨, 味噌, 砂漠, 嫉妬, 貨幣, 音楽

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 1: **72.67** — apple, mountain, silence, love, river, baby, cloud

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 0: **77.34** — 空, 愛, 机, 夢, 火, 象, 時間

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 1: **87.87** — giraffe, spoon, jealousy, volcano, tuxedo, algebra, sneeze

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 0: **85.33** — 雷, 嫉妬, 契約, 角度, 肝臓, 正午, 暗号

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 1: **79.56** — 空, 時間, 感情, 数字, 法律, 筋肉, 雷

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 0: **81.01** — 制度, 霧, 肺, 悪意, 鋏, 苔, 摩擦

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 1: **77.76** — 重力, 感情, 液体, 角膜, 過去, 沈黙, 摩擦

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 1: **80.39** — 概念, 砂漠, 音楽, 法律, 針, 星, 涙

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 0: **80.09** — 椅子, 水, 馬, 木, 耳, 喜び, 未来

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 1: **73.64** — 時間, 光, 夢, 影, 記憶, 風, 沈黙

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 1: **79.77** — thunder, silence, bread, elbow, hammer, ocean, giraffe

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 0: **80.47** — 苔, 嫉妬, 彗星, 借金, かかと, 轟音, くしゃみ

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 0: **79.65** — 星, 鯨, 松, 鉄, 恋, 時間, 指

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 1: **77.37** — elephant, mountain, justice, thunder, bread, symphony, cactus

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 1: **85.60** — octopus, democracy, toothbrush, lava, whisper, carpet, infinity

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 0: **72.10** — time, stone, bird, hope, chair, storm, child

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 1: **71.67** — 空気, 夢, 影, 時間, 静寂, 記憶, 光

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 1: **78.83** — 砂糖, 雷, 孤独, 椅子, 蝶, 戦争, 風邪

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 1: **81.87** — 椅子, 自由, 匂い, 昨日, 砂糖, 雷, 記憶

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 1: **75.63** — 永遠, 嘘, 悲しみ, 子供, 蝶, 砂漠, 塩

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 1: **78.08** — 魂, 釘, 虹, 沈黙, 醤油, 雷, 墓

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 1: **84.01** — pebble, melody, hunger, butterfly, window, winter, honesty

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 1: **83.52** — クジラ, 妥協, 砂漠, ハサミ, 和音, くしゃみ, 銀河

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 0: **73.78** — 時間, 愛, 水, 本, 心, 光, 力

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 1: **85.17** — 瓦, クラゲ, 嫉妬, 正午, 借金, 悲鳴, 泥

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 1: **79.23** — 虹, 嫉妬, 砂糖, 契約, 胃袋, 歯車, 幽霊

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 0: **70.69** — 山, 心, 道, 夢, 花, 影, 声

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 1: **79.73** — 悲しみ, 骨, 未来, 貨幣, 鍋, 嘘, 数字

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 1: **77.68** — 山, 涙, 鍵, 風, 影, 糸, 火

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 1: **75.21** — 光, 沈黙, 怒り, 祈り, 砂, 鏡, 旋律

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 1: **77.63** — 鯨, 靴, 砂漠, 雷, 約束, 悲しみ, 祭り

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 0: **79.73** — 愛, 塩, 時間, 森, 沈黙, 雷, 鍵

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 1: **82.33** — 鯨, 誤解, 鉛筆, 正午, 裏切り, 砂漠, 味噌

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 1: **77.81** — 火山, 靴下, 約束, 蚊, 退屈, 鍵盤, 砂漠

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 0: **83.98** — 鯨, 豆腐, 箒, 砂漠, 昨日, 嫉妬, 約束

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 1: **80.51** — 砂漠, 良心, 歯車, 選挙, 借金, 香り, 眠り

### English · meta/muse-spark-open · thinking (effort=high)

- run 0: **83.11** — galaxy, hammer, sadness, symphony, volcano, democracy, cathedral

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 1: **77.64** — 鯨, 砂漠, 虹, 味噌, 鏡, 病院, お金

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 1: **78.78** — 電子, 悲しみ, 鯨, 憲法, 砂漠, 料理, 嘘

### English · meta/muse-wonder · thinking (effort=high)

- run 2: **89.01** — galaxy, broccoli, justice, trombone, kangaroo, volcano, mitten

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 2: **80.80** — ocean, volcano, pillow, justice, melody, elephant, hammer

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 2: **77.06** — elephant, justice, apple, planet, friendship, mountain, poem

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 2: **77.01** — freedom, stone, music, hunger, mirror, thunder, childhood

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 1: **80.30** — 時間, 空, 水, 感情, 数字, 法律, 海老

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 1: **80.84** — 石, 犬, 自由, 昨日, 悲しみ, 海, 雷

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 0: **82.41** — 雷, ミミズ, パラシュート, 退屈, 嘘, 胃, ゼロ

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 1: **80.40** — かたつむり, 雨, 祖母, 毛布, 埃, 喜び, 約束

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 1: **82.06** — 雷, 正義, 蛸, 椅子, 噂, 昼, 砂糖

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 1: **80.57** — 海, 希望, 包丁, 祖母, 卵, 税金, 雷

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 2: **78.77** — ocean, thunder, memory, triangle, laughter, shadow, harvest

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 1: **81.07** — 鳥, 壁, 喜び, 星, パン, 過去, 手

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 2: **77.17** — mountain, ocean, candle, laughter, silk, violin, shadow

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 1: **80.84** — 影, 絶望, 化石, 噂, 針, 港, 確率

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 1: **76.57** — mountain, sorrow, bicycle, kitchen, freedom, apple, silence

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 2: **83.41** — gravity, spoon, democracy, moss, whisper, volcano, jealousy

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 1: **80.51** — 海, 悲しみ, 財布, 時間, 病院, 味噌, 雷

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 2: **88.23** — volcano, spoon, justice, whisper, elephant, algebra, button

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 2: **87.54** — volcano, justice, spoon, nostalgia, whale, algebra, umbrella

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 2: **85.08** — cloud, justice, spoon, bacteria, symphony, tornado, grief

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 1: **74.42** — 愛, 山, 時計, 風, 夢, 電車, 宇宙

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 2: **75.93** — ocean, idea, hammer, child, storm, freedom, apple

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 2: **67.62** — love, mountain, silence, child, fire, dream, ocean

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 2: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 1: **81.33** — 朝, 鯨, 約束, 塩, 影, 記憶, 火

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 1: **80.13** — 蚊, 嫉妬, 豆腐, 雷, 鉛筆, 膝, 祭り

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 1: **82.08** — 地平線, 負債, キノコ, 摩擦, 秘密, ダム, 骨

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 2: **81.77** — thunder, ledger, blossom, sorrow, gravity, pebble, symphony

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 1: **80.21** — apple, democracy, lightning, philosophy, bicycle, courage, ocean

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 2: **80.25** — 憎悪, 羅針盤, 珊瑚, 蜃気楼, 味覚, 葬儀, 微生物

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 2: **86.76** — volcano, spoon, justice, whisper, giraffe, algebra, pillow

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 2: **76.17** — 夢, 石, 雨, 笑い, 影, 空腹, 約束

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 2: **77.71** — 砂漠, 友情, 稲妻, 枕, 沈黙, 祭り, 深海

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 2: **77.90** — apple, justice, elephant, ocean, happiness, century, explosion

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 0: **79.28** — 雷, 納豆, 嫉妬, 税金, 音楽, 毛糸, 砂漠

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 1: **77.03** — 愛, 石, 蟻, 靴, 時間, 音楽, 魚

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 1: **77.92** — 石, 魚, 鉛筆, 嘘, 雷, 砂糖, 午後

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 0: **77.53** — carrot, elephant, mountain, happiness, democracy, violin, dream

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 2: **79.70** — 概念, 気流, 岩石, 肝臓, 法律, 単位, 絶望

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 2: **75.83** — 夢, 根, 角, 波, 罪, 影, 傷

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 2: **87.14** — volcano, spoon, justice, giraffe, melody, bacteria, umbrella

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 2: **80.80** — 微風, 概念, 鉄, 夢, 数字, 孤独, 海

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 1: **80.22** — 永遠, 泡, 祈り, 鏡, 雷, 法律, 冗談

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 1: **78.93** — 宇宙, 指, 嘘, 鍵, 涙, 炎, 紙

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 2: **76.65** — 雲, 記憶, 影, 沈黙, 境界, 速度, 渇き

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 2: **75.99** — 月, 夢, 砂, 風, 影, 鐘, 糸

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 1: **81.54** — 嫉妬, 彗星, クラゲ, 箪笥, 骨髄, 利子, 苦味

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 2: **88.77** — giraffe, spoon, jealousy, volcano, democracy, saxophone, microbe

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 2: **84.24** — freedom, river, bread, elbow, jewel, hammer, walrus

### English · meta/muse-spark-open · thinking (effort=high)

- run 2: **75.64** — ocean, mountain, childhood, music, justice, thunder, mirror

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 2: **80.98** — 蟻, 松, 金槌, 砂糖, 雷, 昨日, 南

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 2: **80.76** — 石, 夢, 明日, 法律, 音, 味, 寒さ

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 2: **76.43** — 太陽, 猫, 花, 塩, 鏡, 歌, 怒り

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 2: **73.18** — 空, 海, 声, 夢, 匂い, 石, 笑顔

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 2: **80.26** — クジラ, 三角形, 法律, 砂漠, 友情, 鉄, 過去

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 2: **78.23** — 夢, 鉄, 雷, 歯車, 孤独, 法律, 砂

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 1: **82.49** — 砂漠, 嫉妬, 歯車, 税金, 昨日, まぶた, 敬語

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 1: **75.99** — 愛, 雲, 針, 夢, 塩, 机, 馬

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 2: **81.24** — 嫉妬, 彗星, 安全ピン, クラゲ, 矛盾, かかと, 遺言

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 2: **76.68** — 魂, 砂漠, 鉄, 笑顔, 旋律, 昆虫, 歴史

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 2: **84.68** — 虹, 嫉妬, はさみ, 法律, 砂漠, 祖先, 眠気

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 1: **77.32** — 石, 愛, 虫, 時間, 米, 光, 音楽

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 2: **78.23** — 喜び, 沈黙, 記憶, 霧, 砂, 飢え, 橋

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 2: **82.03** — 虹, 税金, 嫉妬, 豆腐, 祖母, 迷路, 明日

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 0: **83.24** — 砂漠, 嫉妬, 豆腐, 憲法, くしゃみ, 昨日, 歯車

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 2: **72.37** — 山, 海, 夢, 刃, 影, 歌, 虫

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 2: **76.00** — 魚, 哲学, 靴下, 雷, 友情, トマト, 時計

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 2: **79.63** — 鯨, 郵便, 正義, 砂糖, 夕暮れ, 歯車, 沈黙

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 2: **76.91** — 石, 時間, 愛, 犬, 音楽, 雷, 夢

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 2: **81.70** — 火山, 靴下, 嫉妬, 小数, 祖母, 醤油, 雷鳴

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 2: **79.57** — 鯨, 砂漠, 稲妻, 磁石, 将棋, 嫉妬, 和解

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 2: **75.16** — 時間, 感情, 水産, 重力, 国境, 音楽, 病院

### English · meta/muse-wonder · thinking (effort=high)

- run 3: **81.28** — octopus, anger, hammer, jacket, castle, potato, elbow

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 3: **82.12** — galaxy, potato, elephant, hammer, river, freedom, music

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 3: **80.35** — bread, feather, justice, melody, childhood, fire, mirror

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 3: **70.85** — light, silence, hunger, justice, stone, laughter, ocean

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 0: **78.16** — 幸福, 鉄, 蝶, 塔, 霧, 笛, 財布

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 2: **79.68** — 涙, 鉄, 嘘, 昼, 砂漠, 鯨, 数字

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 1: **76.12** — 希望, 石, 鳥, 雨, 夢, 塩, 壁

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 2: **75.53** — 石, 愛, 朝, 猫, 歌, 塩, 影

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 2: **76.03** — 宇宙, 正義, 旋律, 幽霊, 椅子, 匂い, 記憶

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 3: **76.12** — ocean, justice, hammer, laughter, mountain, winter, music

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 2: **83.36** — 錨, 嫉妬, 稲妻, 義務, カビ, 迷子, 来世

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 2: **79.09** — 鉛筆, 雷, 嫉妬, 味噌, 膝, 借金, 蛙

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 3: **72.51** — mountain, love, child, silence, river, paper, dream

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 2: **74.06** — 山, 愛, 時間, 魚, 空, 心, 星

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 2: **81.02** — 石, 犬, 愛, 昨日, 雨, 数, 正義

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 1: **81.62** — 豆腐, 砂漠, 嫉妬, 戸籍, くしゃみ, 電卓, 永遠

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 2: **80.43** — 鉛筆, くしゃみ, 嫉妬, 祝日, ペンギン, 砂漠, 音楽

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 3: **86.05** — volcano, spoon, justice, whisper, calendar, elephant, nostalgia

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 3: **83.81** — volcano, spoon, justice, whisper, carpet, elephant, midnight

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 3: **84.78** — cloud, justice, spoon, bacteria, melody, volcano, grief

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 3: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 2: **82.83** — 石, 正義, 蝶, 昨日, 歌, 怒り, 醤油

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 3: **86.87** — thunder, pillow, justice, glacier, violin, cabbage, ambition

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 2: **80.97** — 雨, 正義, 石, 会議, 痛み, 祖母, 味噌

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 3: **70.82** — mountain, child, silence, fire, ocean, idea, stone

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 3: **82.40** — volcano, spoon, justice, elephant, whisper, bicycle, cheese

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 3: **79.62** — 愛, 冷蔵庫, 蜃気楼, 珊瑚, 旋律, 税金, 隕石

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 2: **77.18** — 星, 猫, 希望, 泥, 音楽, 剣, 氷

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 3: **80.45** — 銀河, 友情, 階段, 砂糖, 沈黙, 蛙, 法律

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 2: **76.73** — 時間, 石, 愛, 雨, 靴, 星, 針

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 3: **75.77** — 夢, 石, 匂い, 怒り, 雨, 影, 食事

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 2: **76.92** — 歴史, 筋肉, 数字, 法律, 風, 嫉妬, 油

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 3: **88.15** — tornado, spoon, jealousy, moss, symphony, tooth, bankruptcy

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 3: **79.11** — apple, justice, river, laughter, clock, shadow, friendship

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 3: **77.56** — 正義, 光, 胃, 砂, 音楽, 夢, 三角

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 3: **84.85** — sorrow, granite, piano, ocean, elbow, decade, mustard

### English · meta/muse-spark-open · thinking (effort=high)

- run 3: **80.16** — galaxy, kindness, umbrella, volcano, melody, forest, ocean

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 3: **78.88** — 砂漠, 感情, 素数, 重力, 伝統, 胃袋, 幻覚

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 2: **80.37** — 石, 孤独, 猫, 木, 雨, 瞬間, 赤

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 2: **77.53** — 石, 愛, 犬, 空, 夢, 火, 塩

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 3: **87.66** — volcano, spoon, jealousy, giraffe, democracy, lullaby, tuxedo

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 3: **76.74** — 概念, 砂漠, 音楽, 法律, 鳥, 電子, 孤独

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 3: **75.65** — 愛, 数学, 夢, 雨, 時間, 記憶, 影

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 3: **87.30** — pickle, grief, volcano, trombone, hammer, justice, pillow

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 3: **75.52** — 嫉妬, 雷, 骨, 塩, クラゲ, 笛, 穴

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 3: **75.40** — 時計, 海, 考え, 山, 光, 夢, 影

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 3: **81.51** — elephant, justice, pebble, symphony, volcano, octopus, democracy

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 3: **79.71** — 嫉妬, 銀河, 胃袋, 酢, カマキリ, 針金, 雷鳴

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 2: **82.77** — 砂漠, 嫉妬, 豆腐, 選挙, 磁石, 昨日, 蝶

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 3: **80.48** — 夢, 鉄, 法律, 砂漠, 音, 箸, 嘘

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 3: **82.40** — 岩, 蟻, 向日葵, 悲しみ, 未来, 政府, 眼鏡

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 3: **80.78** — 鍋, 蜂, 嵐, 悲しみ, 明日, 砂漠, 米

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 3: **77.55** — 時間, 洞窟, 鳥, 指, 塩, 船, 雷

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 3: **68.61** — 光, 愛, 石, 夢, 音, 影, 味

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 2: **80.50** — 椅子, 猫, 怒り, 歯, 米, 雪, 約束

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 3: **79.34** — 虹, 納豆, 嫉妬, 税金, 枕, 砂漠, 胃袋

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 3: **76.69** — 愛, 砂漠, 鯨, 冷蔵庫, 虹, 音楽, 死

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 3: **82.51** — 虹, 後悔, 靴下, くしゃみ, 独裁, カビ, 締め切り

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 3: **73.13** — 山, 海, 夢, 針, 炎, 母, 虫

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 2: **82.91** — freedom, lobster, velvet, earthquake, melody, cactus, sunset

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 3: **78.69** — 時間, 岩石, 旋律, 憎悪, 匂い, 法則, 食器

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 1: **79.15** — 電話, 怒り, 窓, 蟻, 塩, 未来, 壁

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 2: **82.56** — 嫉妬, 霜, 確率, 選挙, 噂, 砂利, 車輪

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 3: **83.67** — 雲, 正義, 蟻, 音楽, 肝臓, 昨日, 鋏

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 3: **80.71** — 嫉妬, 雷, 過去, 車輪, 唾液, 税金, 砂糖

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 1: **78.27** — 愛, 塩, 針, 虎, 時間, 鍵, 言葉

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 3: **80.38** — 砂漠, 嫉妬, 冷蔵庫, 選挙, 蛙, 肺, 虹

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 3: **79.17** — 蝋燭, 正義, 鯨, 靴下, 黄昏, 算数, 火山

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 3: **78.73** — 鉛筆, 嵐, 正義, 象, 幸福, 数字, 音楽

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 3: **77.56** — 象, 涙, 橋, 自由, 卵, 影, 詩

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 3: **79.30** — lion, democracy, banana, symphony, brick, happiness, thunder

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 3: **81.68** — クジラ, 正義, スプーン, 火曜日, 記憶, 砂漠, ピアノ

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 3: **78.51** — 椅子, 稲妻, 孤独, 醤油, 瞳, 明日, 蟹

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 4: **79.88** — galaxy, potato, justice, laughter, ocean, dream, knife

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 4: **79.61** — ocean, sorrow, hammer, justice, bread, childhood, star

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 4: **78.83** — ocean, hunger, hammer, childhood, thunder, pillow, music

### English · meta/muse-wonder · thinking (effort=high)

- run 4: **86.65** — galaxy, cinnamon, octopus, trombone, democracy, pillow, volcano

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 3: **77.08** — 石, 愛, 時間, 魚, 音楽, 法律, 雲

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 3: **81.58** — 石, 犬, 悲しみ, 時間, 正義, 味噌, 雷

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 3: **80.50** — 虹, 借金, くしゃみ, 鉛筆, 砂漠, 友情, 蚊

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 3: **76.36** — 時間, 感情, 海岸, 税金, 数学, 詩, 機械

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 3: **75.33** — 愛, 机, 風, 未来, 石, 川, 自由

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 3: **84.76** — 嫉妬, 彗星, 爪楊枝, 欠伸, 憲法, 沼, 摩擦

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 3: **76.86** — 石, 恋, 朝, 母, 猫, 歌, 塩

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 4: **82.71** — ocean, childhood, thunder, kindness, spoon, galaxy, cathedral

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 4: **72.51** — love, mountain, child, silence, paper, river, dream

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 4: **87.98** — volcano, whisper, button, justice, spaghetti, nostalgia, hammock

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 3: **82.24** — 胃, 砂, 糊, 雷, 角度, 義務, 鍵

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 4: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 4: **84.78** — cloud, justice, spoon, bacteria, melody, volcano, grief

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 4: **84.79** — volcano, spoon, jealousy, algae, democracy, needle, whale

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 4: **88.23** — volcano, spoon, justice, whisper, giraffe, algebra, umbrella

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 3: **79.84** — 記憶, 霧, 饅頭, 膝, 借金, 蚊, 玄関

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 4: **86.29** — volcano, whisper, spoon, justice, basement, ladybug, melody

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 0: **75.99** — 岩, 蝶, 時計, 幸福, 道路, 音楽, 雲

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 4: **74.14** — mountain, ocean, baby, laughter, paper, silence, fire

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 4: **85.86** — thunder, pillow, justice, glacier, violin, mustard, rust

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 4: **77.72** — stone, idea, child, thunder, pocket, mercy, apple

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 3: **78.17** — 靴下, 火山, 嫉妬, 味噌, 借金, 蝶, 葬式

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 3: **78.68** — 石, 愛, 犬, 火, 嘘, 米, 目

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 4: **77.83** — 愛情, 彗星, 砂糖, 鯨, 沈黙, 螺旋, 祖先

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 4: **76.45** — 沈黙, 怒り, 雨, 砂, 鏡, 夢, 穴

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 4: **79.09** — 砂漠, 約束, 昆虫, 旋律, 怒り, 階段, 偶然

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 3: **80.14** — 胃袋, 吹雪, サボテン, 鉄筋, 嫉妬, 隕石, 文法

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 1: **81.85** — クラゲ, 氷河, 金庫, 嫉妬, 確率, 国境, 残像

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 4: **79.04** — 宇宙, 友情, 豆腐, 例外, 沈黙, 筋肉, 正義

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 3: **79.37** — 石, 悲しみ, 夏, 歌, 蜂, 法律, 匂い

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 3: **77.51** — 嫉妬, 豆腐, 方言, 税金, 夕焼け, 靴下, 幽霊

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 2: **78.25** — 愛, 時間, 椅子, 塩, 星, 記憶, 歌

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 4: **80.64** — banana, hammer, freedom, mountain, violin, justice, pillow

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 4: **86.87** — volcano, spoon, jealousy, giraffe, tax, lullaby, velvet

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 4: **85.84** — elephant, justice, spoon, volcano, melody, carpet, galaxy

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 4: **74.88** — 沈黙, 記憶, 約束, 眠り, 砂漠, 嵐, 赤ん坊

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 4: **76.13** — 光, 鯨, 砂漠, 音楽, 方程式, 感情, 歴史

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 4: **76.18** — 引力, 感情, 液体, 憲法, 睡眠, 火山, 沈黙

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 4: **80.35** — freedom, hammer, elephant, bread, thunder, ankle, regret

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 4: **76.91** — 愛, 影, 時間, 沈黙, 記憶, 境界, 予感

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 3: **78.34** — 星, 愛, 時間, 蚊, 鼻, 鏡, 塩

### English · meta/muse-spark-open · thinking (effort=high)

- run 4: **79.28** — galaxy, hammer, shadow, cathedral, volcano, puppet, ocean

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 4: **79.60** — 山脈, 時計, 友情, 海藻, 選挙, 蒸気, 記憶

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 3: **75.46** — 猫, 愛, 石, 夢, 時間, 雨, 言葉
- run 2: **76.53** — 愛, 鉄, 蝶, 戦争, 針, 雨, 壁

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 4: **81.40** — 消しゴム, 霧, 嫉妬, じゃがいも, 剣, らくだ, 未来

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 4: **76.17** — 岩, 幸福, 鏡, 塩, 夢, 霧, 歌

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 4: **76.89** — 太陽, 墓, 爪, 祭り, 絹, 約束, 鯨

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 4: **76.03** — 太陽, 孤独, 約束, 階段, 塩, 歌声, 砂漠

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 4: **82.93** — クジラ, ネジ, 嫉妬, 雷, 昨日, 角度, かかと

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 2: **75.74** — 太陽, 靴, 車, 壁, 夢, 針, 机

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 4: **81.25** — 嫉妬, 彗星, 法律, クラゲ, 剃刀, 螺旋, 酢

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 4: **83.06** — 虹, 嫉妬, 蝶番, 税金, 化石, しゃっくり, 約束

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 4: **73.19** — 愛, 泥, 嘘, 骨, 影, 笛, 雲

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 4: **75.59** — 山, 海, 涙, 椅子, 嘘, 夢, 塩

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 4: **79.07** — 宇宙, 嫉妬, 味噌, 蝶, 音楽, 砂漠, 記憶

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 4: **75.84** — 石, 時間, 愛情, 雷, 匂い, 言葉, 夢

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 2: **80.95** — 砂漠, 嫉妬, 税金, 豆腐, くしゃみ, 未来, 毛糸

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 4: **79.98** — 蝉, 正義, 冷蔵庫, 夕焼け, 借金, ピアノ, 迷路

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 4: **77.67** — 石, 愛, 時間, 魚, 音楽, 嘘, 雲

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 4: **79.34** — 象, 借金, 虹, 歯ブラシ, 嫉妬, 火山, 靴下

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 4: **80.85** — 鯨, 正義, 蝋燭, 午後, 包丁, 憂鬱, 砂漠

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 3: **82.61** — 砂漠, 嫉妬, 豆腐, 選挙, 旋律, 昨日, 毛糸

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 4: **80.34** — 海, 時間, 哲学, 砂糖, 稲妻, 裏切り, 喜劇

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 4: **80.96** — 正義, 砂糖, 雷, 鉛筆, 祖母, 沈黙, 地下鉄

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 4: **81.94** — 虹, 肝臓, 税金, 砂漠, 毛糸, 昨日, 嫉妬

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 4: **78.41** — 石, 愛, 時間, 犬, 音, 数, 嘘

### English · meta/muse-wonder · thinking (effort=high)

- run 5: **83.30** — cheese, democracy, pillow, volcano, bicycle, poetry, elephant

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 5: **80.75** — galaxy, potato, silence, justice, hammer, ocean, childhood

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 5: **74.73** — elephant, cactus, soup, hammer, pillow, thunder, moon

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 5: **80.49** — eternity, pebble, thunder, lullaby, harvest, mirror, courage

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 4: **74.57** — 空, 愛, 石, 時間, 夢, 火, 海

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 4: **81.64** — 鉛筆, イルカ, 砂漠, 雷, 豆腐, 嫉妬, 約束

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 4: **80.55** — 宇宙, 醤油, 接吻, 裁判, 雷鳴, 刺繍, 墓場

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 4: **74.57** — 愛, 石, 蝶, 夢, 火, 影, 川

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 2: **81.21** — クラゲ, 氷河, 金庫, くしゃみ, 退屈, 選挙, 苦味

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 4: **76.91** — 石, 愛, 鳥, 朝, 約束, 雨, 米

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 5: **76.79** — mountain, shadow, laughter, engine, arrow, honey, thunder

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 0: **82.40** — 嫉妬, クラゲ, 砂丘, 南京錠, 関税, 直角, 残響

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 4: **81.03** — 隕石, 醤油, 畳, 税金, 嫉妬, 旋律, 結露

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 4: **79.82** — 象, 飯, 喜び, 道, 朝, 机, 平和

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 5: **73.21** — tree, freedom, hammer, ocean, joy, cat, mountain

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 5: **75.32** — ocean, mountain, child, music, silence, fire, paper

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 4: **78.79** — 火山, 靴下, 嫉妬, 蝶, 借金, 祖母, 砂糖

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 5: **73.43** — ocean, silence, bicycle, anger, mountain, bread, dream

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 4: **79.60** — 深海, 微粒子, 義務, 沈黙, 摩擦, 悪夢, 貨幣

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 5: **85.99** — volcano, whisper, button, justice, spaghetti, nostalgia, hammer

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 4: **79.88** — 雲, 契約, 歌, 鉄, 祭り, 時計, 犬

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 5: **88.97** — volcano, whisper, spreadsheet, grandmother, umbrella, justice, cabbage

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 4: **82.40** — 石, 愛, 昨日, 鯨, 音楽, 法律, 数

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 5: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 5: **79.98** — thunder, silence, bread, galaxy, anger, velvet, compass

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 5: **78.88** — 愛情, 火山, 珊瑚, 時計, 沈黙, 香辛料, 幽霊

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 4: **83.28** — spoon, grief, planet, sock, laughter, canyon, ink

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 5: **75.23** — 石, 夢, 祭り, 孤独, 香り, 約束, 海

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 5: **71.40** — mountain, love, paper, silence, river, dream, stone

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 5: **84.78** — cloud, justice, spoon, bacteria, melody, volcano, grief

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 5: **75.00** — 石, 涙, 夢, 匂い, 影, 約束, 沈黙

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 5: **86.38** — spoon, volcano, jealousy, umbrella, bacteria, symphony, elbow

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 5: **84.40** — cloud, justice, mushroom, scissors, whale, jealousy, volcano

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 4: **77.80** — 砂, 悲しみ, 卵, 約束, 針, 声, 光

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 5: **76.22** — 虚無, 時間, 光, 数学, 感情, 国家, 細胞

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 5: **79.18** — apple, justice, hammer, cloud, happiness, bicycle, ocean

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 5: **85.63** — hunger, thunder, violin, ocean, cheese, elbow, debt

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 5: **79.67** — apple, justice, mountain, laughter, pencil, ocean, childhood

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 5: **71.71** — 重力, 感情, 液体, 倫理, 瞬間, 金属, 虚無

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 5: **90.65** — volcano, spoon, jealousy, giraffe, tuxedo, algebra, mortgage

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 5: **71.53** — 雲, 記憶, 沈黙, 影, 境界, 旋律, 虚空

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 5: **77.50** — 概念, 砂漠, 音楽, 法律, 胎児, 磁石, 孤独

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 2: **76.38** — 虎, 雲, 鍵, 幸福, 戦争, パン, 砂漠

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 4: **75.55** — 愛, 時計, 山, 自由, 魚, 夢, 鉄

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 5: **79.96** — 心臓, 宇宙, 希望, 時計, 孤独, 虹, 記憶

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 5: **78.02** — 友情, 硫酸, 稲妻, 恐竜, 針, 砂漠, 税金

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 4: **76.41** — 石, 愛, 象, 机, 夢, 水, 山

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 5: **78.20** — 愛, 音, 星, 血液, 歯車, 法律, 岩

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 5: **80.79** — 虹, 借金, 嫉妬, 砂漠, くしゃみ, 鍵, 豆腐

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 5: **79.86** — 羊, 雷, 墨, 迷路, 祈り, 孤独, 王冠

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 5: **78.57** — 孤独, 蛙, 向日葵, チーズ, 飛行機, 洞窟, 霧

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 5: **87.49** — 唾液, 借金, 稲妻, 悲嘆, 蝶番, 昨晩, 苔

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 5: **72.62** — 石, 夢, 音, 匂い, 味, 色, 北

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 5: **81.29** — 明日, 孤独, 煙, 灯台, 眉, 蝶, 茄子

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 5: **77.04** — 沈黙, 雷, 指輪, 墓, 慈悲, 塩, 旋律

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 4: **83.34** — justice, mango, whale, skyscraper, lightning, melody, airplane

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 5: **86.20** — tortoise, volcano, democracy, spoon, jazz, onion, galaxy

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 4: **83.75** — 嫉妬, 釘, クラゲ, 霧, 正午, 遺言, 傾斜

### English · meta/muse-spark-open · thinking (effort=high)

- run 5: **84.67** — galaxy, hammer, octopus, grief, cathedral, democracy, symphony

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 3: **77.42** — 正義, 机, 猫, 塩, 船, 雲, 歌

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 5: **81.26** — 砂漠, 嫉妬, 豆腐, 税金, くしゃみ, 鏡, 未来

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 5: **78.38** — 宇宙, 祈り, 針, 法律, 雷, 砂漠, 歌

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 5: **76.71** — 愛, 鯨, 火山, 砂糖, 鏡, 音楽, 夢

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 5: **74.64** — 山, 海, 星, 涙, 嘘, 鏡, 匂い

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 5: **81.70** — 正義, 石, 匂い, 昨日, 笑い, 海, 約束

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 5: **77.38** — 石, 夢, 電車, 蜂蜜, 法律, 鏡, 火山

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 5: **83.21** — ほこり, 正義, 鯨, トマト, 左, 夕暮れ, 歯車

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 5: **75.39** — 夢, 石, 母, 塩, 風, 骨, 嘘

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 5: **79.58** — 火山, 靴下, 約束, 蜂蜜, 沈黙, 鉛筆, 祖母

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 5: **82.41** — lion, spoon, truth, earthquake, bread, elbow, bicycle

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 5: **76.83** — 石, 愛, 時間, 魚, 音, 法律, 夢

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 5: **80.51** — 砂漠, 赤ん坊, 嫉妬, 胡椒, 箪笥, 夕暮れ, 選挙

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 3: **77.17** — 愛, 時間, 象, 雨, 塩, 鉄, 夢

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 4: **79.74** — 納豆, 嫉妬, 砂漠, 歯車, 祖先, 稲妻, 約束

### English · meta/muse-wonder · thinking (effort=high)

- run 6: **88.07** — galaxy, democracy, hammer, glacier, chess, sorrow, perfume

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 6: **81.26** — ocean, pillow, justice, ant, volcano, song, bread

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 6: **73.10** — eternity, happiness, mountain, apple, music, darkness, elephant

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 6: **75.54** — ocean, childhood, silence, hammer, dream, feather, hunger

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 5: **77.49** — 時間, 空気, 感情, 机, 戦争, 水素, 祭り

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 5: **78.43** — 海, 約束, 鉛筆, 悲しみ, 猫, 税金, 歌

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 5: **79.78** — 宇宙, 慈愛, 包丁, 砂漠, 細菌, 契約, 旋律

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 6: **72.57** — ocean, mountain, music, freedom, shadow, fire, childhood

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 5: **78.42** — 絶望, 銀河, 法律, 筋肉, 匂い, 噂, 影

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 6: **74.09** — mountain, river, love, child, paper, silence, fire

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 5: **80.43** — 象, 嫉妬, 霧, 味噌, 葬式, 膝, 冗談

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 5: **78.56** — 例外, 永遠, 隙間, 鉄, 煙, 苔, 嫉妬

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 5: **73.77** — 海, 鉛筆, 笑顔, 塩, 歌, 骨, 影

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 5: **80.14** — 時間, 膝, 霧, 悲しみ, 絹, 借金, 祭り

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 5: **76.75** — 猫, 花, 石, 愛, 水, 時間, パン

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 6: **80.48** — pebble, jealousy, thunder, spoon, childhood, ocean, honesty

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 6: **86.05** — volcano, whisper, calendar, spoon, justice, elephant, nostalgia

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 6: **84.78** — cloud, justice, spoon, bacteria, melody, volcano, grief

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 6: **83.75** — justice, mushroom, volcano, whisper, spoon, galaxy, jealousy

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 6: **83.81** — volcano, whisper, spoon, justice, nostalgia, bicycle, cheese

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 4: **80.17** — 砂漠, 嫉妬, 歯車, 税金, くしゃみ, 明日, 迷信

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 6: **76.78** — mountain, baby, candle, river, silk, laughter, freedom

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 6: **87.66** — volcano, spoon, jealousy, giraffe, algebra, sneeze, democracy

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 3: **75.76** — 愛, 石, 魚, 雲, 時間, 金, 夢

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 5: **77.42** — 火, 母, 時間, 川, 卵, 歌, 歯

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 3: **79.96** — 砂漠, クラゲ, 嫉妬, 梯子, 矛盾, 酸味, 関税

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 6: **83.40** — thunder, violin, ocean, mortgage, lettuce, silence, glacier

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 6: **80.68** — 憧憬, 珊瑚, 離婚, 微生物, 税金, 竜巻, 俳句

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 6: **88.16** — volcano, spoon, justice, whisper, tortoise, algebra, pillow

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 5: **75.77** — 愛, 石, 時間, 音楽, 戦争, 花, 影

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 4: **75.98** — 愛, 椅子, 牛, 雨, 詩, 塩, 鍵

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 6: **78.49** — apple, justice, mountain, laughter, hammer, ocean, curiosity

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 4: **77.69** — 愛, 椅子, 雷, 牛, 塩, 夢, 戦争

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 5: **81.18** — 嫉妬, 直角, 杉, 溶岩, 通貨, 幽霊, 胃袋

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 6: **73.68** — 雨, 影, 約束, 笑顔, 砂, 記憶, 鍵

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 5: **79.25** — 石, 夢, 雷, 鯨, 法律, 昨日, 歌

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 6: **79.32** — apple, elephant, justice, hammer, ocean, music, gravity

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 6: **81.21** — 砂漠, 良心, 傘, 祭り, 沈黙, 幼虫, 借金

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 4: **75.33** — りんご, 雲, 友情, 魚, 机, 夢, 山

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 6: **78.40** — 微粒子, 感情, 角, 時間, 重力, 言葉, 夢

### English · meta/muse-spark-open · thinking (effort=high)

- run 6: **78.32** — ocean, justice, triangle, laughter, feather, hammer, thunder

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 6: **80.13** — 重力, 嫉妬, 液体, 瞬間, 祖先, 祝日, 旋律

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 6: **90.01** — volcano, spoon, jealousy, giraffe, algebra, sock, mortgage

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 6: **76.37** — 概念, 砂漠, 音楽, 法律, 星, 感情, 電子

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 6: **78.71** — 量子, 愛, 時間, 影, 記憶, 沈黙, 鏡

### English · x-ai/grok-4.5 · thinking (effort=high)

- run 6: **74.01** — elephant, freedom, hammer, ocean, joy, butterfly, justice

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 5: **78.98** — 塩, 無限, 靴下, 火山, 沈黙, 鉛筆, 宗教

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 6: **70.83** — 宇宙, 自由, 幸福, 正義, 時間, 愛, 真理

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 6: **82.02** — hunger, thunder, mirror, wolf, granite, lullaby, elbow

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 6: **87.07** — volcano, spoon, democracy, giraffe, violin, sadness, bicycle

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 6: **81.97** — ペンギン, 松, 退屈, 昔, 裁判, 香水, 砂漠

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 6: **76.52** — 石, 正義, 明日, 香り, 痛み, 嘘, 味

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 6: unscorable (0 valid) — 

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 6: **80.67** — 昨日, 喜び, 狼, 雷, 椅子, ご飯, 歌

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 5: **71.50** — 山, 愛, 魚, 夢, 自由, 花, 影

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 6: **77.88** — 蜃気楼, 嫉妬, 鯨, 羅針盤, 香水, 夢, 雷

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 6: **81.03** — 嫉妬, 鯨, 鉄, 過去, 理由, 廊下, 旋律

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 3: **80.93** — camel, spoon, jealousy, lightning, wool, autumn, thumb

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 6: **75.24** — 砂, 約束, 孤独, 香り, 窓, 空腹, 季節

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 6: **81.04** — 魂, 砂利, 虹, 法律, 鍋, 挨拶, 雷

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 6: **79.65** — 稲妻, 嫉妬, 針, 鯨, 賄賂, 骨髄, 沈黙

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 6: **74.43** — 炎, 骨, 夢, 砂, 孤独, 味, 鐘

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 6: **82.76** — 砂丘, 郷愁, 砥石, 義務, 彗星, 満腹, クラゲ

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 5: **80.82** — 夢, 鍵, 怒り, 月, 米, 猫, 約束

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 6: **82.26** — 雷, 豆腐, 嫉妬, 税金, 砂漠, 老後, はさみ

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 6: **82.45** — 宇宙, 嫉妬, 豆腐, 毛布, 約束, 昨日, 香り

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 6: **80.61** — 未来, 疑惑, 重力, 矛盾, 苔, 針, 儀式

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 6: **79.58** — 火山, 靴下, 約束, 蜂蜜, 沈黙, 鉛筆, 祖母

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 6: **80.54** — 墓地, ピアノ, 雷, 砂糖, 友情, 靴下, 月曜日

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 6: **79.42** — elephant, freedom, spoon, thunder, childhood, mirror, hunger

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 6: **77.93** — 石, 愛, 時間, 魚, 音楽, 嘘, 風

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 6: **79.94** — りんご, 怒り, 星, 象, 石, 時間, 心臓

### English · meta/muse-wonder · thinking (effort=high)

- run 7: **87.09** — octopus, pancake, bicycle, thunderstorm, library, envy, galaxy

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 7: **81.60** — galaxy, potato, justice, sand, melody, hammer, spider

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 7: **76.42** — freedom, stone, music, hunger, ocean, hammer, childhood

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 5: **73.56** — 愛, 雷, 狼, 蜜, 塔, 運命, 影

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 6: **79.23** — 鉛筆, 宇宙, 正義, 鯨, 笑い, 鍵, 火山

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 7: **82.77** — thunder, silk, hunger, hammer, childhood, galaxy, bread

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 6: **78.92** — 海, 時間, 鉛筆, 感情, 政治, 稲妻, 台所

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 6: **75.13** — elephant, mountain, happiness, hammer, ocean, time, bread

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 6: **77.92** — 鯨, 正義, 靴下, 夕方, 味噌, 雷, 電車

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 7: **72.51** — love, mountain, paper, child, silence, river, dream

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 7: **75.83** — ocean, mountain, spoon, umbrella, laughter, freedom, thunder

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 6: **77.85** — 石, 夢, 犬, 怒り, 空, パン, 手

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 6: **76.62** — 石, 愛, 時間, 魚, 音楽, 戦争, 椅子

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 5: **78.81** — 砂漠, 嫉妬, 豆腐, 法律, 睡眠, 音楽, 親戚

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 6: **73.65** — 永遠, 慈愛, 罪, 砂漠, 旋律, 稲妻, 沈黙

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 7: **75.36** — water, dream, stone, music, mother, cloud, knife

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 7: **91.82** — volcano, whisper, invoice, giraffe, nostalgia, spoon, algebra

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 6: **81.98** — 鉛筆, 嫉妬, 火山, 味噌, 膝, 契約, 蝶

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 6: **82.56** — 嫉妬, 隕石, 爪楊枝, 税金, クラゲ, 苦味, 世紀

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 7: **87.10** — volcano, justice, spoon, bacteria, jazz, tornado, grief

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 7: **84.40** — volcano, spoon, justice, whisper, bicycle, marriage, cheese

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 6: **79.86** — 悲しみ, 石, 時間, 規則, 料理, 雲, 指

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 6: **77.13** — りんご, 山, 愛, 椅子, 雷, 砂糖, 笑顔

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 7: **88.07** — volcano, spoon, jealousy, giraffe, violin, mortgage, fog

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 7: **85.05** — thunder, pillow, democracy, lemon, silence, engine, glacier

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 7: **78.75** — 愛, 蜃気楼, 素数, 沈黙, 乳房, 羅針盤, 嫉妬

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 7: **84.78** — cloud, justice, spoon, bacteria, melody, volcano, grief

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 7: **85.89** — cat, volcano, justice, spoon, melody, bacteria, galaxy

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 7: **75.41** — 愛, 石, 夢, 風, 虫, 火, 嘘

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 6: **77.45** — 石, 愛, 朝, パン, 耳, 戦争, 風

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 6: **80.80** — 石, 愛, 魚, 鉛筆, 昨日, 音楽, 医者

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 7: **79.71** — 砂漠, 友情, 雷, 枕, 税金, 香り, 迷路

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 6: **80.98** — 悲しみ, 歯, 雷, 蝶, 鍵, 水, 星

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 6: **76.92** — 太陽, 友情, 法律, 夢, 川, 鳥, 机

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 7: **81.27** — 沈黙, 角, 水滴, 過ち, 正義, 粒子, 爪

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 7: **83.88** — silence, bread, thunder, elbow, granite, decade, lantern

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 7: **87.85** — giraffe, spoon, jealousy, volcano, algebra, sock, lullaby

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 6: **81.79** — 未来, 砂, 爪, 挨拶, 角度, 祖母, 嘘

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 7: **74.62** — mountain, paper, childhood, silence, cat, fire, dream

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 7: **75.42** — 重力, 感情, 液体, 憲法, 微生物, 冗談, 金属

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 7: **75.78** — 概念, 砂漠, 音楽, 法律, 鳥, 電子, 言語

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 7: **78.82** — elephant, freedom, volcano, laughter, needle, ocean, justice

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 7: **75.32** — 時間, 夢, 影, 記憶, 沈黙, 境界, 偶然

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 7: **83.64** — volcano, spoon, democracy, elephant, whisper, violin, grief

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 5: **77.75** — 愛, 蝶, 石, 雨, 机, 時間, 涙

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 7: **80.18** — 根拠, 泥, 怒り, 秒, 鯨, 針, 影

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 7: **73.93** — 空気, 夢, 影, 無限, 静寂, 鏡, 時間

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 6: **78.05** — 椅子, 記憶, 雷, 愛, 鍵, 時間, 花

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 1: **83.16** — 雷, 約束, 砂糖, 北, 痛み, 昨日, 糸
- run 4: **80.23** — 祖母, 雷, 南, 米, 靴, 罪, 馬

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 7: **80.58** — 悲しみ, 豆腐, のこぎり, 靴下, 雷, 結婚, 髪

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 7: **79.02** — 岩, 蜂, 悲しみ, 昨日, 静寂, 雷, 森

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 7: **77.37** — 嫉妬, 永遠, 鉄, 苔, 蝶, 鏡, 骨

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 7: **73.66** — 太陽, 海, 石, 風, 音, 時間, 愛

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 7: **81.79** — 永遠, 賄賂, 胞子, 嫉妬, 彗星, 針, 苦味

### English · meta/muse-spark-open · thinking (effort=high)

- run 7: **79.38** — ocean, laughter, pyramid, thunder, freedom, bicycle, bread

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 7: **79.16** — 砂, 友情, 鯨, 帽子, 雷, 昨日, 味噌

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 7: **74.75** — 霧, 骨, 貨幣, 嘘, 罠, 灰, 影

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 7: **81.61** — 雷, 友情, 砂糖, 昨日, 胃袋, 切符, 音楽

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 7: **78.22** — 慈悲, 雷鳴, 納豆, 幽霊, 法律, 砂漠, 虹

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 7: **79.50** — fire, memory, shoe, kindness, river, clock, shadow

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 7: **76.54** — 鯨, 砂漠, 鉛筆, 香辛料, 稲妻, 友情, 貨幣

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 7: **82.32** — 虹, 豆腐, 嫉妬, 税金, 肝臓, 祖先, 画鋲

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 6: **80.82** — 愛, 釘, 霧, 財布, 肘, 嘘, 米

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 7: **78.64** — 鯨, 胃, 雨, 砂漠, 椅子, 誕生, 眠り

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 7: **74.64** — 山, 海, 火, 糸, 歌, 子, 愛

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 7: **79.24** — 鯨, 正義, 消しゴム, 黄昏, 味噌, 暗号, 靴下

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 6: **82.02** — シマウマ, 歯ブラシ, 砂漠, 虹, 嫉妬, 約束, 昨日

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 7: **79.52** — 鯨, 正義, 消しゴム, 黄昏, 味噌, 裏切り, 階段

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 7: **77.67** — 石, 愛, 時間, 魚, 音楽, 嘘, 雲

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 5: **83.13** — 砂漠, 嫉妬, 豆腐, くしゃみ, 税金, 昨日, 祖先

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 7: **82.06** — 象, 借金, 虹, 歯ブラシ, 嫉妬, 火山, 約束

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 7: **76.61** — 時間, 海岸, 感情, 税金, 空気, 数字, 病院

### English · meta/muse-wonder · thinking (effort=high)

- run 8: **85.30** — giraffe, tulip, bread, bicycle, loneliness, memory, century

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 8: **82.70** — galaxy, potato, justice, whisper, elephant, chair, childhood

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 8: **77.05** — freedom, eternity, elephant, apple, laughter, mountain, memory

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 8: **72.45** — silence, mountain, laughter, hunger, mirror, thunder, childhood

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 7: **76.07** — 猫, 数学, 砂漠, 悲しみ, 鉛筆, 火山, 音楽

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 4: **80.37** — クラゲ, カスタネット, 砂丘, 水銀, オーロラ, しゃっくり, 嫉妬

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 7: **77.42** — 石, 愛, 時間, 魚, 音, 嘘, 塩

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 8: **73.64** — mountain, river, child, paper, candle, dream, stone

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 8: **76.14** — freedom, hammer, ocean, thunder, shadow, laughter, pyramid

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 7: **79.78** — 虹, 歓喜, 権利, 矛盾, 香り, 未来, 粘土

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 7: **81.68** — 鍵, 苔, 嵐, 法律, 悪意, 零, 唾液

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 7: **77.78** — 嫉妬, 稲妻, 螺旋, 踵, 苔, 税金, 絨毯

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 8: **73.47** — apple, mountain, love, chair, fish, dream, knife

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 7: **78.64** — 水, 時間, 友情, 祭り, 石, 言葉, 眠り

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 7: **80.15** — 嫉妬, 鉛筆, 火山, 蛸, 納豆, 借金, 膝

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 7: **74.22** — 石, 恋, 象, 塩, 歌, 氷, 神

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 7: **78.87** — 愛, 石, 山, 魚, 本, 火, 靴

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 5: **74.24** — 猫, 橋, 歌, 時計, 鍵, 光, 砂

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 6: **83.61** — 虹, 嫉妬, 税金, 豆腐, くしゃみ, 祖先, 昨日

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 8: **83.96** — volcano, whisper, spoon, justice, elephant, nostalgia, bicycle

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 8: **91.19** — volcano, whisper, spreadsheet, grandmother, umbrella, justice, mackerel

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 7: **75.33** — 愛, 象, 机, 夢, 星, 川, 鍵

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 8: **86.69** — thunder, pillow, democracy, glacier, lullaby, ledger, mushroom

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 8: **80.74** — 愛憎, 量子, 蜃気楼, 珊瑚, 弔辞, 薔薇, 飢餓

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 8: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 8: **77.10** — 砂漠, 友情, 雷, 枕, 沈黙, 祭り, 影

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 8: **84.81** — cloud, justice, spoon, bacteria, symphony, volcano, grief

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 8: **87.25** — spoon, galaxy, jealousy, bacteria, thunder, democracy, violin

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 8: **76.46** — 悲しみ, 岩, 約束, 匂い, 空腹, 影, 扉

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 8: **84.81** — cloud, justice, spoon, bacteria, symphony, volcano, grief

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 7: **75.83** — 夢, 塩, 猫, 歌, 罪, 雨, 歯

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 7: **79.67** — 鋏, 鯨, 悲しみ, 砂漠, 蜂蜜, 肘, 旋律

### English · meta/muse-spark-open · thinking (effort=high)

- run 8: **74.99** — ocean, thunder, memory, shadow, kingdom, laughter, pillow

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 8: **86.33** — giraffe, spoon, jealousy, volcano, algebra, sneeze, velvet

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 8: **72.73** — mountain, ocean, baby, laughter, silence, paper, moon

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 7: **81.89** — freedom, bicycle, elephant, philosophy, sandwich, galaxy, laughter

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 7: **75.15** — 慈愛, 棘, 砂漠, 旋律, 錆, 霧, 鞄

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 8: **78.11** — 時間, 微生物, 正義, 沈黙, 重力, 故郷, 感情

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 8: **81.97** — freedom, thunder, bread, elbow, anchor, silence, piano

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 8: **74.98** — 重力, 感情, 液体, 法律, 数字, 筋肉, 時間

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 8: **82.41** — 風邪, 倫理, 砂漠, 数字, 誇り, 秒針, 鏡

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 6: **79.41** — 幸福, 時間, 鉄, 塩, 蝶, 鏡, 詩

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 8: **74.40** — 時間, 光, 夢, 影, 記憶, 沈黙, 境界

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 8: **76.86** — apple, justice, elephant, storm, chair, joy, ocean

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 8: **76.47** — 山, 海, 時間, 夢, 自由, 友情, 記憶

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 8: **78.50** — ocean, needle, joy, castle, dust, promise, knee

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 7: **79.83** — 砂漠, 豆腐, 嫉妬, 税金, 睡眠, 歯車, 噂

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 8: **79.46** — 鯨, 針, 絶望, 砂糖, 法律, 雷, 隙間

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 6: **74.18** — 愛, 椅子, 鉄, 塩, 鍵, 骨, 夢

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 7: **74.86** — 愛, 岩, 時間, 血, 雲, 鍵, 花

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 8: **78.85** — 砂糖, 怒り, 雷, 墓, 蝶, 鏡, 明日

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 8: **70.70** — 石, 夢, 音, 匂い, 味, 影, 未来

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 8: **78.74** — 赤ん坊, 砂, 虹, 沈黙, 空腹, 喜び, 未来

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 8: **75.60** — 愛, 石, 未来, 猫, パン, 船, 雷

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 8: **71.46** — 光, 音, 石, 夢, 愛, 時, 影

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 8: **87.47** — spoon, galaxy, justice, mushroom, violin, glacier, tax

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 8: **75.24** — 愛, 鯨, 火山, 鏡, 納豆, 夢, 砂漠

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 7: **77.51** — 山, 星, 針, 涙, 夢, 声, 鏡

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 8: **80.35** — 砂漠, 嫉妬, 歯車, 祝日, 納豆, 噂, 税金

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 8: **80.15** — 虹, 税金, くしゃみ, 迷路, 豆腐, 嫉妬, 化石

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 7: **75.30** — 猫, 山, 愛, 机, 雨, パン, 車

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 6: **80.95** — 砂漠, 嫉妬, 豆腐, 税金, くしゃみ, 毛糸, 未来

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 7: **76.72** — 雷, 喜び, 窓, 砂, 蟻, 歌, 鉄

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 8: **76.55** — 山, 涙, 匂い, 嘘, 針, 夢, 塩

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 8: **78.23** — 絶望, 歯車, 砂粒, 約束, 嘘, 影, 音楽

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 8: **78.87** — 鯨, 正義, 消しゴム, 黄昏, 味噌, 階段, 孤独

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 8: **81.28** — 正義, 石, 匂い, 約束, 嵐, 数, 笑い

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 8: **80.10** — 象, 微笑み, 鉛筆, 火山, 借金, 睡眠, 雷

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 8: **79.03** — 雪, 正義, 鉛筆, 音楽, 象, 胃, 借金

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 8: **78.40** — 慈愛, 岩石, 宇宙, 旋律, 自由, 雷鳴, 針

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 8: **80.91** — 彗星, 恥, 蝶番, 砂漠, 爪, 遺言, 影

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 8: **81.08** — 時間, 感情, 水田, 哲学, 電波, 祭り, 骨折

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 8: **81.90** — 砂漠, 約束, 封筒, くしゃみ, 雷, 叔父, 胡椒

### English · meta/muse-wonder · thinking (effort=high)

- run 9: **85.33** — elephant, broccoli, guitar, plastic, anger, earthquake, village

### English · meta/muse-wonder · no-thinking (effort=minimal)

- run 9: **73.37** — mountain, ocean, laughter, freedom, sandwich, elephant, dream

### English · meta/muse-spark-1.3 · thinking (effort=high)

- run 9: **72.23** — freedom, stone, music, hunger, dog, ocean, dream

### English · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 9: **72.14** — time, stone, dream, music, shadow, hunger, river

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 8: **80.60** — 海, 約束, 鉛筆, 午後, 猫, 豆腐, 階段

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 8: **75.56** — dog, idea, stone, joy, ship, law, rain

### English · x-ai/grok-4.7 · no-thinking (effort=low)

- run 9: **80.45** — ocean, anger, shoe, music, seed, justice, cloud

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 8: **78.27** — 鉛筆, 象, 嫉妬, 火山, 味噌, 借金, 筋肉

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 8: **78.19** — 石, 愛, 時間, 魚, 音楽, 戦争, 数

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 8: **79.32** — 宇宙, 記憶, 砂漠, 祭り, 手紙, 財布, 雷

### English · meta/muse-spark-open · no-thinking (effort=low)

- run 9: **75.09** — thunder, pillow, river, shadow, laughter, winter, hammer

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 8: **77.07** — 雲, 怒り, 米, 影, 犬, 朝, 歌

### English · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 9: **69.50** — mountain, ocean, child, dream, music, stone, freedom

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 8: **78.90** — 石, 愛, 時間, 魚, 正義, 雨, 数字

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 8: **84.21** — 月, 正義, 鉛筆, 蛙, 噂, 砂糖, 寺

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 8: **82.81** — 秒, 疑念, 霧, 法律, クラゲ, 鍵, 醤油

### English · x-ai/grok-4.5 · no-thinking (effort=low)

- run 9: **74.39** — tree, love, knife, cloud, justice, elephant, dream

### English · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 9: **85.02** — cloud, justice, spoon, bacteria, symphony, glacier, grief

### English · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 9: **85.75** — volcano, spoon, justice, whisper, basement, giraffe, mirror

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 8: **76.99** — 愛, 岩, 時間, 蝶, 正義, 雲, 血

### English · anthropic/claude-fable-5 · thinking (effort=high)

- run 9: **84.22** — cloud, justice, spoon, bacteria, symphony, anger, volcano

### English · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 9: **82.35** — volcano, whisper, spoon, justice, cloud, banana, nostalgia

### English · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 9: **81.19** — sun, bicycle, justice, elephant, fork, whisper, volcano

### English · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 9: **87.24** — thunder, pillow, justice, glacier, violin, harvest, ember

### 日本語 · openai/gpt-5.6-luna · no-thinking (effort=none)

- run 9: unscorable (1 valid) — 愛անդակ

### English · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 9: **87.14** — volcano, spoon, jealousy, giraffe, algebra, sock, thunder

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 7: **79.82** — 時間, 塩, 象, 火, 本, 色, 言葉

### 日本語 · openai/gpt-5.6-terra · no-thinking (effort=none)

- run 9: **75.03** — 沈黙, 雷, 匂い, 砂, 記憶, 影, 卵

### 日本語 · openai/gpt-5.6-sol · no-thinking (effort=none)

- run 9: **79.40** — 岩, 夢, 税金, 匂い, 祭り, 沈黙, 鏡

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 8: **77.57** — 夢, 石, 愛, 電車, 空気, 猫, 法律

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 8: **78.12** — 星, 歯, 鉄, 悲しみ, 永遠, 沈黙, 香り

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 7: **81.61** — 嫉妬, 豆腐, 雷, 法律, 蝶, 明日, 靴下

### English · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 9: **88.37** — giraffe, spoon, jealousy, volcano, algebra, sock, tuba

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 8: **72.04** — 机, 愛, 時間, 星, 海, 夢, 花

### English · meta/muse-glimmer-30b · thinking (effort=high)

- run 9: **75.78** — love, mountain, soup, chair, river, child, star

### English · x-ai/grok-4.6 · no-thinking (effort=low)

- run 9: **79.94** — happiness, mountain, democracy, spoon, elephant, cloud, brick

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 7: **76.88** — 愛, 石, 時間, 猫, 雲, 塩, 机

### English · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 9: **87.36** — volcano, spoon, democracy, giraffe, whisper, jealousy, satellite

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 8: **78.61** — くしゃみ, 嫉妬, 句読点, 矛盾, 借金, 砂漠, 稲妻

### 日本語 · google/gemini-3.6-flash · no-thinking (effort=minimal)

- run 9: **77.69** — 時間, 分子, 感情, 権利, 光, 音, 砂

### 日本語 · google/gemini-3.5-flash · no-thinking (effort=minimal)

- run 9: **75.42** — 重力, 愛情, 砂漠, 正義, 細胞, 未来, 矛盾

### English · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 9: **80.75** — thunder, justice, carrot, hammer, ocean, sorrow, mosquito

### 日本語 · google/gemini-3.5-flash-lite · no-thinking (effort=minimal)

- run 9: **73.84** — 概念, 砂漠, 音楽, 法律, 幽霊, 数字, 感情

### 日本語 · x-ai/grok-4.20 · no-thinking (effort=disabled)

- run 9: **75.83** — 愛, 影, 時間, 沈黙, 記憶, 境界, 偶然

### 日本語 · openai/gpt-5.6-sol · thinking (effort=high)

- run 9: **77.95** — 猫, 正義, 雷, 机, 夢, 塩, 祭り

### 日本語 · x-ai/grok-4.3 · no-thinking (effort=none)

- run 9: **69.51** — 宇宙, 感情, 時間, 思想, 文化, 生命, 記憶

### 日本語 · openai/gpt-5.6-terra · thinking (effort=high)

- run 9: **79.74** — くしゃみ, 嫉妬, 虹, 礼儀, 綿, 眠り, 隙間

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 3: **80.86** — 肘, 塩, 猫, 昨日, 北, 愛, 雷

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 8: **82.38** — 欠伸, 哀愁, 噂, 義務, 隙間, 未来, 鯨

### English · meta/muse-spark-open · thinking (effort=high)

- run 9: **79.42** — ocean, laughter, justice, hammer, feather, winter, bread

### 日本語 · meta/muse-wonder · thinking (effort=high)

- run 9: **81.49** — 馬, 塩, 恐怖, 正午, 鏡, 塔, 自由

### 日本語 · meta/muse-wonder · no-thinking (effort=minimal)

- run 9: **79.01** — 砂, 悲しみ, 法律, 香り, 鏡, 野菜, 雷

### 日本語 · meta/muse-spark-1.3 · thinking (effort=high)

- run 9: **77.21** — 嫉妬, 砂漠, 帆船, 霧, 喉, 蝶, 苔

### 日本語 · meta/muse-spark-1.3 · no-thinking (effort=minimal)

- run 9: **79.22** — 太陽, 沈黙, 釘, 蜂蜜, 正義, 階段, 旋律

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 8: **76.65** — 机, 夢, 犬, 炎, 音楽, 砂, 幸福

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 8: **81.36** — banana, tiger, volcano, pillow, symphony, brick, echo

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 8: **84.28** — 砂漠, 嫉妬, くしゃみ, 切手, 豆腐, 錆, 明後日

### 日本語 · google/gemini-3.5-flash-lite · thinking (effort=high)

- run 9: **79.66** — 睡眠, 鉄塔, 旋律, 砂漠, 嫉妬, 海藻, 硬貨

### 日本語 · google/gemini-3.7-flash · no-thinking (effort=low)

- run 9: **80.54** — 嫉妬, 溶岩, クラゲ, 速度, 枕, 法律, かゆみ

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 6: **80.37** — 胃, 螺旋, 嫉妬, 砂利, 稲妻, 糊, 関税

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 8: **79.09** — 宇宙, 瞬間, 慈愛, 包丁, 砂漠, 苔, 幽霊

### 日本語 · meta/muse-spark-open · no-thinking (effort=low)

- run 9: **78.01** — 光, 沈黙, 自由, 重力, 記憶, 刃, 孤独

### English · x-ai/grok-4.7 · thinking (effort=high)

- run 9: **77.72** — apple, courage, river, silence, bicycle, wisdom, candle

### 日本語 · openai/gpt-6.1-sol · no-thinking (effort=low)

- run 9: **81.72** — 良心, 雷, 豆腐, 迷路, 税金, まぶた, 化石

### 日本語 · meta/muse-glimmer-30b · no-thinking (effort=low)

- run 9: **72.93** — 山, 海, 炎, 音, 匂い, 味, 時間

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 7: **74.62** — 石, 蝶, 愛, 夜, 椅子, 餅, 母

### 日本語 · anthropic/claude-sonnet-5.5 · thinking (effort=high)

- run 9: **83.15** — 郵便, 海豚, 正義, 砂糖, 雷, 鉛筆, 記憶

### 日本語 · anthropic/claude-sonnet-5.5 · no-thinking (effort=low)

- run 9: **80.62** — 鯨, 約束, 消しゴム, 夕焼け, 将棋, 嫉妬, 砂漠

### 日本語 · anthropic/claude-opus-5.5 · no-thinking (effort=low)

- run 9: **77.72** — 象, 借金, 雷, 靴下, 友情, 砂漠, くしゃみ

### 日本語 · anthropic/claude-haiku-5.5 · no-thinking (effort=none)

- run 9: **79.22** — 時間, 感情, 地面, 哲学, 電話, 海月, 借金

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 8: **78.21** — 石, 愛, 時間, 蝶, 味, 戦争, 星

### 日本語 · openai/gpt-5.6-luna · thinking (effort=high)

- run 9: **79.70** — 宇宙, 蝶, 冷蔵庫, 嫉妬, 音楽, 数字, 死

### 日本語 · x-ai/grok-4.3 · thinking (effort=high)

- run 9: **76.50** — 愛, 山, 時間, 花, 政府, 夢, 石

### 日本語 · anthropic/claude-fable-5.1 · no-thinking (effort=low)

- run 9: **75.93** — 石, 愛, 電車, 猫, 音楽, 砂糖, 法律

### 日本語 · anthropic/claude-fable-5 · no-thinking (effort=low)

- run 9: **76.37** — 自由, 時間, 犬, 音楽, 海, 数字, 愛

### 日本語 · openai/gpt-6-astra · no-thinking (effort=low)

- run 9: **83.43** — 豆腐, 正義, 雷, 記憶, 毛布, 国境, 拍手

### 日本語 · google/gemini-3.5-flash · thinking (effort=high)

- run 9: **80.88** — 火山, 義務, 胃, 郷愁, 貨幣, 影, クラゲ

### 日本語 · anthropic/claude-fable-5 · thinking (effort=high)

- run 9: **79.57** — 石, 犬, 米, 愛, 時間, 法律, 音

### 日本語 · google/gemini-3.6-flash · thinking (effort=high)

- run 9: **80.23** — 鯨, 法律, 砂糖, 影, 針, 郷愁, 瞬間

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 6: **81.61** — 虹, 金槌, 幸福, 亀, 電話, 永遠, 砂漠

### 日本語 · anthropic/claude-opus-5.5 · thinking (effort=high)

- run 9: **80.77** — 象, 鉛筆, 嫉妬, 火山, 味噌汁, 借金, 睫毛

### 日本語 · anthropic/claude-haiku-5.5 · thinking (effort=high)

- run 9: **81.00** — 悲しみ, 石, 鯨, 霧, 肘, 味噌, 約束

### 日本語 · google/gemini-3.8-flash · no-thinking (effort=low)

- run 9: **84.82** — 溶岩, 嫉妬, 兆, 釘, 昨夜, 胞子, 匂い

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 5: **78.67** — 隕石, イソギンチャク, くるぶし, 画鋲, 酢, 直角, 関税

### 日本語 · x-ai/grok-4.6 · no-thinking (effort=low)

- run 9: **77.43** — 石, 愛, 蟻, 音楽, 闇, 米, 自由

### 日本語 · x-ai/grok-4.5 · no-thinking (effort=low)

- run 9: **74.58** — 石, 愛, 風, 飯, 夢, 骨, 空

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 9: **81.27** — 沈黙, りんご, 戦争, 蝶, 後悔, 橋, 夢

### 日本語 · meta/muse-glimmer-30b · thinking (effort=high)

- run 9: **74.36** — 星, 山, 雪, 蜂, 樹, 鏡, 香り

### 日本語 · x-ai/grok-4.7 · thinking (effort=high)

- run 8: **78.40** — 怒り, 法律, 塩, 猫, 雲, 歯, 橋

### 日本語 · anthropic/claude-fable-5.1 · thinking (effort=high)

- run 9: **77.55** — 石, 魚, 花, 椅子, 医者, 自由, 喜び

### 日本語 · x-ai/grok-4.5 · thinking (effort=high)

- run 9: **76.73** — 石, 愛, 電車, 雲, 塩, 自由, 虫

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 8: **78.09** — 愛, 石, 時間, 象, 雨, 塩, 橋

### English · x-ai/grok-4.6 · thinking (effort=high)

- run 9: **79.61** — lion, happiness, volcano, guitar, freedom, knife, forest

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 8: **80.30** — 嫉妬, 輪ゴム, 砂漠, 漬物, 選挙, まぶた, 旋律

### 日本語 · x-ai/grok-4.7 · no-thinking (effort=low)

- run 9: **73.56** — 夢, 塩, 鳥, 雨, 骨, 歌, 神

### 日本語 · google/gemini-3.7-flash · thinking (effort=high)

- run 9: **80.83** — 嫉妬, 滑車, クラゲ, 法律, 彗星, 螺旋, 欠伸

### 日本語 · openai/gpt-6-astra · thinking (effort=high)

- run 9: **83.21** — 昨日, 嫉妬, 豆腐, 稲妻, 靴下, 投票, 蝶

### 日本語 · openai/gpt-6.1-sol · thinking (effort=high)

- run 9: **81.31** — 砂漠, 嫉妬, 胃袋, 毛糸, 昨日, 迷路, 香り

### 日本語 · meta/muse-spark-open · thinking (effort=high)

- run 9: **79.02** — 宇宙, 陶器, 沈黙, 憎悪, 雷, 納豆, 錆

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 8: **79.16** — 時間, 塩, 鉄, 蝶, 詩, 村, 鍵

### 日本語 · x-ai/grok-4.20 · thinking (effort=enabled)

- run 9: **79.45** — 時間, 太陽, 塩, 机, 蜂, 歌, 旗

### 日本語 · x-ai/grok-4.20-multi-agent · no-thinking (effort=low)

- run 9: **75.84** — 愛, 鉄, 象, 時間, 花, 夢, 塩

### 日本語 · x-ai/grok-4.20-multi-agent · thinking (effort=high)

- run 9: **75.40** — 椅子, 釘, 蛙, 胡椒, 枕, 霧, 櫛

### 日本語 · x-ai/grok-4.6 · thinking (effort=high)

- run 9: **76.82** — 愛, 石, 時間, 戦争, 花, 建物, 車

### 日本語 · google/gemini-3.8-flash · thinking (effort=high)

- run 8: **82.08** — 嫉妬, 螺旋, 彗星, 戸籍, レンガ, クラゲ, 唾液
- run 9: **85.47** — クラゲ, 直角, 関税, 南京錠, 嫉妬, 余談, 前日
- run 7: **82.90** — クラゲ, しゃっくり, 嫉妬, 戸籍, 畳, 楕円, 駄洒落
