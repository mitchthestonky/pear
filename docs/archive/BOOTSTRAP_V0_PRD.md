# Pear Bootstrap v0 — Product Requirements Document

> Version: 0.3 (Dogfood Bootstrap)
> Last updated: February 2026
> Author: Mitch
>
> **Purpose:** Build a functional prototype of Pear that Mitch can use *while building Pear itself*. Close enough to the real product that dogfooding generates genuine product insights. **The core experience is proactive** — Pear watches you code and teaches in the gaps, not just when you ask.

---

## Why Bootstrap First

1. **Learn Go by building Go, with Pear teaching you Go.** The bootstrap tool watches your changes, explains Go patterns you just used (or misused), and asks you questions — exactly what the real product does.
2. **Generate real teaching interactions.** Every session produces data: what prompts work, what explanations land, what context is actually useful. This becomes golden-file test fixtures and marketing material.
3. **Experience the UX firsthand.** Friction you feel becomes features you build. You'll know what's missing because you'll miss it.
4. **Ship story writes itself.** "I built this tool with itself" is a genuine narrative for launch.

---

## What Bootstrap v0 Is

A Go binary (`pear`) with two primary modes:

1. **Watch mode (proactive)** — `pear watch`: runs in a terminal pane alongside your editor. Monitors your repo for changes. When it detects a pause in your work (no file changes for ~30 seconds), it proactively reviews what you just wrote and teaches you about it. Also triggers on every commit. You don't ask it anything — it's always paying attention.

2. **Interactive mode (reactive)** — `pear`: the REPL from before. Ask questions, get teaching responses, follow up. This is for when *you* want to drive the conversation.

Watch mode is the default experience. It's what makes Pear feel like a pair programmer who's actually watching your screen, not a chatbot waiting for input.

---

## The Core Loop: Watch Mode

```
You code in your editor (VS Code, Cursor, vim, whatever)
        │
        ▼
Pear watches for file changes (polling git status every 5s)
        │
        ▼
You pause for ~30 seconds (natural break — thinking, reading docs, etc.)
        │
        ▼
Pear detects the pause, snapshots the diff since last review
        │
        ▼
Pear proactively streams a teaching review to your terminal
        │
        ▼
You read it, learn something, keep coding
        │
        ▼
You commit → Pear reviews the commit automatically
        │
        ▼
Loop continues until you stop the session
```

**The key insight:** developers have natural pauses. Between writing a function and starting the next one. After a compile error. While waiting for tests. Pear teaches in those gaps without breaking flow.

---

## Scope

### In Scope

| Component | Details |
|---|---|
| **Watch mode (`pear watch`)** | Runs in foreground in a terminal pane. Polls `git status`/`git diff` every 5 seconds. Detects "settled" state (no new changes for ~30s). On settle: snapshots diff since last review, sends teaching review, streams to terminal. On commit (detected via HEAD change): reviews the commit. User can type during watch mode to ask questions (hybrid). |
| **Interactive REPL (`pear`)** | Default command. User-driven conversation with multi-turn history, `@file` support, streaming responses. |
| **Teach command** | `pear teach` (no args): review current diff and teach about it (same as proactive review, but on-demand). `pear teach <topic>`: deep dive on a concept (e.g., `pear teach goroutines`) using examples from the user's actual codebase — not abstract textbook examples. |
| **One-shot commands** | `pear ask "question"` (one-shot question with repo context), `pear review` (one-shot diff review) |
| **Post-commit teaching** | Watch mode detects when HEAD changes (new commit). Automatically reviews the committed diff and teaches. Also installable as a git hook via `pear hooks install` for repos where you want commit reviews even without watch mode running. |
| **Pause detection** | Tracks last file-change timestamp. When no changes detected for `settle_time` seconds (default 30, configurable), triggers a proactive review of accumulated changes since last review. Debounced: won't trigger again until new changes appear. |
| **Diff tracking** | Maintains a "last reviewed" state so it only teaches about *new* changes since the last proactive review. Stored in memory (resets on session exit). Implemented via snapshotting `git diff` output and comparing. |
| **Context injection** | Auto-attaches: git diff (new changes since last review), file tree, branch name, list of changed files. Shown as context summary line. |
| **`@file` attachment** | In both watch and interactive mode: type `@path/to/file` to attach file contents. |
| **Teaching prompt engine** | Same pedagogical framing as before. Proactive reviews use a slightly different prompt frame (see below): shorter, more focused, "here's what I noticed" tone rather than answering a question. |
| **Concept tagging & tracking** | LLM tags concepts per response. Stored in `~/.pear/learning.json`. `pear progress` shows history. |
| **Config file** | `~/.pear/config.toml`: `api_key`, `model`, `name`, `languages`, `level`, `settle_time` (default 30s), `watch_interval` (default 5s). |
| **`pear init` wizard** | First-run: Name, languages, level, API key → config.toml → runs doctor. |
| **`pear doctor`** | Checks: git, config, API key validity. |
| **`pear hooks install`** | Installs a `post-commit` git hook in the current repo that runs `pear review --commit HEAD` automatically. Lightweight alternative to watch mode. |
| **Smart review throttling** | Won't trigger a proactive review if the diff is < 5 lines (too small to teach on). Won't review more than once per 2 minutes (avoid spam on rapid iteration). |
| **Session awareness** | Watch mode prints a subtle status line on start: what it's watching, how long it's been running, how many reviews it's given this session. |

