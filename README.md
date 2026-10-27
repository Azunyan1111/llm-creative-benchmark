# llm-creative-benchmark

LLM の創造性を、Divergent Association Task（DAT、発散的連想課題）で測るベンチマークです。OpenRouter 経由で 26 モデルに「互いに意味が離れた名詞を 10 個」挙げさせ、単語の埋め込み同士の距離からスコアを計算し、思考あり・思考なしの 2 つのアームと英語版・日本語版の 2 言語で順位を付けます。手法の詳細と使い方は、結果のあとにあります。

## 結果

2026 年 10 月 27 日に 26 モデル × 2 言語 × 2 アーム × 10 回の 1,040 試行を計測しました（費用 $15.71）。各アームの上位 10 モデル（平均 ± SD、10 回）は次のとおりです。スコアは人間の平均が約 78、95 が非常に高い値です。

| 順位 | 英語・思考あり | 英語・思考なし | 日本語・思考あり | 日本語・思考なし |
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

全体の表、順位表、各試行で実際に挙がった単語は [results/ranking.md](results/ranking.md) にあります。集計値は [results/summary.json](results/summary.json)、生データは [results/trials.jsonl](results/trials.jsonl) です。

読むときの注意点です。

- 英語ではほとんどのモデルが人間の平均（約 78）を上回り、上位は 86〜89 に集中しています。日本語は全体に 5〜7 点低く、言語間で比べられない前提どおりの結果です。
- 思考ありと思考なしの差は多くのモデルで SD の範囲に収まっており、この課題では思考の有無が創造性スコアをほとんど変えません。
- `meta/muse-wonder` と `meta/muse-spark-1.3` は、最初の計測では OpenRouter アカウントの ZDR（Zero Data Retention）設定でエンドポイントが除外され、全試行がエラーでした。設定を切ってから同じ `run` コマンドを再実行し、失敗した 80 試行だけを再試行しています。Muse の中では `muse-wonder` が最も高く、英語の思考ありで 85.3（10 位）でした。
- `claude-sonnet-5.5` の思考ありは、effort を `high` にしても reasoning トークンが 0 と報告されています。`claude-opus-5.5` と `claude-fable-5` の思考なしは 10 回のうち 4 通りの単語リストしか出さず、SD が 0.1〜0.3 と極端に小さくなっています。
- `google/gemini-4-argon`（Gemini 4 Argon）は、ほかの 25 モデルを計測したあとに追加し、同じ `run` コマンドで 40 試行を計測しました（費用 $1.46）。Argon は reasoning が必須で、effort は `low`・`medium`・`high` だけを受け付けるため、思考なしは `low` で測っています。`low` では reasoning トークンが 0 と報告されました。英語の思考ありは 90.3 で 1 位、日本語の思考なしは 83.5 で 1 位です。ただし思考ありは 20 回のうち 5 回が採点不能でした。4 回（英語 1 回、日本語 3 回）は推論だけで 16,000 トークンの上限を使い切って本文が空になり、1 回（日本語）はプロバイダが finish_reason `error` を返して本文がありませんでした。日本語の思考ありの平均 80.5 は残り 6 回の値です。
- 採点不能の試行が 7 件あり、平均から除いています。Argon の 5 件のほか、日本語の `gpt-5.6-luna` 思考なしは単語ではない文字列を返し、日本語の `meta/muse-spark-1.3` 思考ありは推論だけで 16,000 トークンの上限を使い切り、本文が空でした。

## 測定方法

DAT は Olson らが 2021 年に PNAS で発表した創造性の指標です（Olson JA, Nahas J, Chmoulevitch D, Cropper SJ, Webb ME. Naming unrelated words predicts creativity. PNAS 118(25): e2022340118. doi:10.1073/pnas.2022340118）。回答者が挙げた単語のうち有効な先頭 7 語を取り、21 通りの組み合わせすべてについて単語ベクトルのコサイン距離を求め、その平均に 100 を掛けた値がスコアです。離れた単語を挙げられるほどスコアが高くなり、人間の平均は約 78（SD 6.35）、50 は低く、95 は非常に高い値とされています。

