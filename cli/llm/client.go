package llm

import (
	"context"
	"fmt"
	"time"
)

// ProviderDetail holds API key and model for a provider.
// Defined here until the config package is implemented.
type ProviderDetail struct {
	APIKey string `toml:"api_key"`
	Model  string `toml:"model"`
}

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
	Code    string // "rate_limit", "auth", "network", "unknown"
	Message string
	Retry   bool
	After   time.Duration
}

func (e *LLMError) Error() string {
	return fmt.Sprintf("llm error (%s): %s", e.Code, e.Message)
}

// StreamWithRetry wraps a client's Stream call with up to 2 retries for
// retryable errors (rate limits, transient network errors).
func StreamWithRetry(ctx context.Context, client LLMClient, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error) {
	const maxRetries = 2
	var lastErr error
	for attempt := 0; attempt <= maxRetries; attempt++ {
		resp, err := client.Stream(ctx, messages, opts, onChunk)
		if err == nil {
			return resp, nil
		}
		lastErr = err
		llmErr, ok := err.(*LLMError)
		if !ok || !llmErr.Retry {
			return nil, err
		}
		if attempt == maxRetries {
			break
		}
		wait := llmErr.After
		if wait == 0 {
			wait = time.Duration(attempt+1) * 2 * time.Second
		}
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(wait):
		}
	}
	return nil, lastErr
}

func NewClient(provider string, detail ProviderDetail) (LLMClient, error) {
	switch provider {
	case "anthropic":
		return NewAnthropicClient(detail), nil
	case "openai":
		return NewOpenAIClient(detail), nil
	case "openrouter":
		return NewOpenRouterClient(detail), nil
	default:
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}
