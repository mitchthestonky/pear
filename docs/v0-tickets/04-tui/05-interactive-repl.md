# 04-05: Interactive REPL Mode

## Summary
Wire up the root command (`pear` with no args) to launch Bubble Tea in interactive mode. Full multi-turn conversation.

## Event Model Refs
- E5: Interactive Mode Startup
- E7: User Input flow (without watcher)

## Files to Edit
- `cli/cmd/root.go` — replace stub, launch TUI in interactive mode

## Behavior
1. Print header: `🍐 Pear v0 · interactive · {provider}/{model}`
2. Launch Bubble Tea program with `mode: "interactive"`, `triggers: nil`
3. User types question → context collected → prompt assembled → LLM streams → response rendered
4. Multi-turn: full conversation history maintained
5. @file autocomplete works
6. Slash commands work (/help, /clear, /exit, /settings, /provider, /model, /key)
7. Ctrl+C or /exit → session end summary

## Acceptance Criteria
- `pear` launches TUI with input field and empty viewport
- User can ask questions and get streaming responses
- Follow-up questions include conversation history
- /clear resets history
- /exit prints session summary and quits
- @file resolves and attaches file contents

## Dependencies
- 04-01 (app shell)
- 04-02 (input)
- 04-03 (output)
- 04-04 (styles)
- 03-01 (collector)
- 03-02 (assembler)
- 02-01 + at least one provider

## Notes
- This is the first TUI "it works" moment
- Watch mode (04-01 watcher integration) is ticket 05-03
