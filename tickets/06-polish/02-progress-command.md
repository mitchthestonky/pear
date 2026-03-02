# 06-02: Progress Command

## Summary
Wire up `pear progress` to display concept learning history.

## Event Model Refs
- E11: Tracker.Display

## Files to Edit
- `cli/cmd/progress.go` — replace stub

## Behavior
1. `learning.Load("~/.pear/learning.json")`
2. `learning.Display(os.Stdout)`
3. Exit 0

## Output Format
```
🍐 Concepts Pear has taught you:

  goroutines            ████████░░  4 sessions
    → channels, select, context.Context
  strings.Builder       ██████░░░░  3 sessions
    → string immutability, performance

  12 concepts across 6 sessions
```

No TUI — straight stdout.

## Acceptance Criteria
- `pear progress` renders concept list sorted by frequency
- Shows relationship edges indented below each concept
- Shows total summary line
- Handles empty learning.json: "No concepts tracked yet. Start a session with `pear watch`."

## Dependencies
- 06-01 (tracker)
