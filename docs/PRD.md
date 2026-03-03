# Pear v0 — Product Requirements Document

> Version: 1.0
> Last updated: March 2026
> Author: Mitch
>
> **Purpose:** Build a functional prototype of Pear that Mitch can use *while building Pear itself*. The core experience is proactive — Pear watches you code and teaches in the gaps, not just when you ask.

---

## Why Bootstrap First

1. **Learn Go by building Go, with Pear teaching you Go.** The bootstrap tool watches your changes, explains patterns you just used (or misused), and asks you questions — exactly what the real product does.
2. **Generate real teaching interactions.** Every session produces data: what prompts work, what explanations land, what context is actually useful. This becomes golden-file test fixtures and marketing material.
3. **Experience the UX firsthand.** Friction you feel becomes features you build.
4. **Ship story writes itself.** "I built this tool with itself."

---

## Modes

### 1. Watch Mode (`pear watch`) — Primary Experience

Runs in a terminal pane alongside your editor. Monitors your repo via fsnotify for instant file change detection, backed by git for diff content. When it detects a pause in your work (~30s of no file changes), it proactively reviews what you just wrote and teaches you about it. Also triggers on every commit.

You don't ask it anything — it's always paying attention.

```
You code in your editor
        │
        ▼
fsnotify detects file changes instantly, git provides diff content
        │
        ▼
You pause for ~30 seconds (natural break)
        │
        ▼
Pear snapshots the diff since last review
        │
        ▼
Pear streams a teaching review to your terminal (Bubble Tea TUI)
        │
        ▼
You read it, learn something, keep coding
        │
        ▼
You commit → Pear reviews the commit automatically
        │
        ▼
Loop continues until you stop
```

Users can also type questions during watch mode (hybrid input). Recent proactive reviews stay in conversation history for continuity.

### 2. Interactive Mode (`pear`) — On-Demand

Standard REPL with Bubble Tea TUI. User-driven conversation with multi-turn history, `@file` support, streaming responses. For when you want to drive the conversation.

### 3. Teach Command (`pear teach`)

- `pear teach` (no args): review current diff and teach about it — on-demand version of a proactive review.
- `pear teach <topic>`: deep dive on a concept using examples from the user's actual codebase.

### 4. One-Shot Commands

- `pear ask "question"` — single question with repo context
- `pear review` — one-shot diff review
- `pear review --commit HEAD` — review a specific commit

### 5. Git Hooks (`pear hooks install`)

Installs a `post-commit` git hook that runs `pear review --commit HEAD` automatically. Lightweight alternative to watch mode for always-on commit reviews.

- `pear hooks install` — writes hook to `.git/hooks/post-commit`
- `pear hooks uninstall` — removes Pear's lines from the hook

---

## Watch Mode Details

### File Change Detection (Hybrid)

- **fsnotify** for instant file change detection — no polling delay
- **git** for diff content and commit detection (polls `git rev-parse HEAD`)
- fsnotify triggers reset the "last change" timestamp; git provides the actual diff for review

### Pause Detection

- Tracks last file-change timestamp (updated by fsnotify events)
- When no changes for `settle_time` seconds (default 30, configurable), triggers proactive review
- Debounced: won't trigger again until new changes appear

### Smart Throttling

- Skips diffs < 5 lines (too small to teach on)
- Minimum 2 minutes between proactive reviews
- Both thresholds configurable

### Commit Detection

- Polls `git rev-parse HEAD` — when it changes, a commit happened
- Automatically reviews the committed diff

### Diff Tracking

- Maintains "last reviewed" state so it only teaches about *new* changes
- Stored in memory (resets on session exit)

---

## LLM Integration

### Multi-Provider Architecture

v0 ships with an extensible LLM interface supporting three providers via per-provider client implementations:

| Provider | API Format | Notes |
|---|---|---|
| **Anthropic** | Native Messages API | Direct Anthropic client |
| **OpenAI** | Chat Completions API | Native OpenAI client |
| **OpenRouter** | OpenRouter API | Access to many models via single key |

All providers implement a common `LLMClient` interface for streaming responses. Each has its own client handling auth, request format, and response parsing.

### Provider Configuration

Configured at app level in `~/.pear/config.toml`. Users pick one provider during `pear init`, then switch providers and models from within the app via slash commands.

---

## Config & Settings

### Global Config (`~/.pear/config.toml`)

```toml
name = "Mitch"
languages = "Go, TypeScript, React"
level = "mid"

[provider]
active = "openrouter"

[provider.anthropic]
api_key = "sk-ant-..."
model = "claude-haiku-4-5"

[provider.openai]
api_key = "sk-..."
model = "gpt-4o"

[provider.openrouter]
api_key = "sk-or-..."
model = "anthropic/claude-3.5-sonnet"

[watch]
settle_time = 30
interval = 5
min_diff_lines = 5
cooldown = 120
```

### Per-Codebase Config (`~/.pear/codebases/<name>.toml`)

All config lives under `~/.pear/` — nothing pollutes the user's repo. Per-codebase overrides (language, level, watch settings) are stored under `~/.pear/codebases/` keyed by codebase name. Pear detects the current repo and loads the appropriate overrides.

