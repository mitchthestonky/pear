# 03-04: Teach Command

## Summary
Wire up `pear teach` (diff-based) and `pear teach <topic>` (deep dive with auto-selected files).

## Event Model Refs
- E8: One-Shot [teach, no topic] and [teach, with topic]

## Files to Edit
- `cli/cmd/teach.go` — replace stub

## Behavior

### No args (`pear teach`)
1. `repocontext.GitDiff()` → current diff
2. `Collector.Build(diff)` → RepoContext
3. `Assembler.Proactive(repoContext, noHistory)` → use teach variant (adds Socratic framing)
4. Stream to stdout

### With topic (`pear teach goroutines`)
1. Grep codebase for topic: `git grep -l "goroutines"` → file list
2. Pick top 3 most relevant files (by match count)
3. Print: `📎 Context: auto-selected {files}`
4. Read selected files via `repocontext.ReadFile()`
5. `Collector.Build(topic, autoFiles)` → RepoContext
6. `Assembler.DeepDive(repoContext, topic)` → system prompt + messages
7. Stream to stdout

## Acceptance Criteria
- `pear teach` reviews current diff with teaching framing
- `pear teach goroutines` auto-selects files, shows what it picked
- Deep dive responses are longer and more thorough than proactive reviews
- Auto-selection uses `git grep` and shows results in context line

## Dependencies
- 03-01 (collector)
- 03-02 (assembler — all three prompt variants)
