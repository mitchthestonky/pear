# Pear — Product Design: Free, Pro & Teams

**Date:** 2026-03-14
**Author:** Mitch + Claude (product analysis session)
**Status:** Draft

---

## Overview

Pear is a CLI pair programmer that watches you code and teaches during natural pauses. The product is structured across three tiers, each serving a distinct identity and buyer:

| Tier | Identity | Buyer | Price |
|------|----------|-------|-------|
| Free | Your AI pair programmer | Individual developer | $0 (BYOK) |
| Pro | Your personal learning engine | Developer investing in growth | $15/mo ($120/yr) |
| Teams | Engineering intelligence platform | Eng managers, VP Eng, L&D | $25/user/mo |

The progression is **tool → system → platform**. Free gets people in the door. Pro turns casual usage into a learning flywheel. Teams surfaces engineering competency data to leadership.

---

## Free — "Your AI pair programmer"

### What It Does

A genuinely useful real-time code companion. Catches mistakes, explains patterns, answers questions in context. Think "senior engineer over your shoulder."

### Included

- **Watch mode** — auto-reviews on code settle (30s pause) and git commits. Unlimited.
- **Interactive Q&A** — REPL with multi-turn conversation, @file support, slash commands.
- **One-shot commands** — `pear ask`, `pear review`, `pear teach`. Direct to stdout.
- **Concept tagging** — concepts tagged in every response (`Concepts: [X, Y]`).
- **Single-session memory** — Pear doesn't repeat itself within a session, but forgets between sessions.
- **Basic `pear progress`** — local flat list of concepts encountered, stored in `~/.pear/learning.json`.
- **All LLM providers** — Anthropic, OpenAI, OpenRouter. BYOK (bring your own key).
- **Git hooks** — `pear hooks install` for post-commit auto-reviews.

### What It Doesn't Do

- No cross-session memory. Each session starts fresh.
- No gap analysis or curriculum.
- No web dashboard.
- No sync across machines.

### Why It's Free

- BYOK means zero infrastructure cost per user.
- The free tier is the distribution vehicle. Every free user is a potential Pro subscriber and a potential Teams champion inside their company.
- A crippled free tier doesn't generate word-of-mouth. A great free tier does.

---

## Pro ($15/mo or $120/yr) — "Your personal learning engine"

### What It Does

Pear remembers what you know, finds what you don't, and builds a curriculum from your real code. The CLI becomes a persistent learning companion, and a web dashboard at pear.dev gives you visibility into your growth.

### Included (everything in Free, plus)

#### Cross-Session Memory
- Learning state persists across sessions. Pear remembers what it's taught you.
- Memory builds from **watch mode and interactive REPL sessions only**. One-shot commands (`ask`, `review`, `teach`) are transactional — they help in the moment but don't contribute to your learning profile.
- Rationale: sustained engagement signals real learning. Quick lookups don't.

#### Progressive Depth Teaching
- Pear adapts its teaching depth based on concept exposure:
  - **Introduce** — first encounter, explain the concept and why it matters.
  - **Deepen** — seen it a few times, surface edge cases, failure modes, trade-offs.
  - **Connect** — link to related concepts, show patterns across domains.
  - **Challenge** — ask harder questions, test understanding, push to mastery.
- This is what makes Pear feel intelligent over time vs. a stateless prompt.

#### AI-Generated Gap Analysis
- Pear analyzes your code patterns and surfaces what you *don't* know.
- Example: "You've used channels 12 times but never `select` — here's what you're missing."
- Gaps are detected from real code, not self-assessment.

#### Hybrid Curriculum Engine
- **AI-generated micro-lessons**: personalized to your observed gaps, delivered through the CLI during natural pauses.
- **Curated learning paths** (over time): hand-authored structured journeys (e.g., "Go concurrency fundamentals", "React patterns for backend engineers"). Pear recommends the right path based on what it sees.
- Pear is the primary teaching resource. It teaches directly — not a referral engine.
- Free official resources (Go blog, MDN, official docs) recommended when relevant.
- Paid course partnerships TBD for the future.

#### Concept Graph
- Visual map of concepts and their relationships, not a flat list.
- Shows clusters, connections, and progression over time.

