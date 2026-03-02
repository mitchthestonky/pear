# Pear v0 — Staff Engineer Audit

> Generated 2026-03-02. All 25 tickets implemented and committed.

---

## Top 10 Issues (by severity)

### P0 — Broken

**1. Streaming chunks never reach the TUI** (`cli/tui/app.go` — `startStream` method)
- The `onChunk` callback in `startStream` is a no-op: `func(chunk string) { }`
- The user sees a frozen screen for the entire LLM response (potentially 30+ seconds), then the full response appears at once via `StreamDoneMsg`
- Fix requires piping chunks to the TUI via `tea.Program.Send()` stored on the Model, or a dedicated chunk channel that a `tea.Cmd` listens on
- This affects BOTH interactive mode and watch mode
- **This is the single most important fix. Without it, the product feels broken.**

### P1 — Will Cause User Pain

**2. No retry logic for rate limits** (`cli/llm/anthropic.go`, `cli/llm/openai.go`)
- `LLMError.Retry` and `LLMError.After` fields are populated by `mapHTTPError` but never consumed anywhere
- Rate-limited users see a cryptic error with no recovery
- EVENT_MODEL E13 specifies up to 2 retries with backoff — completely missing
- Fix: add retry wrapper in `startStream` or a shared `llm.StreamWithRetry` helper

**3. Concept tracking is never invoked**
- `learning.Extract()` and `ConceptStore.Record()` exist in `cli/learning/tracker.go` but are never called
- `StreamDoneMsg` handler in `cli/tui/app.go` does not extract concepts from responses
- One-shot commands (`ask.go`, `review.go`, `teach.go`) also never call learning functions
- `pear progress` will always show "No concepts tracked yet"
- Fix: after each `StreamDoneMsg`, call `learning.Extract(response)` and `store.Record()`

**4. Logger is never used (except watch.go)**
- `cli/logging/logger.go` is fully implemented with rotation, structured JSON, thread safety
- But it's never instantiated in interactive mode, one-shot commands, or TUI
- `watch.go` creates a logger but only passes it to the watcher
- Fix: create logger in root PersistentPreRun, pass through to TUI model

**5. Config file permissions too open** (`cli/config/config.go:70`)
- `os.Create(tmp)` creates config with mode `0666` (minus umask, typically `0644`)
- File contains plaintext API keys — should be `0600`
- Fix: `os.OpenFile(tmp, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)`

### P2 — Should Fix

**6. `pear ask` doesn't collect repo context** (`cli/cmd/ask.go`)
- Sends bare question with generic system prompt only
- Per EVENT_MODEL E8, `ask` should call `repocontext.Build()` and inject context
- "Explain my error handling" has no access to the user's code
- Fix: add `repocontext.Build` call, use `prompt.Reactive` instead of hardcoded system prompt

**7. Wasted printf before alt screen** (`cli/cmd/root.go:43`)
- Prints `🍐 Pear v0 · interactive · ...` before launching TUI with `tea.WithAltScreen()`
- Alt screen immediately clears this output — user never sees it
- The welcome banner inside the TUI properly handles this
- Fix: remove the `fmt.Printf` line

**8. `loadFileCache()` doesn't set `cmd.Dir`** (`cli/tui/input.go` — `loadFileCache` function)
- `git ls-files` runs without `cmd.Dir` set
- @file autocomplete breaks if user's cwd is a repo subdirectory
- Fix: pass repo root to InputModel, set `cmd.Dir` in `loadFileCache`

**9. Missing `f.Sync()` in config.Save** (`cli/config/config.go`)
- Write-tmp-rename is good for atomicity, but without `f.Sync()` before `f.Close()`, data can be lost on power failure
- Fix: add `f.Sync()` between encode and close

**10. Custom `repeat` instead of `strings.Repeat`** (`cli/tui/styles.go:95-101`)
- Hand-rolled O(n²) string concatenation loop when stdlib has O(n) `strings.Repeat`
- Ironic for a Go teaching tool
- Fix: replace with `strings.Repeat`

---

## Architecture Assessment

### Good
- Clean package separation with clear single responsibilities
- No circular dependencies
- `repocontext` naming to avoid stdlib `context` shadow (per CLAUDE.md)
- Atomic config writes via tmp+rename pattern
- SSE parsing handles all event types correctly for each provider
- `exec.Command` with separate args prevents shell injection on all git calls
- API keys masked in `/settings` display
- Prompt engineering is solid — proactive/reactive/deep-dive variants properly calibrated

