# 02-02: Anthropic Client

## Summary
Claude Messages API client with SSE streaming.

## Event Model Refs
- E6c, E7, E8: LLM.Stream()

## Files to Create
- `cli/llm/anthropic.go`

## Implementation
- `POST https://api.anthropic.com/v1/messages`
- Headers: `x-api-key`, `anthropic-version: 2023-06-01`, `content-type: application/json`
- System prompt goes in top-level `system` field (NOT in messages array)
- Request body: `{ model, system, messages, max_tokens, temperature, stream: true }`
- SSE parsing: handle `content_block_delta` events → extract `delta.text` → call `onChunk`
- On `message_stop`: extract `usage.input_tokens`, `usage.output_tokens` → build Response
- Error mapping: 401 → auth, 429 → rate_limit (parse `retry-after` header), 5xx → network

## Acceptance Criteria
- Implements `LLMClient` interface
- Streams a real response when given a valid API key (manual test)
- Returns typed `LLMError` for auth failures and rate limits
- Returns token counts in Response

## Dependencies
- 02-01 (LLM interface)

## Notes
- No external SDK — hand-rolled HTTP client with `net/http`
- Use `bufio.Scanner` with custom split function for SSE lines
