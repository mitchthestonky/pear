# 05-02: Watch Mode Command (TUI + Watcher Integration)

## Summary
Wire up `pear watch` to launch Bubble Tea with the watcher goroutine. This is the full proactive experience.

## Event Model Refs
- E4: Watch Mode Startup (header, dirty diff prompt, watcher start)
- E6c: TUI receives trigger → context → prompt → stream
- State Machine: IDLE ⇄ STREAMING with queued triggers

## Files to Edit
- `cli/cmd/watch.go` — replace stub
- `cli/tui/app.go` — add watcher integration to Update loop

## Behavior

### Startup
1. Print header: `🍐 Pear v0 · watching · {provider}/{model}`
2. Init watcher
3. Check dirty diff exists:
   - Yes → prompt "You have uncommitted changes (N lines). Review now? [y/N]"
   - y → trigger immediate proactive review
   - N → baseline diff, continue
4. Start watcher goroutine → get triggers channel
5. Launch Bubble Tea with triggers channel set

### TUI Integration
- Issue `waitForTrigger(triggers)` Cmd on start and after each trigger processed
- On `ReviewTriggerMsg`:
  - If STREAMING → queue (store in `queuedTrig`)
  - If paused → drop, re-issue waitForTrigger
  - Else → run proactive review flow (E6c)
- On `StreamDoneMsg` → check `queuedTrig`, process if exists, then re-issue waitForTrigger
- `/pause` → set paused flag
- `/resume` → clear paused flag

### waitForTrigger pattern
```go
func waitForTrigger(ch <-chan watcher.ReviewTrigger) tea.Cmd {
    return func() tea.Msg {
        trigger, ok := <-ch
        if !ok { return nil }
        return ReviewTriggerMsg(trigger)
    }
}
```

## Acceptance Criteria
- `pear watch` starts watcher and TUI
- File changes trigger proactive reviews after settle time
- Commits trigger commit reviews
- User can type questions during watch mode (hybrid input)
- Follow-ups reference last 3 proactive reviews
- Queued triggers auto-play after current stream
- /pause stops triggers, /resume restarts
- Ctrl+C stops watcher and TUI cleanly
- Dirty diff prompt works on startup

## Dependencies
- 05-01 (watcher core)
- 04-01 (TUI app shell)
- 04-02 (input)
- 04-03 (output)
- 03-01 (collector)
- 03-02 (assembler)

## Notes
- This is the hardest integration ticket. The watcher channel → Bubble Tea bridge is the critical path.
- Test with a real repo: edit files, wait 30s, verify review streams.
