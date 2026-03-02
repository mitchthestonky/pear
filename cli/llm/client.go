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

func NewClient(provider string, detail ProviderDetail) (LLMClient, error) {
	switch provider {
	case "anthropic":
		return NewAnthropicClient(detail), nil
	case "openai":
		return NewOpenAIClient(detail), nil
	case "openrouter":
		return nil, fmt.Errorf("openrouter provider not yet implemented")
	default:
		return nil, fmt.Errorf("unknown provider: %s", provider)
	}
}
