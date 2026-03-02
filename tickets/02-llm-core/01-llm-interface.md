# 02-01: LLM Client Interface

## Summary
Define the common interface all LLM providers implement, plus shared types.

## Event Model Refs
- E6c, E7, E8: all use LLM.Stream()

## Files to Create
- `cli/llm/client.go`

## Types

```go
type LLMClient interface {
    Stream(ctx context.Context, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error)
}

type Message struct {
    Role    string // "user", "assistant"
    Content string
}

type StreamOptions struct {
    SystemPrompt string
    MaxTokens    int
    Temperature  float64
}

type Response struct {
    Content      string
    InputTokens  int
    OutputTokens int
}

type LLMError struct {
    Code    string        // "rate_limit", "auth", "network", "unknown"
    Message string
    Retry   bool
    After   time.Duration
}
```

## Functions
- `NewClient(provider string, detail config.ProviderDetail) (LLMClient, error)` — factory that returns the right client based on provider name

## Acceptance Criteria
- Interface compiles
- `NewClient("anthropic", ...)`, `NewClient("openai", ...)`, `NewClient("openrouter", ...)` return correct types (once provider tickets are done, initially return error)
- `LLMError` implements `error` interface

## Dependencies
- 01-02 (config types for ProviderDetail)

## Notes
- System prompt is in StreamOptions, NOT in messages. Each provider handles placement internally.
- `NewClient` is used in E1 bootstrap and when /provider or /model triggers LLM.Reinit
