// Package openrouter implements usecase.Generator and usecase.Catalog on the
// OpenRouter chat completions API.
package openrouter

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

// DefaultBaseURL is the public API root.
const DefaultBaseURL = "https://openrouter.ai/api/v1"

// Client calls OpenRouter.
type Client struct {
	APIKey      string
	BaseURL     string
	HTTP        *http.Client
	MaxAttempts int           // first try included; retries on 429/5xx/network
	Backoff     time.Duration // base delay, doubled per retry

	catOnce sync.Once
	catalog map[string]modelInfo
	catErr  error
}

// New creates a client. An empty baseURL means DefaultBaseURL.
func New(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return &Client{
		APIKey:      apiKey,
		BaseURL:     baseURL,
		HTTP:        &http.Client{Timeout: 10 * time.Minute},
		MaxAttempts: 4,
		Backoff:     2 * time.Second,
	}
}

// NewFromEnv builds a client from OPENROUTER_API_KEY.
func NewFromEnv() (*Client, error) {
	key := os.Getenv("OPENROUTER_API_KEY")
	if key == "" {
		return nil, errors.New("OPENROUTER_API_KEY is not set")
	}
	return New(key, ""), nil
}

type reasoning struct {
	Effort  string `json:"effort,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
}

// reasoningFor maps the plan's effort to OpenRouter's reasoning object.
func reasoningFor(effort string) *reasoning {
	switch effort {
	case "":
		return nil
	case bench.EffortEnabled:
		on := true
		return &reasoning{Enabled: &on}
	case bench.EffortDisabled:
		off := false
		return &reasoning{Enabled: &off}
	}
	return &reasoning{Effort: effort}
}

type message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type request struct {
	Model     string     `json:"model"`
	Messages  []message  `json:"messages"`
	MaxTokens int        `json:"max_tokens,omitempty"`
	Reasoning *reasoning `json:"reasoning,omitempty"`
	Usage     struct {
		Include bool `json:"include"`
	} `json:"usage"`
}

type apiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *apiError) Error() string { return fmt.Sprintf("openrouter error %d: %s", e.Code, e.Message) }

type response struct {
	Provider string `json:"provider"`
	Choices  []struct {
		FinishReason string `json:"finish_reason"`
		Message      struct {
			Content   string `json:"content"`
			Reasoning string `json:"reasoning"`
		} `json:"message"`
	} `json:"choices"`
	Usage struct {
		PromptTokens            int     `json:"prompt_tokens"`
		CompletionTokens        int     `json:"completion_tokens"`
		Cost                    float64 `json:"cost"`
		CompletionTokensDetails struct {
			ReasoningTokens int `json:"reasoning_tokens"`
		} `json:"completion_tokens_details"`
	} `json:"usage"`
	Error *apiError `json:"error,omitempty"`
}

// Generate implements usecase.Generator.
func (c *Client) Generate(ctx context.Context, req usecase.GenerateRequest) (usecase.Generation, error) {
	body := request{
		Model:     req.ModelID,
		Messages:  []message{{Role: "user", Content: req.Prompt}},
		MaxTokens: req.MaxTokens,
	}
	body.Usage.Include = true
	body.Reasoning = reasoningFor(req.Effort)
	raw, err := json.Marshal(body)
	if err != nil {
		return usecase.Generation{}, err
	}
	attempts := c.MaxAttempts
	if attempts < 1 {
		attempts = 1
	}
	var lastErr error
	for i := 0; i < attempts; i++ {
		if i > 0 {
			select {
			case <-ctx.Done():
				return usecase.Generation{}, ctx.Err()
			case <-time.After(c.Backoff * time.Duration(1<<uint(i-1))):
			}
		}
		g, retry, err := c.once(ctx, raw)
		if err == nil {
			return g, nil
		}
		lastErr = err
		if !retry {
			break
		}
	}
	return usecase.Generation{}, lastErr
}

func (c *Client) once(ctx context.Context, raw []byte) (usecase.Generation, bool, error) {
	hr, err := http.NewRequestWithContext(ctx, http.MethodPost, c.BaseURL+"/chat/completions", bytes.NewReader(raw))
	if err != nil {
		return usecase.Generation{}, false, err
	}
	hr.Header.Set("Authorization", "Bearer "+c.APIKey)
	hr.Header.Set("Content-Type", "application/json")
	hr.Header.Set("HTTP-Referer", "https://github.com/Azunyan1111/llm-creative-benchmark")
	hr.Header.Set("X-Title", "llm-creative-benchmark")
	res, err := c.HTTP.Do(hr)
	if err != nil {
		return usecase.Generation{}, true, err
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return usecase.Generation{}, true, err
	}
	if res.StatusCode == 429 || res.StatusCode >= 500 {
		return usecase.Generation{}, true, fmt.Errorf("http %d: %s", res.StatusCode, truncate(data))
	}
	if res.StatusCode != http.StatusOK {
		return usecase.Generation{}, false, fmt.Errorf("http %d: %s", res.StatusCode, truncate(data))
	}
	var out response
	if err := json.Unmarshal(data, &out); err != nil {
		return usecase.Generation{}, false, fmt.Errorf("decode: %w (%s)", err, truncate(data))
	}
	if out.Error != nil {
		// Provider errors can arrive with HTTP 200; treat rate limits and
		// upstream failures as transient.
		return usecase.Generation{}, out.Error.Code == 429 || out.Error.Code >= 500, out.Error
	}
	if len(out.Choices) == 0 {
		return usecase.Generation{}, true, fmt.Errorf("no choices: %s", truncate(data))
	}
	ch := out.Choices[0]
	return usecase.Generation{
		Content:          ch.Message.Content,
		Reasoning:        ch.Message.Reasoning,
		Provider:         out.Provider,
		FinishReason:     ch.FinishReason,
		PromptTokens:     out.Usage.PromptTokens,
		CompletionTokens: out.Usage.CompletionTokens,
		ReasoningTokens:  out.Usage.CompletionTokensDetails.ReasoningTokens,
		CostUSD:          out.Usage.Cost,
	}, false, nil
}

func truncate(b []byte) string {
	if len(b) > 300 {
		return string(b[:300]) + "..."
	}
	return string(b)
}

// --- catalogue ---------------------------------------------------------

type modelInfo struct {
	ID                  string   `json:"id"`
	SupportedParameters []string `json:"supported_parameters"`
	Reasoning           *struct {
		Mandatory        bool     `json:"mandatory"`
		SupportedEfforts []string `json:"supported_efforts"`
	} `json:"reasoning"`
}

// Lookup implements usecase.Catalog; the catalogue is fetched once.
func (c *Client) Lookup(ctx context.Context, id string) (bench.ModelInfo, bool, error) {
	c.catOnce.Do(func() { c.catalog, c.catErr = c.fetchCatalog(ctx) })
	if c.catErr != nil {
		return bench.ModelInfo{}, false, c.catErr
	}
	mi, ok := c.catalog[id]
	if !ok {
		return bench.ModelInfo{}, false, nil
	}
	out := bench.ModelInfo{ID: mi.ID}
	if mi.Reasoning != nil {
		out.Mandatory = mi.Reasoning.Mandatory
		out.SupportedEfforts = mi.Reasoning.SupportedEfforts
	}
	for _, p := range mi.SupportedParameters {
		if p == "reasoning" {
			out.ReasoningToggle = true
		}
	}
	return out, true, nil
}

func (c *Client) fetchCatalog(ctx context.Context) (map[string]modelInfo, error) {
	hr, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/models", nil)
	if err != nil {
		return nil, err
	}
	res, err := c.HTTP.Do(hr)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET /models: http %d", res.StatusCode)
	}
	var out struct {
		Data []modelInfo `json:"data"`
	}
	if err := json.NewDecoder(res.Body).Decode(&out); err != nil {
		return nil, err
	}
	m := make(map[string]modelInfo, len(out.Data))
	for _, mi := range out.Data {
		m[mi.ID] = mi
	}
	return m, nil
}
