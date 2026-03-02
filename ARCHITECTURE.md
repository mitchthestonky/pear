# Pear v0 — Technical Architecture

> Last updated: March 2026

---

## Project Structure

```
cli/
├── main.go                     # Entry point, Cobra root command setup
├── cmd/
│   ├── root.go                 # Root command (launches interactive REPL)
│   ├── watch.go                # `pear watch` — proactive watch mode
│   ├── teach.go                # `pear teach [topic]`
│   ├── ask.go                  # `pear ask "question"`
│   ├── review.go               # `pear review [--commit HEAD]`
│   ├── progress.go             # `pear progress` — concept display
│   ├── init.go                 # `pear init` — first-run wizard
│   ├── doctor.go               # `pear doctor` — health checks
│   └── hooks.go                # `pear hooks install/uninstall`
├── watcher/
│   └── watcher.go              # fsnotify + git polling hybrid, pause detection
├── repocontext/
│   └── collector.go            # Git diff, file tree, @file reading, diff tracking
├── prompt/
│   └── assembler.go            # System prompt + context + history → messages
├── llm/
│   ├── client.go               # LLMClient interface
│   ├── anthropic.go            # Anthropic Messages API client
│   ├── openai.go               # OpenAI Chat Completions client
│   └── openrouter.go           # OpenRouter client
├── config/
│   └── config.go               # ~/.pear/ config read/write, in-app settings
├── learning/
│   └── tracker.go              # Concept extraction, relationship graph, learning.json
├── hooks/
│   └── hooks.go                # Git hook install/uninstall
├── tui/
│   ├── app.go                  # Bubble Tea app model, update, view
│   ├── input.go                # Input component (prompt, @file autocomplete)
│   ├── output.go               # Output component (streaming markdown render)
│   └── styles.go               # Lipgloss styles, colors, separators
└── go.mod
```

---

## Package Responsibilities

### `cmd/` — Cobra Commands

Each subcommand is a separate file. Cobra handles arg parsing, help text, and flag definitions. The root command (no args) launches the interactive REPL.

### `watcher/` — The Proactive Engine

Hybrid file watching: fsnotify for instant change detection, git for diff content and commit detection.

```go
type Watcher struct {
    settleTime  time.Duration  // pause threshold (default 30s)
    minDiffSize int            // minimum diff lines to trigger review (default 5)
    cooldown    time.Duration  // minimum time between reviews (default 2min)

    lastChangeTime time.Time   // last fsnotify event timestamp
    lastReviewTime time.Time   // when we last triggered a review
    lastHEAD       string      // last known commit hash
    lastReviewDiff string      // diff snapshot at last review
    settled        bool        // whether changes have settled
}

type ReviewTrigger struct {
    Type    string // "settle" or "commit"
    Diff    string // the diff to review
    Summary string // "3 files, +47 lines" or "commit: feat: add collector"
}
```

**Event loop:**

```
fsnotify event (file changed):
  → Update lastChangeTime
  → Mark settled = false

Ticker (every 5s):
  → git rev-parse HEAD → check for new commits
  → If !settled AND time.Since(lastChangeTime) > settleTime:
    → git diff HEAD (minus last reviewed snapshot)
    → If diff >= minDiffSize AND cooldown elapsed
    → Send ReviewTrigger{Type: "settle"} on channel

Commit detected (HEAD changed):
  → git diff oldHEAD..newHEAD
  → Send ReviewTrigger{Type: "commit"} on channel
```

**Communication:** Sends `ReviewTrigger` on a channel. Main loop selects on this channel and Bubble Tea input events.

### `repocontext/` — Context Collection (renamed from `context/`)

> **Flag:** Original name `context/` shadows Go's stdlib `context` package, causing import collisions in every file that needs both. Renamed to `repocontext/`.


```go
type RepoContext struct {
    Diff         string
    ChangedFiles []string
    FileTree     string
    Branch       string
    Files        map[string]string  // @file path → contents
    TriggerType  string             // "settle", "commit", "user"
    TriggerInfo  string             // summary string
}
```

