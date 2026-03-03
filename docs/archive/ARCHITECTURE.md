# Pear — Product Vision & Technical Architecture

> Original vision document — the technical blueprint.
> See `PRD.md` for the full product requirements.

---

## Product Vision

**One-liner:** Voice-native pair programming for the terminal — talk to your codebase, not your keyboard.

**Thesis:** The best developers think out loud. Current AI coding tools force developers into a text-chat paradigm that breaks flow state. Voice is the natural interface for *thinking through problems* — but no one has built a voice layer that's native to the terminal, context-aware, and model-agnostic.

**Pear** starts as a voice-first CLI companion that makes any LLM a better teacher and pair programmer by automatically injecting rich repo context into every prompt. It evolves into a voice-first AI tutor that teaches software engineering *in situ* — inside real codebases, during real work.

### Why This Wins

- **Wedge:** Voice input + teaching-first prompt engine is a genuine UX gap in CLI AI tools. No one owns it.
- **Moat:** Closed source protects the teaching engine, prompt templates, and voice integration. The integrated product (voice + context + pedagogy) is hard to replicate as a feature toggle.
- **Teaching gap:** AI tools make developers faster but not better. Nobody is addressing the learning crisis AI copilots are creating.
- **Business model:** BYOK-first means ~100% margin on subscriptions. Users pay for the teaching tool, not LLM access.
- **Acquisition value:** A voice + pedagogy layer is complementary to every existing copilot. Claude Code, Cursor, Windsurf — they all want this but won't build it themselves.

---

## CLI UX — Slash Command System

Inspired by Claude Code's slash command pattern, Pear uses `/` commands for in-session control. This gives users a discoverable, consistent interface without leaving the REPL.

### Command Reference

Type `/` to see all available commands. Type `/` + letters to filter.

#### Session


| Command          | Description                                                 |
| ---------------- | ----------------------------------------------------------- |
| `/help`          | Show available commands and usage                           |
| `/status`        | Show session info: model, mode, context stats, voice status |
| `/clear`         | Clear conversation history, keep session open               |
| `/exit`          | End the session                                             |
| `/history`       | Browse past sessions                                        |
| `/export [file]` | Export current session to markdown                          |
| `/copy`          | Copy last response to clipboard                             |


#### Voice


| Command        | Description                                 |
| -------------- | ------------------------------------------- |
| `/voice`       | Toggle voice input on/off (default: on)     |
| `/speak`       | Toggle TTS audio output on/off              |
| `/mic`         | Show mic status, test recording             |
| `/lang <code>` | Set voice language (ISO 639-1, default: en) |


#### Mode & Model


| Command                     | Description                                                     |
| --------------------------- | --------------------------------------------------------------- |
| `/mode [teach|mentor|pair]` | Switch prompt mode                                              |
| `/model [name]`             | Switch LLM model (e.g., `/model claude-sonnet-4-5-20250929`)    |
| `/role [name]`              | Switch role frame (e.g., `/role security`, `/role performance`) |
| `/verbose`                  | Toggle showing the assembled prompt before sending              |


#### Context


| Command       | Description                                   |
| ------------- | --------------------------------------------- |
| `@<file>`     | Add file(s) to active context                 |
|               |                                               |
|               |                                               |
| `/context`    | Show full context summary (what will be sent) |
| `/log <file>` | Attach an error log file to context           |


#### Config


| Command   | Description                                            |
| --------- | ------------------------------------------------------ |
| `/config` | Open config editor                                     |
| `/init`   | First-time setup wizard (API keys, preferences)        |
| `/doctor` | Check system health (sox, mic, API keys, connectivity) |
| `/keys`   | Manage API keys                                        |
| `/cost`   | Show token usage and estimated costs this session      |


#### Teaching (v1.7+)


| Command           | Description                                      |
| ----------------- | ------------------------------------------------ |
| `/explain <file>` | Full walkthrough of a file's design and patterns |
| `/review`         | Teaching-oriented review of staged changes       |
| `/learn <topic>`  | On-demand concept deep-dive using your codebase  |
| `/concepts`       | Show concepts Pear has taught you in this repo   |


### Input Patterns


| Pattern        | Behavior                                   |
| -------------- | ------------------------------------------ |
| `/`            | Trigger command autocomplete (filterable)  |
| `@`            | File path autocomplete (attach to context) |
| `[Space] hold` | Push-to-talk voice recording               |
| Text + Enter   | Send as text prompt (no voice needed)      |


### Example Session