### Out of Scope (Build in Real CLI)

- Multiple LLM providers (Claude only for v0)
- Voice input/output
- Cobra (hand-rolled arg parsing)
- `/mode` switching (teach mode only)
- `/cost` token tracking
- TUI framework (Bubble Tea is v1.5)
- Filesystem watcher (fsnotify) — v0 uses git polling, v1.5 upgrades to fsnotify for instant detection
- Per-repo config

---

## User Flows

### First Run

```
$ pear
🍐 Welcome to Pear! Let's get you set up.

  What's your name? Mitch
  What languages do you work with? Go, TypeScript, React
  How would you describe your level? (junior / mid / senior) mid
  Anthropic API key? sk-ant-...

  ✓ Config saved to ~/.pear/config.toml
  Running pear doctor...
  ✓ git: installed
  ✓ config: found
  ✓ API key: valid
  All checks passed.

🍐 Ready. Run `pear watch` to start a session, or `pear` for interactive mode.
```

### Watch Mode (Primary Experience)

```
# Terminal pane 1: your editor (VS Code, Cursor, vim)
# Terminal pane 2:

$ pear watch
🍐 Pear v0 · watching · claude-sonnet-4-5-20250929 · teach mode
   Monitoring for changes... (settle time: 30s)

# You write code in your editor for a few minutes...
# You pause to think about what to do next...

🍐 Pear noticed you made changes (3 files, +47 lines)

📎 Context: git diff (47 lines), branch: feat/context-collector | teach mode

━━━ Pear ━━━
I see you just built the context collector. Here's what stood out:

1. **collector.go:15 — exec.Command without Dir**
   You're calling `exec.Command("git", "diff")` without setting
   `cmd.Dir`. This works from the repo root but breaks from
   subdirectories. Resolve the repo root first:

   ```go
   root, _ := exec.Command("git", "rev-parse", "--show-toplevel").Output()
   cmd.Dir = strings.TrimSpace(string(root))
   ```

   Every CLI tool that wraps git does this — `gh`, `lazygit`, etc.

2. **collector.go:38 — String concatenation in a loop**
   `result += line + "\n"` allocates a new string every iteration.
   Go strings are immutable. Use `strings.Builder` — it's the
   idiomatic way to build strings incrementally.

📚 Concepts: [exec.Command, cmd.Dir, strings.Builder]

🤔 Your collector reads the diff but doesn't handle the case where
the repo has no commits yet (fresh `git init`). What would
`git diff HEAD` return in that case, and how would you handle it?
━━━━━━━━━━━━

# You keep coding... make a few more changes... pause again...

🍐 Pear noticed you made changes (1 file, +12 lines)

📎 Context: git diff (12 lines), branch: feat/context-collector | teach mode

━━━ Pear ━━━
Good addition — you added stderr capture to the exec calls.

1. **collector.go:22 — Separate stdout/stderr buffers**
   This is the right pattern. `cmd.Output()` swallows stderr, which
   means if git fails you get "exit status 128" instead of
   "fatal: not a git repository". Your approach gives useful errors.
   This is production-grade Go.

📚 Concepts: [stderr capture, exec.ExitError]

🤔 You're now capturing stderr but still using `cmd.Output()`. That
method only populates the ExitError's stderr field, not your buffer.
Do you need to switch to `cmd.Run()` to use your custom buffers?
━━━━━━━━━━━━

# You commit your work...

🍐 Commit detected: "feat: add context collector"

📎 Context: commit diff (59 lines), branch: feat/context-collector | teach mode

━━━ Pear ━━━
Nice commit. One thing I want to highlight about the full picture
now that I can see it as a unit:

1. **Package structure** — You've got `context/collector.go` as its
   own package. Good instinct. In Go, packages are the unit of
   encapsulation. But notice your package is named `context` — that
   shadows the stdlib `context` package. Anyone importing both will
   hit a name collision. Consider `repocontext` or `collector`.

📚 Concepts: [package naming, stdlib shadowing]

🤔 Why does Go not have a way to alias package imports at the
declaration site, like TypeScript's `import { x as y }`? What
does Go offer instead?
━━━━━━━━━━━━

# You can also type questions during watch mode:

> Why did you say cmd.Output() doesn't use my buffers?

📎 Context: (follow-up, watching) | teach mode

━━━ Pear ━━━
Good question. Here's the subtlety...
━━━━━━━━━━━━

# Ctrl+C to stop watching

🍐 Session ended. 3 reviews, 6 concepts taught. Run `pear progress` to see your learning.
```