| Context | Command | Truncation |
|---|---|---|
| Git diff (since last review) | `git diff HEAD` minus snapshot | 300 lines |
| Commit diff | `git diff {old}..{new}` | 300 lines |
| Changed files | Parsed from diff headers | — |
| File tree | `git ls-files` | Depth 2, 100 files |
| Branch | `git branch --show-current` | — |
| `@file` contents | `os.ReadFile` | 200 lines/file |

### `prompt/` — Prompt Assembly

Three prompt variants, all assembled from templates:

1. **Proactive** (watch mode, settle/commit) — "I noticed you..." tone, 2-3 teaching points, one Socratic question
2. **Reactive** (interactive mode, user-driven) — full teaching prompt for explicit questions
3. **Deep dive** (`pear teach <topic>`) — thorough concept explanation using the user's actual codebase

All prompts inject user profile (name, languages, level) and request concept tagging with relationships.

### `llm/` — Multi-Provider LLM Clients

```go
type LLMClient interface {
    Stream(ctx context.Context, messages []Message, opts StreamOptions, onChunk func(string)) (*Response, error)
}

type Message struct {
    Role    string // "user", "assistant" — system handled separately
    Content string
}

type StreamOptions struct {
    SystemPrompt string  // handled per-provider (Anthropic: top-level param, OpenAI: role message)
    MaxTokens    int     // varies by use case (proactive: 1024, deep dive: 4096)
    Temperature  float64 // default 0.7
}

type Response struct {
    Content     string
    InputTokens  int
    OutputTokens int
}

type LLMError struct {
    Code    string // "rate_limit", "auth", "network", "unknown"
    Message string
    Retry   bool          // whether caller should retry
    After   time.Duration // suggested wait (for rate limits)
}

func (e *LLMError) Error() string { return e.Message }
```

> **Flag:** System messages are handled differently per provider — Anthropic uses a top-level `system` param, OpenAI/OpenRouter use a `role: "system"` message. The `StreamOptions.SystemPrompt` field lets each provider handle placement internally, keeping the interface clean. Token counts are returned for logging/cost tracking during dogfooding.

Three implementations:

**`anthropic.go`** — Anthropic Messages API
- `POST https://api.anthropic.com/v1/messages`
- SSE streaming (`stream: true`)
- Handles `content_block_delta` events

**`openai.go`** — OpenAI Chat Completions API
- `POST https://api.openai.com/v1/chat/completions`
- SSE streaming (`stream: true`)
- Handles `chat.completion.chunk` events

**`openrouter.go`** — OpenRouter API
- `POST https://openrouter.ai/api/v1/chat/completions`
- OpenAI-compatible format with OpenRouter-specific headers
- `HTTP-Referer`, `X-Title` headers

Provider is selected based on `config.Provider.Active` and instantiated at startup.

### `config/` — Configuration Management

**File layout:**

```
~/.pear/
├── config.toml              # Global config (profile, providers, watch settings)
├── learning.json            # Concept tracking data
└── codebases/
    └── <codebase-name>.toml # Per-codebase overrides
```

**Codebase detection:** Uses the git repo root absolute path, slugified as the key (e.g., `/Users/mitch/Documents/Pear-v0` → `Users-mitch-Documents-Pear-v0.toml`). Full path avoids collisions between repos with the same name in different directories. When Pear starts in a repo, it checks for a matching codebase config and merges overrides onto the global config.

**In-app editing:** `/settings`, `/provider`, `/model`, `/key` commands read and write the config file directly. Changes take effect immediately (no restart needed).

### `learning/` — Concept Tracking & Graph

```go
type ConceptStore struct {
    Concepts map[string]*Concept `json:"concepts"`
}

type Concept struct {
    Count    int       `json:"count"`
    Sessions []string  `json:"sessions"`  // ISO timestamps
    Related  []string  `json:"related"`   // related concept names
}
```

