# Pear — Product Requirements Document

> Version: 1.5 (MVP)
> Last updated: February 2026
> Author: Mitch
>
> **Note:** This PRD describes the original vision and full product scope. For the current MVP scope and all decisions that override this document, see `PRODUCT_DECISIONS.md` — that is the authoritative source. Where these documents conflict, Product Decisions wins.

### Builder Context

**Solo founder.** Strongest in TypeScript and React; junior-to-intermediate in Go. Pear is both a product and a Go learning project — building it deepens understanding of the language it teaches others about.

**AI-assisted development.** v1.5 is being built with Claude Code and Cursor as primary development tools. Target: ship in ~5 weeks. The web layer (Next.js, TypeScript, React) is a weekend's work. The CLI (Go) is the learning surface and the time investment.

**Selective open source (pending).** The teaching engine (prompts, context collection, concept tracking) is proprietary. Some infrastructure components (CLI framework, MCP transport, config system) may be open-sourced for GitHub credibility and HN goodwill. Exact scope is a pending decision — see `PRODUCT_DECISIONS.md`.

**Implication for architecture:** Lean toward simplicity in Go code. Avoid clever abstractions. Prefer explicit, readable patterns that a junior-intermediate Go developer can maintain and debug. Let the AI tools handle boilerplate; invest human attention in the teaching engine and the parts that require taste.

---

## Product Vision

**One-liner:** The only learning tool that teaches while you work.

**Thesis:** AI coding tools are making developers faster but not better. Junior and mid-level engineers are shipping code they don't fully understand — accepting Copilot suggestions, copy-pasting Claude output, vibe-coding through Cursor. They're productive but not *learning*. The gap between "code that works" and "code you understand" is widening.

**Pear closes that gap.** It's a CLI companion that sits alongside your existing AI coding tools and teaches you *why* — why that pattern exists, why the LLM suggested that approach, why the alternative is dangerous, why this matters at scale. It uses your actual codebase, your real diffs, your live errors as the teaching surface.

**Text-first, voice as an enhancer.** The teaching engine must deliver full value through text alone. Voice input and TTS output are Pro feature unlocks that enhance the experience but are not the product. This de-risks launch: if voice latency sucks, text mode still works. Some users will always prefer text — that's fine.

**Pear is not a replacement for Claude Code or Cursor. It's the mentor those tools will never be.**

---

## Problem Statement

### The AI-Assisted Learning Crisis

1. **AI tools optimize for output, not understanding.** Copilot autocompletes, Claude generates, Cursor agents execute. None of them pause to ask "do you understand what just happened?"

2. **Developers are losing the learning reps.** The struggle of figuring things out — debugging, reading docs, reasoning through architecture — is where deep understanding forms. AI tools skip that struggle entirely.

3. **Senior engineers aren't scaling.** Every team has a 10:1 ratio of juniors who need mentorship to seniors who can provide it. AI tools exacerbate this by creating an illusion of competency.

4. **Existing learning resources are disconnected from real work.** Udemy courses, documentation, and tutorials exist in a vacuum. They don't activate *in the moment* when a developer encounters a real pattern in their real codebase.

### Who this is for

**Primary:** Junior and mid-level developers (1-5 years experience) who use AI coding tools daily and want to genuinely understand what they're building, not just ship it. These developers already pay for Cursor, Claude API, OpenAI API, and online courses (Udemy, etc.) — they have established spending habits for tools and education.

**Secondary:** Software businesses and engineering teams that want to offer Pear as a perk and L&D tool for their engineers — and free up Senior/Lead/Staff Engineers from mentoring time. Team tier unlocks this audience.

**Tertiary:** Self-taught developers, bootcamp graduates, and career-switchers who lack the CS fundamentals that AI tools paper over.

---

## Solution — v1.5 MVP

### What Pear Is

A Go CLI binary (`pear`) that:

1. **Enriches** — automatically injects repo context (git diff, file tree, error logs, active files) and pedagogical framing into the prompt
2. **Teaches** — responds with structured explanations grounded in your actual code, not abstract examples
3. **Remembers** — tracks concepts taught per session, builds a local learning profile
4. **Integrates** — exposes an MCP server so Claude Code, Cursor, and other MCP-capable tools can invoke Pear's teaching capabilities natively
5. **Listens** (Pro) — optional push-to-talk voice input in the terminal
6. **Speaks** (Pro) — optional TTS audio output so you can learn while your hands stay on the keyboard

### The Teaching Difference

When a developer says *"why is this auth middleware returning 403?"*, a pair programmer says:

> "The bug is on line 47 — you're using `==` against a slice. Use `slices.Contains()`."

Pear says:

> "The bug is on line 47. You're comparing a string to a slice with `==`, which always returns false in Go — slices aren't comparable with `==`. The fix is `slices.Contains()`. **This is a common Go gotcha.** Unlike Python where `in` works on lists, Go requires explicit membership checks. This pattern matters because RBAC logic like this is a security boundary — a silent `false` here means unauthorized access could slip through. Want me to explain how Go's type system handles equality, or should we move on?"

Same diagnosis. But Pear *teaches the concept*, *explains why it matters*, and *offers to go deeper*.

### User Flow

