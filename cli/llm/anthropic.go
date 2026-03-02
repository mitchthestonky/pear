package llm

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const anthropicAPIURL = "https://api.anthropic.com/v1/messages"

type AnthropicClient struct {
	apiKey string
	model  string
	client *http.Client
}

func NewAnthropicClient(detail ProviderDetail) *AnthropicClient {
	return &AnthropicClient{
		apiKey: detail.APIKey,
		model:  detail.Model,
		client: &http.Client{Timeout: 5 * time.Minute},
	}
}

type anthropicRequest struct {
	Model       string            `json:"model"`
	System      string            `json:"system,omitempty"`
	Messages    []anthropicMsg    `json:"messages"`
	MaxTokens   int               `json:"max_tokens"`
	Temperature float64           `json:"temperature"`
	Stream      bool              `json:"stream"`
}

type anthropicMsg struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (c *AnthropicClient) Stream(ctx context.Context, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error) {
	msgs := make([]anthropicMsg, len(messages))
	for i, m := range messages {
		msgs[i] = anthropicMsg{Role: m.Role, Content: m.Content}
	}

	maxTokens := opts.MaxTokens
	if maxTokens == 0 {
		maxTokens = 4096
	}

	reqBody := anthropicRequest{
		Model:       c.model,
		System:      opts.SystemPrompt,
		Messages:    msgs,
		MaxTokens:   maxTokens,
		Temperature: opts.Temperature,
		Stream:      true,
	}

	body, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, anthropicAPIURL, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &LLMError{Code: "network", Message: err.Error(), Retry: true}
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, mapHTTPError(resp)
	}

	return parseSSEStream(resp.Body, onChunk)
}

func mapHTTPError(resp *http.Response) *LLMError {
	switch resp.StatusCode {
	case 401:
		return &LLMError{Code: "auth", Message: "invalid API key"}
	case 429:
		e := &LLMError{Code: "rate_limit", Message: "rate limited", Retry: true}
		if after := resp.Header.Get("Retry-After"); after != "" {
			if secs, err := strconv.Atoi(after); err == nil {
				e.After = time.Duration(secs) * time.Second
			}
		}
		return e
	default:
		if resp.StatusCode >= 500 {
			return &LLMError{Code: "network", Message: fmt.Sprintf("server error: %d", resp.StatusCode), Retry: true}
		}
		return &LLMError{Code: "unknown", Message: fmt.Sprintf("HTTP %d", resp.StatusCode)}
	}
}

func parseSSEStream(body interface{ Read([]byte) (int, error) }, onChunk func(string)) (*Response, error) {
	scanner := bufio.NewScanner(body)
	var content strings.Builder
	var inputTokens, outputTokens int

	for scanner.Scan() {
		line := scanner.Text()

		if !strings.HasPrefix(line, "data: ") {
			continue
		}
		data := strings.TrimPrefix(line, "data: ")
		if data == "[DONE]" {
			break
		}

		var event struct {
			Type    string `json:"type"`
			Delta   struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"delta"`
			Message struct {
				Usage struct {
					InputTokens  int `json:"input_tokens"`
					OutputTokens int `json:"output_tokens"`
				} `json:"usage"`
			} `json:"message"`
			Usage struct {
				OutputTokens int `json:"output_tokens"`
			} `json:"usage"`
		}
		if err := json.Unmarshal([]byte(data), &event); err != nil {
			continue
		}

		switch event.Type {
		case "content_block_delta":
			if event.Delta.Text != "" {
				content.WriteString(event.Delta.Text)
				if onChunk != nil {
					onChunk(event.Delta.Text)
				}
			}
		case "message_delta":
			if event.Usage.OutputTokens > 0 {
				outputTokens = event.Usage.OutputTokens
			}
		case "message_start":
			if event.Message.Usage.InputTokens > 0 {
				inputTokens = event.Message.Usage.InputTokens
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, &LLMError{Code: "network", Message: fmt.Sprintf("stream read error: %v", err), Retry: true}
	}

	return &Response{
		Content:      content.String(),
		InputTokens:  inputTokens,
		OutputTokens: outputTokens,
	}, nil
}