**Extraction:** Regex parses `📚 Concepts: [...]` and `🔗 Related: [x → y, ...]` from LLM responses. Updates counts, appends session timestamps, merges relationships.

**Display (`pear progress`):** Sorted by frequency, shows relationship edges beneath each concept.

### `hooks/` — Git Hook Management

- `pear hooks install`: writes/appends to `.git/hooks/post-commit` with a marker comment (`# pear-hook`)
- `pear hooks uninstall`: removes Pear's lines, deletes hook if empty
- Hook runs `pear review --commit HEAD` — one-shot proactive review

### `tui/` — Bubble Tea Terminal UI

```go
// app.go — main Bubble Tea model
type Model struct {
    input    InputModel      // text input with @file autocomplete
    output   OutputModel     // streaming response display
    mode     string          // "interactive" or "watch"
    watching bool            // proactive reviews active
    history  []llm.Message   // conversation history
    stats    SessionStats    // reviews given, concepts taught, uptime
}
```

**Key components:**

- **InputModel**: text input field with `@file` path autocomplete, slash command detection
- **OutputModel**: streaming markdown rendering (code blocks, bold, lists), scrollable viewport
- **Styles**: Lipgloss styles for separators (`━━━`), context lines (`📎`), concept tags (`📚`), Socratic questions (`🤔`)

**Watch mode hybrid:** The TUI runs a Bubble Tea program that listens for both keyboard input (user questions) and watcher channel events (proactive triggers). Uses `tea.Cmd` functions that block on the watcher channel and return `tea.Msg` values to the update loop.

> **Flag:** `tea.Sub` does not exist in Bubble Tea. The correct pattern is a `tea.Cmd` that listens on the watcher channel:
> ```go
> func waitForTrigger(triggers <-chan ReviewTrigger) tea.Cmd {
>     return func() tea.Msg {
>         return <-triggers // blocks until watcher sends
>     }
> }
> ```
> Re-issue this command after each trigger to keep listening.

---

## TUI State Machine

> **Flag:** This is the hardest engineering problem in v0. Streaming output + user input + watcher events in the same TUI must be explicitly managed.

```
                    ┌──────────────────┐
                    │      IDLE        │
                    │  input focused   │
                    │  waiting for     │
                    │  user or watcher │
                    └───┬──────────┬───┘
                        │          │
            user submits│          │watcher trigger
                        ▼          ▼
              ┌──────────┐  ┌──────────┐
              │STREAMING │  │STREAMING │
              │(reactive)│  │(proactive)│
              │input     │  │input     │
              │disabled  │  │disabled  │
              └────┬─────┘  └────┬─────┘
                   │              │
              stream ends    stream ends
                   │              │
                   ▼              ▼
              ┌──────────────────────┐
              │       IDLE           │
              │  re-focus input      │
              │  re-issue watcher cmd│
              └──────────────────────┘
```

**Key rules:**
- Input is **disabled** while streaming (keystrokes buffered, not sent)
- If a watcher trigger arrives during streaming, it's **queued** (buffered channel, size 1) and processed after current stream completes
- `/pause` sets a `paused` flag — watcher triggers are silently dropped while paused
- Viewport auto-scrolls during streaming, user can scroll up to review (locks auto-scroll until they scroll back to bottom)

---

## Data Flow

### Watch Mode — Proactive Review

```
fsnotify detects file change
        │
        ▼
Watcher updates lastChangeTime, marks unsettled
        │
        ▼
30s passes with no new changes
        │
        ▼
Watcher runs git diff, checks throttle rules
        │
        ▼
Sends ReviewTrigger on channel
        │
        ▼
Bubble Tea receives trigger via tea.Sub
        │
        ▼
Context collector builds RepoContext (diff, tree, branch)
        │
        ▼
Prompt assembler builds proactive prompt + history
        │
        ▼
LLM client streams response → TUI renders chunks
        │
        ▼
Learning tracker extracts concepts + relationships
        │
        ▼
Updates learning.json, session stats
```

### Interactive Mode — User Question