### Concerns
- `llm.ProviderDetail` duplicated in both `llm/client.go` and `config/config.go` — `cmd/stream_helpers.go` manually bridges them
- Bubble Tea value receiver `Update` method calls pointer receiver mutation methods (`handleTrigger`, `handleUserInput`, etc.) — works because modified fields are heap-allocated (slices, maps, pointers) but is fragile
- `output.content` (strings.Builder) grows unbounded — multi-hour sessions accumulate all rendered text with no pruning
- `history` slice grows unbounded — proactive prompt caps to last 3, but full history stays in memory

---

## Security

- **API keys:** Stored plaintext in `~/.pear/config.toml` — standard for CLI tools but needs `0600` permissions (currently `0644`)
- **Command injection:** Safe — all git calls use `exec.Command` with separate args, no shell invocation
- **Path traversal:** `@/etc/passwd` would read arbitrary files via @file resolution. Low risk for local CLI but should constrain to repo root
- **Process list:** Keys sent via HTTP headers, not CLI args — not visible in `ps`

---

## Performance

- **Watcher:** Uses fsnotify + git polling hybrid. Ignores `.git/`, `node_modules/`, hidden dirs. Should be efficient on normal repos. No inotify limit handling for very large repos.
- **TUI responsiveness:** Currently broken due to P0 streaming issue. Once fixed, should be responsive since chunks arrive via tea.Msg.
- **Memory:** `history` and `output.content` grow unbounded. For multi-hour watch sessions, consider pruning viewport content and capping history.
- **Glamour renderer:** Re-created on resize (fine, resizes are rare).

---

## Missing Features vs PRD

| Feature | PRD/Event Model Ref | Status |
|---|---|---|
| Streaming chunks to TUI | Core UX | **Broken** (P0) |
| Retry on rate limit | E13 | Missing |
| Concept extraction integration | E6c, E7, E8 | Code exists, never called |
| Logging integration | All events | Package exists, barely used |
| Codebase config resolution | E1 | `ResolveCodebase` exists, only called in `watch.go` |
| Doctor LLM test (check 4) | E3 | Stubbed with TODO |
| `pear ask` with repo context | E8 | Sends bare question only |
| Session end summary on /exit | E12 | Partially implemented (/exit shows stats but no concept summary) |
| Binary file detection in @file | Robustness | Missing — binary content sent to LLM |
| History pruning | Performance | Missing — grows unbounded |

---

## Files Reference

```
cli/
├── main.go
├── cmd/
│   ├── root.go          # Entry + interactive REPL launch
│   ├── watch.go         # Watch mode with dirty diff prompt
│   ├── ask.go           # One-shot ask (missing context)
│   ├── review.go        # One-shot review with context
│   ├── teach.go         # One-shot teach with auto-file selection
│   ├── doctor.go        # Health checks (LLM check stubbed)
│   ├── init_cmd.go      # First-run wizard
│   ├── hooks.go         # Git hook install/uninstall
│   ├── progress.go      # Learning progress display
│   └── stream_helpers.go # Shared LLM client creation + formatting
├── config/config.go     # TOML config read/write
├── llm/
│   ├── client.go        # Interface + factory
│   ├── anthropic.go     # Anthropic SSE streaming
│   ├── openai.go        # OpenAI SSE streaming
│   ├── openai_sse.go    # Shared SSE parser
│   └── openrouter.go    # OpenRouter (delegates to OpenAI parser)
├── repocontext/collector.go  # Git diff, file tree, @file reading
├── prompt/assembler.go  # System prompt assembly (3 variants)
├── learning/tracker.go  # Concept extraction + persistence
├── hooks/hooks.go       # Post-commit hook management
├── logging/logger.go    # Structured JSON logging
├── watcher/watcher.go   # fsnotify + git polling
└── tui/
    ├── app.go           # Bubble Tea model + state machine
    ├── input.go         # Text input + @file autocomplete
    ├── output.go        # Scrollable viewport + markdown
    ├── styles.go        # Lipgloss styles + separators
    └── welcome.go       # Startup welcome banner
```
