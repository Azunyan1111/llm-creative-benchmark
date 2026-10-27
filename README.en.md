# llm-creative-benchmark

A creativity benchmark for LLMs based on the Divergent Association Task (DAT). Through OpenRouter, 26 models are asked for 10 nouns that are as different from each other as possible; the score is the mean cosine distance between word embeddings of the first 7 valid words × 100. Each model is ranked in two arms (thinking / no-thinking) and two languages (English / Japanese). Method and usage follow the results.

## Results

Measured on 2026-10-27: 26 models × 2 languages × 2 arms × 10 runs = 1,040 trials, $15.71 in total. Top 10 per arm (mean ± SD over 10 runs; humans average ≈ 78, 95 is very high):

| # | EN thinking | EN no-thinking | JA thinking | JA no-thinking |
|---:|---|---|---|---|
| 1 | gemini-4-argon 90.3 ± 1.6 | gemini-3.5-flash 88.0 ± 2.0 | gpt-6-astra 82.0 ± 1.5 | gemini-4-argon 83.5 ± 2.0 |
| 2 | gemini-3.8-flash 89.2 ± 2.8 | gemini-3.6-flash 87.8 ± 4.1 | gemini-3.8-flash 81.5 ± 1.9 | gemini-3.8-flash 82.8 ± 2.6 |
| 3 | claude-opus-5.5 88.1 ± 1.4 | gemini-3.8-flash 87.6 ± 3.0 | gemini-3.5-flash 81.3 ± 1.5 | gpt-6.1-sol 81.6 ± 0.9 |
| 4 | gemini-3.7-flash 87.7 ± 1.2 | claude-opus-5.5 87.3 ± 0.3 | gemini-3.7-flash 81.1 ± 1.9 | gpt-6-astra 81.6 ± 1.7 |
| 5 | gpt-6.1-sol 87.3 ± 2.0 | gemini-3.7-flash 87.1 ± 2.9 | gpt-6.1-sol 80.7 ± 1.9 | gemini-3.7-flash 80.6 ± 1.7 |
| 6 | claude-sonnet-5.5 86.6 ± 2.3 | gemini-4-argon 86.4 ± 1.7 | gemini-3.6-flash 80.7 ± 2.3 | claude-sonnet-5.5 80.5 ± 1.6 |
| 7 | gemini-3.6-flash 86.6 ± 2.1 | gpt-6-astra 86.3 ± 2.5 | gpt-5.6-sol 80.6 ± 2.0 | claude-opus-5.5 79.7 ± 1.4 |
| 8 | gpt-6-astra 86.4 ± 1.5 | claude-sonnet-5.5 86.1 ± 2.9 | gemini-4-argon 80.5 ± 1.8 | gpt-5.6-luna 79.6 ± 1.0 |
| 9 | claude-fable-5.1 86.1 ± 1.5 | claude-fable-5.1 85.9 ± 2.3 | muse-wonder 80.5 ± 1.4 | claude-fable-5 79.5 ± 2.5 |
| 10 | muse-wonder 85.3 ± 3.4 | gpt-6.1-sol 85.2 ± 2.4 | claude-fable-5.1 80.5 ± 2.2 | gemini-3.6-flash 78.3 ± 1.6 |

Full tables, per-arm rankings and the words each model produced are in [results/ranking.md](results/ranking.md); aggregates in [results/summary.json](results/summary.json); raw data in [results/trials.jsonl](results/trials.jsonl).