### Interactive Mode (On-Demand)

```
$ pear
🍐 Pear v0 · claude-sonnet-4-5-20250929 · teach mode
   Type your question or /help for commands.

> @cli/llm/claude.go How does my streaming implementation look?

📎 Context: git diff (12 lines), cli/llm/claude.go (89 lines),
   4 files, branch: feat/streaming | teach mode

━━━ Pear ━━━
...
━━━━━━━━━━━━

> /exit
🍐 Session ended. 1 review, 3 concepts taught.
```

### Teach Command

```
# No args: teach on current diff (on-demand version of proactive review)
$ pear teach

📎 Context: git diff (47 lines), branch: feat/context-collector | teach mode

━━━ Pear ━━━
Looking at your current changes, here's what's worth understanding:

1. **collector.go:15 — exec.Command without Dir** ...
2. **collector.go:38 — String concatenation in a loop** ...

📚 Concepts: [exec.Command, cmd.Dir, strings.Builder]

🤔 Your collector calls three separate git commands sequentially.
What would change if you ran them concurrently with goroutines?
━━━━━━━━━━━━

# With topic: deep dive using your actual codebase as examples
$ pear teach goroutines

📎 Context: file tree (12 files), branch: feat/watcher,
   @cli/watcher/watcher.go (auto-selected — contains goroutines) | teach mode

━━━ Pear ━━━
Let's talk about goroutines using your watcher code as the example.

In `watcher.go:34`, you launch the polling loop with `go w.pollLoop()`.
This is a goroutine — a lightweight thread managed by Go's runtime,
not the OS. Here's what's actually happening:

1. **`go` keyword** — This doesn't call the function and wait. It
   schedules it to run concurrently. Your main function continues
   immediately after this line. The poll loop runs in the background.

2. **Why this matters for your watcher:** Your main.go has a `select`
   loop that reads from both the watcher channel and stdin. Without
   goroutines, you'd have to choose: either listen for file changes
   OR listen for user input. Goroutines let you do both.

3. **The channel pattern you're using:**
   ```go
   triggers := make(chan ReviewTrigger)
   ```
   This is how goroutines communicate. Your watcher goroutine sends
   triggers, your main loop receives them. This is safer than shared
   memory — Go's mantra is "don't communicate by sharing memory,
   share memory by communicating."

4. **What you're missing:** Your goroutine has no way to stop. When
   the user hits Ctrl+C, the main function exits but the goroutine
   keeps running briefly. Use a `context.Context` with cancellation:

   ```go
   ctx, cancel := context.WithCancel(context.Background())
   defer cancel()
   go w.pollLoop(ctx)
   ```

   Then in pollLoop, check `ctx.Done()` in your select/ticker loop.

📚 Concepts: [goroutines, channels, go keyword, context.Context, cancellation]

🤔 Your watcher sends on a channel, but what happens if the main
loop is busy streaming a response and isn't reading from the channel?
What would you do to prevent the watcher from blocking?
━━━━━━━━━━━━
```