#### Knowledge Gap Detection
- Distinct from gap analysis: this is the persistent, evolving view of what you know vs. don't.
- Surfaces in the dashboard and influences what the CLI teaches next.

#### Web Dashboard (pear.dev)
- **Read-only at launch**: view your concept graph, browse session history, see gaps and growth over time.
- **Interactive later** (based on user feedback): set learning goals, mark concepts as known, pin focus topics.
- Cloud-hosted. Learning data syncs to Pear servers.

#### Session Archive & Search
- Full history of what Pear taught you, searchable.
- "What did Pear teach me about error handling last Tuesday?"

#### Shareable Progress Cards
- Generate a visual card of your learning progress.
- Share on Twitter/X, LinkedIn, GitHub profile.
- Think GitHub contribution graph but for learning.

#### Cross-Machine Sync
- Learning state travels with you. Use Pear on work laptop and personal machine, same profile.
- Cloud-hosted, no self-hosted option at launch.

#### Export
- Markdown, Obsidian, Notion export of learning data and session history.

### Data & Privacy
- All Pro learning data is cloud-hosted (syncs from CLI to Pear servers).
- Required for web dashboard, cross-machine sync, and cloud-side analysis.
- Self-hosted option not planned at launch. Enterprise self-hosted TBD based on demand.

---

## Teams ($25/user/mo) — "Engineering intelligence platform"

*Includes everything in Pro.*

### What It Does

Surfaces engineering competency data to leadership. Managers see where their team is growing, where the gaps are, and where upskilling is needed — all derived from real code, not self-reported surveys.

### Timeline

Demand-driven. Not on a fixed roadmap. Focus is on free + Pro adoption first. Teams gets built when companies ask for it.

### Buyer Journey

The sale is a flywheel, not a single motion:

```
Individual dev adopts free Pear (bottom-up)
    → Upgrades to Pro (self-serve)
        → Champions it to their eng manager
            → Manager wants team visibility
                → VP Eng / CTO signs org deal
                    → L&D budget pays for it
```

### Included (everything in Pro, plus)

#### Skill-Based Competency Graphs
- Per-engineer competency graph showing concept mastery over time.
- **Growth-oriented framing**: "Sarah improved 40% in concurrency this month" — not "Sarah doesn't know concurrency."
- Managers see where each person is improving and where they need guidance.

#### Team-Wide Gap Analysis
- Aggregate view: "3 engineers need upskilling in error handling."
- Team concept heatmap: what the team collectively knows vs. doesn't.
- Identifies systemic gaps (e.g., "nobody on the team understands observability patterns").

#### Individual Growth Tracking
- Full visibility into each engineer's learning trajectory.
- Track progress against concepts relevant to the team's codebase and stack.
- Useful for 1:1s, performance conversations, and growth planning.

#### Onboarding Tracking
- Secondary use case, not the headline pitch.
- Track how new hires are ramping on codebase concepts.
- "New hire reached codebase fluency in X weeks" as a measurable metric.

#### Architecture: CLI Reports Up, Cloud Analyzes

