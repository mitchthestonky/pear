# Pear v0 — User Journeys

> Derived from UX decisions made during design review. These are the canonical user flows.

---

## Journey 1: First Run

Any `pear` command without a valid config **blocks and forces the init wizard**. User cannot proceed until setup is complete.

```
$ pear watch
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
  ✓ config: valid
  ✓ API key: valid
  All checks passed.

🍐 Ready. Starting watch mode...
```

**Rules:**
- Init wizard runs automatically on any command if no valid config exists
- Doctor runs automatically at the end of init (not a separate step)
- After init completes, the originally requested command continues
- No optional fields — every prompt is required

---

## Journey 2: Watch Mode (Daily Coding Session)

### Startup

```
$ pear watch
🍐 Pear v0 · watching · openrouter/claude-3.5-sonnet
```

One-line header: emoji, version, provider/model. User can `/status` for details.

If dirty diff exists at startup:

```
🍐 You have uncommitted changes (47 lines). Review them now? [y/N] y

📎 Context: git diff (47 lines), branch: main | teach mode

━━━ Pear ━━━
...
━━━━━━━━━━━━
```

### Core Loop

```
Terminal 1: Editor              Terminal 2: pear watch
─────────────────               ──────────────────────
                                🍐 Pear v0 · watching · openrouter/claude-3.5-sonnet
                                >                          ← always-on input field

Write code, save...             (fsnotify detects changes, resets settle timer)

Pause to think...               (30s passes...)

                                🍐 Pear noticed you made changes (3 files, +47 lines)
                                📎 Context: git diff (47 lines), branch: feat/collector
                                ━━━ Pear ━━━
                                1. **collector.go:15 — exec.Command without Dir**
                                   ...
                                2. **collector.go:38 — String concatenation in a loop**
                                   ...
                                📚 Concepts: [exec.Command, cmd.Dir, strings.Builder]
                                🔗 Related: [exec.Command → cmd.Dir]
                                🤔 What would `git diff HEAD` return on a fresh repo?
                                ━━━━━━━━━━━━
                                >
```

### Follow-Up Questions

Input field is **always visible** at the bottom. User just types:

```
                                > why did you suggest strings.Builder over bytes.Buffer?
                                📎 Context: (follow-up, watching)
                                ━━━ Pear ━━━
                                Good question. Both work, but here's the key difference...
                                ━━━━━━━━━━━━
                                >
```

LLM has the last 3 proactive reviews in history for continuity.

### Commit Detection

```
git commit -m "feat: collector" (HEAD changes)

                                🍐 Commit detected: "feat: collector"
                                📎 Context: commit diff (59 lines), branch: feat/collector
                                ━━━ Pear ━━━
                                Nice commit. One thing about the full picture...
                                ━━━━━━━━━━━━
                                >
```

### Pause / Resume

```
                                > /pause
                                🍐 Proactive reviews paused. Type /resume to restart.
                                >

(user codes without interruption... watcher triggers silently dropped)

                                > /resume
                                🍐 Proactive reviews resumed.
                                >
```

### Queued Reviews

If a watcher settle trigger fires while the LLM is streaming a response, it's **silently queued** and **auto-plays** after the current stream ends. No notification — it just flows naturally.

### Session End

```
                                Ctrl+C (or /exit)
                                🍐 Session ended. 4 reviews, 11 concepts taught.
                                   Run `pear progress` to see your learning history.
```

---

## Journey 3: Interactive Mode (REPL)

```
$ pear
🍐 Pear v0 · interactive · openrouter/claude-3.5-sonnet
>
```

Same TUI as watch mode but no watcher. User-driven only.

```
> @cli/llm/claude.go How does my streaming implementation look?
📎 Context: git diff (12 lines), cli/llm/claude.go (89 lines), branch: feat/streaming
━━━ Pear ━━━
...
━━━━━━━━━━━━

> What about error handling?
📎 Context: (follow-up)
━━━ Pear ━━━
...
━━━━━━━━━━━━

> /clear
🍐 History cleared.

> /exit
🍐 Session ended. 2 reviews, 5 concepts taught.
   Run `pear progress` to see your learning history.
```

### @file Autocomplete

Typing `@` triggers a **live dropdown** of matching files from the repo. Arrow keys to select, Enter to confirm. Filters as user types.

