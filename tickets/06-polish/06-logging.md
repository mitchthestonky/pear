# 06-06: Structured Logging

## Summary
Session-scoped structured logging to `~/.pear/logs/` for dogfooding diagnostics.

## Files to Create
- `cli/logging/logger.go`

## Events to Log

| Event | Fields |
|---|---|
| `llm.request` | provider, model, input_tokens, output_tokens, latency_ms |
| `llm.error` | provider, error_code, retry, message |
| `watcher.trigger` | type, diff_lines, files_changed |
| `watcher.skip` | reason (too_small, cooldown, paused) |
| `context.collect` | diff_lines, tree_files, attached_files, truncated |
| `concept.extract` | concepts_found, relationships_found, parse_failures |

## Implementation
- One log file per session: `~/.pear/logs/{ISO-timestamp}.log`
- JSON lines format (one JSON object per line)
- Keep last 5 session logs, delete older on startup
- `Log(event string, fields map[string]any)` function
- Init with `NewLogger()` at app startup

## Acceptance Criteria
- Log file created on session start
- All event types write valid JSON lines
- Log rotation keeps max 5 files
- Logs include timestamps
- Logger is safe for concurrent use (watcher goroutine + main goroutine)

## Dependencies
- None (pure utility package)

## Notes
- Not user-facing. This is for debugging prompt quality, watcher timing, and concept extraction during dogfooding.
- Use `sync.Mutex` or channels for thread safety.