```
User types question (with optional @file refs)
        │
        ▼
TUI parses input, resolves @file paths
        │
        ▼
Context collector builds RepoContext
        │
        ▼
Prompt assembler builds reactive prompt + full history
        │
        ▼
LLM client streams response → TUI renders chunks
        │
        ▼
Learning tracker extracts concepts
        │
        ▼
Response added to conversation history
```

---

## Concurrency Model

```
Main goroutine
  └── Bubble Tea event loop (handles input + rendering)

Watcher goroutine
  ├── fsnotify listener (file events → lastChangeTime)
  └── Ticker (5s) — git HEAD check, settle detection
      └── Sends ReviewTrigger on channel

tea.Sub bridges watcher channel → Bubble Tea Msg
```

**Cancellation:** `context.Context` with cancel propagates shutdown from Ctrl+C through the watcher goroutine.

**Streaming:** LLM response streaming runs in a goroutine, sending chunks via `tea.Sub` to the Bubble Tea update loop for incremental rendering.

---

## Dependencies

| Package | Purpose |
|---|---|
| `github.com/spf13/cobra` | CLI command framework |
| `github.com/charmbracelet/bubbletea` | Terminal UI framework |
| `github.com/charmbracelet/lipgloss` | TUI styling |
| `github.com/charmbracelet/glamour` | Markdown rendering in terminal |
| `github.com/fsnotify/fsnotify` | Filesystem change notifications |
| `github.com/BurntSushi/toml` | Config file parsing |

No external LLM SDKs — HTTP clients are hand-rolled for each provider to keep control over streaming and error handling.

---

## Error Handling

- **Git not in repo:** `pear watch` and context-dependent commands fail fast with a clear message
- **API key invalid:** `pear doctor` catches this; runtime errors show provider + HTTP status
- **Network errors during streaming:** Partial response is displayed, error message appended
- **fsnotify limits:** Falls back to git-only polling if fsnotify can't watch (too many files, unsupported FS)

---

## Evolution Path

```
v0 (bootstrap)             → v1.5 (MVP)
────────────────────────────────────────────────────
Cobra commands               → same (carry forward)
Bubble Tea TUI               → same (carry forward)
fsnotify + git polling       → same (carry forward)
3 LLM providers              → LLM adapter registry (pluggable)
in-memory history            → same (server persistence in v1.6)
regex concept extract        → structured LLM output parsing
basic concept graph          → full knowledge graph + spaced repetition
~/.pear/codebases/           → same (no per-repo files)
no auth                      → license key / trial check
fixed settle time            → adaptive timing
```

Each v0 package carries forward. The code gets extended, not rewritten.

---

## Testing Strategy

| Layer | Approach |
|---|---|
| **Prompt assembler** | Golden-file tests: fixed inputs → expected prompt output, compared with `testdata/*.golden` files |
| **LLM clients** | Record/replay: capture real API responses as fixtures, replay in tests. Mock HTTP transport. |
| **Watcher** | Synthetic test: create temp git repo, make file changes, assert correct triggers on channel with timeouts |
| **TUI** | `teatest` package for Bubble Tea — send key events, assert view output |
| **Context collector** | Test repos in `testdata/`: known git state → expected RepoContext |
| **Concept extraction** | Unit tests with sample LLM responses (varied formatting) → expected concepts |
| **Config** | Round-trip tests: write config → read config → assert equal |

---

## Logging & Observability

Structured logging to `~/.pear/logs/` for dogfooding diagnostics:

| Event | Fields |
|---|---|
| `llm.request` | provider, model, input_tokens, output_tokens, latency_ms |
| `llm.error` | provider, error_code, retry, message |
| `watcher.trigger` | type (settle/commit), diff_lines, files_changed |
| `watcher.skip` | reason (too_small, cooldown, paused) |
| `context.collect` | diff_lines, tree_files, attached_files, truncated |
| `concept.extract` | concepts_found, relationships_found, parse_failures |

Log rotation: keep last 5 session logs. One file per session (`~/.pear/logs/2026-03-02T10:00:00.log`).
