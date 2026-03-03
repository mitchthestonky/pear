# 03-01: Context Collector

## Summary
Builds `RepoContext` from git state and file contents. Used by every command that sends context to the LLM.

## Event Model Refs
- E6c: Collector.Build(trigger.Diff)
- E7: Collector.Build(input, @files)
- E8: Collector.Build for all one-shot commands

## Files to Create
- `cli/repocontext/collector.go`

## Types

```go
type RepoContext struct {
    Diff         string
    ChangedFiles []string
    FileTree     string
    Branch       string
    Files        map[string]string // @file path → contents
    TriggerType  string            // "settle", "commit", "user"
    TriggerInfo  string            // summary string
}
```

## Functions
- `Build(opts CollectOpts) (*RepoContext, error)` — main entry point
- `GitDiff() (string, error)` — `git diff HEAD`, truncate to 300 lines
- `GitDiffRange(from, to string) (string, error)` — `git diff from..to`, truncate to 300 lines
- `GitFileTree() (string, error)` — `git ls-files`, depth 2, max 100 entries
- `GitBranch() (string, error)` — `git branch --show-current`
- `ReadFile(path string) (string, error)` — read file, truncate to 200 lines
- `ParseChangedFiles(diff string) []string` — extract file paths from diff headers
- `DiffSummary(diff string) string` — "3 files, +47 lines"
- `RepoRoot() (string, error)` — `git rev-parse --show-toplevel`

## All git commands must set `cmd.Dir` to repo root.

## Acceptance Criteria
- `Build` returns populated RepoContext from a real git repo
- Truncation works: 300 lines for diffs, 200 for files, 100 for tree
- `ParseChangedFiles` extracts paths from unified diff format
- `DiffSummary` produces human-readable summary
- Works from subdirectories (uses repo root)
- Returns meaningful error when not in a git repo

## Dependencies
- 01-01 (go.mod exists)

## Notes
- Package is `repocontext` not `context` (avoids stdlib shadow)
- All exec.Command calls must set `cmd.Dir` to repo root
- This is a pure data package — no TUI, no LLM dependency