```
$ pear init                          # one-time: 3-question wizard
🍐 Welcome to Pear! Let's get you set up.
  What's your name? Mitch
  What languages do you work with? Go, TypeScript, React
  How would you describe your level? mid (3-5 years)
  ✓ Profile saved. Enter your API key to get started...
  Provider? openrouter
  API key? sk-or-...
  ✓ Config saved. Running pear doctor...
  ✓ All checks passed.

$ pear                               # start session in current repo
🍐 Pear active. Type your question or use /help.

> I just accepted a bunch of Claude suggestions in this file.
  Can you walk me through what changed and whether it's solid?

📎 Context: git diff (3 files changed), src/auth/rbac.go,
   file tree | teaching mode

━━━ Pear ━━━
Looking at your diff, Claude made 3 changes to rbac.go. Let me
walk through each:

1. **Line 23: Switched from sync.Mutex to sync.RWMutex** — This is
   a good change. Your original mutex blocked all readers even when
   no write was happening. RWMutex lets multiple goroutines read
   concurrently. This matters when you have high read traffic on
   the RBAC cache, which you do based on your middleware pattern.

2. **Line 41: Added context.WithTimeout** — Also good. Without this,
   a slow DB call in role lookup could hang your request indefinitely.
   The 5s timeout is reasonable for a DB call. One thing to watch:
   make sure you're checking `ctx.Err()` after the call, not just
   the error return...

3. **Line 58: Changed error handling from panic to log+return** —
   This is critical. Panics in HTTP handlers crash the goroutine and
   can bring down your server under load. The log+return pattern is
   standard in production Go. This is an SRE best practice...

Want me to explain any of these patterns deeper?
━━━━━━━━━━━━
📚 Concepts: [sync.RWMutex, context.WithTimeout, panic recovery]
```

---

## MVP Scope

### What Ships in v1.5 (MVP — Ruthless Cut)

> **Scope authority:** `PRODUCT_DECISIONS.md` MVP Scope (Round 3). Ship the smallest thing that makes someone message a friend. CLI + BYOK + teach mode + context injection. No billing, no dashboard, no hosted mode, no voice at launch.

| Feature | Details |
|---|---|
| **Teaching-first prompt engine** | System prompt instructs LLM to explain *why*, cite patterns, reference best practices, offer to go deeper. Structured format: diagnosis → concept → production context → depth offer. Not just "fix the bug" — "fix the bug and teach the concept." |
| **Context injection** | Auto-attaches: git diff (staged + unstaged), error logs (stderr capture or `pear log <file>`), file tree (.gitignore-aware, depth-limited), active file(s) if specified. All context types attached on every request — the LLM determines relevance. Context budget and per-type allocations will be refined during development and beta testing. |
| **Multi-turn conversation** | In-memory conversation history within a session. Users can ask follow-up questions ("explain that deeper", "wait, why?"). Cleared on `/clear` or session exit. No server-side persistence in v1.5. |
| **Local concept tracking** | LLM tags concepts per response, stored locally in `~/.pear/learning.json`. `pear progress` shows concepts with frequency + session streak. |
| **BYOK LLM support** | Adapters for Claude API, OpenAI API, Gemini API, and **OpenRouter API**. Config in `~/.pear/config.toml`. BYOK is the only usage mode at launch — users bring their own LLM keys and pay Pear for the teaching tool, not for LLM access. **OpenRouter gives users access to dozens of cheaper, performant models (Llama, Mistral, DeepSeek, Qwen, etc.) — teaching concepts and explaining patterns doesn't require frontier-model inference. A $3/M-token model teaches goroutines just as well as a $15/M-token model.** |
| **Setup wizard (`pear init`)** | 3-question wizard on first launch (~30 seconds): (1) Name, (2) Primary languages/frameworks, (3) Self-declared level + years of experience. This personalizes the first teaching response — pear "knows you" from the first interaction. After wizard: drop into the prompt. First question gets a response calibrated to their level. Profile stored locally (`~/.pear/profile.json`). `pear doctor` runs automatically on first install to validate setup. |
| **14-day free trial** | Full Pro experience (BYOK, all modes, voice, concept tracking, unlimited questions). No credit card required. Converts users who experience the full product before deciding. After trial ends, free users can upgrade to Pro ($30/mo) or request a trial extension for specific use cases. |
| **Pro tier ($30/mo)** | BYOK with all modes, voice (when shipped), concept tracking, full features. Annual: $300/yr (save $60/yr). $30/mo is the price. Raise later when retention data justifies it. Matches the professional tier for developer tools (Cursor Pro, Claude Pro). |
| **CLI commands** | `pear` (interactive session), `pear ask "question"` (one-shot text), `pear progress` (learning profile), `pear doctor` (system health check), `pear init` (setup wizard). |
| **Installation** | `curl -fsSL https://pearcode.dev/install.sh \| sh` (primary). Homebrew tap: `brew install pearcode/tap/pear` (secondary, macOS). Detects OS + architecture, downloads compiled Go binary. No runtime dependencies at launch (sox only needed when voice ships). Installer runs `pear doctor` automatically after install. |
| **Marketing website** | `pearcode.dev` — landing page with positioning, demo video/screenshot, pricing, docs, blog, and install command. Already built with Next.js + Tailwind, deployed on Vercel. Optimised for Product Hunt and HN launches. |
| **Platform** | CLI: macOS (Linux in v1.6). Web: all browsers. |

### What Ships in v1.5.1 (MCP + Public Launch)

| Feature | Details |
|---|---|
| **MCP server mode** | Pear runs as an MCP server (`pear mcp`), exposing teaching tools to Claude Code, Cursor, and any MCP-capable client. Tools: `pear_teach`, `pear_review`, `pear_explain`. Ships AFTER closed beta validates teaching quality. |
| **Voice input/output** | Push-to-talk voice input (sox + Whisper API) + optional TTS output (OpenAI TTS). Feature-flagged, ships when ready. Pro feature unlock. Text-based teaching must stand on its own first. |

### Explicitly Deferred from MVP

