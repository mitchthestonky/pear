# 04-01: Bubble Tea App Shell

## Summary
Main Bubble Tea model with state machine, message routing, and layout. This is the container for input and output components.

## Event Model Refs
- E4: TUI.StartWatchMode → TUI.Run
- E5: TUI.StartInteractive → TUI.Run
- State Machine: IDLE ⇄ STREAMING → IDLE → QUIT

## Files to Create
- `cli/tui/app.go`

## Model

```go
type Model struct {
    input      InputModel
    output     OutputModel
    mode       string          // "interactive" or "watch"
    state      string          // "idle", "streaming"
    paused     bool            // watch mode: proactive reviews paused
    history    []llm.Message
    stats      SessionStats
    config     *config.Config
    llmClient  llm.LLMClient
    triggers   <-chan watcher.ReviewTrigger // nil in interactive mode
    queuedTrig *watcher.ReviewTrigger       // buffered trigger during streaming
    width      int
    height     int
}

type SessionStats struct {
    StartTime    time.Time
    Reviews      int
    Concepts     int
}
```

## Messages (tea.Msg types)
- `ReviewTriggerMsg` — from watcher channel via waitForTrigger Cmd
- `ChunkMsg{Text string}` — LLM streaming chunk
- `StreamDoneMsg{Response *llm.Response}` — stream complete
- `StreamErrorMsg{Err error}` — stream failed
- `tea.KeyMsg` — keyboard input
- `tea.WindowSizeMsg` — terminal resize

## State Machine Rules
- **IDLE**: input enabled, watcher triggers processed, slash commands processed
- **STREAMING**: input disabled (keystrokes buffered), watcher triggers queued, viewport auto-scrolls
- On StreamDoneMsg: process queued trigger if any, re-issue waitForTrigger, state → IDLE
- On StreamErrorMsg: render inline error, state → IDLE

## Layout
```
┌─────────────────────────────┐
│ Scrollable output viewport  │
│ (chat log — appends)        │
│                             │
│                             │
├─────────────────────────────┤
│ > input field               │
└─────────────────────────────┘
```

## Acceptance Criteria
- `tea.NewProgram(model).Run()` launches and renders
- State transitions work: IDLE → STREAMING → IDLE
- Window resize updates layout
- Ctrl+C triggers graceful quit
- Mode string ("interactive"/"watch") affects behavior

## Dependencies
- 02-01 (llm types)
- Deps: `bubbletea`, `lipgloss`

## Notes
- Input and output are separate components (tickets 04-02, 04-03)
- This ticket wires them together and handles message routing
- The `waitForTrigger` Cmd pattern: `func() tea.Msg { return <-triggers }`
