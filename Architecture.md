# Pear — Architecture Overview

> Last updated: March 2026

---

## Monorepo Structure

```
Pear-v0/
├── cli/                    # Go CLI (Cobra + Bubble Tea + fsnotify)
├── website/                # Next.js 15 marketing site (Vercel)
├── docs/
│   ├── cli/                # CLI-specific docs (architecture, events, user journeys)
│   ├── archive/            # Historical strategy/analysis docs
│   └── v0-tickets/         # Completed implementation tickets
├── Vision.md               # Product vision and positioning
├── PRD.md                  # Product requirements
├── Architecture.md         # This file
├── CLAUDE.md               # Claude Code project instructions
├── install.sh              # CLI install script
└── README.md               # CLI-focused README
```

---

## CLI Architecture

**Language:** Go
**Framework:** Cobra (CLI) + Bubble Tea (TUI) + fsnotify (file watching)
**LLM:** Hand-rolled HTTP clients for Anthropic, OpenAI, OpenRouter (no SDKs)

For detailed CLI architecture, see [docs/cli/ARCHITECTURE.md](docs/cli/ARCHITECTURE.md).

### Packages

| Package | Responsibility |
|---------|---------------|
| `cmd/` | Cobra command definitions (root, watch, ask, review, teach, doctor, progress, hooks, init) |
| `tui/` | Bubble Tea app, input/output components, styles, slash commands, settings |
| `watcher/` | Hybrid fsnotify + git polling, pause detection, commit detection |
| `llm/` | LLMClient interface + Anthropic, OpenAI, OpenRouter streaming clients |
| `prompt/` | System prompt assembly (proactive, reactive, deep-dive variants) |
| `repocontext/` | Git diff, file tree, @file resolution, context building |
| `config/` | ~/.pear/config.toml read/write, per-codebase overrides |
| `learning/` | Concept extraction from responses, learning.json persistence |
| `hooks/` | Git post-commit hook install/uninstall |
| `logging/` | Structured JSON logs to ~/.pear/logs/ |

### Key Design Decisions

- **No `context` package name** — uses `repocontext` to avoid stdlib shadow
- **No `tea.Sub`** — doesn't exist in Bubble Tea. Channel-based commands instead.
- **System prompts via `StreamOptions`** — not in the message array. Each provider handles placement internally.
- **Buffered channels (size 1)** — for watcher triggers. Prevents blocking, latest-wins semantics.
- **Atomic file writes** — write to temp file, then rename. For config and learning.json.

### TUI State Machine

```
IDLE → STREAMING → IDLE (→ process queued trigger if any)

IDLE: input enabled, watcher triggers processed
STREAMING: input disabled, watcher triggers queued
PAUSED: input enabled, watcher triggers dropped
```

### User Data

All stored locally under `~/.pear/`:

```
~/.pear/
├── config.toml              # Global config (provider, model, API key, user profile)
├── learning.json            # Concept tracking data
├── codebases/<path-slug>.toml  # Per-repo overrides
└── logs/<timestamp>.log     # Session logs (structured JSON)
```

---

## Website Architecture

**Framework:** Next.js 15 (App Router, Turbopack)
**Runtime:** Node.js
**Hosting:** Vercel
**Styling:** Tailwind CSS 4, shadcn/ui, Framer Motion

### Structure

```
website/
├── app/
│   ├── (marketing)/        # Public pages (home, blog, pricing, docs, about, FAQ, compare)
│   ├── (dashboard)/        # Future authenticated dashboard
│   └── api/waitlist/       # Waitlist API endpoint
├── components/             # React components (hero, pricing, problem, solution, etc.)
├── emails/                 # Resend email templates (waitlist welcome, admin notify)
├── lib/                    # Utilities, theme provider
├── public/                 # Static assets (logos, images, llms.txt)
└── docs/                   # Website-specific docs (positioning, design guidelines)
```

### Services

- **Resend** — Transactional email (waitlist welcome, admin notifications)
- **Upstash Redis** — KV store for waitlist data, rate limiting
- **Vercel Analytics** — Traffic and performance monitoring

---

## Shared Concepts

### Concept Model

Both CLI and website reference the same concept model:

- **Concepts** are tagged in LLM responses (e.g., `[concept:error-handling]`)
- **Relationships** between concepts are tracked (e.g., `[related:interfaces→polymorphism]`)
- Currently stored in local `~/.pear/learning.json`
- Future: synced to server for Pro tier, aggregated for Teams

### Config Format

TOML-based configuration. Global config at `~/.pear/config.toml`, per-repo overrides at `~/.pear/codebases/<slug>.toml`.

---

## Infrastructure

```
Developer's machine          Vercel                   External
┌─────────────────┐    ┌─────────────────┐    ┌──────────────┐
│  pear CLI        │    │  website/        │    │  Anthropic   │
│  ~/.pear/        │    │  Vercel hosting  │    │  OpenAI      │
│  (all user data) │────│  Upstash Redis   │    │  OpenRouter   │
│                  │    │  Resend email    │    │              │
└────────┬─────────┘    └─────────────────┘    └──────┬───────┘
         │                                            │
         └──────── LLM API calls (user's key) ────────┘
```

No server-side user data storage in v0-v1. CLI talks directly to LLM providers using the user's API key. Website is marketing + waitlist only.