### One-Shot Modes

```
$ pear review
📎 Context: git diff (73 lines), branch: feat/prompt | teach mode
━━━ Pear ━━━
...
━━━━━━━━━━━━

$ pear ask "what's the idiomatic way to handle errors in Go?"
📎 Context: 4 files, branch: main | teach mode
━━━ Pear ━━━
...
━━━━━━━━━━━━
```

### Git Hook (Always-On Commit Reviews)

```
$ pear hooks install
✓ Installed post-commit hook in .git/hooks/post-commit
  Pear will review every commit in this repo.

# Later, after any commit:
$ git commit -m "fix: handle empty repo case"

🍐 Post-commit review:

📎 Context: commit diff (23 lines), branch: feat/edge-cases | teach mode

━━━ Pear ━━━
...
━━━━━━━━━━━━
```

### Concept Progress

```
$ pear progress

🍐 Concepts Pear has taught you:

  exec.Command          ████████░░  4 sessions
  strings.Builder       ██████░░░░  3 sessions
  error wrapping        ██████░░░░  3 sessions
  fmt.Errorf            ████░░░░░░  2 sessions
  exec.ExitError        ████░░░░░░  2 sessions
  stderr capture        ██░░░░░░░░  1 session
  cmd.Dir               ██░░░░░░░░  1 session
  package naming        ██░░░░░░░░  1 session
  stdlib shadowing      ██░░░░░░░░  1 session

  9 concepts across 5 sessions
```

---

## Technical Spec

### Project Structure

```
cli/
├── main.go                # Entry point, arg routing, REPL loop
├── watcher/
│   └── watcher.go         # File change polling, pause detection, commit detection
├── context/
│   └── collector.go       # Git diff, file tree, @file reading, diff-since-last-review
├── prompt/
│   └── assembler.go       # System prompt + context + history → messages (reactive + proactive variants)
├── llm/
│   └── claude.go          # Claude Messages API client with streaming
├── config/
│   └── config.go          # ~/.pear/config.toml read/write, pear init wizard
├── learning/
│   └── tracker.go         # Concept extraction from responses, ~/.pear/learning.json
├── hooks/
│   └── hooks.go           # Git hook install/uninstall
└── go.mod
```

Eight packages. Each becomes a real package in v1.5.

### watcher/watcher.go — The Proactive Engine

This is the new core package. It's what makes v0 feel like Pear.

```go
type Watcher struct {
    interval    time.Duration  // poll interval (default 5s)
    settleTime  time.Duration  // pause threshold (default 30s)
    minDiffSize int            // minimum diff lines to trigger review (default 5)
    cooldown    time.Duration  // minimum time between reviews (default 2min)

    lastDiff       string      // snapshot of diff at last review
    lastChangeTime time.Time   // when we last saw a change
    lastReviewTime time.Time   // when we last triggered a review
    lastHEAD       string      // last known commit hash (for commit detection)
    settled        bool        // whether changes have settled
}
```

**Poll loop (runs in a goroutine):**

```
Every `interval`:
  1. Run `git diff HEAD` → currentDiff
  2. Run `git rev-parse HEAD` → currentHEAD

  If currentHEAD != lastHEAD:
    → New commit detected
    → Get commit diff: `git diff lastHEAD..currentHEAD`
    → Trigger commit review
    → Update lastHEAD

  If currentDiff != lastDiff:
    → Files changed since last check
    → Update lastChangeTime = now
    → Update lastDiff = currentDiff
    → Mark settled = false

  If !settled AND time.Since(lastChangeTime) > settleTime:
    → Changes have settled (user paused)
    → If diff since last review is >= minDiffSize lines
    → AND time.Since(lastReviewTime) > cooldown
    → Trigger proactive review
    → Update lastReviewTime = now
    → Mark settled = true
    → Snapshot "last reviewed" diff
```

