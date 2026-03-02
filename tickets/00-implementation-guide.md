# Implementation Guide for Claude Agents

## Before You Start

1. **Read `CLAUDE.md`** at the project root. It has critical rules that will save you from known pitfalls.
2. **Read your assigned ticket(s)** in full, including the Event Model refs.
3. **Check dependencies.** If your ticket depends on another ticket, check `tickets/completed/` to confirm it's done. If the dependency isn't complete, check if the types/interfaces you need are defined yet — you may be able to work against the interface and stub the import.
4. **Read the referenced docs** (PRD.md, ARCHITECTURE.md, USER_JOURNEYS.md, EVENT_MODEL.md) for sections your ticket references.

## Working Directory

All Go code goes under `cli/`. The module is `github.com/pearcode/pear`.

```
Pear-v0/
├── cli/           ← all Go source code here
├── docs files     ← PRD, ARCHITECTURE, etc.
└── tickets/       ← your assignments
```

## How to Implement

1. **Start coding immediately.** Don't spend time exploring or planning. Your ticket has everything you need.
2. **Follow the ticket's acceptance criteria** as your checklist.
3. **Keep it minimal.** Implement exactly what the ticket says. No extra features, no extra abstractions, no extra comments.
4. **Run `go build ./...` frequently** to catch compile errors early.
5. **Run `go vet ./...`** before marking complete.

## When You're Done

After completing each ticket:

1. Run `go build ./...` — must compile
2. Run `go vet ./...` — must pass
3. Create a completion file at `tickets/completed/<ticket-id>.md` with this format:

```markdown
# <ticket-id>: <ticket name>

## Status: COMPLETE

## Files Created/Modified
- cli/path/to/file.go

## Notes
<any decisions made, deviations from ticket, or issues for downstream tickets>
```

4. **Commit** with message: `feat(<package>): <what you built>` — e.g., `feat(config): config package with TOML read/write and codebase resolution`

## Dependency Check

Before starting a ticket, check if its dependencies are done:

```bash
ls tickets/completed/
```

If a dependency ticket isn't complete but you only need its types/interfaces, you can:
- Define the types yourself inline and refactor later
- Or skip and move to a non-blocked ticket

## Parallel Work Rules

Multiple agents may be working simultaneously on different tickets. To avoid conflicts:
- **Only modify files in your ticket's package.** Don't edit files owned by another ticket.
- **If you need a type from another package that doesn't exist yet**, define a minimal version locally or in the package you're creating. The integration agent will resolve conflicts.
- **Don't modify `go.mod` unless you're the first agent** (ticket 01-01). If you need a new dependency, add it with `go get` — the lockfile handles conflicts.

## Common Mistakes to Avoid

- Don't name any package `context` — use `repocontext`
- Don't use `tea.Sub` — it doesn't exist. Use `tea.Cmd` blocking on a channel
- Don't put system prompts in `llm.Message` — use `StreamOptions.SystemPrompt`
- Don't use external LLM SDKs — hand-roll HTTP clients
- Don't forget `cmd.Dir` on git exec.Command calls
- Don't use unbuffered channels for watcher triggers — use buffered (size 1)
