# 04-03: Output Component

## Summary
Scrollable viewport that renders streaming markdown responses in a chat-log style.

## Event Model Refs
- E6c: TUI renders chunks, appends reviews
- E7: TUI renders reactive responses
- Display Rules: chat log append, auto-scroll during streaming

## Files to Create
- `cli/tui/output.go`

## Model

```go
type OutputModel struct {
    viewport viewport.Model  // charmbracelet/bubbles
    content  strings.Builder  // accumulated content
    renderer *glamour.TermRenderer
    autoScroll bool
}
```

## Behavior
- Chat-log style: new content appends below old content
- Markdown rendering via glamour (code blocks, bold, lists)
- During streaming: receives ChunkMsg, appends to content, re-renders, auto-scrolls to bottom
- User can scroll up (mouse or keys) → disables auto-scroll
- Scrolling back to bottom → re-enables auto-scroll
- Renders styled elements:
  - `🍐` trigger headers
  - `📎` context lines
  - `━━━ Pear ━━━` / `━━━━━━━━━━━━` separators
  - `📚 Concepts` and `🔗 Related` lines
  - `🤔` Socratic questions
  - `⚠` error/warning lines

## Functions
- `AppendHeader(text string)` — adds styled header (trigger notification, context line)
- `StartStream()` — begins a new response block with separator
- `AppendChunk(text string)` — adds streaming chunk, re-renders markdown
- `EndStream()` — adds closing separator
- `AppendError(text string)` — adds warning-styled inline error
- `AppendSystem(text string)` — adds system message (slash command output)

## Acceptance Criteria
- Streaming text renders incrementally with markdown formatting
- Code blocks render with syntax highlighting
- Scrolling up pauses auto-scroll
- Old reviews stay visible above new ones
- Separators render correctly
- Terminal resize adjusts viewport

## Dependencies
- Deps: `bubbletea`, `bubbles/viewport`, `glamour`, `lipgloss`