**Communication with main loop:** Watcher sends review triggers through a channel:

```go
type ReviewTrigger struct {
    Type    string // "settle" or "commit"
    Diff    string // the diff to review
    Summary string // "3 files, +47 lines" or "commit: feat: add collector"
}

triggers := make(chan ReviewTrigger)
```

Main loop selects on this channel AND stdin (for user questions in watch mode).

### main.go — Entry Point & Modes

```
Route subcommand:
  no args / ""     → launch interactive REPL
  "watch"          → launch watch mode (listening)
  "teach"          → no args: teach on current diff. with args: deep dive on topic
  "ask"            → one-shot question
  "review"         → one-shot diff review
  "progress"       → concept progress
  "init"           → config wizard
  "doctor"         → health checks
  "hooks"          → install/uninstall git hooks

Watch mode loop:
  Print session header
  Start watcher goroutine → receives triggers channel
  Start stdin reader goroutine → receives user input channel
  Select loop:
    case trigger from watcher:
      Collect full context
      Assemble proactive prompt
      Stream response
      Extract concepts
      Track in conversation history (limited — keep last 3 proactive reviews for follow-ups)
    case input from stdin:
      Handle slash commands
      Handle @file references
      Collect context
      Assemble reactive prompt (with recent proactive history for continuity)
      Stream response
      Extract concepts

Interactive REPL loop:
  Same as v0.2 — no watcher, user-driven only
```

### context/collector.go — Context Collection

Same as before, plus:

| Context | Command | Truncation |
|---|---|---|
| Git diff (since last review) | `git diff HEAD` minus last snapshot | First 300 lines |
| Commit diff | `git diff {old}..{new}` | First 300 lines |
| Changed files list | Parsed from diff header | N/A |
| File tree | `git ls-files` | Depth 2, max 100 files |
| Branch name | `git branch --show-current` | N/A |
| Active files (`@file`) | Direct file read | First 200 lines per file |

```go
type RepoContext struct {
    Diff         string
    ChangedFiles []string         // list of files in the diff
    FileTree     string
    Branch       string
    Files        map[string]string // @file path → contents
    TriggerType  string            // "settle", "commit", "user"
    TriggerInfo  string            // "3 files, +47 lines" or commit message
}
```

### prompt/assembler.go — Prompt Assembly

Two prompt variants:

**Proactive system prompt (watch mode, settle/commit triggers):**

```
You are Pear, a teaching-first coding companion running in watch mode.
You are observing the user code in real-time and proactively teaching
when they pause or commit.

## About the user
- Name: {name}
- Primary languages: {languages}
- Level: {level}

## Your role in watch mode
You just noticed the user {paused after making changes / committed code}.
Review their changes and teach them — they didn't ask you to, so keep
your tone natural and conversational, like a senior dev glancing at their
screen and saying "hey, nice — one thing though..."

Guidelines:
1. Lead with what you noticed: "I see you just..." / "Looking at what
   you changed..."
2. Focus on the 1-3 most teachable moments in the diff. Don't review
   every line — pick what matters.
3. If they did something well, say so first. Then flag improvements.
4. Keep it shorter than a full review — this is a nudge, not a lecture.
   Aim for 2-3 teaching points max.
5. End with exactly ONE Socratic question (🤔)
6. Tag concepts: 📚 Concepts: [...]

## Rules
- Be specific: file names, line numbers, variable names from the diff
- Teach the concept, not just the fix
- Show better code when suggesting improvements
- Match their level — skip basics, focus on Go idioms and production patterns
- Be concise. This is ambient teaching, not a code review they requested.
- If the diff is a commit, you can comment on the commit message quality too
```

**Teach deep-dive system prompt (`pear teach <topic>`):**