```
> @cli/w
  ┌──────────────────────────┐
  │ cli/watcher/watcher.go   │ ← highlighted
  │ cli/watcher/watcher_test │
  └──────────────────────────┘
```

---

## Journey 4: One-Shot Commands

No TUI. Streams directly to stdout. Exits when done.

### pear ask

```
$ pear ask "what's the idiomatic way to handle errors in Go?"
📎 Context: 4 files, branch: main
━━━ Pear ━━━
...
📚 Concepts: [error wrapping, fmt.Errorf, sentinel errors]
━━━━━━━━━━━━
$
```

### pear review

```
$ pear review
📎 Context: git diff (73 lines), branch: feat/prompt
━━━ Pear ━━━
...
━━━━━━━━━━━━

$ pear review --commit HEAD
📎 Context: commit diff (23 lines), branch: feat/edge-cases
━━━ Pear ━━━
...
━━━━━━━━━━━━
```

### pear teach

```
# No args — teach on current diff
$ pear teach
📎 Context: git diff (47 lines), branch: feat/collector
━━━ Pear ━━━
Looking at your current changes...
━━━━━━━━━━━━

# With topic — deep dive, auto-selects relevant files
$ pear teach goroutines
📎 Context: auto-selected cli/watcher/watcher.go (contains goroutines), file tree
━━━ Pear ━━━
Let's talk about goroutines using your watcher code as the example...
━━━━━━━━━━━━
```

Auto-selection: Pear greps the codebase for the topic, picks the most relevant files, and **shows what it picked** in the context line. User can override with `@file` on re-run.

---

## Journey 5: Check Progress

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

No TUI — straight stdout output.

---

## Journey 6: In-App Config Management

### /settings

```
> /settings

🍐 Current configuration:

  1. Name:       Mitch
  2. Languages:  Go, TypeScript, React
  3. Level:      mid
  4. Provider:   openrouter
  5. Model:      anthropic/claude-3.5-sonnet
  6. API Key:    sk-or-...****

  Watch settings:
  7. Settle time:    30s
  8. Min diff lines: 5
  9. Cooldown:       120s

  Enter a number to edit, or press Esc to close.

> 4

  Choose your LLM provider:
    1. Anthropic (Claude)
    2. OpenAI
    3. OpenRouter ← current
  > 1

  Anthropic API key? sk-ant-...
  ✓ Switched to anthropic (claude-sonnet-4-5-20250929)
```

Numbered quick actions — user types a number to edit that field. Changes saved immediately.

### Quick commands

```
> /provider              ← jump straight to provider picker
> /model claude-haiku    ← change model inline
  ✓ Model changed to claude-haiku (anthropic)
> /key                   ← update API key for active provider
  New API key? sk-ant-...
  ✓ API key updated.
```

---

## Journey 7: Git Hooks

```
$ pear hooks install
✓ Installed post-commit hook in .git/hooks/post-commit

$ git commit -m "fix: handle empty repo"
🍐 Post-commit review:
📎 Context: commit diff (23 lines), branch: feat/edge-cases
━━━ Pear ━━━
...
━━━━━━━━━━━━

$ pear hooks uninstall
✓ Removed Pear hook from .git/hooks/post-commit
```

---

## Error States

All errors appear **inline in the chat log**, styled as warnings:

```
⚠ Network error: connection refused. Check your internet and try again.
>

⚠ Rate limited by openrouter. Retrying in 5s...
━━━ Pear ━━━
[response after retry]
━━━━━━━━━━━━

⚠ API key invalid for anthropic. Run /key to update.
>

⚠ Stream interrupted. Partial response above.
>

⚠ Not a git repository. Run from a project with git initialized.
```

---

## Display Rules

| Rule | Detail |
|---|---|
| New reviews append below old ones | Chat log style, scroll up to see history |
| Input field always visible | Bottom of screen, always ready for input |
| Input disabled during streaming | Keystrokes buffered, input re-activates after stream ends |
| Queued watcher triggers auto-play | No notification, streams immediately after current response |
| Session header is one line | `🍐 Pear v0 · watching · provider/model` |
| `/status` for full details | Uptime, reviews given, concepts taught, config summary |