| Component | Ships in | Why defer |
|---|---|---|
| Voice (push-to-talk + TTS) | v1.5.1 | Text must stand on its own. Voice is UX sugar, not the product. |
| Billing / Stripe / dashboard | v1.6 | Invoice first 50 users manually. No billing infrastructure at launch. |
| Hosted LLM mode | v1.6+ | Solo founder can't subsidize LLM costs pre-revenue. BYOK only. |
| Go backend (api.pearcode.dev) | v1.6 | No backend needed for BYOK-only, local-first MVP. |
| Auth system (GitHub OAuth) | v1.6 | Local-first for beta. Server sync when auth ships. |
| Web dashboard (app.pearcode.dev) | v1.6 | No dashboard needed without billing. |
| Mentor and pair modes | v1.5.1+ | Teach mode must prove value first. |
| MCP server mode | v1.5.1 | Distribution play. Ships after core CLI is stable and teaching quality is validated. |
| Intent detection | v1.7 | Attach all context, let LLM sort it out. Build classifier when there's usage data. |
| Linux support | v1.6 | macOS-only cuts platform surface area for launch. |
| Per-repo config (`pear.toml`) | v1.7 | Ship with `~/.pear/config.toml` only. |
| Concept tracking server sync | v1.6 | Local tracking first. Server sync when auth ships. |

### Prompt Modes

Users can set the prompt mode to control how Pear responds. **MVP ships with teach mode only.** Mentor and pair modes ship in v1.5.1+.

| Mode | Behavior | Who it's for | Ships |
|---|---|---|---|
| `teach` (default) | Explain the concept, cite patterns, offer to go deeper | Juniors, learners | **MVP** |
| `mentor` | Give the answer + one key insight, stay concise | Mid-level devs | v1.5.1 |
| `pair` | Direct, minimal answers — just solve the problem | Seniors, power users | v1.5.1 |

Set via `~/.pear/config.toml` or in-session with `/mode teach`.

---

## Non-Goals (v1.5 MVP)

- **NOT a code editor or executor.** Pear reads your repo. It does not write files, run commands, or modify code. It's advisory and educational.
- **NOT a replacement for Claude Code / Cursor / Aider.** It's the teaching layer those tools lack. Use them together — Pear's MCP server mode (v1.5.1) is designed for exactly this.
- **NOT a structured course.** No lessons, no modules, no quizzes. Teaching is contextual — triggered by your real code, in the moment. Structured curriculum is v2.
- **NOT voice-first at launch.** Voice is a Pro feature unlock shipping in v1.5.1. The text-based teaching engine must stand on its own.
- **No billing at launch.** Invoice first 50 users manually. Stripe and dashboard are v1.6.
- **No backend at launch.** BYOK-only, local-first. No hosted mode, no server-side anything.
- **No IDE plugin.** Terminal only.
- **No Windows or Linux at launch.** macOS only. Linux in v1.6.
- **No intent detection.** All context types are attached on every request. The LLM determines what's relevant. Intent-based context selection is deferred until real usage data shows it's needed.

---

## Risks

| Risk | Severity | Mitigation |
|---|---|---|
| **"Why not just ask Claude directly?"** | HIGH | The teaching framing + context injection must produce visibly better educational responses than raw ChatGPT/Claude. Show prompt enrichment by default for the first 3 sessions so users see the value. Demo this relentlessly. The "aha moment" is seeing a response calibrated to *your level* using *your code*. |
| **Teaching quality is inconsistent** | HIGH | The system prompt is the product. Invest heavily in prompt engineering for pedagogical quality. Collect user feedback on response quality early. Iterate prompts in CLI releases (server-side iteration in v1.6). |
| **"Just a wrapper" perception** | HIGH | Lead with teaching + context injection + concept tracking + personalised learning profile. The pitch is "AI tutor that knows your code and your skill level," not "ChatGPT in your terminal." |
| **Developers don't think they need to learn** | MEDIUM | Position around the Dunning-Kruger gap AI tools create. Marketing angle: "AI makes you fast. Pear makes you good." |
| **Claude Code / Cursor adds teaching mode** | HIGH | Move fast. Ship in ~5 weeks. The pedagogical engine, context injection, concept tracking, and personalised learning profile as an integrated product are hard to replicate as a feature toggle. |
| **OpenRouter model quality varies** | MEDIUM | Test teaching prompts against 5+ OpenRouter models. Provide recommended models in docs. The prompt assembler may need model-specific tuning (e.g., smaller models need more explicit structure in the system prompt). |
| **Trial conversion rate too low / too high** | MEDIUM | 14-day free trial with full Pro access is the starting point. Monitor conversion rate from trial-end → paid. If <10%, extend trial or expand trial capabilities. If >30%, shorten trial or tier features differently. Trial window is easy to adjust server-side in v1.6. |
| **No billing infrastructure at launch** | MEDIUM | Manual invoicing for first 50 users. Stripe ships in v1.6. Risk: some users won't pay via manual invoice. Mitigation: only need 20-50 paying users to validate, manual process is fine at that scale. |

---

## Success Metrics (First 90 Days)

> **Detailed phase-by-phase metrics:** See `DISTRIBUTION_GUIDE.md` Part 6.

**Must-hit (product validation):**

| Metric | Target | Why it matters |
|---|---|---|
| Weekly active CLI users | 50+ | Real adoption — people using the product regularly |
| Sessions / user / week | 3+ | Retention — are people coming back? |
| Time-to-first-value | <90 seconds | Install → first useful response. Critical onboarding metric. |
| Day 7 retention | >25% | Are users coming back after the first week? |
| Qualitative feedback | Teaching quality noticeably better than raw Claude | Core product hypothesis validation |

**Stretch (business validation):**

| Metric | Target | Why it matters |
|---|---|---|
| Pro conversions | 20+ paying users | Revenue validation |
| MRR | $900+ | Sustainable signal at ~$30/user (target: 30 paying users) |
| pearcode.dev → install conversion | >25% | Is the landing page doing its job? |
| HN/PH front page | At least one | Launch distribution |
| K-factor (virality coefficient) | >0.5 | Are users sharing organically? |

---

