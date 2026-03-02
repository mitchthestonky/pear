# 05-01: Watcher Core (fsnotify + Git Polling)

## Summary
Hybrid file watcher: fsnotify for instant change detection, git polling for diff content and commit detection. Runs in a goroutine, sends ReviewTrigger on a channel.

## Event Model Refs
- E4: Watcher.Init, Watcher.Start
- E6a: Settle Detection
- E6b: Commit Detection

## Files to Create
- `cli/watcher/watcher.go`

## Types

```go
type Watcher struct {
    settleTime   time.Duration
    minDiffSize  int
    cooldown     time.Duration

    lastChangeTime time.Time
    lastReviewTime time.Time
    lastHEAD       string
    lastReviewDiff string
    settled        bool

    fsWatcher    *fsnotify.Watcher
    triggers     chan ReviewTrigger
    repoRoot     string
}

type ReviewTrigger struct {
    Type    string // "settle" or "commit"
    Diff    string
    Summary string
}
```

## Functions
- `New(cfg config.WatchConfig, repoRoot string) (*Watcher, error)` — init fsnotify, set baseline HEAD and diff
- `Start(ctx context.Context) <-chan ReviewTrigger` — launch goroutine, return triggers channel
- `Stop()` — called via context cancellation

## Goroutine Logic
Two concurrent listeners in one goroutine:

1. **fsnotify events**: on any Write/Create event → update `lastChangeTime`, set `settled = false`. Ignore `.git/` directory events.

2. **5s ticker**:
   - Check HEAD for commit detection (E6b)
   - Check settle condition (E6a): `!settled && time.Since(lastChangeTime) > settleTime && diff >= minDiffSize && cooldown elapsed`
   - Send ReviewTrigger on channel

## fsnotify Setup
- Watch repo root recursively (walk directories, add each)
- Ignore: `.git/`, `node_modules/`, hidden dirs
- Fallback: if fsnotify fails (too many watches), log warning and rely on git polling only

## Acceptance Criteria
- fsnotify detects file changes in a test repo
- Settle triggers after configured seconds of inactivity
- Commit triggers when HEAD changes
- Skips diffs < minDiffSize (logs skip)
- Respects cooldown between reviews (logs skip)
- Context cancellation stops goroutine cleanly
- Triggers channel is buffered (size 1)

## Dependencies
- 01-02 (config.WatchConfig)
- 03-01 (repocontext — for git commands, or inline git exec)
- Dep: `github.com/fsnotify/fsnotify`

## Notes
- This is the core concurrency package. Use `select` with fsnotify.Events, ticker.C, and ctx.Done().
- The triggers channel is read by the TUI via `waitForTrigger` Cmd (ticket 05-02).
