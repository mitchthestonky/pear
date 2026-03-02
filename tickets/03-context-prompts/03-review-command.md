# 03-03: Review Command

## Summary
Wire up `pear review` and `pear review --commit HEAD` with context + prompts.

## Event Model Refs
- E8: One-Shot [review]

## Files to Edit
- `cli/cmd/review.go` — replace stub

## Behavior
1. If `--commit` flag: `repocontext.GitDiffRange("HEAD~1", "HEAD")`
2. Else: `repocontext.GitDiff()` (working tree diff)
3. `Collector.Build(diff)` → RepoContext
4. `Assembler.Proactive(repoContext, noHistory)` → system prompt + messages
5. Stream to stdout (no TUI), same pattern as `pear ask`
6. Tracker.Extract + Save (if tracker ready, else skip)

## Acceptance Criteria
- `pear review` reviews current dirty diff
- `pear review --commit HEAD` reviews last commit
- Context line printed: `📎 Context: git diff (N lines), branch: X`
- Streaming output with separators (`━━━ Pear ━━━`)

## Dependencies
- 03-01 (collector)
- 03-02 (assembler)
- 02-05 (ask command pattern — reuse streaming-to-stdout)
