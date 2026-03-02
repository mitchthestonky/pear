# Pear — Positioning & GTM Strategy

*Last updated: 2 March 2026*

---

## One-Liner

**A pair programmer that watches your code and tells you what matters.**

Top-of-funnel sells the free tool people will install today. The learning engine positioning ("Pear remembers what you know and teaches you what you don't") is the Pro upsell story.

---

## The Problem

AI is making engineers faster. Not better.

- 19% slower with AI — developers just feel 20% faster (METR, 2025)
- 67% drop in junior engineering roles (Stanford Digital Economy Lab)
- 66% of developers fix "almost-right" AI code (Stack Overflow, 2025)
- 4x increase in code duplication (GitClear analysis)

AI coding agents (Claude Code, Cursor, Aider, OpenCode) generate massive diffs that developers barely read, let alone understand. Code ships faster, understanding doesn't follow. Mentorship can't scale. The gap widens.

---

## Product Tiers

### Free (OSS)
**What it is:** A pair programmer that watches your code and surfaces insights — with a teaching edge.

- `pear watch` — auto-reviews code changes in real-time
- Interactive Q&A with full codebase context (@file references)
- BYO API key (Anthropic, OpenAI, OpenRouter — any provider)
- `/review` — one-shot review of current git diff
- **Subtle teaching moments** — concept tags on responses, occasional "here's why this pattern matters," light context on *why* not just *what*. Practical-first, but visibly smarter than raw Claude.
- Fully local, no account needed
- Free forever

**Who it's for:** Any developer who wants a second pair of eyes on their code, especially when AI agents are writing it.

**Why the teaching tone matters:** Free Pear must be *visibly different* from "just ask Claude." The subtle teaching moments — structured responses, concept tagging, the occasional insight — are what make someone screenshot it and share it. Without this, free Pear is a generic code reviewer in a crowded field.

### Pro — $20/mo ($130/yr — 4 months free)
**What it is:** Pear remembers what you understand and teaches you what you don't.

Everything in Free, plus:
- **Learning state memory** — Pear tracks what you *understand*, not just project context. Never re-explains what you've already mastered. This is fundamentally different from Claude Code's CLAUDE.md or Cursor Rules, which remember project preferences — Pear remembers *your knowledge*.
- **`pear teach`** — adaptive learning tied to your actual codebase
- **Concept tracking & progress** — see what you've learned over time
- **Personal knowledge gap visibility** — identify patterns you keep missing
- **Cross-machine sync** — learning profile follows you everywhere

**The upgrade moment:** Free Pear teaches in the moment but forgets between sessions. Pro Pear builds a model of your understanding over time. The difference becomes obvious after a few weeks of use — Pro users get responses calibrated to exactly what they know.

### Teams — $30/seat/mo ($200/seat/yr — 4 months free) — 2-seat minimum
**What it is:** Visibility into code quality and knowledge across your engineering team.

Everything in Pro, plus:
- **Team dashboard** — code smell trends, common anti-patterns surfaced
- **Team knowledge gap visibility** — "3 engineers struggled with context cancellation this week"
- **Custom codebase rules** — "always flag missing error handling in Go", "enforce our API conventions"
- **Junior ramp-up metrics** — track onboarding progress with real data
- **Admin controls** — manage seats, set team-wide teaching preferences

**Who buys it:** Engineering managers and CTOs. The pitch: "You're paying for AI to write code faster. Are you sure your team understands what it's writing?"

---

## Install Flow — Zero Friction

Install is zero-friction. No email gate, no signup, no waitlist. Anyone can install from GitHub, Homebrew, or the website install command immediately.

```
Landing page → Install command shown directly (+ GitHub link)
    or
GitHub/Homebrew → brew install pear / go install
```

### Optional Email Capture (In-CLI, Post-Value)
After the user's first successful `pear watch` session or Q&A interaction, the CLI optionally prompts:

> "Want updates on new features? Enter your email (optional):"

This captures higher-quality leads — people who've already experienced value. No gate, no friction before first use.

### Email Sequence (for users who opt in):
- **Immediate:** Quick tips + feature highlights
- **Day 5:** "Did you know you can use /review for one-shot reviews?"
- **Day 14:** "You've used Pear for 2 weeks. Here's what Pro adds." (upgrade nudge)
- **Day 30:** "Free Pear teaches in the moment. Pro Pear remembers what you know." (learning state memory hook)

---

## Landing Page Structure

### Hero
> **AI is making engineers faster. Not better.**
>
> Pear is a pair programmer that watches your code changes and surfaces what matters — patterns, problems, and insights you'd otherwise miss.
>
> `curl -sSL https://pearcode.dev/install | sh`
>
> [View on GitHub] — free, open source

### Section 2: The Problem
Keep current stats and three-point problem framing. It works.