```
$ pear
🍐 Pear v1.5 · Claude claude-sonnet-4-5-20250929 · teach mode
   Hold [Space] to talk. Type /help for commands.

/status
  Model:   claude-sonnet-4-5-20250929 (via BYOK)
  Mode:    teach
  Role:    senior-staff
  Voice:   on (mic: Built-in Microphone)
  Speak:   off
  Context: git diff (47 lines), tree (23 files), no errors attached
  Session: 0 turns, 0 tokens used

/mode mentor
  Switched to mentor mode (concise answers + one key insight)

@src/auth/rbac.go
  Added src/auth/rbac.go (142 lines) to context

[Space held → voice] "What's the deal with this RWMutex here?"
🎤 Transcribed → enriching...
📎 Context: git diff, src/auth/rbac.go, file tree
🧠 Sending to Claude...

━━━ Pear ━━━
The RWMutex on line 12 protects concurrent access to the
role cache. It's used instead of a regular Mutex because
your middleware reads roles on every request but only
writes on cache refresh — RWMutex lets all those reads
happen in parallel.

💡 Key insight: If you ever add role write-back (e.g.,
   admin role changes), make sure to use `Lock()` not
   `RLock()` for writes, or you'll get a race condition
   that only shows up under load.
━━━━━━━━━━━━

/cost
  This session: 1,247 input tokens, 189 output tokens
  Estimated cost: $0.004 (BYOK Claude claude-sonnet-4-5-20250929)
```

---

## Technical Architecture

### High-Level Diagram

```
┌─────────────────────────────────────────────────────┐
│                    pear CLI (Go)                     │
│                                                      │
│  ┌───────────┐  ┌──────────────┐  ┌──────────────┐  │
│  │ Voice I/O │  │   Context    │  │   Prompt     │  │
│  │  Engine   │  │  Collector   │  │  Assembler   │  │
│  │           │  │              │  │              │  │
│  │ • Record  │  │ • git diff   │  │ • mode frame │  │
│  │ • STT     │  │ • file tree  │  │ • role frame │  │
│  │ • TTS     │  │ • errors     │  │ • pedagogy   │  │
│  │ • Hotkey  │  │ • active     │  │ • context    │  │
│  │           │  │   files      │  │ • user query │  │
│  └─────┬─────┘  └──────┬───────┘  └──────┬───────┘  │
│        │               │                 │           │
│        ▼               ▼                 ▼           │
│  ┌───────────────────────────────────────────────┐   │
│  │              LLM Adapter Layer                │   │
│  │  ┌─────────┐  ┌────────┐  ┌────────┐         │   │
│  │  │ Claude  │  │ OpenAI │  │ Gemini │   ...   │   │
│  │  └─────────┘  └────────┘  └────────┘         │   │
│  └────────────────────┬──────────────────────────┘   │
│                       │                              │
│  ┌────────────────────┴──────────────────────────┐   │
│  │            Billing / Auth Layer                │   │
│  │  • BYOK passthrough (direct to provider)      │   │
│  │  • Hosted proxy (api.pearcode.dev)                │   │
│  └───────────────────────────────────────────────┘   │
└─────────────────────────────────────────────────────┘
```

### Component Breakdown

#### Voice I/O Engine (`internal/voice/`)

```
Strategy: exec-based, not cgo.

Record:  shell out to `sox` (rec) → write to temp .wav file
         Fallback: `ffmpeg` from mic input
         macOS: can also use `say` for TTS, `sox` for record
STT:     POST .wav to Whisper API or Deepgram
         Interface: type Transcriber interface { Transcribe(audio []byte) (string, error) }
TTS:     POST text to OpenAI TTS → stream .mp3 to `afplay` (macOS) / `mpv` (linux)
         TTS reads a SPOKEN SUMMARY (2-3 sentences of teaching points),
         not the full response. Full detail + code blocks stay in text output.
         The LLM generates both: text response + spoken_summary field.
         Interface: type Speaker interface { Speak(text string) error }
Hotkey:  Raw terminal mode — detect spacebar hold/release.
         Use golang.org/x/term for raw mode.
         Alternative: configurable key via config.
```

**Why exec-based for audio:** cgo + PortAudio is a build/distribution nightmare for a single-binary Go CLI. Requiring `sox` (or `brew install sox`) is a one-time friction that saves weeks of cross-compilation pain. Document it clearly, check on startup, offer `/doctor` to verify deps.

#### Context Collector (`internal/context/`)