## Technical Architecture

### High-Level

> **MVP architecture is local-first.** The Go backend, dashboard, billing, and auth are all deferred to v1.6. At launch, pear is a CLI binary that calls LLM APIs directly via BYOK.

```
┌────────────────────────────────────────────────────────────┐
│                     pear CLI (Go) — MVP                     │
│                                                             │
│  ┌──────────────┐  ┌──────────────┐  ┌──────────────┐      │
│  │   Context     │  │   Prompt     │  │   Concept    │      │
│  │  Collector    │  │  Assembler   │  │   Tracker    │      │
│  │               │  │              │  │              │      │
│  │ • git diff    │  │ • mode frame │  │ • tag per    │      │
│  │ • file tree   │  │ • role frame │  │   response   │      │
│  │ • errors      │  │ • pedagogy   │  │ • learning   │      │
│  │ • active      │  │ • context    │  │   .json      │      │
│  │   files       │  │ • user query │  │ • progress   │      │
│  │ • profile     │  │ • profile    │  │   command    │      │
│  └──────┬───────┘  └──────┬───────┘  └──────────────┘      │
│         │                 │                                  │
│         ▼                 ▼                                  │
│  ┌─────────────────────────────────────────────────┐        │
│  │              LLM Adapter Layer (BYOK)           │        │
│  │  ┌─────────┐ ┌────────┐ ┌────────┐ ┌─────────┐ │        │
│  │  │ Claude  │ │ OpenAI │ │ Gemini │ │OpenRouter│ │        │
│  │  └─────────┘ └────────┘ └────────┘ └─────────┘ │        │
│  └────────────────────┬────────────────────────────┘        │
│                       │                                      │
│  ┌─────────────────────────────────────────────────┐        │
│  │       Voice I/O Engine — v1.5.1 (Pro only)      │        │
│  │  • Record (sox) • STT (Whisper) • TTS (OpenAI)  │        │
│  └─────────────────────────────────────────────────┘        │
│                                                              │
│  ┌─────────────────────────────────────────────────┐        │
│  │          MCP Server (stdio) — v1.5.1            │        │
│  │  • pear_teach, pear_review, pear_explain        │  ◄──── │── Claude Code
│  │  • Reuses Context Collector + Prompt Assembler  │        │   Cursor
│  └─────────────────────────────────────────────────┘        │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│         pearcode.dev (Next.js + shadcn) — MVP                │
│                                                               │
│  • Landing page (positioning, demo, pricing, install)        │
│  • Blog (content marketing, SEO)                             │
│  • Docs (MDX — getting started, commands, MCP setup)         │
│                                                               │
│  Vercel · No auth/dashboard at launch                        │
└──────────────────────────────────────────────────────────────┘

┌──────────────────────────────────────────────────────────────┐
│        v1.6 — Backend + Billing (deferred)                    │
│                                                               │
│  api.pearcode.dev — Go Backend (Fly.io)                      │
│  • Auth (GitHub OAuth → JWT)                                  │
│  • Billing (Stripe subscriptions)                            │
│  • LLM Proxy (hosted mode)                                   │
│  • Analytics ingest                                           │
│  • Dashboard API                                              │
│  Postgres (Fly.io)                                           │
│                                                               │
│  app.pearcode.dev — User Dashboard (Next.js)                 │
│  • Plan overview, usage, billing management                  │
│  • Stripe Customer Portal                                    │
└──────────────────────────────────────────────────────────────┘
```

### Key Technical Decisions

| Decision | Choice | Rationale |
|---|---|---|
| **CLI (MVP)** | | |
| CLI language | Go | Single binary, fast startup, good CLI ecosystem |
| Terminal UI | Bubble Tea | Best Go TUI library, handles raw mode + hotkeys |
| CLI framework | Cobra | Standard, well-documented |
| Config | TOML (`~/.pear/config.toml`) | Human-readable, standard for CLI tools. Per-repo config deferred to v1.7. |
| LLM adapters | Claude, OpenAI, Gemini, **OpenRouter** | 4 adapters covers ~99% of users. **OpenRouter is key — it gives access to dozens of cheaper models (Llama, Mistral, DeepSeek, Qwen) at $1-5/M tokens instead of $10-15/M for frontier models. Teaching concepts doesn't need heavy-duty inference.** A $3/M-token model explains goroutines, RBAC patterns, and Go idioms just as well as Claude Opus. Users who want frontier-quality analysis can still use Claude/OpenAI directly. |
| Audio capture (v1.5.1) | `sox` subprocess | Avoids cgo/PortAudio build complexity. Ships with voice feature. |
| STT (v1.5.1) | Whisper API (OpenAI) | Best accuracy, most users already have an OpenAI key |
| TTS (v1.5.1) | OpenAI TTS API | Low-latency, good voice quality |
| MCP transport (v1.5.1) | stdio (JSON-RPC) | Standard MCP transport, works with Claude Code and Cursor out of the box |
| Telemetry | PostHog (opt-in) | Lightweight, open-source friendly |
| **Web (MVP)** | | |
| Web framework | Next.js 15 App Router, React, TypeScript | SSR for marketing SEO, type safety throughout |
| UI components | shadcn/ui + Tailwind | Beautiful defaults, fully customizable. Developer-friendly aesthetic out of the box. |
| Web hosting | Vercel | Zero-config Next.js deploy, preview deployments, free tier for launch |
| **Backend (v1.6 — deferred)** | | |
| Backend language | Go | Same language as CLI — shared types, shared knowledge, one toolchain |
| Backend hosting | Fly.io | Simple Go deploy, global edge, built-in Postgres, cheap at low scale |
| Backend DB | Postgres (Fly.io) | Users, subscriptions, usage logs. Single DB. |
| Auth | GitHub OAuth → JWT | Backend owns all auth when it ships. Single system for CLI + web. |
| Billing | Stripe (upfront tiers) | Flat-rate subscriptions, Cursor-style. No metered billing. |
| **Repos** | | |
| Structure | Single private monorepo | CLI ships as compiled binary. Web deploys on Vercel. Backend deploys on Fly.io when it ships. |
| Distribution | `curl` installer + Homebrew tap | `curl -fsSL https://pearcode.dev/install.sh \| sh` (primary). `brew install pearcode/tap/pear` (macOS). No npm, no runtime dependencies. |

