package openrouter

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"

	"github.com/Azunyan1111/llm-creative-benchmark/internal/domain/bench"
	"github.com/Azunyan1111/llm-creative-benchmark/internal/usecase"
)

func TestGenerateSendsReasoningAndMapsUsage(t *testing.T) {
	var gotBody map[string]any
	var gotAuth string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotAuth = r.Header.Get("Authorization")
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &gotBody)
		w.Write([]byte(`{"id":"x","model":"m","provider":"P","choices":[{"finish_reason":"stop","message":{"role":"assistant","content":"cat\ndog","reasoning":"hmm"}}],"usage":{"prompt_tokens":10,"completion_tokens":50,"cost":0.0012,"completion_tokens_details":{"reasoning_tokens":40}}}`))
	}))
	defer srv.Close()
	c := New("KEY", srv.URL)
	g, err := c.Generate(context.Background(), usecase.GenerateRequest{ModelID: "m", Prompt: "p", Effort: "high", MaxTokens: 99})
	if err != nil {
		t.Fatal(err)
	}
	if gotAuth != "Bearer KEY" {
		t.Errorf("auth: %q", gotAuth)
	}
	if gotBody["model"] != "m" || gotBody["max_tokens"].(float64) != 99 {
		t.Errorf("body: %v", gotBody)
	}
	if r := gotBody["reasoning"].(map[string]any); r["effort"] != "high" {
		t.Errorf("reasoning: %v", r)
	}
	if u := gotBody["usage"].(map[string]any); u["include"] != true {
		t.Errorf("usage include: %v", gotBody["usage"])
	}
	msgs := gotBody["messages"].([]any)
	if len(msgs) != 1 || msgs[0].(map[string]any)["content"] != "p" {
		t.Errorf("messages: %v", msgs)
	}
	if g.Content != "cat\ndog" || g.Reasoning != "hmm" || g.Provider != "P" || g.FinishReason != "stop" ||
		g.PromptTokens != 10 || g.CompletionTokens != 50 || g.ReasoningTokens != 40 || g.CostUSD != 0.0012 {
		t.Errorf("generation: %+v", g)
	}
}

func TestGenerateOmitsReasoningWhenEffortEmpty(t *testing.T) {
	var gotBody map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		json.Unmarshal(b, &gotBody)
		w.Write([]byte(`{"choices":[{"message":{"content":"x"}}],"usage":{}}`))
	}))
	defer srv.Close()
	c := New("KEY", srv.URL)
	if _, err := c.Generate(context.Background(), usecase.GenerateRequest{ModelID: "m", Prompt: "p"}); err != nil {
		t.Fatal(err)
	}
	if _, ok := gotBody["reasoning"]; ok {
		t.Errorf("reasoning should be omitted: %v", gotBody)
	}
}

func TestGenerateRetriesTransientAndSurfacesEmbeddedErrors(t *testing.T) {
	var n int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch atomic.AddInt32(&n, 1) {
		case 1:
			w.WriteHeader(429)
			w.Write([]byte(`{"error":{"code":429,"message":"slow down"}}`))
		case 2:
			w.Write([]byte(`{"error":{"code":502,"message":"upstream"}}`)) // HTTP 200 with error body
		default:
			w.Write([]byte(`{"choices":[{"message":{"content":"ok"}}],"usage":{}}`))
		}
	}))
	defer srv.Close()
	c := New("KEY", srv.URL)
	c.Backoff = 0
	g, err := c.Generate(context.Background(), usecase.GenerateRequest{ModelID: "m", Prompt: "p"})
	if err != nil || g.Content != "ok" || atomic.LoadInt32(&n) != 3 {
		t.Errorf("g=%+v err=%v calls=%d", g, err, n)
	}

	srv2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(400)
		w.Write([]byte(`{"error":{"code":400,"message":"bad effort"}}`))
	}))
	defer srv2.Close()
	c2 := New("KEY", srv2.URL)
	c2.Backoff = 0
	if _, err := c2.Generate(context.Background(), usecase.GenerateRequest{ModelID: "m", Prompt: "p"}); err == nil {
		t.Error("400 must not be retried into success")
	}
}

func TestLookupMapsCatalogue(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/models" {
			w.WriteHeader(404)
			return
		}
		w.Write([]byte(`{"data":[{"id":"a","reasoning":{"mandatory":true,"supported_efforts":["high","low"]}},{"id":"b"}]}`))
	}))
	defer srv.Close()
	c := New("", srv.URL)
	mi, found, err := c.Lookup(context.Background(), "a")
	if err != nil || !found || !mi.Mandatory || len(mi.SupportedEfforts) != 2 {
		t.Errorf("a: %+v %v %v", mi, found, err)
	}
	mi, found, err = c.Lookup(context.Background(), "b")
	if err != nil || !found || mi.Mandatory || len(mi.SupportedEfforts) != 0 {
		t.Errorf("b: %+v %v %v", mi, found, err)
	}
	if _, found, _ := c.Lookup(context.Background(), "zzz"); found {
		t.Error("zzz should not be found")
	}
}

func TestGenerateMapsToggleEfforts(t *testing.T) {
	var bodies []map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		var m map[string]any
		json.Unmarshal(b, &m)
		bodies = append(bodies, m)
		w.Write([]byte(`{"choices":[{"message":{"content":"x"}}],"usage":{}}`))
	}))
	defer srv.Close()
	c := New("KEY", srv.URL)
	c.Generate(context.Background(), usecase.GenerateRequest{ModelID: "m", Prompt: "p", Effort: bench.EffortEnabled})
	c.Generate(context.Background(), usecase.GenerateRequest{ModelID: "m", Prompt: "p", Effort: bench.EffortDisabled})
	on := bodies[0]["reasoning"].(map[string]any)
	off := bodies[1]["reasoning"].(map[string]any)
	if on["enabled"] != true || on["effort"] != nil || off["enabled"] != false || off["effort"] != nil {
		t.Errorf("toggle mapping wrong: on=%v off=%v", on, off)
	}
}

func TestLookupDetectsToggleOnlyReasoning(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"data":[{"id":"g","supported_parameters":["reasoning","max_tokens"],"reasoning":{"mandatory":false,"default_enabled":false}},{"id":"plain","supported_parameters":["max_tokens"]}]}`))
	}))
	defer srv.Close()
	c := New("", srv.URL)
	mi, _, _ := c.Lookup(context.Background(), "g")
	if !mi.ReasoningToggle || len(mi.SupportedEfforts) != 0 {
		t.Errorf("g: %+v", mi)
	}
	mi, _, _ = c.Lookup(context.Background(), "plain")
	if mi.ReasoningToggle {
		t.Errorf("plain: %+v", mi)
	}
}
