package llm

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const openaiAPIURL = "https://api.openai.com/v1/chat/completions"

type OpenAIClient struct {
	apiKey string
	model  string
	client *http.Client
}

func NewOpenAIClient(detail ProviderDetail) *OpenAIClient {
	return &OpenAIClient{
		apiKey: detail.APIKey,
		model:  detail.Model,
		client: &http.Client{Timeout: 5 * time.Minute},
	}
}

type openaiRequest struct {
	Model         string          `json:"model"`
	Messages      []openaiMsg     `json:"messages"`
	MaxTokens     int             `json:"max_tokens"`
	Temperature   float64         `json:"temperature"`
	Stream        bool            `json:"stream"`
	StreamOptions *openaiStreamOp `json:"stream_options,omitempty"`
}

type openaiMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiStreamOp struct {
	IncludeUsage bool `json:"include_usage"`
}

func (c *OpenAIClient) Stream(ctx context.Context, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error) {
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

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, openaiAPIURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.apiKey)

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