### Project Structure (MVP)

> **MVP is CLI + website only.** The `/api` backend directory ships in v1.6 when auth and billing are needed.

```
pear/
├── cli/                                # ── Go CLI (compiled binary) ──
│   ├── cmd/pear/main.go               # CLI entrypoint (Cobra)
│   ├── internal/
│   │   ├── context/
│   │   │   ├── collector.go           # parallel context gathering
│   │   │   ├── git.go                 # git diff, status, branch
│   │   │   ├── tree.go                # file tree generation
│   │   │   └── errors.go             # error log capture
│   │   ├── prompt/
│   │   │   ├── assembler.go           # prompt construction
│   │   │   ├── modes.go               # teach mode (mentor/pair in v1.5.1)
│   │   │   └── roles.go               # role frame templates
│   │   ├── llm/
│   │   │   ├── adapter.go             # LLM interface
│   │   │   ├── claude.go
│   │   │   ├── openai.go
│   │   │   ├── gemini.go
│   │   │   └── openrouter.go          # OpenRouter — cheap models for teaching
│   │   ├── learning/
│   │   │   ├── tracker.go             # concept tagging per response
│   │   │   └── progress.go            # pear progress command
│   │   ├── profile/
│   │   │   └── profile.go             # user profile (from pear init wizard)
│   │   ├── session/
│   │   │   └── history.go             # multi-turn conversation history
│   │   ├── config/
│   │   │   └── config.go              # config parsing
│   │   └── tui/
│   │       └── session.go             # interactive terminal UI
│   ├── prompts/
│   │   └── teach.txt                  # teaching mode system prompt
│   └── go.mod
│
├── website/                            # ── Next.js (Vercel) ──
│   ├── app/
│   │   ├── page.tsx                   # landing page
│   │   ├── pricing/page.tsx           # pricing (no Stripe yet — waitlist/manual)
│   │   ├── docs/                      # documentation (MDX)
│   │   ├── blog/                      # blog (MDX) — SEO + content marketing
│   │   └── layout.tsx
│   ├── components/
│   ├── lib/
│   ├── content/                       # MDX docs + blog posts
│   ├── public/
│   ├── package.json
│   └── vercel.json
│
├── scripts/
│   └── install.sh                     # curl installer script
├── Makefile
└── README.md
```

**Deferred directories (v1.6):**
```
├── cli/internal/voice/                # Voice I/O (v1.5.1)
├── cli/internal/mcp/                  # MCP server (v1.5.1)
├── cli/internal/auth/                 # pear login (v1.6)
├── api/                               # Go Backend (v1.6)
```

**How the components connect (MVP):**
```
CLI (Go binary on user's machine)
  └── BYOK mode: calls LLM APIs directly (Claude, OpenAI, Gemini, OpenRouter)
      No server calls. No auth. Fully local-first.

Web (Next.js on Vercel)
  └── Marketing: pearcode.dev (landing, pricing, docs, blog)
      No auth. No dashboard. Pure marketing + content.

Environment variables:
  CLI:     None required (API keys in ~/.pear/config.toml)
  Web:     NEXT_PUBLIC_PLAUSIBLE_DOMAIN (analytics)
```

### Marketing Site — pearcode.dev

**Purpose:** Launch surface, SEO, conversion funnel. Must be live before PH/HN launch.

**Pages:**

| Page | Purpose |
|---|---|
| `/` (landing) | Hero with tagline + demo screenshot, problem statement, how it works (3-step), feature highlights, pricing CTA, install command |
| `/pricing` | Free vs Pro comparison table, FAQ. Manual invoicing at launch (no Stripe integration yet). |
| `/docs` | Getting started, install, `pear init` walkthrough, command reference, config reference |
| `/docs/mcp` | Dedicated MCP integration guide (ships with v1.5.1) |
| `/blog` | Launch post, "AI makes you fast, Pear makes you good" essay, SEO-optimized articles mapped to content pillars |

**Design:** Clean, developer-focused. Dark mode default. Inspired by Linear, Vercel, Cursor marketing sites. Ship with a single strong demo screenshot/video and iterate.

**Tech:** Next.js 15 App Router + TypeScript + Tailwind + shadcn/ui + Vercel. MDX for docs and blog. No auth at launch — no dashboard, no Stripe integration. This is the founder's strongest stack (TypeScript/React) — a weekend build.

### User Dashboard & Billing (v1.6 — Deferred)

> **Not at launch.** The first 50 Pro users will be invoiced manually. Dashboard, Stripe integration, and auth ship in v1.6 when there's revenue and user demand justifying the infrastructure. See `PRODUCT_DECISIONS.md` for the full billing architecture when it ships.

### Prompt Assembly (Teaching Mode)

> **The system prompt is the product.** This is a simplified representation. The actual prompts will be refined extensively during development and beta.

