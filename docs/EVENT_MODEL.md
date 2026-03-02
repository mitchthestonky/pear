# Pear v0 — Event Model

> Every command, trigger, and state change in the system. Grounded in USER_JOURNEYS.md decisions.

---

## Actors

| Actor | Description |
|---|---|
| **User** | Types commands, writes code, commits |
| **CLI** | Cobra command router |
| **Config** | Reads/writes ~/.pear/config.toml, codebase overrides |
| **Doctor** | Health checks (git, config, API key) |
| **TUI** | Bubble Tea app — input, output, state machine |
| **Watcher** | fsnotify + git polling goroutine |
| **Collector** | Builds RepoContext from git + files |
| **Assembler** | Builds prompt from context + history + template |
| **LLM** | Provider client, streams response |
| **Tracker** | Extracts concepts, reads/writes learning.json |

---

## E1: Application Bootstrap

```
User runs any `pear` command
  │
  ├── Config.Load("~/.pear/config.toml")
  │     ├── [not found or invalid] → BLOCK → E2: Init Wizard
  │     │     └── after wizard completes → Config.Load again
  │     └── [valid] → continue
  │
  ├── Config.ResolvCodebase
  │     ├── Git.RepoRoot → slug path
  │     ├── Check ~/.pear/codebases/<slug>.toml
  │     └── Merge overrides onto global config
  │
  ├── LLM.Init(config.provider.active) → instantiate client
  ├── Tracker.Load("~/.pear/learning.json")
  │
  └── CLI.RouteCommand(os.Args)
        ├── "" (no args)  → E5: TUI Interactive
        ├── "watch"       → E4: TUI Watch Mode
        ├── "ask"         → E8: One-Shot
        ├── "review"      → E8: One-Shot
        ├── "teach"       → E8: One-Shot
        ├── "progress"    → E11: Display Progress
        ├── "init"        → E2: Init Wizard (explicit re-run)
        ├── "doctor"      → E3: Doctor
        └── "hooks"       → E10: Hook Management
```

---

## E2: Init Wizard

```
CLI.RunInitWizard
  │
  ├── Prompt: "What's your name?" → name (required)
  ├── Prompt: "What languages?" → languages (required)
  ├── Prompt: "Level?" → enum(junior, mid, senior) (required)
  ├── Prompt: "Choose provider" → enum(anthropic, openai, openrouter)
  ├── Prompt: "API key?" → string (required)
  ├── Prompt: "Default model?" → string (with provider-specific default)
  │
  ├── Config.Write("~/.pear/config.toml")
  │     └── mkdir -p ~/.pear/codebases/ ~/.pear/logs/
  │
  └── E3: Doctor (auto-runs, not optional)
        ├── [all pass] → "✓ All checks passed." → return to original command
        └── [fail] → show which check failed, prompt user to fix, re-run doctor
```

---

## E3: Doctor

```
CLI.RunDoctor
  │
  ├── Check 1: exec.LookPath("git")
  │     ├── [pass] → "✓ git: installed"
  │     └── [fail] → "✗ git: not found. Install git first."
  │
  ├── Check 2: Config.Load → parses without error
  │     ├── [pass] → "✓ config: valid"
  │     └── [fail] → "✗ config: invalid. Run `pear init`."
  │
  ├── Check 3: config.provider[active].api_key != ""
  │     ├── [pass] → "✓ API key: set"
  │     └── [fail] → "✗ API key: missing for [provider]."
  │
  └── Check 4: LLM.TestRequest (minimal prompt, e.g., "respond with OK")
        ├── [pass] → "✓ API key: valid"
        └── [fail] → "✗ API key: invalid for [provider]. HTTP [status]."
```

---

## E4: Watch Mode Startup