- Engineers ship code without understanding it
- Senior engineers can't mentor at scale
- Learning happens outside the point of execution

### Section 3: How It Works
> 1. You code (or your AI agent codes for you)
> 2. Pear watches your changes in real-time
> 3. Pear surfaces insights — what's good, what's risky, what you should know
> 4. You ask follow-up questions with full codebase context

**Embed 60-second terminal demo here.** This is the single highest-leverage marketing asset.

### Section 4: How Pear Is Different

**Pear vs. Claude Code / Cursor / OpenCode**

| | Claude Code / Cursor / OpenCode | Pear (Free) | Pear (Pro) |
|---|---|---|---|
| Generates code for you | Yes | No — Pear reviews, not writes | No |
| Auto-reviews your changes | No | **Yes** (`pear watch`) | **Yes** |
| Explains *why*, not just *what* | Sometimes, if you ask | **Yes — by default** | **Yes** |
| Concept tagging on responses | No | **Yes** | **Yes** |
| Remembers project context | Yes (CLAUDE.md, rules files) | No | No |
| Remembers *what you understand* | No | No | **Yes — learning state memory** |
| Adapts teaching to your level | No | No | **Yes — adaptive pedagogy** |
| Tracks your knowledge gaps | No | No | **Yes** |
| Shows your growth over time | No | No | **Yes — concept tracking** |

**Pear doesn't replace your AI coding tool. It's the layer that makes sure you understand what your AI coding tool writes.**

### Section 5: Pricing

| | Free (OSS) | Pro $20/mo | Teams $30/seat/mo |
|---|---|---|---|
| Watch mode — auto-reviews changes | Yes | Yes | Yes |
| Interactive Q&A with codebase context | Yes | Yes | Yes |
| BYO any LLM provider | Yes | Yes | Yes |
| Subtle teaching moments & concept tags | Yes | Yes | Yes |
| **Learning state memory** | — | Yes | Yes |
| **Adaptive learning (pear teach)** | — | Yes | Yes |
| **Concept tracking & progress** | — | Yes | Yes |
| **Personal knowledge gap visibility** | — | Yes | Yes |
| **Custom codebase rules** | — | — | Yes |
| **Team dashboard & code smell trends** | — | — | Yes |
| **Team knowledge gap visibility** | — | — | Yes |
| **Admin controls** | — | — | Yes |
| Annual pricing | Free forever | $130/yr (4mo free) | $200/seat/yr (4mo free) |
| | | | 2-seat minimum |

### Section 6: Who It's For

1. **Developers using AI agents** — Claude Code, Cursor, Aider, OpenCode generate huge diffs. Pear tells you what just happened.
2. **Self-taught developers** — Fill gaps in fundamentals while you work.
3. **Junior engineers** — Get the mentorship your team can't provide at scale.
4. **Bootcamp graduates** — Bridge the gap between bootcamp patterns and production code.
5. **Founder-coders** — Ship fast with AI, but understand what you're shipping.
6. **Senior engineers learning new stacks** — Pear calibrates to your level.

### Section 7: The Crisis (social proof / urgency)
Keep current stats. They validate the problem and create urgency.

### Section 8: FAQ
- **Is it really free?** Yes. The CLI is open source and free forever with your own API key. Pro adds learning state memory and adaptive teaching.
- **What LLM providers work?** Anthropic (Claude), OpenAI, and any OpenRouter-compatible provider.
- **How much does the API cost?** Typical usage runs $0.05-0.15 per session depending on provider and model.
- **Can I use it without watch mode?** Yes. `pear` launches interactive mode for Q&A. `pear watch` adds automatic code review.
- **What about my code privacy?** Pear runs locally. Your code goes directly to your LLM provider via your own API key. We never see it.
- **How is Pear different from just asking Claude/ChatGPT?** Free Pear gives you structured, context-aware responses with concept tagging — visibly better than raw LLM output. Pro Pear remembers what you understand across sessions and adapts its teaching to your level. No LLM does that.
- **Do I need an API key?** Yes — Pear uses your own LLM provider key (BYO). A managed option where Pear handles LLM access is planned for the future.

---

## Key Messaging Changes from Current Website

| Current (pearcode.dev) | New |
|---|---|
| "The only learning engine for AI-accelerated builders" | "A pair programmer that watches your code and tells you what matters" |
| "Join the waitlist" | Install command shown directly + GitHub link |
| "14-day free trial" | "Free forever (OSS). Pro $20/mo for learning state memory + adaptive teaching." |
| 5 learning pillars as headline features | Lead with watch mode + subtle teaching. Pillars become "Why upgrade to Pro" |
| $300/yr (save 2 months) | $130/yr (4 months free) |
| $30/mo Pro | $20/mo Pro |
| Teams "coming soon" | Teams $30/seat/mo (2-seat min) with specific features |
| No mention of open source | Open source is the headline distribution strategy |
| Target personas missing AI agent users | Lead persona: "developers using AI agents" |
| No comparison to Claude Code / Cursor | Explicit comparison section showing what Pear does that they don't |
| Email gate before install | Zero-friction install, optional email capture in-CLI after first value |
| "Memory across sessions" (sounds like CLAUDE.md) | "Learning state memory" — remembers what you *understand*, not project context |