```
[SYSTEM]
You are a senior staff engineer pair-programming with {name}. They are a
{level}-level developer ({years} years), working primarily in {languages}.

Your primary role is to TEACH — not just solve problems, but help {name}
understand the concepts behind the solution.

When responding:
- Diagnose the immediate issue first (be useful)
- Explain the underlying concept or pattern (be educational)
- Note why this matters in production / at scale (be practical)
- Reference established best practices when relevant (SRE, OWASP,
  Go proverbs, language idioms, etc.)
- Offer to go deeper on any concept ("want me to explain X?")
- Keep explanations grounded in their actual code, not abstract examples
- Calibrate your explanation depth to {level} level

After your response, tag the key concepts you taught (as a structured
JSON block) for learning tracking.

You are working in the repo: {repo_name} on branch: {branch}.
Respond in a conversational, direct tone — like a knowledgeable colleague
at the next desk, not a textbook.

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
{user_input}
```

### MCP Server Mode

Pear ships as both a standalone CLI and an MCP (Model Context Protocol) server. This is a first-class feature, not an afterthought — MCP is the distribution strategy that puts Pear inside every developer's existing AI coding workflow.

**Why MCP matters:**
- MCP support was the #1 feature request in the Aider community (375 combined reactions across two issues)
- Claude Code, Cursor, and other tools already support MCP tool discovery
- Users can invoke Pear's teaching from inside Claude Code without switching contexts
- A developer can be mid-flow in Claude Code, realize they don't understand a change, and ask Pear to teach the concept — all without leaving their session

**Starting the MCP server:**
```bash
# Register Pear with Claude Code
claude mcp add pear -- pear mcp

# Or run standalone for any MCP client
pear mcp
```

**Exposed MCP tools:**

| Tool | Description | Parameters |
|---|---|---|
| `pear_teach` | Explain a concept in the context of the current codebase. Detects the pattern, explains why it matters, offers to go deeper. | `query` (string): what to teach about. Optional: `files` (string[]): specific files for context. |
| `pear_review` | Teaching-oriented code review. Walks through each change, explains the pattern, flags issues, teaches the underlying concept. | `diff` (string): git diff or inline diff. Optional: `mode` (teach\|mentor\|pair). |
| `pear_explain` | Full walkthrough of a file's design, patterns, and architecture decisions. | `file` (string): file path to explain. Optional: `focus` (string): what aspect to focus on. |

**Architecture:** The MCP server reuses the same Context Collector and Prompt Assembler as the interactive CLI. The only difference is the transport — stdio JSON-RPC instead of terminal TUI. This means MCP responses have the same teaching quality as direct Pear sessions.

**Example: Pear inside Claude Code**
```
You: Fix the authentication bug in rbac.go

Claude Code: [fixes the bug, commits the change]

You: /pear_teach "What did that RWMutex change do and why does it matter?"

Pear (via MCP): The RWMutex change on line 12 replaced your
original sync.Mutex. Here's why this matters...
[full teaching response with context]
```

### Context Strategy (v1.5)

In v1.5, all available context is attached on every request — no intent-based filtering. The LLM determines what's relevant based on the user's question. This avoids the misclassification risk of a naive keyword-based intent detector.

**Context types attached on every request:**
- Git diff (staged + unstaged)
- File tree (.gitignore-aware, depth-limited)
- Active files (if referenced via `@file`)
- Error logs (if attached via `/log`)
- Conversation history (last N turns)

**Context budget and per-type token allocations will be refined during development and beta testing.** Modern models (Sonnet, GPT-4o, Gemini Pro) handle 128k+ context easily, so the budget can be generous. When truncation is needed, use middle-out (keep first and last N lines) with a `[...truncated {N} lines...]` signal.

**OpenRouter consideration:** When using cheaper models via OpenRouter (e.g., Llama, Mistral, DeepSeek), context windows may be smaller. The context collector should respect the model's advertised context limit and truncate intelligently. This is especially relevant because teaching-mode prompts don't need frontier-model context windows — most teaching responses reference a single function or pattern, not the entire repo.

Intent-based context selection is deferred to v1.7 when real usage data can inform a proper classifier.

### `pear doctor` — System Health Check

`pear doctor` runs on first install and on-demand. It validates every dependency Pear needs and reports pass/fail with actionable fix instructions.

**MVP checks (no backend required):**

| Check | How | Pass | Fail action |
|---|---|---|---|
| **Profile configured** | Check `~/.pear/profile.json` exists | Profile found | "No profile found. Run `pear init` to set up." |
| **API key configured** | Read `~/.pear/config.toml` | At least one LLM key present | "No API key configured. Run `pear init` to set up." |
| **API key valid** | Lightweight test call (e.g., 5-token completion) | 200 response | "API key for {provider} is invalid or expired. Check your key." |
| **Git available** | `which git` | git found at path | "Git not found. Install git to enable context collection." |
| **Pear version** | Compare local version to GitHub releases API | Up to date or within 1 minor | "Update available: {version}. Run `curl -fsSL https://pearcode.dev/install.sh \| sh`" |

**Additional checks (added in v1.5.1 when voice ships):**

| Check | How | Pass | Fail action |
|---|---|---|---|
| **sox installed** | `which sox` | sox found at path | "Install sox: `brew install sox`" |
| **sox can record** | Record 1s test clip | .wav file produced, >0 bytes | "sox can't access mic. Check System Settings → Privacy → Microphone." |
| **Mic accessible** | Parse sox test output for device | Device name detected | "No microphone found. Connect a mic and try again." |

**Output format:**
```
$ pear doctor
🍐 Pear Doctor — checking system health...

  ✓ Profile configured (Mitch, Go/TypeScript, mid-level)
  ✓ API key configured (OpenRouter — sk-or-...XXXX)
  ✓ API key valid (openrouter, mistralai/mistral-large)
  ✓ Git available (git 2.43.0)
  ✓ Pear up to date (v1.5.0)

All checks passed. You're good to go.
```

### Config (v1.5 — Keep It Simple)

v1.5 ships with `pear init` writing a flat config. No per-repo config, no custom roles, no TOML nesting.

