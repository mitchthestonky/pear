# Pear v0 — Claude Code Instructions

## What This Is
A Go CLI teaching tool that watches you code and proactively teaches during natural pauses. Built with Cobra, Bubble Tea, fsnotify. Multi-provider LLM support (Anthropic, OpenAI, OpenRouter).

## Project Structure
All code lives under `cli/`. Module: `github.com/pearcode/pear`

```
cli/
├── main.go              # Entry point
├── cmd/                  # Cobra commands (one file per subcommand)
├── watcher/              # fsnotify + git polling, runs in goroutine
├── repocontext/          # Git diff, file tree, @file reading (NOT "context/" — avoids stdlib shadow)
├── prompt/               # System prompt assembly (proactive, reactive, deep-dive variants)
├── llm/                  # LLMClient interface + per-provider clients
├── config/               # ~/.pear/config.toml read/write
├── learning/             # Concept extraction, learning.json persistence
├── hooks/                # Git hook install/uninstall
├── tui/                  # Bubble Tea app, input, output, styles
└── logging/              # Structured JSON logging to ~/.pear/logs/
```

## Critical Rules

### Do NOT
- Name any package `context` — use `repocontext` (stdlib shadow)
- Use `tea.Sub` — it doesn't exist in Bubble Tea. Use `tea.Cmd` that blocks on a channel:
  ```go
  func waitForTrigger(ch <-chan watcher.ReviewTrigger) tea.Cmd {
      return func() tea.Msg { return <-ch }
  }
  ```
- Put system prompts in the `llm.Message` array — use `StreamOptions.SystemPrompt`. Each provider handles placement internally (Anthropic: top-level `system` param, OpenAI/OpenRouter: prepend as role "system")
- Use external LLM SDKs — all HTTP clients are hand-rolled with `net/http` for streaming control
- Over-engineer. This is a one-night build. Get it working, then polish.
- Add features not in the tickets. If it's not specced, skip it.
- Create test files unless the ticket explicitly mentions tests

### Do
- Read your ticket file AND the referenced event model sections before starting
- Check dependency tickets are complete before starting yours
- Set `cmd.Dir` to repo root on ALL `exec.Command` calls to git
- Use buffered channel (size 1) for watcher triggers
- Handle `context.Context` cancellation in all goroutines
- Use atomic file writes (write tmp → rename) for config and learning.json
- Keep error handling simple: return typed `LLMError` for LLM failures, `fmt.Errorf` for everything else

## Key Interfaces

```go
// llm/client.go
type LLMClient interface {
    Stream(ctx context.Context, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error)
}

type StreamOptions struct {
    SystemPrompt string
    MaxTokens    int
    Temperature  float64
}
```

## Dependencies
```
github.com/spf13/cobra
github.com/charmbracelet/bubbletea
github.com/charmbracelet/bubbles
github.com/charmbracelet/lipgloss
github.com/charmbracelet/glamour
github.com/fsnotify/fsnotify
github.com/BurntSushi/toml
```

## TUI State Machine
```
IDLE → STREAMING → IDLE (→ process queued trigger if any)

IDLE: input enabled, watcher triggers processed
STREAMING: input disabled, watcher triggers queued (buffered chan, size 1)
PAUSED: input enabled, watcher triggers dropped silently
```

## Config Location
All config under `~/.pear/`. Nothing in the user's repo except optional git hooks.
```
~/.pear/
├── config.toml
├── learning.json
├── codebases/<path-slug>.toml
└── logs/<session-timestamp>.log
```

## Reference Docs (read before implementing)
- `PRD.md` — product requirements, phased build order, risk flags
- `ARCHITECTURE.md` — packages, interfaces, data flows, concurrency model
- `USER_JOURNEYS.md` — exact UX flows and display rules
- `EVENT_MODEL.md` — every event flow with pseudocode
- `tickets/` — individual implementation tickets with deps and acceptance criteria

## Commit Style
- Conventional commits: `feat:`, `fix:`, `refactor:`, `docs:`
- No co-authored-by lines
- Commit after each ticket is complete