```
You are Pear, a teaching-first coding companion. The user wants to
learn about a specific concept. Teach it using their actual codebase
as the primary example — not abstract textbook examples.

## About the user
- Name: {name}
- Primary languages: {languages}
- Level: {level}

## Your approach for topic deep dives
1. Find where the concept appears (or should appear) in their codebase
2. Use their real code as the teaching example
3. Explain the concept from first principles, grounded in their code
4. Show what good usage looks like vs what they currently have
5. Connect to the broader ecosystem and production patterns
6. If the concept doesn't appear in their code yet, use their project
   structure to create relevant examples they could actually use
7. End with a Socratic question (🤔)
8. Tag concepts: 📚 Concepts: [...]

Keep it thorough but focused — this is an on-demand deep dive, so
the user is actively choosing to learn. Go deeper than a proactive
review would.
```

**Reactive system prompt (interactive mode, user-driven):**

Same as v0.2 — full teaching prompt for when the user explicitly asks.

**Continuity in watch mode:** When the user asks a follow-up question during watch mode, include the last 2-3 proactive reviews in the conversation history so the LLM has context. E.g., user asks "why did you say that about cmd.Dir?" and the LLM can reference its earlier proactive review.

### llm/claude.go — Claude API Client

Same as v0.2. No changes needed — the streaming client works for both proactive and reactive responses.

### config/config.go — Config & Init Wizard

**Config file:** `~/.pear/config.toml`

```toml
api_key = "sk-ant-..."
model = "claude-sonnet-4-5-20250929"
name = "Mitch"
languages = "Go, TypeScript, React"
level = "mid"
settle_time = 30       # seconds of inactivity before proactive review
watch_interval = 5     # seconds between git polls
min_diff_lines = 5     # minimum diff size to trigger review
review_cooldown = 120  # minimum seconds between proactive reviews
```

### learning/tracker.go — Concept Tracking

Same as v0.2. Works for both proactive and reactive responses.

### hooks/hooks.go — Git Hook Management

**`pear hooks install`:**
1. Check `.git/hooks/post-commit` exists
2. If exists, check if it already has pear's marker comment
3. Write/append: `#!/bin/sh\n# pear-hook\npear review --commit HEAD`
4. `chmod +x`

**`pear hooks uninstall`:**
1. Remove pear's lines from post-commit hook
2. If hook is now empty, delete it

**`pear review --commit HEAD`:**
- Gets the diff for the latest commit: `git diff HEAD~1..HEAD`
- Runs a one-shot proactive-style review
- Used by the git hook (runs automatically, no watch mode needed)

---

## Slash Commands (v0)

| Command | Behavior |
|---|---|
| `/help` | Print available commands |
| `/clear` | Clear conversation history, stay in session |
| `/exit` | Exit session (watch or interactive) |
| `/pause` | Pause proactive reviews (watch mode only, keep session open) |
| `/resume` | Resume proactive reviews after pause |
| `/status` | Show watch session stats: uptime, reviews given, concepts taught |

---

## Teaching Prompt Design Principles

These carry forward to the real product:

1. **Teach the concept, not just the fix.** "Use `strings.Builder`" is a fix. "Strings are immutable in Go, so `+=` copies the entire string on every iteration" is teaching.

2. **Connect to the broader ecosystem.** "This is how the standard library's `fmt.Fprintf` works internally" grounds the lesson.

3. **Reinforce good patterns.** Don't only flag problems. When the user does something right, tell them *why* it's right.

4. **Socratic over informational.** End with a question that requires reasoning. "What would happen if..." > "Do you understand?"

5. **Ambient over interruptive.** Proactive reviews should feel like a helpful glance, not a mandatory code review. Short, focused, useful. The user should *want* to read them.

6. **Teach in the gaps.** Don't interrupt active coding. Wait for natural pauses. This is the core UX principle.

7. **Calibrate to the developer.** v0 reads level from config. The real product builds a knowledge graph.

---

## What You'll Learn Building This