```toml
# ~/.pear/config.toml — written by pear init
provider = "openrouter"              # claude | openai | gemini | openrouter
api_key = "sk-or-..."               # provider API key
model = "anthropic/claude-3.5-sonnet" # model identifier (provider-specific)
mode = "teach"                       # teach only in MVP
```

```json
// ~/.pear/profile.json — written by pear init wizard
{
  "name": "Mitch",
  "languages": ["go", "typescript", "react"],
  "level": "mid",
  "years": 5
}
```

**OpenRouter model examples** (for users who want cheap, fast teaching):
- `mistralai/mistral-large` — strong reasoning, ~$2/M tokens
- `meta-llama/llama-3.1-70b-instruct` — excellent for code, ~$0.80/M tokens
- `deepseek/deepseek-chat` — strong Go knowledge, ~$0.14/M tokens
- `anthropic/claude-3.5-sonnet` — frontier quality via OpenRouter's unified API

Per-repo config (`pear.toml`) and custom role overrides are deferred to v1.7.

### Testing Strategy

The prompt assembler is the product. It gets the most test coverage:

- **Golden-file tests for prompt assembly:** Given a fixed set of inputs (diff, tree, error log, active files, profile, conversation history), assert the exact assembled prompt. One golden file for teach mode at MVP. These tests catch regressions when prompts are edited.
- **Context collector tests:** Mock git/tree/file operations. Assert correct truncation behavior at budget boundaries.
- **LLM adapter tests:** Mock HTTP responses. Assert correct request formatting per provider (Claude, OpenAI, Gemini, OpenRouter).
- **Concept tracker tests:** Assert that concept tags are correctly parsed from LLM responses and stored in `learning.json`.
- **Integration test:** End-to-end test with a mock LLM that verifies: config load → profile load → context collection → prompt assembly → LLM call → response display → concept tracking.

No unit tests for trivial code. Focus test investment on the prompt pipeline, concept tracking, and LLM adapter correctness (especially OpenRouter, which has provider-specific routing).

---

## Roadmap

> **Scope authority:** `PRODUCT_DECISIONS.md` controls what ships when. This roadmap reflects the ruthlessly cut MVP scope.

### v1.5 — MVP: CLI + BYOK + Teach (Ship in ~5 weeks)

Built with AI-assisted development (Claude Code + Cursor). Solo founder. Local-first, no backend.

**CLI (weeks 1-3):**
- Teaching-first prompt engine (teach mode only) — the core differentiator
- Context injection (git, tree, errors, files — all attached, no intent filtering)
- Multi-turn conversation (in-memory history)
- 4 LLM adapters (Claude, OpenAI, Gemini, OpenRouter) + BYOK only
- Local concept tracking — LLM tags concepts per response, stores in `~/.pear/learning.json`
- `pear progress` — show learning profile, concepts, streaks
- Setup wizard (`pear init`) — 3-question personalisation (name, languages, level)
- `pear doctor` system health check
- Structured response format: diagnosis → concept → why it matters → offer to go deeper
- Golden-file tests for prompt assembler
- 14-day free trial: full Pro experience with no credit card required

**Website (weeks 3-5):**
- Marketing site (`pearcode.dev`): landing page, pricing, docs (MDX), install guide, blog
- Deploy on Vercel
- No auth, no dashboard, no Stripe at launch

**Closed Beta (week 5):**
- curl installer script hosted on pearcode.dev
- Homebrew tap for macOS
- Invite waitlist users in batches (20-50 at a time)
- Validate: teaching quality, second-session rate, conversion intent
- Invoice first Pro users manually (no billing infrastructure)
- Iterate prompts based on real feedback
- Content: LinkedIn (5k followers), Substack, new X account — document the beta

### v1.5.1 — MCP + Voice + Public Launch (weeks 6-7)

**MCP server (`pear mcp`):**
- stdio JSON-RPC transport (standard MCP)
- Tools: `pear_teach`, `pear_review`, `pear_explain`
- Reuses Context Collector + Prompt Assembler from core CLI
- Listed on MCP directories (mcp.so, PulseMCP) for discoverability

**Voice (Pro feature unlock):**
- Push-to-talk input (sox + Whisper API)
- Optional TTS output (OpenAI TTS)
- Feature-flagged — text must already be proven

**Additional prompt modes:**
- Mentor mode (concise answer + one key insight)
- Pair mode (direct, minimal — just solve the problem)

**Public launch (week 7):**
- Product Hunt + Hacker News (staggered, not simultaneous)
- 60-second demo video/screenshot
- Launch blog post: "AI makes you fast. Pear makes you good."
- Email full waitlist with install link

### v1.6 — Backend + Billing + Retention (weeks 8-11)

- **Go backend (api.pearcode.dev):** Auth (GitHub OAuth → JWT), billing (Stripe), LLM proxy (hosted mode), analytics ingest
- **Dashboard (app.pearcode.dev):** Plan overview, usage, billing management
- **Stripe:** $30/mo Pro, $300/yr annual. Flat-rate subscription.
- **Hosted mode:** Included hosted requests for Pro users who want zero-config
- Linux support
- `pear history` — searchable past sessions (server-side storage)
- `pear pipe` — pipe stderr directly: `make build 2>&1 | pear pipe`
- Server-side concept tracking sync
- Latency optimizations

### v1.7 — Teaching Intelligence (weeks 12-17)