```go
type RepoContext struct {
    GitDiff      string   // git diff HEAD (staged + unstaged)
    GitStatus    string   // short status
    FileTree     string   // depth-limited, .gitignore-aware
    ErrorLog     string   // last N lines from specified log or stderr
    ActiveFiles  []string // explicitly referenced files, included in full
    RepoName     string
    Branch       string
}

// Collector gathers context with configurable limits
type Collector struct {
    MaxDiffLines   int  // default 200, truncate with "[truncated]"
    MaxTreeDepth   int  // default 3
    MaxErrorLines  int  // default 30
    MaxFileLines   int  // default 500 per file
}
```

Context is collected in parallel (git diff, file tree, error log are independent). Use `errgroup` for concurrent collection with a timeout.

#### Prompt Assembler (`internal/prompt/`)

Prompt structure varies by mode:

**Teach mode (default):**

```
[SYSTEM]
You are a senior staff engineer pair-programming with the developer.
Your primary role is to TEACH — not just solve problems, but help the
developer understand the concepts behind the solution.

When responding:
- Diagnose the immediate issue first (be useful)
- Explain the underlying concept or pattern (be educational)
- Note why this matters in production / at scale (be practical)
- Reference established best practices when relevant
- Offer to go deeper on any concept
- Keep explanations grounded in their actual code

Repo: {repo_name} Branch: {branch}

[USER]
## Repository Context
### Git Diff
{git_diff}
### File Tree
{file_tree}
### Error Log
{error_log}
### Active Files
{active_file_contents}

## Developer Question
{transcribed_voice_input}
```

**Mentor mode:**

```
[SYSTEM]
You are a senior staff engineer pair-programming with the developer.
Give the direct answer, then add one key insight or lesson learned.
Be concise — respect their time. Cite specific files and line numbers.
```

**Pair mode:**

```
[SYSTEM]
You are a senior staff engineer pair-programming with the developer.
Be direct. Cite specific files and line numbers. Suggest concrete fixes.
No explanations unless asked.
```

#### LLM Adapter Layer (`internal/llm/`)

```go
type LLM interface {
    Chat(ctx context.Context, req ChatRequest) (*ChatResponse, error)
    StreamChat(ctx context.Context, req ChatRequest) (<-chan StreamChunk, error)
    Name() string
    EstimateCost(inputTokens, outputTokens int) float64
}

// BYOK adapters (primary — user's own API keys, direct to provider)
type ClaudeAdapter struct  { apiKey string; model string }
type OpenAIAdapter struct   { apiKey string; model string }
type GeminiAdapter struct   { apiKey string; model string }

// Hosted adapter (secondary — 50 requests/month via api.pearcode.dev)
type PearProxyAdapter struct { sessionToken string }
```

Use streaming everywhere. UX: user sees response tokens while TTS buffers the first sentence.

#### Billing / Auth Layer (`internal/billing/`)

```
BYOK mode (primary):
  Keys stored in ~/.pear/config.toml (file perm 0600)
  Direct API calls to LLM providers — Pear never sees the key or the code.
  Unlimited usage — user pays their LLM provider directly.

Hosted mode (secondary, 50 requests/month with Pro):
  - pear login → GitHub OAuth (opens browser → Go backend)
  - Receives session JWT
  - Hosted LLM calls route through api.pearcode.dev proxy
  - Proxy: lightweight Go service on Fly.io
  - Validates JWT → checks hosted request quota → proxies to upstream LLM → logs usage
  - Stripe subscription via web dashboard
  - Pro: $30/mo or $300/yr
  - Usage tracking: hosted request count per billing cycle, cap at 50/month
  - Additional hosted usage tiers available for purchase
```

#### Slash Command System (`internal/commands/`)

```go
type Command struct {
    Name        string
    Aliases     []string
    Description string
    Args        string // e.g., "[teach|mentor|pair]"
    Category    string // session, voice, mode, context, config, teaching
    Run         func(args []string, session *Session) error
}

// Registry
type CommandRegistry struct {
    commands map[string]*Command
}

func (r *CommandRegistry) Match(input string) (*Command, []string)
func (r *CommandRegistry) Autocomplete(prefix string) []*Command
func (r *CommandRegistry) ByCategory() map[string][]*Command
```

Commands are registered at startup. The TUI intercepts lines starting with `/` and routes to the command registry instead of the LLM. Autocomplete triggers on `/` keypress in raw terminal mode, showing a filterable list (Bubble Tea list component).

### Project Structure

