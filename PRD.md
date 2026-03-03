# Pear — Product Requirements Document

> Last updated: March 2026

---

## Current State

The v0 CLI is feature-complete. 26 implementation tickets across 6 phases, all done. The marketing website is live at pearcode.dev. Pre-launch.

For v0-specific requirements, see [docs/cli/PRD.md](docs/cli/PRD.md).

---

## Product Principles

1. **Teach, don't just answer.** Every response should leave the developer understanding more than before.
2. **Alongside, not instead of.** Pear complements existing AI tools. Never compete with Claude Code or Cursor.
3. **Context is everything.** Teaching grounded in the developer's actual code beats abstract examples every time.
4. **Invisible until valuable.** Watch mode should feel like a thoughtful colleague, not a nagging tool.
5. **Ship small, learn fast.** Optimize for real user feedback over speculative features.

---

## Feature Overview

### CLI (`cli/`)

| Feature | Command | Status |
|---------|---------|--------|
| Interactive REPL | `pear` | v0 done |
| Watch mode (proactive teaching) | `pear watch` | v0 done |
| One-shot ask | `pear ask "question"` | v0 done |
| Diff review | `pear review [--commit]` | v0 done |
| Topic teaching / deep dive | `pear teach [topic]` | v0 done |
| Health check | `pear doctor` | v0 done |
| Learning progress | `pear progress` | v0 done |
| Git hooks | `pear hooks install/uninstall` | v0 done |
| First-run setup | `pear init` | v0 done |
| Concept tracking | Automatic via response parsing | v0 done |
| Multi-provider LLM (Anthropic, OpenAI, OpenRouter) | Config-based | v0 done |
| Slash commands (/help, /settings, /review, etc.) | In-REPL | v0 done |

### Website (`website/`)

| Feature | Status |
|---------|--------|
| Marketing landing page | Done |
| Blog (6 posts) | Done |
| Docs (commands, config, providers, usage) | Done |
| Pricing page | Done |
| Waitlist + email capture | Done |
| About page | Done |
| SEO (sitemap, robots, OG images) | Done |
| FAQ | Done |
| Comparison page | Done |

---

## Pricing Tiers

### Free (BYOK)

- $0 forever
- Bring your own API key (Anthropic, OpenAI, OpenRouter)
- Watch mode (auto-reviews changes)
- Interactive Q&A with codebase context
- Subtle teaching moments and concept tags

### Pro — $20/month ($130/year)

- Everything in Free
- Learning state memory
- Adaptive teaching (`pear teach`)
- Concept tracking and progress
- Personal knowledge gap visibility
- Cross-machine sync

### Teams — $30/seat/month

- Everything in Pro
- Team-wide learning metrics
- Shared billing
- Contact: mitch@pearcode.dev

---

## Roadmap

### Pre-launch (now)

- Smoke test CLI end-to-end
- Recreate missing .goreleaser.yaml for binary distribution
- Verify install.sh works with GitHub releases
- Demo video (60-90s, real unscripted teaching moment)

### v1 (post-launch iteration)

- Prompt quality iteration based on real user feedback
- Session summary at exit
- Improved concept extraction accuracy
- Usage analytics (opt-in)

### v1.5 (learning engine foundations)

- Concept graph persistence (~50-60 macro concepts per language)
- Knowledge tracing (what you've seen vs what you've mastered)
- Spaced retrieval prompts
- Logging-only proficiency model

### v2.0 (full learning engine)

- Misconception modeling
- Adaptive difficulty
- Cross-session continuity
- Team tier with shared analytics
- B2B dashboard

---

## Infrastructure

- **CLI distribution:** goreleaser → GitHub Releases → install.sh
- **Website:** Next.js on Vercel
- **Email:** Resend
- **KV store:** Upstash Redis (waitlist, rate limiting)
- **Analytics:** Vercel Analytics
- **User data:** Local only (~/.pear/). No server-side storage in v0-v1.