```
TUI.StartWatchMode
  │
  ├── Render: "🍐 Pear v0 · watching · {provider}/{model}"
  │
  ├── Watcher.Init(config.watch)
  │     ├── fsnotify.NewWatcher → add repo root recursively
  │     ├── Git.RevParseHEAD → lastHEAD
  │     └── Git.DiffHEAD → baseline diff
  │
  ├── Check: baseline diff has content?
  │     ├── [yes] → Prompt: "You have uncommitted changes ({n} lines). Review now? [y/N]"
  │     │     ├── [y] → E6: Proactive Review (with baseline diff)
  │     │     └── [N] → baseline becomes lastReviewDiff, continue
  │     └── [no] → continue
  │
  ├── Watcher.Start(ctx) → launches goroutine
  │     ├── fsnotify event loop
  │     └── 5s ticker loop
  │
  └── TUI.Run (Bubble Tea program)
        ├── Issue: waitForTrigger(triggers) Cmd
        ├── Render: always-on input field at bottom
        └── Enter main event loop → E6 / E7
```

---

## E5: Interactive Mode Startup

```
TUI.StartInteractive
  │
  ├── Render: "🍐 Pear v0 · interactive · {provider}/{model}"
  │
  └── TUI.Run (Bubble Tea program)
        ├── No watcher, no waitForTrigger
        ├── Render: always-on input field at bottom
        └── Enter main event loop → E7 only
```

---

## E6: Proactive Review (Watcher → TUI → LLM)

### 6a: Watcher Settle Detection

```
Watcher ticker fires (every 5s)
  │
  ├── time.Since(lastChangeTime) > settleTime?
  │     └── [no] → return
  │
  ├── settled already?
  │     └── [yes] → return (no new changes)
  │
  ├── Git.DiffHEAD → currentDiff
  ├── Subtract lastReviewDiff → newDiff
  │
  ├── CountLines(newDiff) >= minDiffSize?
  │     └── [no] → log(watcher.skip, too_small), return
  │
  ├── time.Since(lastReviewTime) > cooldown?
  │     └── [no] → log(watcher.skip, cooldown), return
  │
  ├── settled = true
  ├── lastReviewDiff = currentDiff
  ├── lastReviewTime = now
  │
  └── triggers <- ReviewTrigger{Type: "settle", Diff: newDiff, Summary: "3 files, +47 lines"}
```

### 6b: Watcher Commit Detection

```
Watcher ticker fires
  │
  ├── Git.RevParseHEAD → currentHEAD
  │
  ├── currentHEAD == lastHEAD?
  │     └── [yes] → return
  │
  ├── Git.Diff(lastHEAD + ".." + currentHEAD) → commitDiff
  ├── Git.LogOneline(currentHEAD) → commitMsg
  ├── lastHEAD = currentHEAD
  ├── lastReviewDiff = "" (reset baseline)
  │
  └── triggers <- ReviewTrigger{Type: "commit", Diff: commitDiff, Summary: commitMsg}
```

### 6c: TUI Receives Trigger

```
TUI.Update receives ReviewTriggerMsg (via waitForTrigger Cmd)
  │
  ├── state == STREAMING?
  │     └── [yes] → queue in buffered channel (size 1), return
  │
  ├── paused == true?
  │     └── [yes] → drop trigger, re-issue waitForTrigger, return
  │
  ├── state → STREAMING
  ├── Render trigger header: "🍐 Pear noticed you made changes ({summary})"
  │   or "🍐 Commit detected: {commitMsg}"
  │
  ├── Collector.Build(trigger.Diff) → RepoContext
  ├── Assembler.Proactive(repoContext, last3History) → messages + systemPrompt
  │
  ├── Issue streamCmd:
  │     └── goroutine: LLM.Stream(ctx, messages, StreamOptions{
  │           SystemPrompt: proactivePrompt,
  │           MaxTokens: 1024,
  │           Temperature: 0.7,
  │         }, onChunk)
  │           ├── onChunk(text) → program.Send(ChunkMsg{text})
  │           └── done → program.Send(StreamDoneMsg{response})
  │
  └── [on StreamDoneMsg]:
        ├── Tracker.Extract(response.Content) → concepts, relationships
        ├── Tracker.Save()
        ├── History.Append(assistant response, capped at 3 proactive)
        ├── Stats.reviews++
        ├── Stats.concepts += len(concepts)
        ├── log(llm.request, provider, model, input_tokens, output_tokens, latency)
        ├── state → IDLE
        ├── Check queued trigger → if exists, process it (loop back to top)
        └── Re-issue waitForTrigger Cmd
```

