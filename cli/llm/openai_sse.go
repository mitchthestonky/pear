package llm

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
)

// openaiSSEChunk represents a chat.completion.chunk SSE event from OpenAI-compatible APIs.
type openaiSSEChunk struct {
	Choices []struct {
		Delta struct {
			Content string `json:"content"`
		} `json:"delta"`
	} `json:"choices"`
	Usage *struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
	} `json:"usage"`
}

// parseOpenAISSEStream parses an OpenAI-compatible SSE stream, calling onChunk for each content delta.
func parseOpenAISSEStream(body interface{ Read([]byte) (int, error) }, onChunk func(string)) (*Response, error) {
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

		var chunk openaiSSEChunk
		if err := json.Unmarshal([]byte(data), &chunk); err != nil {
			continue
		}

		if len(chunk.Choices) > 0 {
			text := chunk.Choices[0].Delta.Content
			if text != "" {
				content.WriteString(text)
				if onChunk != nil {
					onChunk(text)
				}
			}
		}

		if chunk.Usage != nil {
			inputTokens = chunk.Usage.PromptTokens
			outputTokens = chunk.Usage.CompletionTokens
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