### In-App Configuration

All settings are viewable and editable from within the app — no need to manually edit files. Follows Claude Code UX patterns:

| Command | Behavior |
|---|---|
| `/settings` | Show current config |
| `/provider` | Switch active provider |
| `/model <name>` | Change model for current provider |
| `/key` | Update API key for current provider |

---

## First Run (`pear init`)

```
$ pear
🍐 Welcome to Pear! Let's get you set up.

  What's your name? Mitch
  What languages do you work with? Go, TypeScript, React
  How would you describe your level? (junior / mid / senior) mid

  Choose your LLM provider:
    1. Anthropic (Claude)
    2. OpenAI
    3. OpenRouter
  > 3

  OpenRouter API key? sk-or-...
  Default model? (anthropic/claude-3.5-sonnet) ↵

  ✓ Config saved to ~/.pear/config.toml
  Running pear doctor...
  ✓ git: installed
  ✓ config: found
  ✓ API key: valid
  All checks passed.

🍐 Ready. Run `pear watch` to start a session, or `pear` for interactive mode.
   Tip: Use /settings to change providers or /model to switch models anytime.
```

---

## `pear doctor`

Checks:
1. Git installed and accessible
2. Config file exists and parses
3. Active provider API key is set
4. API key is valid (test request)

---

## Context Injection

| Context | Source | Truncation |
|---|---|---|
| Git diff (since last review) | `git diff HEAD` minus last snapshot | First 300 lines |
| Commit diff | `git diff {old}..{new}` | First 300 lines |
| Changed files list | Parsed from diff header | N/A |
| File tree | `git ls-files` | Depth 2, max 100 files |
| Branch name | `git branch --show-current` | N/A |
| Active files (`@file`) | Direct file read | First 200 lines per file |

---

## Concept Tracking & Knowledge Graph

### Concept Extraction

- LLM is prompted to tag concepts per response: `📚 Concepts: [...]`
- Also prompted to tag relationships: `🔗 Related: [goroutines → channels, channels → select]`
- Extracted via regex from responses

### Storage (`~/.pear/learning.json`)

```json
{
  "concepts": {
    "goroutines": {
      "count": 4,
      "sessions": ["2026-03-01T10:00:00Z", "..."],
      "related": ["channels", "select", "context.Context"]
    },
    "channels": {
      "count": 3,
      "sessions": ["..."],
      "related": ["goroutines", "select", "buffered channels"]
    }
  }
}
```

### `pear progress`

```
$ pear progress

🍐 Concepts Pear has taught you:

  goroutines            ████████░░  4 sessions
    → channels, select, context.Context
  strings.Builder       ██████░░░░  3 sessions
    → string immutability, performance
  error wrapping        ██████░░░░  3 sessions
    → fmt.Errorf, sentinel errors

  12 concepts across 6 sessions
```

---

## Slash Commands

| Command | Behavior | Mode |
|---|---|---|
| `/help` | Print available commands | All |
| `/clear` | Clear conversation history | All |
| `/exit` | Exit session | All |
| `/pause` | Pause proactive reviews | Watch |
| `/resume` | Resume proactive reviews | Watch |
| `/status` | Session stats: uptime, reviews, concepts | Watch |
| `/settings` | Show current config | All |
| `/provider` | Switch active provider | All |
| `/model <name>` | Change model | All |
| `/key` | Update API key | All |

---

## Output Style

Polished Bubble Tea TUI with emoji markers, styled separators, and context summary lines:

```
🍐 Pear noticed you made changes (3 files, +47 lines)

📎 Context: git diff (47 lines), branch: feat/context-collector | teach mode

━━━ Pear ━━━
I see you just built the context collector. Here's what stood out:

1. **collector.go:15 — exec.Command without Dir**
   You're calling `exec.Command("git", "diff")` without setting
   `cmd.Dir`. This works from the repo root but breaks from
   subdirectories.

   ```go
   root, _ := exec.Command("git", "rev-parse", "--show-toplevel").Output()
   cmd.Dir = strings.TrimSpace(string(root))
   ```

2. **collector.go:38 — String concatenation in a loop**
   Go strings are immutable. `+=` copies the entire string on every
   iteration. Use `strings.Builder`.

📚 Concepts: [exec.Command, cmd.Dir, strings.Builder]
🔗 Related: [exec.Command → cmd.Dir, strings.Builder → string immutability]

🤔 Your collector reads the diff but doesn't handle the case where
the repo has no commits yet. What would `git diff HEAD` return in
that case, and how would you handle it?
━━━━━━━━━━━━
```

---

## Teaching Prompt Principles

1. **Teach the concept, not just the fix.** "Strings are immutable in Go, so `+=` copies the entire string" > "use `strings.Builder`"
2. **Connect to the broader ecosystem.** Ground lessons in real-world patterns.
3. **Reinforce good patterns.** When the user does something right, say why it's right.
4. **Direct over Socratic.** Deliver insights directly — no quizzing, no "what do you think?" endings. Be a pair programmer, not a tutor.
5. **Ambient over interruptive.** Proactive reviews are nudges, not lectures. 2-3 teaching points max.
6. **Teach in the gaps.** Wait for natural pauses. Never interrupt active coding.
7. **Calibrate to the developer.** Read level from config, match explanation depth.
8. **Tag concepts.** End responses with `📚 Concepts:` and `🔗 Related:` tags for learning tracking.