---

## E7: User Input (Interactive + Watch)

```
TUI.Update receives KeyMsg(Enter)
  │
  ├── state == STREAMING?
  │     └── [yes] → buffer keystroke, return (input disabled)
  │
  ├── input = TUI.InputValue()
  ├── TUI.ClearInput()
  │
  ├── input starts with "/"?
  │     └── [yes] → E9: Slash Command, return
  │
  ├── Parse @file references in input
  │     └── For each @path: resolve against repo root
  │
  ├── Collector.Build(input, @files) → RepoContext
  ├── Assembler.Reactive(repoContext, fullHistory) → messages + systemPrompt
  │
  ├── state → STREAMING
  ├── Render: "📎 Context: {contextSummary}"
  │
  ├── Issue streamCmd:
  │     └── goroutine: LLM.Stream(ctx, messages, StreamOptions{
  │           SystemPrompt: reactivePrompt,
  │           MaxTokens: 2048,
  │           Temperature: 0.7,
  │         }, onChunk)
  │
  └── [on StreamDoneMsg]:
        ├── Tracker.Extract(response.Content)
        ├── Tracker.Save()
        ├── History.Append(user message + assistant response)
        ├── Stats update
        ├── log(llm.request, ...)
        ├── state → IDLE
        └── Re-issue waitForTrigger Cmd (if watch mode)
```

### @file Autocomplete Sub-Flow

```
TUI.Update receives KeyMsg('@')
  │
  ├── Enter autocomplete mode
  ├── Git.LsFiles → file list (cached per session)
  │
  └── On each subsequent KeyMsg:
        ├── Filter file list by typed prefix
        ├── Render dropdown below input field
        │     ├── Arrow keys: navigate
        │     ├── Enter: confirm selection, insert path
        │     ├── Esc: cancel autocomplete
        │     └── Keep typing: narrow filter
        └── On confirm: append file path to input, exit autocomplete mode
```

---

## E8: One-Shot Commands (No TUI)

```
CLI.OneShot(commandType, args)
  │
  ├── [ask] → question = args[0]
  │     ├── Collector.Build(question) → RepoContext
  │     └── Assembler.Reactive(repoContext, noHistory) → messages
  │
  ├── [review] →
  │     ├── --commit flag? → Git.Diff("HEAD~1..HEAD")
  │     ├── else → Git.DiffHEAD
  │     ├── Collector.Build(diff) → RepoContext
  │     └── Assembler.Proactive(repoContext, noHistory) → messages
  │
  ├── [teach, no topic] →
  │     ├── Git.DiffHEAD → diff
  │     ├── Collector.Build(diff) → RepoContext
  │     └── Assembler.Proactive(repoContext, noHistory) → messages (teach variant)
  │
  ├── [teach, with topic] →
  │     ├── Grep codebase for topic → auto-select top 3 files
  │     ├── Render: "📎 Context: auto-selected {files}"
  │     ├── Collector.Build(topic, autoFiles) → RepoContext
  │     └── Assembler.DeepDive(repoContext, topic) → messages
  │
  ├── Render: "📎 Context: {summary}"
  ├── LLM.Stream(ctx, messages, opts, func(chunk) { fmt.Print(chunk) })
  │
  ├── Tracker.Extract(response.Content)
  ├── Tracker.Save()
  ├── log(llm.request, ...)
  │
  └── os.Exit(0)
```

---

## E9: Slash Commands

```
TUI.HandleSlashCommand(input)
  │
  ├── /help → Render command table
  │
  ├── /clear → History.Clear(), TUI.ClearViewport()
  │     └── Render: "🍐 History cleared."
  │
  ├── /exit → E12: Session End
  │
  ├── /pause → paused = true
  │     └── Render: "🍐 Proactive reviews paused. Type /resume to restart."
  │
  ├── /resume → paused = false
  │     └── Render: "🍐 Proactive reviews resumed."
  │
  ├── /status → Render: uptime, reviews, concepts, provider, model, branch
  │
  ├── /settings → Render numbered config display
  │     └── User types number → edit that field → Config.Write
  │         └── If provider/model/key changed → LLM.Reinit
  │
  ├── /provider → Render provider picker (numbered)
  │     └── User selects → prompt for API key if new → Config.Write → LLM.Reinit
  │
  ├── /model <name> → Config.SetModel(name) → Config.Write → LLM.Reinit
  │     └── Render: "✓ Model changed to {name} ({provider})"
  │
  └── /key → Prompt: "New API key?"
        └── Config.SetKey(key) → Config.Write → LLM.Reinit
              └── Render: "✓ API key updated."
```

