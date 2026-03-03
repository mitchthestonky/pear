# 02-04: OpenRouter Client

## Summary
OpenRouter client with SSE streaming. OpenAI-compatible format with OpenRouter-specific headers.

## Event Model Refs
- E6c, E7, E8: LLM.Stream()

## Files to Create
- `cli/llm/openrouter.go`

## Implementation
- `POST https://openrouter.ai/api/v1/chat/completions`
- Headers: `Authorization: Bearer {key}`, `HTTP-Referer: https://pearcode.dev`, `X-Title: Pear`, `content-type: application/json`
- System prompt: same as OpenAI (role: "system" in messages)
- Request/response format: identical to OpenAI Chat Completions
- SSE parsing: same as OpenAI client
- Error mapping: same codes

## Acceptance Criteria
- Implements `LLMClient` interface
- Streams response with valid OpenRouter key (manual test)
- OpenRouter-specific headers included in requests
- Returns typed `LLMError` for failures

## Dependencies
- 02-01 (LLM interface)

## Notes
- Consider extracting shared OpenAI-format SSE parsing into a helper used by both 02-03 and 02-04, since the response format is identical. Only the URL and headers differ.