---

## Success Criteria

- [ ] `pear init` creates config with provider selection and validates setup
- [ ] `pear doctor` checks git, config, and API key
- [ ] `pear watch` monitors a repo via fsnotify + git and proactively reviews on pause
- [ ] `pear watch` detects commits and reviews them automatically
- [ ] User can type questions during watch mode (hybrid input via Bubble Tea)
- [ ] `/pause` and `/resume` control proactive reviews
- [ ] `/status` shows session stats
- [ ] `/settings`, `/provider`, `/model`, `/key` manage config in-app
- [ ] `pear` launches interactive REPL with Bubble Tea TUI and multi-turn history
- [ ] `pear ask "question"` works as one-shot
- [ ] `pear review` and `pear review --commit HEAD` work
- [ ] `pear teach` and `pear teach <topic>` work
- [ ] `pear hooks install/uninstall` manages post-commit hook
- [ ] `@file` attaches file contents in both modes
- [ ] Smart throttling: skips tiny diffs, respects cooldown
- [ ] LLM provider switching works (Anthropic, OpenAI, OpenRouter)
- [ ] Concepts extracted with relationships and saved to `~/.pear/learning.json`
- [ ] `pear progress` shows concepts with relationship graph
- [ ] Session end summary: reviews given, concepts taught
- [ ] Used watch mode for at least 5 sessions while building v0
- [ ] Saved at least 5 proactive interactions as golden-file fixtures

---

## Phased Build Order

This is a one-night build. Phases are sequential — each produces a working binary.

### Phase 1: Skeleton (commits: init, config, doctor)
- Cobra CLI with all subcommands stubbed
- `pear init` wizard (one provider, writes config.toml)
- `pear doctor` (git, config, API key checks)
- Config read/write with `~/.pear/` directory structure

### Phase 2: LLM Core (commits: llm interface, providers, streaming)
- `LLMClient` interface with `StreamOptions`
- Anthropic, OpenAI, OpenRouter clients with streaming
- `pear ask "question"` works end-to-end (one-shot, no TUI)

### Phase 3: Context + Prompts (commits: collector, assembler)
- Context collector (git diff, file tree, branch, @file)
- Prompt assembler (proactive, reactive, deep-dive templates)
- `pear review` and `pear teach` work end-to-end

### Phase 4: TUI Shell (commits: bubble tea app, input, output, styles)
- Bubble Tea app with input + streaming output
- Interactive REPL (`pear`) with multi-turn history
- Slash command routing
- Markdown rendering via glamour

### Phase 5: Watcher (commits: fsnotify, git polling, integration)
- fsnotify + git polling hybrid watcher
- Pause detection, commit detection, throttling
- `pear watch` with hybrid input (watcher channel → Bubble Tea)
- TUI state machine (idle → streaming → input-focused)

### Phase 6: Learning + Polish (commits: tracker, progress, hooks)
- Concept extraction + relationship graph → learning.json
- `pear progress` display
- `pear hooks install/uninstall`
- Session end summary
- In-app config commands (`/settings`, `/provider`, `/model`, `/key`)

---

## Command Surface Clarification

| Command | When to use | Distinct from |
|---|---|---|
| `pear watch` | Primary workflow — leave running alongside editor | Everything else |
| `pear` (no args) | When you want to drive the conversation yourself | Watch mode (no proactive reviews) |
| `pear ask "q"` | Quick one-off from shell, scripting, piping | REPL (no session, no history) |
| `pear review` | On-demand diff review, also used by git hooks | Watch mode settle trigger (same output, manual trigger) |
| `pear teach` | On-demand diff teaching (adds Socratic question) | `pear review` (review is assessment, teach is pedagogy) |
| `pear teach <topic>` | Deep dive on a concept using your codebase | Everything else (topic-driven, not diff-driven) |

---

## Known Risks & Flags

> Staff review findings — address during implementation.

1. **Scope is 3-4x original bootstrap.** Cobra + Bubble Tea + fsnotify + 3 providers + concept graph. This is a small product, not a 10-day prototype. The phased build order above mitigates this — each phase produces a working binary.

2. **Bubble Tea + Watcher integration is the hardest problem.** Streaming output + user input + watcher events in the same TUI. Needs a state machine (see ARCHITECTURE.md). Don't underestimate this.

3. **Concept extraction via regex is fragile.** LLMs vary formatting of `📚 Concepts: [...]` lines. Accept this for v0 — fix edge cases as hit. Structured output is v1.5.

4. **VISION.md says "premier model usage" in Pro tier but v0 is BYOK-only.** No conflict for v0 (dogfood), but align docs before public launch.

---

## Out of Scope (v1.5+)

- Voice input/output
- Adaptive timing (learns your coding rhythm)
- Full knowledge graph with spaced repetition
- Server-side persistence
- Auth / license keys
- Per-repo config files (all config lives in `~/.pear/`)