Caveats: most models beat the human average (≈ 78) in English; Japanese scores run 5–7 points lower and are not comparable across languages. Thinking vs. no-thinking differences are mostly within one SD. `meta/muse-wonder` and `meta/muse-spark-1.3` failed every trial in the first pass because the account's ZDR (Zero Data Retention) setting excluded their only endpoint; after the setting was switched off, re-running `run` retried only those 80 trials (`muse-wonder` is the strongest Muse model: 85.3 in English thinking, rank 10). `claude-sonnet-5.5` reported 0 reasoning tokens even at effort `high`; `claude-opus-5.5` and `claude-fable-5` without thinking produced only 4 distinct word lists in 10 runs (SD 0.1–0.3). `google/gemini-4-argon` (Gemini 4 Argon) was added after the other 25 models and measured with the same `run` command (40 trials, $1.46). Its reasoning is mandatory and only `low`/`medium`/`high` are accepted, so the no-thinking arm uses `low`, which reported 0 reasoning tokens. It ranks first in English thinking (90.3) and in Japanese no-thinking (83.5), but 5 of its 20 thinking trials are unscorable: 4 (1 English, 3 Japanese) spent the whole 16,000-token budget on reasoning and returned empty content, and 1 Japanese trial came back from the provider with finish_reason `error` and no content; its Japanese thinking mean (80.5) is over the remaining 6 runs. Seven trials are unscorable in total: the 5 Argon trials above, a Japanese `gpt-5.6-luna` no-thinking reply that was a non-word string, and a Japanese `meta/muse-spark-1.3` thinking reply that spent the whole 16,000-token budget on reasoning and returned empty content.

## Method

DAT: Olson JA, Nahas J, Chmoulevitch D, Cropper SJ, Webb ME. *Naming unrelated words predicts creativity.* PNAS 118(25): e2022340118 (2021). doi:10.1073/pnas.2022340118. Humans average ≈ 78 (SD 6.35); 50 is low, 95 very high.

The scorer is a Go port of the reference implementation ([jayolson/divergent-association-task](https://github.com/jayolson/divergent-association-task), MIT). It reproduces the paper's examples (50 / 78 / 95 → 50.31 / 77.90 / 95.22).

| Language | Embeddings | Dictionary |
|---|---|---|
| English | GloVe 840B-300d (as in the paper) | reference `words.txt` (Hunspell-derived, in `assets/`) |
| Japanese | fastText cc.ja.300 (Common Crawl + Wikipedia) | the 300k most frequent kana/kanji-only tokens of the fastText vocabulary |

Scores are comparable within a language only.

**Arms.** `thinking` sends `reasoning.effort = high`. `no-thinking` sends `none` where the provider allows switching reasoning off, else `minimal`, else the weakest allowed effort (`low`). Many models (Claude 5.5, GPT-6, Gemini 3.5+ Flash, Grok 4.5+, Meta Muse) cannot disable reasoning at all; the effort actually used and the reasoning token count are recorded with every trial. `datbench plan` prints the assignment.

The prompt is the paper's participant instruction plus a one-word-per-line output format; the Japanese prompt is a translation. No temperature or sampling parameters are sent.

## Models

GPT-5.6 (luna, sol, terra), GPT-6 astra, GPT-6.1 sol, Gemini 3.5–3.8 Flash (+3.5 Flash Lite), Gemini 4 Argon, Grok 4.20 / 4.20-multi-agent / 4.3 / 4.5 / 4.6 / 4.7, Meta Muse (wonder, spark-1.3, spark-open, glimmer-30b), Claude Fable 5 / 5.1, Claude Opus / Sonnet / Haiku 5.5.

## Setup and usage

Requires Go ≥ 1.27 and `OPENROUTER_API_KEY`.

```bash
go build -o bin/datbench ./cmd/datbench
curl -L -o data/glove/glove.840B.300d.zip https://huggingface.co/stanfordnlp/glove/resolve/main/glove.840B.300d.zip
curl -L -o data/fasttext/cc.ja.300.vec.gz https://dl.fbaipublicfiles.com/fasttext/vectors-crawl/cc.ja.300.vec.gz
./bin/datbench prepare --lang en
./bin/datbench prepare --lang ja
./bin/datbench plan
./bin/datbench run --runs 10        # resumable; --models, --family, --lang to filter
./bin/datbench report               # results/ranking.md, results/summary.json
./bin/datbench score cat dog --min 2
```

## Layout

Clean-architecture style, dependencies pointing inward (domain ← usecase ← adapter ← cmd); every package was written test-first.

- `internal/domain/dat` – DAT scoring, English/Japanese word cleaning
- `internal/domain/bench` – entities, effort planning, statistics, prompts
- `internal/usecase` – ports and the run / report / prepare use cases
- `internal/adapter/{openrouter,vectors,jsonl,markdown}` – implementations of the ports
- `cmd/datbench` – CLI composition root