---

## Competitive Positioning

**What Pear is NOT:**
- Not a code generation tool (Cursor, Copilot, Claude Code, OpenCode do that)
- Not a linter or static analysis tool (ESLint, golangci-lint do that)
- Not a course platform (Udemy, Pluralsight do that)
- Not another CLAUDE.md / Cursor Rules (those remember project preferences — Pear remembers *your knowledge*)

**What Pear IS:**
- The pair programmer that watches what your AI agent writes and tells you what matters
- The only tool that bridges "code generated" and "code understood"
- Free: a smart code reviewer with a teaching edge. Pro: a code reviewer that knows what you know.

**Category:** Developer understanding tool. Not dev productivity, not dev education. Understanding.

**Voice:** Planned as a future UX enhancement once teaching quality and core UX are validated. Not in launch positioning. Text-first Pear must stand on its own.

---

## Distribution Funnel

```
Content (LinkedIn, X, HN, YouTube)
    "AI is making engineers faster. Not better."
                    ↓
     Landing page / GitHub repo
    "Pair programmer for your code"
                    ↓
    Zero-friction install (brew/curl/go install)
                    ↓
      pear init → BYO key → pear watch
                    ↓
        Daily use. Genuinely useful.
        Subtle teaching creates "this is different" moments.
                    ↓
    Optional email capture after first value moment
                    ↓
    Learning state gap felt over weeks
    (free Pear teaches well but forgets between sessions)
                    ↓
         Upgrade to Pro ($20/mo)
                    ↓
    Dev shares with team → manager sees value
                    ↓
        Teams ($30/seat/mo)
```

### Viral Loops
1. **Screenshot loop** — Pear's structured responses with concept tags look *different* from raw Claude output → dev screenshots → shares on Slack/X/LinkedIn
2. **Team word-of-mouth** — "Look what Pear found in my code" → teammates install free version
3. **Content loop** — Build in public → demo clips → developers install to try

### Key Metrics
- **Install rate** from landing page / GitHub (target: high — no gate to suppress it)
- **Second-session rate** (target: 40%+ — product-market fit signal)
- **Optional email capture rate** (target: 20%+ of active users)
- **Free → Pro conversion** (measure reality, don't target a number — build something people want to pay for)
- **Time to upgrade** (expected: 3-6 weeks of regular use)
- **Team expansion rate** (one dev installs → how many teammates follow)

---

## Future Revenue Lever: Managed Pear

Post-launch, add a managed option where LLM calls route through Pear's infrastructure at a margin. This solves two problems:

1. **BYOK friction** — juniors and bootcamp grads who don't have API keys can pay for convenience
2. **Additional revenue** — margin on LLM usage on top of subscription revenue

This is not a launch feature. Launch is BYOK-only. Managed mode ships when revenue supports the infrastructure costs.

---

## The 60-Second Demo Video

This is the single highest-leverage marketing asset. Requirements:

1. No intro, no logo animation — start with action
2. Show: developer makes code changes → Pear auto-detects → surfaces 2-3 insights with concept tags → developer asks follow-up → Pear answers with codebase context
3. The response must look *visibly different* from raw Claude — structured format, concept tags, teaching tone
4. End with install command
5. Record when TUI is polished (pixel-art pear, clean layout, smooth streaming)
6. Post everywhere: landing page hero, X, LinkedIn, YouTube, PH launch

---

## Launch Sequence

### Pre-Launch (Weeks 1-2)
- Establish "AI learning crisis" narrative on LinkedIn/X (no product mention)
- Collect emails via waitlist on pearcode.dev (transition to direct install when ready)
- Target: 100+ emails

### Reveal (Weeks 3-4)
- First screenshots, technical decisions, building-in-public content
- Early access to hand-picked beta testers (10-20)
- Target: 300+ emails, real user feedback

### Public Launch (Week 5)
- Product Hunt launch (Day 0)
- Show HN (Day +1)
- LinkedIn + X + Reddit + email blast
- 60-second demo video everywhere
- Target: 500+ installs first week, 25+ Pro conversions month 1

### Post-Launch (Weeks 6-12)
- Daily content (user stories, metrics, learnings)
- Weekly YouTube (full Pear sessions)
- Podcast appearances (Changelog, Syntax, devtools.fm)
- Iterate based on usage data and feedback
- Build upgrade nudges based on real conversion data (don't over-engineer pre-launch)