- Intent detection from accumulated usage data (smart context selection)
- Per-repo teaching memory (Pear remembers what it's taught you)
- "Did you know?" micro-lessons triggered by code context
- Custom prompt templates in per-repo `pear.toml`
- Teaching frequency insights ("you've asked about mutexes 3 times — here's a deep dive")

### v1.8 — Knowledge Foundation (weeks 18-24)

- Concept graph: topics → prerequisites → best practices → sources
- Grounded in: Go proverbs, OWASP, AWS Well-Architected, Google Engineering Practices, SRE doctrine, language-specific idioms
- `pear learn <topic>` — on-demand concept deep-dive using your codebase as examples

### v2.0 — Full AI Tutor + Team Tier (weeks 25-34)

- Structured learning paths within real codebases
- Skill assessment from code patterns
- Progress tracking dashboard (web)
- Spaced repetition: resurface concepts
- **Team tier ($TBD/seat/mo):** seat management, shared billing, team learning analytics
- VS Code extension

---

## What to Build Now vs. Defer

> See the "Explicitly Deferred from MVP" table in the MVP Scope section above for the full deferral list with target versions and rationale.

---

## Positioning & Messaging

### Tagline Options

- "AI makes you fast. Pear makes you good."
- "The only learning tool that teaches while you work."
- "Learn software engineering while you code."
- "Your senior engineer, always on call."

### Key Differentiator (vs. Aider, Cursor, Claude Code)

Those tools are **power tools** — they make you faster at producing code.
Pear is a **learning tool** — it makes you better at understanding code.

They answer *"how do I fix this?"*
Pear answers *"why did this break, what's the pattern, and how do you avoid it next time?"*

### Marketing Angle

The AI coding tools market is saturated with "write code faster" products. Nobody is addressing the growing concern that AI tools are creating developers who can *ship* but can't *reason*. Pear is positioned at the intersection of developer tools and developer education — a category that doesn't exist yet.

---

## Distribution Strategy

> **Full tactical playbook:** See `DISTRIBUTION_GUIDE.md` for the complete 4-phase distribution strategy, content pillars, launch timeline, and metrics.

### Summary

Distribution pillars: **Content + ecosystem partnerships + selective open source + community presence.** GitHub virality is not the primary channel but open-sourcing select components may be considered for credibility (pending decision — see `PRODUCT_DECISIONS.md`).

### Channels

1. **LinkedIn** (5,000 followers) — founder's existing audience, highest-leverage channel at launch
2. **Substack** — long-form essays on the AI-assisted learning crisis
3. **X (Twitter)** — demo videos, developer community engagement
4. **Product Hunt + Hacker News** — staggered launch (not simultaneous)
5. **MCP directories** — mcp.so, PulseMCP (v1.5.1)
6. **Reddit** — r/programming, r/golang, r/learnprogramming, r/ExperiencedDevs, r/cursor
7. **Discord/Slack** — Cursor, Claude dev servers, Indie Hackers, Gophers
8. **Dev.to / Hashnode** — SEO-optimized cross-posts with canonical URLs to pearcode.dev
9. **pearcode.dev blog** — SEO strategy mapped to content pillars

### Installation

```bash
# Primary — works on macOS (Linux in v1.6)
curl -fsSL https://pearcode.dev/install.sh | sh

# Secondary — macOS via Homebrew
brew install pearcode/tap/pear

# Tertiary — Go users
go install github.com/pearcode/pear@latest
```

No npm. No runtime dependencies at launch (sox only needed when voice ships in v1.5.1).

### Trust (CLI Reading User Code)

- **BYOK sends directly to LLM provider.** Pear sends context directly to Claude/OpenAI/Gemini/OpenRouter APIs. Pear never sees the code. No hosted mode at launch.
- **Transparency documentation.** Publish exactly what data is collected, where it goes, what's stored. Clear privacy policy on pearcode.dev.
- **Precedent.** Users already trust Cursor, Claude Code, GitHub Copilot, and Aider with their code. One more CLI tool is not a novel trust barrier for this audience.

---

## Business Model (BYOK-Only at Launch)

Users pay for the teaching tool — the context injection, pedagogical prompt engine, concept tracking, and learning profile. They do not pay for LLM access. BYOK is the only usage mode at launch. Hosted mode ships in v1.6.

| Tier | Monthly | Annual | Includes |
|---|---|---|---|
| **14-day free trial** | $0 | — | Full Pro experience for 14 days. No credit card required. Includes BYOK, all modes, voice (when shipped), concept tracking, `pear progress`, unlimited questions. Converts users by letting them experience the full product. |
| **Pro** | $30/mo | $300/yr (annual saves $60) | BYOK with all modes, voice (when shipped), concept tracking, `pear progress`, unlimited questions. $30/mo is the price. Raise later when retention data justifies it. |
| **Team** (v2) | TBD/seat/mo | TBD | Everything in Pro + team learning analytics + admin dashboard. Waitlist on launch site. |

**Why BYOK-only at launch:** Zero upstream LLM costs. Pear's margin is ~100% on the subscription (only infrastructure costs: Vercel for website). No backend needed, no LLM proxy, no quota enforcement. The simplest possible billing model for a solo founder.

**Why $30/mo:** Positions Pear as a professional learning tool (not a cheap taster). Matches the professional tier for developer tools (Cursor Pro, Claude Pro). High enough to signal serious value, accessible for individual credit cards. Can raise later when retention data and feature depth justify it.

**Why 14-day trial (not permanent free tier):** Full product access for 14 days lets users experience the true value of teaching mode, concept tracking, and all features. Most conversion happens within the first week of use. Removes friction for onboarding while creating a natural deadline for upgrade decisions.

**OpenRouter changes the unit economics for users:** With OpenRouter, users can access high-quality teaching at $0.14-2/M tokens instead of $10-15/M. This means Pear's $30/mo subscription cost is often *more* than the user's monthly LLM spend for teaching use cases. That's intentional — users are paying for the teaching intelligence, not the model access.

**Unit economics at scale:** BYOK-only means the only costs are: Vercel (free tier → $20-50/mo), domain ($12/yr), PostHog (free tier), trial system infrastructure. At 30 Pro users ($900 MRR), hosting costs are <$50/mo = 94%+ margin. Backend infrastructure costs only appear in v1.6 when hosted mode ships.