```
                          DEVELOPER'S MACHINE
                    ┌─────────────────────────────┐
                    │                             │
                    │   Editor (VS Code, Neovim)  │
                    │         │ saves file        │
                    │         ▼                   │
                    │   ┌───────────┐             │
                    │   │ pear watch│             │
                    │   │  (CLI)    │             │
                    │   └─────┬─────┘             │
                    │         │                   │
                    │    ┌────┴────┐              │
                    │    │         │              │
                    │  detects   receives         │
                    │  pause/    teaching          │
                    │  commit    strategy          │
                    │    │         ▲              │
                    │    ▼         │              │
                    │  ┌──────────┴──────────┐   │
                    │  │  Local Teaching      │   │
                    │  │  Engine              │   │
                    │  │                      │   │
                    │  │ • Collects diff/ctx  │   │
                    │  │ • Streams LLM review │   │
                    │  │ • Tags concepts      │   │
                    │  │ • Adapts depth based  │   │
                    │  │   on cloud strategy   │   │
                    │  └──────────┬──────────┘   │
                    │             │               │
                    └─────────────┼───────────────┘
                                  │
                    reports learning data
                    (concepts, exposure counts,
                     session context, gaps)
                                  │
                    ══════════════╪═══════════════
                         NETWORK BOUNDARY
                    ══════════════╪═══════════════
                                  │
                                  ▼
               ┌─────────────────────────────────────┐
               │         PEAR CLOUD (pear.dev)       │
               │                                     │
               │  ┌───────────────────────────────┐  │
               │  │      Analysis Engine           │  │
               │  │                                │  │
               │  │  • Evaluates competency        │  │
               │  │  • Maps concept relationships  │  │
               │  │  • Detects knowledge gaps      │  │
               │  │  • Generates curriculum        │  │
               │  │  • Tracks progression          │  │
               │  │  • Compares across sessions    │  │
               │  └──────────────┬────────────────┘  │
               │                 │                    │
               │      ┌─────────┴─────────┐         │
               │      │                   │         │
               │      ▼                   ▼         │
               │  ┌────────┐    ┌──────────────┐    │
               │  │Teaching│    │  Competency   │    │
               │  │Strategy│    │  Data Store   │    │
               │  │Push    │    │              │    │
               │  │        │    │ Per-engineer: │    │
               │  │• Depth │    │ • Concepts    │    │
               │  │  level │    │ • Mastery %   │    │
               │  │• Focus │    │ • Gaps        │    │
               │  │  areas │    │ • Growth rate │    │
               │  │• Curric│    │ • Sessions    │    │
               │  │  ulum  │    │              │    │
               │  └───┬────┘    └──────┬───────┘    │
               │      │               │             │
               │      │    ┌──────────┴──────────┐  │
               │      │    │                     │  │
               │      │    ▼                     ▼  │
               │      │  ┌─────────┐  ┌──────────┐ │
               │      │  │ Pro     │  │ Teams    │ │
               │      │  │Dashboard│  │Dashboard │ │
               │      │  │         │  │          │ │
               │      │  │ MY view │  │TEAM view │ │
               │      │  └─────────┘  └──────────┘ │
               └──────┼───────────────────┼─────────┘
                      │                   │
               ═══════╪═══════════════════╪══════════
                      │                   │
                      ▼                   ▼

              ┌──────────────┐  ┌─────────────────────┐
              │  DEVELOPER   │  │  ENG MANAGER         │
              │  (Pro user)  │  │  (Teams admin)       │
              │              │  │                      │
              │ Sees:        │  │ Sees:                │
              │ • My concept │  │ • Team competency    │
              │   graph      │  │   heatmap            │
              │ • My gaps    │  │ • Per-engineer       │
              │ • My growth  │  │   growth graphs      │
              │   over time  │  │ • Team-wide gaps     │
              │ • Session    │  │   ("3 people need    │
              │   history    │  │    error handling")  │
              │ • Curriculum │  │ • Onboarding         │
              │   progress   │  │   velocity           │
              │              │  │ • Improvement trends │
              │ Actions:     │  │                      │
              │ • Share card │  │ Actions:             │
              │ • Export     │  │ • Manage seats       │
              │              │  │ • Set focus areas    │
              └──────────────┘  │ • View individual    │
                                │   growth details     │
                                └─────────────────────┘


  THE FEEDBACK LOOP (what makes it work):

  Dev codes ──► CLI detects + teaches ──► Reports to cloud
                     ▲                          │
                     │                          ▼
                     │                    Cloud analyzes:
                     │                    "This dev has seen
                     │                     channels 12x but
                     │                     never select"
                     │                          │
                     │                          ▼
                     └──── Cloud pushes:  "Next time they
                           touch channels, teach select.
                           Use 'deepen' depth level."
```

- The CLI is the data collection + teaching engine.
- The cloud app is the brain: it does the analysis, evaluation, progression mapping, and curriculum generation.
- The CLI doesn't make competency judgments locally — it reports up and receives guidance back.

#### Admin & Infrastructure
- SSO integration
- Admin controls (manage seats, configure team settings)
- Audit log
- Team-wide provider/API key management (optional)

---

## Pricing Summary

