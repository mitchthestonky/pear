# 02-03: OpenAI Client

## Summary
OpenAI Chat Completions client with SSE streaming.

## Event Model Refs
- E6c, E7, E8: LLM.Stream()

## Files to Create
- `cli/llm/openai.go`

## Implementation
- `POST https://api.openai.com/v1/chat/completions`
- Headers: `Authorization: Bearer {key}`, `content-type: application/json`
- System prompt goes as `{ role: "system", content: systemPrompt }` prepended to messages array
- Request body: `{ model, messages, max_tokens, temperature, stream: true, stream_options: { include_usage: true } }`
- SSE parsing: `chat.completion.chunk` events → extract `choices[0].delta.content` → call `onChunk`
- Final chunk has `usage` field → extract token counts
- Error mapping: 401 → auth, 429 → rate_limit, 5xx → network

## Acceptance Criteria
- Implements `LLMClient` interface
- Streams response with valid key (manual test)
- Returns typed `LLMError` for failures
- Returns token counts

## Dependencies
- 02-01 (LLM interface)

## Notes
- Can be developed in parallel with 02-02 and 02-04
