# 06-05: Session End Summary

## Summary
Print summary stats when a session ends (Ctrl+C or /exit).

## Event Model Refs
- E12: Session End

## Files to Edit
- `cli/tui/app.go` — add shutdown handler

## Behavior
On quit:
1. Cancel context (stops watcher goroutine)
2. Final `tracker.Save()` flush
3. Print:
```
🍐 Session ended. {N} reviews, {M} concepts taught.
   Run `pear progress` to see your learning history.
```

## Acceptance Criteria
- Summary prints on Ctrl+C
- Summary prints on /exit
- Review and concept counts are accurate
- Tracker is saved before exit (no data loss)
- Watcher goroutine exits cleanly (no goroutine leak)

## Dependencies
- 04-01 (TUI app — quit handling)
- 06-01 (tracker — save)