採点の手順は公式実装（[jayolson/divergent-association-task](https://github.com/jayolson/divergent-association-task)、MIT ライセンス）を Go に移植したものです。単語の正規化、複合語の扱い（cul de sac → cul-de-sac）、辞書による有効語の判定、先頭 7 語の選び方をそのまま再現しています。論文の例に対する本実装のスコアは次のとおりです。

| 単語リスト | 論文 | 本実装 |
|---|---:|---:|
| arm, eyes, feet, hand, head, leg, body | 50 | 50.31 |
| bag, bee, burger, feast, office, shoes, tree | 78 | 77.90 |
| hippo, jumper, machinery, prickle, tickets, tomato, violin | 95 | 95.22 |

### 言語ごとの埋め込み

| 言語 | 埋め込み | 辞書 |
|---|---|---|
| 英語 | GloVe 840B-300d（論文と同じ） | 公式実装の words.txt（Hunspell 由来、assets/ に同梱） |
| 日本語 | fastText cc.ja.300（Common Crawl と Wikipedia で学習） | fastText 語彙のうち、かな・漢字だけからなる上位 30 万語 |

日本語には GloVe がないため、同じ Common Crawl 系の fastText を使います。距離の分布が言語ごとに異なるので、スコアは同じ言語の中でだけ比較してください。

### プロンプト

論文で参加者に示した指示文をそのまま使い、機械的に解析できるように「1 行に 1 語」という出力形式の指定だけを足しています。日本語版は同じ指示文の日本語訳です。人間向けの 4 分という制限時間は使いません。

### 思考あり・思考なしの定義

OpenRouter の `reasoning.effort` で制御します。

| アーム | effort |
|---|---|
| thinking | `high` |
| no-thinking | `none`。使えないモデルは `minimal`、それもなければそのモデルが許す最低の effort（`low`） |

Claude 5.5 系、GPT-6 系、Gemini 3.5 以降の Flash、Grok 4.5 以降、Meta Muse は推論を完全には切れません（OpenRouter のカタログで `mandatory: true`）。これらの no-thinking は「最低 effort」であり、実際に使った effort と reasoning トークン数を結果に記録しています。effort の段階を持たず on/off だけのモデル（grok-4.20）は `enabled: true/false` で切り替えます。割り当ては `datbench plan` で確認できる仕組みです。

温度などのサンプリング設定は送らず、各プロバイダの既定値に任せています。

## 対象モデル

| ファミリー | モデル（OpenRouter ID） |
|---|---|
| GPT-5.6 | openai/gpt-5.6-luna, gpt-5.6-sol, gpt-5.6-terra |
| GPT-6 / 6.1 | openai/gpt-6-astra, gpt-6.1-sol |
| Gemini Flash | google/gemini-3.8-flash, 3.7-flash, 3.6-flash, 3.5-flash, 3.5-flash-lite |
| Gemini Argon | google/gemini-4-argon |
| Grok 4.2 以上 | x-ai/grok-4.20, grok-4.20-multi-agent, grok-4.3, grok-4.5, grok-4.6, grok-4.7 |
| Meta Muse | meta/muse-wonder, muse-spark-1.3, muse-spark-open, muse-glimmer-30b |
| Claude | anthropic/claude-fable-5, claude-fable-5.1, claude-opus-5.5, claude-sonnet-5.5, claude-haiku-5.5 |

## セットアップ

Go 1.27 以上と、環境変数 `OPENROUTER_API_KEY` が必要です。

```bash
go build -o bin/datbench ./cmd/datbench
```

埋め込みファイルを `data/` に置き、各言語の採点用ファイルを作ります。

```bash
curl -L -o data/glove/glove.840B.300d.zip https://huggingface.co/stanfordnlp/glove/resolve/main/glove.840B.300d.zip
curl -L -o data/fasttext/cc.ja.300.vec.gz https://dl.fbaipublicfiles.com/fasttext/vectors-crawl/cc.ja.300.vec.gz
./bin/datbench prepare --lang en
./bin/datbench prepare --lang ja
```

`data/` は Git 管理の対象外です。

## 使い方

```bash
./bin/datbench plan                      # モデルごとの effort 割り当てを表示
./bin/datbench run --runs 10             # 全モデル × 2 言語 × 2 アーム × 10 回を results/trials.jsonl に追記
./bin/datbench report                    # results/ranking.md と results/summary.json を生成
./bin/datbench score cat dog --min 2     # 任意の単語リストを採点（--lang ja で日本語）
```

`run` は途中で止めても、同じファイルに対して再実行すれば未完了の組み合わせだけを続けます。`--models` でモデル ID を、`--family` でファミリー名を、`--lang` で言語を絞れます。

## 構成

クリーンアーキテクチャ風に、依存の向きを domain ← usecase ← adapter ← cmd の一方向にそろえています。

| 層 | パッケージ | 役割 |
|---|---|---|
| domain | `internal/domain/dat` | DAT の採点（公式実装の移植）と言語別の単語正規化 |
| domain | `internal/domain/bench` | 言語・モード・モデル・試行の型、effort の決定、統計、プロンプト |
| usecase | `internal/usecase` | ポート（Generator, Catalog, VectorRepository, TrialRepository など）と、計測・集計・埋め込み準備の各ユースケース |
| adapter | `internal/adapter/openrouter` | OpenRouter クライアント（Generator と Catalog を実装） |
| adapter | `internal/adapter/vectors` | GloVe zip と fastText gz の読み込み、辞書によるフィルタ、バイナリ保存 |
| adapter | `internal/adapter/jsonl` | 試行の JSONL 保存 |
| adapter | `internal/adapter/markdown` | ランキングの Markdown 出力 |
| cmd | `cmd/datbench` | CLI と依存の組み立て |

各パッケージはテストを先に書いてから実装しています（`go test ./...`）。
