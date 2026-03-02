package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const openrouterAPIURL = "https://openrouter.ai/api/v1/chat/completions"

type OpenRouterClient struct {
	apiKey string
	model  string
	client *http.Client
}

func NewOpenRouterClient(detail ProviderDetail) *OpenRouterClient {
	return &OpenRouterClient{
		apiKey: detail.APIKey,
		model:  detail.Model,
		client: &http.Client{Timeout: 5 * time.Minute},
	}
}

func (c *OpenRouterClient) Stream(ctx context.Context, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error) {
	var msgs []openaiMsg

	if opts.SystemPrompt != "" {
		msgs = append(msgs, openaiMsg{Role: "system", Content: opts.SystemPrompt})
	}
	for _, m := range messages {
		msgs = append(msgs, openaiMsg{Role: m.Role, Content: m.Content})
	}

	maxTokens := opts.MaxTokens
	if maxTokens == 0 {
		maxTokens = 4096
	}

	reqBody := openaiRequest{
		Model:         c.model,
		Messages:      msgs,
		MaxTokens:     maxTokens,
		Temperature:   opts.Temperature,
		Stream:        true,
		StreamOptions: &openaiStreamOp{IncludeUsage: true},
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, openrouterAPIURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Set("HTTP-Referer", "https://pearcode.dev")
	req.Header.Set("X-Title", "Pear")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &LLMError{Code: "network", Message: err.Error(), Retry: true}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, mapHTTPError(resp)
	}

	return parseOpenAISSEStream(resp.Body, onChunk)
}