| | Free | Pro | Teams |
|---|---|---|---|
| **Price** | $0 | $15/mo ($120/yr) | $25/user/mo |
| **LLM cost** | User pays (BYOK) | User pays (BYOK) | User pays (BYOK) |
| **Infrastructure cost to Pear** | None | Cloud sync + dashboard | Cloud sync + dashboard + team analytics |
| **Buyer** | Individual dev | Individual dev | Eng manager / VP Eng / L&D |
| **Install** | `brew install pear` | CLI + pear.dev account | CLI + pear.dev team account |

### Why $15/mo for Pro

- At $20, you compete in mental budget with Copilot ($10) and Cursor Pro ($20).
- At $15, it's "I'll try it" territory for individual devs.
- LLM costs are zero on BYOK — it's pure margin on sync, storage, and compute for analysis.
- Teams at $25/user is where volume revenue lives.

### Why $25/user/mo for Teams

- Includes all of Pro (no confusing add-on pricing).
- Manager buys seats, everyone gets the full experience.
- Competitive with engineering L&D spend (Pluralsight ~$29/user/mo, much less useful).
- Simple: one price, everything included.

---

## Key Design Decisions

| Decision | Choice | Rationale |
|---|---|---|
| What builds your learning profile? | Watch + interactive sessions only | Sustained engagement signals real learning. One-shots are quick lookups. |
| Curriculum approach | Hybrid (AI-generated + curated paths) | Ship AI gap analysis first, layer curated paths over time. |
| External course recommendations | Pear teaches directly. Free official resources when relevant. Paid partnerships TBD. | Pear is the learning resource, not a referral engine. |
| Dashboard interactivity | Read-only first, interactive later | Ship the graph, let users tell you what controls they want. |
| Data residency | Cloud only | Simplifies architecture. Self-hosted is an enterprise play if demand emerges. |
| Team visibility | Full individual-level competency graphs | Growth-oriented framing. Managers see improvement and gaps, not judgments. |
| Teams wedge | Ongoing gap analysis & competency mapping | Not just "learning" — engineering intelligence. Onboarding is secondary. |
| Team admin CLI control | CLI reports up, cloud pushes strategy down | CLI is the teaching engine. Cloud is the brain. Dashboard is the control surface. |
| Seat model | Teams includes Pro | $25/user, everything included. No pricing math for the buyer. |
| Teams timeline | Demand-driven | Ship free + Pro. Let companies tell you when they want Teams. |

---

## Build Sequencing

```
Phase 1: Fix the Foundation (NOW)
├── Fix P0 streaming bug (core UX)
├── Wire up concept tracking (learning data starts flowing)
├── Deduplicate reviews (session-level fingerprinting)
└── Ship free tier that feels great

Phase 2: Pro Learning Engine
├── Cross-session memory (persistent learning.json → cloud sync)
├── Progressive depth teaching (tiered prompts based on exposure)
├── AI-generated gap analysis
├── Session archive
└── Basic web dashboard (read-only: concept graph, history, gaps)

Phase 3: Pro Polish
├── Hybrid curriculum (AI micro-lessons + first curated paths)
├── Shareable progress cards
├── Export (markdown, obsidian, notion)
├── Dashboard iteration based on feedback
└── Free resource recommendations in CLI

Phase 4: Teams (when demand pulls)
├── Team dashboard (competency graphs, gap heatmaps)
├── Individual growth tracking for managers
├── CLI → cloud reporting pipeline
├── Cloud → CLI teaching strategy push
├── SSO, admin, audit log
└── Onboarding tracking
```

---

## Open Questions

- **Pricing experiments**: Should Pro have a free trial period, or does the free tier serve as the trial?
- **Curated paths**: Who authors them? Mitch solo? Community contributions? Contracted curriculum designers?
- **Competency benchmarks**: Can Pear eventually say "you're in the top 20% of Go developers for concurrency"? Is that valuable or creepy?
- **Team privacy controls**: Should engineers be able to opt-out of individual visibility to managers? Or is that a condition of the Teams plan?
- **Pear-hosted inference**: Should Pro/Teams offer a no-BYOK option where Pear provides the LLM? Simplifies onboarding but adds cost.
