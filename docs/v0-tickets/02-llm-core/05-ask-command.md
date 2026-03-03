# 02-05: Ask Command (End-to-End One-Shot)

## Summary
Wire up `pear ask "question"` as the first end-to-end command. No TUI — streams directly to stdout.

## Event Model Refs
- E8: One-Shot [ask]

## Files to Edit
- `cli/cmd/ask.go` — replace stub

## Behavior
1. Read config, init LLM client
2. Build minimal context (just the question, no git context yet — that's ticket 03-01)
3. Create messages: `[{role: "user", content: question}]`
4. Use a basic system prompt: "You are Pear, a teaching-first coding companion. Teach concepts, not just answers."
5. `LLM.Stream(ctx, messages, opts, func(chunk) { fmt.Print(chunk) })`
6. Print newline at end
7. Exit 0

## Acceptance Criteria
- `pear ask "what is a goroutine?"` streams a response to stdout
- Works with all 3 providers (switch via config)
- Errors display inline (auth, network, rate limit)

## Dependencies
- 01-02 (config)
- 02-01 (interface)
- At least one of 02-02/02-03/02-04 (a working provider)

## Notes
- This is the first "it works" moment. Keep it minimal.
- Context injection (git diff, file tree) comes in ticket 03-01 and will be wired in later.