Single private monorepo. See PRD.md for the full directory layout. Types are declared independently in each Go module — no shared types package.

```
pear/                                   # private monorepo
├── cli/                               # Go CLI (compiled binary)
│   ├── cmd/pear/main.go              # CLI entrypoint (Cobra)
│   ├── internal/
│   │   ├── voice/                    # mic capture, STT, TTS
│   │   ├── context/                  # parallel context gathering
│   │   ├── prompt/                   # prompt construction + modes
│   │   ├── llm/                      # LLM adapters (Claude, OpenAI, Gemini, hosted)
│   │   ├── session/                  # multi-turn conversation history
│   │   ├── mcp/                      # MCP server — v1.5.1 (stdio JSON-RPC)
│   │   ├── commands/                 # slash command registry
│   │   ├── auth/                     # pear login client
│   │   ├── config/                   # ~/.pear/config.toml parsing
│   │   └── tui/                      # interactive terminal UI (Bubble Tea)
│   └── prompts/                      # teaching mode system prompts
├── api/                               # Go Backend (Fly.io)
│   ├── cmd/api/main.go
│   └── internal/                     # auth, proxy, billing, usage, analytics, db
├── web/                               # Next.js (Vercel)
│   ├── app/                          # marketing + dashboard
│   └── components/
├── scripts/install.sh                 # curl installer
└── Makefile
```

### Key Technical Decisions


| Decision       | Choice                     | Rationale                                                                           |
| -------------- | -------------------------- | ----------------------------------------------------------------------------------- |
| Language       | Go                         | Single binary, fast startup, you know it                                            |
| Audio capture  | `sox` subprocess           | Avoid cgo. One `brew install sox` vs. weeks of PortAudio                            |
| Terminal UI    | Bubble Tea (charmbracelet) | Best Go TUI library. Handles raw mode, hotkeys, styled output                       |
| CLI framework  | Cobra                      | Standard, well-documented                                                           |
| Config         | TOML                       | Human-readable, dual-level: `~/.pear/config.toml` (global) + `pear.toml` (per-repo) |
| STT primary    | Whisper API (OpenAI)       | Best accuracy, most users already have an OpenAI key                                |
| TTS            | OpenAI TTS API             | Low-latency, good quality                                                           |
| Hosting        | Fly.io                     | Simple Go deploy, global edge, cheap at low scale                                   |
| Billing        | Stripe (upfront tiers)     | BYOK-first. Pro $30/mo or $300/yr. 50 hosted requests included. Stripe webhooks direct to Go backend. |
| Telemetry      | PostHog (opt-in)           | Lightweight, open-source friendly, Go SDK                                           |
| Slash commands | Custom registry            | Discoverable, consistent, Claude Code-inspired UX                                   |


---

## Prompt Refinement Strategy

This is the core IP. The gap between "raw voice transcript" and "high-quality teaching prompt" is where Pear creates value.

### Refinement Pipeline (v1.5)

```
Raw voice transcript
       │
       ▼
┌─────────────────┐
│ 1. Cleanup       │  Fix speech artifacts: "um", "like",
│                  │  repeated words, false starts
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ 2. Attach all    │  All context types attached on every request:
│    context       │  diff + tree + errors + active files + history
│                  │  LLM determines what's relevant
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ 3. Mode frame    │  Apply teaching/mentor/pair system prompt
│                  │  based on current /mode setting
└────────┬────────┘
         │
         ▼
┌─────────────────┐
│ 4. Assemble      │  System prompt + context + conversation
│                  │  history + cleaned query
└────────┬────────┘
         │
         ▼
     Final prompt
```

**No intent detection in v1.5.** All available context is attached on every request. Modern LLMs handle large context well — the risk of misclassifying intent with a naive keyword matcher outweighs the cost savings of selective context. Intent-based context selection is deferred to v1.7 when real usage data can inform a proper classifier.

### Context Budget Management

```
Total token budget: ~25k tokens for context

All requests get the same context (truncated to budget):
  diff:           up to 8k tokens
  tree:           up to 3k tokens
  active_files:   up to 8k tokens
  error_log:      up to 2k tokens
  history:        up to 4k tokens (last N turns, token-aware)

Budget is generous — modern models (Sonnet, Opus, GPT-4o, Gemini Pro)
handle 128k+ context. Better to send too much than truncate useful info.

Truncation: from the middle, keeping first and last N lines.
Signal: "[...truncated {N} lines...]"
```

### Role Frames (Configurable)