| Package | Go Concepts |
|---|---|
| `main.go` | Package structure, `os.Args`, `bufio.Scanner` (stdin), goroutines for select loop, channels |
| `watcher/` | **Goroutines**, `time.Ticker`, channels, `select` statement, `time.Since`, state management |
| `context/` | `os/exec`, `bytes.Buffer`, `strings.Builder`, `os.ReadFile`, struct design |
| `prompt/` | `fmt.Sprintf` or `text/template`, JSON marshaling, slice/map construction |
| `llm/` | `net/http`, JSON encode/decode, streaming HTTP (SSE), `bufio.Scanner`, error handling |
| `config/` | File I/O, TOML parsing (`BurntSushi/toml`), `bufio` for wizard input |
| `learning/` | JSON file read/write, regex (`regexp`), map operations, `sort.Slice` |
| `hooks/` | File permissions (`os.Chmod`), file read/write/append, string manipulation |

**Key upgrade from v0.2:** The watcher package forces you to learn **goroutines, channels, and select** — which are the Go concurrency primitives you need for the real CLI's parallelized context collection. v0.2 would have skipped these entirely.

---

## Success Criteria

Bootstrap v0 is done when:

- [ ] `pear init` creates config and validates setup
- [ ] `pear doctor` checks git, config, and API key
- [ ] `pear watch` monitors a repo and proactively reviews on pause (~30s inactivity)
- [ ] `pear watch` detects commits and reviews them automatically
- [ ] User can type questions during watch mode and get contextual answers
- [ ] `/pause` and `/resume` control proactive reviews
- [ ] `/status` shows session stats
- [ ] `pear` launches interactive REPL with multi-turn history
- [ ] `pear ask "question"` works as one-shot
- [ ] `pear review` works as one-shot (and `pear review --commit HEAD`)
- [ ] `pear hooks install` sets up post-commit hook
- [ ] `@file` attaches file contents to context in both modes
- [ ] Smart throttling: skips tiny diffs, respects cooldown
- [ ] Concepts are extracted and saved to `~/.pear/learning.json`
- [ ] `pear progress` shows learned concepts
- [ ] Session end summary: reviews given, concepts taught
- [ ] You've used watch mode for at least 5 coding sessions while building v0
- [ ] You've saved at least 5 proactive teaching interactions as golden-file fixtures

---

## Evolution Path

```
v0 (bootstrap)             → v1.5 (MVP)
────────────────────────────────────────────────────
os.Args routing            → Cobra commands
git polling (5s interval)  → fsnotify + git polling hybrid
hardcoded teach prompt     → prompt templates + modes
Claude only                → LLM adapter interface
in-memory history          → same (server persistence in v1.6)
raw stdin + channels       → Bubble Tea TUI
regex concept extract      → structured LLM output parsing
simple learning.json       → concept graph + knowledge state
no auth                    → license key / trial check
settle_time config         → adaptive timing (learns your rhythm)
pear hooks install         → automatic hook management
```

Each v0 package becomes a real package. The code gets refactored, not rewritten.

---

## Timeline

**Target: 7-10 days.** The watcher adds real concurrency work, but it's where the Go learning happens.

| Day | Milestone |
|---|---|
| 1 | `config/` — init wizard, doctor, config read/write. Pear boots and greets you. |
| 2 | `context/` — collector works, prints context summary. Diff snapshotting works. |
| 3 | `llm/` — Claude streaming works end-to-end. `pear ask` works. |
| 4 | `main.go` interactive REPL — multi-turn conversation, `/clear`, `/exit`. |
| 5 | `prompt/` — both proactive and reactive prompts. `pear review` works. `@file` works. |
| 6-7 | `watcher/` — polling loop, pause detection, commit detection, channels. **This is the Go learning surface.** |
| 8 | `main.go` watch mode — select loop over watcher channel + stdin. Hybrid input works. |
| 9 | `learning/` — concept extraction, learning.json, `pear progress`. `hooks/` — git hook install. |
| 10 | Polish, dogfood watch mode on real coding sessions, save golden-file fixtures. |

After day 10, you've built a concurrent Go program with goroutines, channels, and select. You've used it to teach yourself Go while building it. You *are* a Pear user. Start v1.5.