---

## E10: Hook Management

```
Hooks.Install
  │
  ├── Git.RepoRoot → check we're in a repo
  ├── Read .git/hooks/post-commit
  │     ├── [exists, has # pear-hook] → "Already installed.", return
  │     ├── [exists, no marker] → append pear block
  │     └── [not exists] → create new file
  │
  ├── Write: "#!/bin/sh\n# pear-hook-start\npear review --commit HEAD\n# pear-hook-end"
  ├── os.Chmod(path, 0755)
  └── Render: "✓ Installed post-commit hook"

Hooks.Uninstall
  │
  ├── Read .git/hooks/post-commit
  ├── Remove lines between # pear-hook-start and # pear-hook-end
  ├── If remaining content is empty → delete file
  ├── Else → write back remaining content
  └── Render: "✓ Removed Pear hook"
```

---

## E11: Display Progress

```
Tracker.Display
  │
  ├── Load ~/.pear/learning.json
  ├── Sort concepts by count (descending)
  │
  ├── For each concept:
  │     ├── Render: "{name}  {bar}  {count} sessions"
  │     └── Render: "  → {related concepts}"
  │
  └── Render: "{total} concepts across {sessions} sessions"
```

---

## E12: Session End

```
TUI.Quit (Ctrl+C or /exit)
  │
  ├── Cancel context → Watcher.Stop (goroutine exits)
  ├── Tracker.Save() (final flush)
  │
  ├── Render: "🍐 Session ended. {reviews} reviews, {concepts} concepts taught."
  ├── Render: "   Run `pear progress` to see your learning history."
  │
  └── os.Exit(0)
```

---

## E13: LLM Error Handling

```
LLM.Stream returns LLMError
  │
  ├── error.Code == "rate_limit" AND error.Retry == true
  │     ├── Render inline: "⚠ Rate limited by {provider}. Retrying in {After}s..."
  │     ├── time.Sleep(error.After)
  │     └── Retry stream (max 2 retries)
  │
  ├── error.Code == "auth"
  │     └── Render inline: "⚠ API key invalid for {provider}. Run /key to update."
  │
  ├── error.Code == "network"
  │     └── Render inline: "⚠ Network error: {message}. Check your connection."
  │
  └── Partial response received before error
        └── Render inline: "⚠ Stream interrupted. Partial response above."

All errors render inline in chat log, styled as warnings.
State returns to IDLE after error. Input re-enabled.
```

---

## State Machine

```
States: INIT → IDLE ⇄ STREAMING → IDLE → ... → QUIT

IDLE:
  - Input field: enabled, focused
  - Watcher triggers: processed immediately
  - Slash commands: processed immediately

STREAMING:
  - Input field: disabled (keystrokes buffered)
  - Watcher triggers: queued (buffered channel, size 1)
  - Auto-scrolls viewport
  - On complete → IDLE (process queued trigger if any)

PAUSED (watch mode sub-state):
  - Input field: enabled
  - Watcher triggers: silently dropped
  - /resume → unpause
```

---

## Channel & Goroutine Map

```
Main goroutine:
  └── Bubble Tea event loop

Watcher goroutine (ctx-cancellable):
  ├── fsnotify.Events listener → updates lastChangeTime
  └── time.Ticker (5s) → settle check, HEAD check
      └── sends on: triggers chan ReviewTrigger (buffered, 1)

LLM streaming goroutine (per-request, ctx-cancellable):
  └── sends via: program.Send(ChunkMsg) and program.Send(StreamDoneMsg)

waitForTrigger Cmd (re-issued after each trigger):
  └── blocks on triggers channel, returns ReviewTriggerMsg
```