```toml
# ~/.pear/config.toml
[prompt]
mode = "teach"                    # teach | mentor | pair
role = "senior-staff"             # or "security", "performance", "architecture"

# Custom role override
[prompt.custom_role]
name = "my-style"
system = "You are a pragmatic Go engineer who values simplicity over cleverness..."
```

Built-in roles ship with Pear. Users override via config or `/role` command.

---

## Distribution Strategy

Pear is closed source. Distribution comes from content marketing and launch events, not GitHub virality.

1. **LinkedIn** (5,000 followers) — founder's existing developer audience
2. **Substack** — essays on the AI-assisted learning crisis
3. **X (Twitter)** — new account, demo videos, dev community engagement
4. **Product Hunt launch** — "AI tutor in your terminal" angle
5. **Hacker News** — technical blog post + launch discussion
6. **Dev communities** — r/programming, r/ChatGPTCoding, relevant Discord servers

### Installation

```bash
curl -fsSL https://pearcode.dev/install.sh | sh    # primary
brew install pearcode/tap/pear                       # macOS Homebrew
```

### The Demo That Gets Attention

60-second screen recording: developer holds space, says "I just accepted a bunch of Claude suggestions — walk me through what changed and tell me if it's solid." Pear shows the prompt enrichment in real-time (context injection visible), then responds with a structured breakdown of each change, explains the patterns, flags a subtle issue, and offers to go deeper. The viewer *sees* the difference between "here's the answer" and "here's the answer *and here's what you should learn from it*."

---

## Latency Budget

Voice-first products live or die on latency. Target: <3s from releasing spacebar to first visible response token.

**Key optimization: start context collection during recording, not after.**

```
User holds spacebar → recording starts
                   ↓ (parallel)
         ┌─────────────────────────────┐
         │ Context collector runs      │  git diff, tree, active files
         │ concurrently with recording │  all gathered while user talks
         └─────────────────────────────┘
User releases spacebar → recording stops
                   ↓
         ┌─────────────────────────────┐
         │ sox writes wav              │  ~100ms
         └─────────┬───────────────────┘
                   ↓
         ┌─────────────────────────────┐
         │ Whisper API transcription   │  ~800ms (turbo model, <15s audio)
         └─────────┬───────────────────┘
                   ↓ (context already collected)
         ┌─────────────────────────────┐
         │ Assemble prompt + send      │  ~100ms
         └─────────┬───────────────────┘
                   ↓
         ┌─────────────────────────────┐
         │ LLM first token             │  ~1000ms (Sonnet, streaming)
         └─────────────────────────────┘

Total to first visible token: ~2.0s
TTS first audio chunk: +800ms after first sentence completes
```

Context collection (300-500ms) happens entirely during recording and costs zero latency on the critical path.

---

## Voice Error Recovery

Every voice failure mode must degrade gracefully to text input. The session should never crash.

| Failure | Detection | Recovery |
|---|---|---|
| sox not installed | `pear doctor` + check on first voice attempt | Print install command, fall back to text-only mode |
| sox crashes mid-recording | Process exit code != 0 | "Recording failed. Try again or type your question." Discard partial wav. |
| Whisper returns empty/garbage | Transcript empty or under 3 words | "I didn't catch that. Try again?" Don't send to LLM. |
| Mic stolen by another process | sox returns error or silence | "Mic appears unavailable. Check if another app is using it." |
| Utterance too short (<1s) | File size check before Whisper | Skip silently (treat as accidental tap) |
| Utterance too long (>60s) | Timer during recording | Auto-stop at 60s, warn user, process what you have |
| Whisper API timeout | HTTP timeout >5s | "Transcription failed. Try again or type your question." |
| Network down during STT | Connection error | "No network. Voice requires internet. Type your question instead." |

---

## Moat & Evolution Timeline

```
v1.5 (weeks 1-5):  Closed source launch. Teaching engine, voice UX, context
                    injection, BYOK-first model. Moat = proprietary code +
                    first-mover in voice AI tutor category.
v1.5.1 (week 6):   MCP server mode. Distribution play — ships after core
                    CLI is stable and has initial users.
v1.6 (weeks 6-9):  Server-side session storage + prompt versioning.
                    Data flywheel begins. Linux support.
v1.7 (weeks 10-15): Intent detection from real data. Concept tracking +
                     per-repo memory. Meaningful data moat.
v1.8 (weeks 16-22): Knowledge graph. Deep moat — years to replicate.
v2.0 (weeks 23-32): Progress tracking + skill assessment + team tier.
                     Full proprietary education platform.
```