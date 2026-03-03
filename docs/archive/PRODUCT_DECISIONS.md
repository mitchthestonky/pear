# Pear — Product Decisions Log

> Decisions made during product design sessions, February 2026.
> These override or refine the PRD where they conflict.

---

## Interaction Model

**Decision:** Smart context-based follow-ups.
- Short, direct answers just end — no forced follow-up prompt
- Complex explanations that introduce new concepts offer a follow-up ("Want me to go deeper on X?")
- Pear decides based on response complexity, not the user
- Rule of thumb: if the response mentions a concept the user hasn't encountered before (tracked in local learning log), offer to go deeper

**Rationale:** This makes pear feel like a thoughtful tutor, not a search engine (silent) or a chatbot (always prompting). It respects the user's flow.

---

## Pricing Model (Refined)

**Decision:** Hybrid flat + usage, BYOK subscription model.

| Tier | Price | Includes |
|------|-------|----------|
| **14-Day Free Trial** | Free | Full Pro experience (BYOK with all modes, voice, concept tracking), no credit card required. 14 days to try everything. |
| **Pro** | $30/mo | BYOK with all modes, voice, concept tracking, full features. Annual: $300/yr, 2 months free. |
| **Team** (v2) | TBD/seat/mo | Everything in Pro + proficiency mapping + team analytics + admin dashboard. |

**Key changes (Round 3):**
- 14-day free trial gives full Pro experience with no credit card required — lets users experience everything before committing
- No hosted-only free tier — users provide their own API keys (BYOK). Cost to serve trial users: $0.
- Pro price is $30/mo flat — this is just the price, not "early adopter" framing. Positioned as justified learning premium above $20 code-generation competitors.
- $30/mo is slightly above the $20 impulse threshold (Cursor, Claude Pro, Bolt) but justified by being a learning system, not a code generation tool. Retains impulse buy viability.
- Voice is a paid feature unlock, not default — text-only pear must stand on its own
- BYOK users still pay subscription — they're paying for pedagogy + context + workflow, not LLM access

**Margin protection:**
- BYOK is the margin play — zero LLM cost for most users
- 14-day trial costs $0 to serve (user's own API key, local CLI, no hosted compute)
- Paid BYOK user costs ~$0.16/mo → 99% margin at $30/mo
- Add hosted convenience tier later when revenue covers it

---

## Retention Strategy

**Decision:** Three-layer retention.

1. **Infinite depth** — There's always more to learn. Security, performance, architecture, system design, language idioms. The curriculum is as deep as software engineering itself.
2. **Team/org expansion** — Individual users graduate but recommend pear for their team's juniors. Growth comes from org seats, not just individual retention.
3. **Usage-based within tiers** — Users who learn less use less, pay less (within their tier), and have no reason to churn. Heavy users hit limits and upgrade.

**The proficiency mapping angle:** Track which concepts users encounter, which explanations lead to retention vs re-asks, and surface this as a learning profile. For individuals: "concepts you've mastered." For teams: "knowledge gaps across your engineering org." This is the enterprise unlock and the data flywheel.

---

## Voice Strategy

**Decision:** Voice is an enhancer, not the product.

- Text-only pear must deliver full value on its own
- Voice input + audio responses are a Pro feature unlock
- This de-risks launch: if voice latency sucks, text mode still works
- Some users will always prefer text — that's fine
- Voice activation in paid plans helps justify the price premium

---

## Text Differentiation (vs "just ask Claude")

**Decision:** The compound advantage of context + pedagogy + workflow.

All three together create the moat:
1. **Context** — Pear reads your actual codebase, git history, errors. ChatGPT needs copy-paste.
2. **Pedagogy** — Pear doesn't just explain, it teaches patterns, connects concepts, builds on previous questions.
3. **Workflow** — In your terminal, in your flow. No tab switching.

No single feature is defensible alone. The integration is the product.

---

## Defensibility / Moat Strategy

**Decision:** Layered moat, built over time.

| Timeline | Moat Layer |
|----------|-----------|
| Launch | Pedagogy — teaching well is genuinely hard. Purpose-built prompts + context beats "add explain button." |
| 6 months | Data flywheel — track which explanations stick vs get re-asked. Proprietary learning effectiveness data. |
| 12 months | Proficiency mapping — individual learning profiles + team knowledge gap analysis. Enterprise L&D unlock. |
| 18 months | Community/content — crowd-sourced concept explanations, learning paths, pattern libraries. |
| 24 months | Platform play — pear becomes the teaching API other tools integrate. |

**The proficiency mapping moat is the most defensible.** Once engineering orgs use pear to evaluate and upskill their engineers as part of L&D and retention, switching costs are enormous.

---

## GTM / Distribution Strategy (Refined)

**Decision:** Phased launch funnel.

1. **Phase 1 (now):** Waitlist + scarcity → beta invite pipeline. Build anticipation.
2. **Phase 2 (launch):** 60-second demo video showing a real bug → pear teaching → understanding. Every post links to this video.
3. **Phase 3 (post-launch):** Open public beta. Remove gates. Grow through organic word-of-mouth.

**The metric that matters:** Second session rate. If 40%+ of users come back the next day, distribution will solve itself.

**Funding angle:** Bootstrap to profitability first. If it attracts VC attention or acquihire interest, engage from a position of strength with real revenue and retention data.

---

## Business Ambition

**Decision:** Bootstrap to profitability, open to opportunistic funding/acquisition.

- Goal: profitable from day one
- Default: lifestyle SaaS ($1-3M ARR), solo or tiny team
- If metrics are exceptional: open to pre-seed/seed from position of strength
- The "pear" name is intentionally positioned to attract Apple's attention
- No decisions optimized for VC scale until the market proves it

---

## First-Run Experience (Refined — Round 3)

**Decision:** Quick setup wizard → personalized first response.

- `pear` starts with a 3-question setup wizard (~30 seconds):
  1. Name
  2. Primary languages / frameworks
  3. Self-declared level + years of experience
- This personalizes the very first teaching response — pear "knows you" from the first interaction
- After wizard: drop into the prompt. First question gets a response calibrated to their level.
- Profile stored locally (`~/.pear/profile.json`) at launch. Server sync when auth ships.
- `pear doctor` runs automatically on first install to validate setup

**Note:** Originally decided "zero-friction first question" but the wizard adds 30 seconds and unlocks personalization that makes the first response visibly better. The wizard IS the onboarding.

---

## Response Defaults

**Decision:** Concise by default, depth on demand.

- Default responses: 3-5 sentences + code fix
- Deeper explanation available via follow-up prompt or `/deeper`
- Users can configure default depth in `pear.toml`
- TTS reads spoken summary only (2-3 sentences), full detail in text
- Walls of text are the #1 UX risk for a teaching tool

---

## Growth Strategy Insights (from market research)

### Common Patterns Across High-Growth AI Products (Cursor, Lovable, Bolt, Claude Code, v0)

1. **Product-led growth is non-negotiable** — Every $100M+ ARR AI tool grew through the product itself, not paid marketing. Zero ad spend.

2. **The wow moment must be instant and shareable** — First experience produces visible value in under 60 seconds. For pear: ask a question about your real code → get a teaching response that's visibly better than ChatGPT.

3. **Free trial as conversion funnel** — Let users experience the full Pro product for 14 days, then upgrade. By the time trial ends, they're invested.

4. **$30/mo is the learning premium** — Cursor, Lovable, Bolt, Claude Pro all converge on $20/mo for code generation. Pear is $30/mo because it's a learning system, not a coding tool. Still within impulse-buy range and justified by pedagogy differentiation.

5. **Usage-based pricing is the endgame** — Every product starts flat, moves to credits/tokens as they scale. Plan for this migration.

6. **Controversy/hype without product = fragile** — Cluely got 70K signups via rage-bait but growth plateaued. Quality sustains; hype doesn't.

7. **Enterprise is the revenue multiplier** — Individual devs drive adoption; enterprise contracts drive revenue. Pear's proficiency mapping is the enterprise bridge.

8. **Speed of execution defines winners** — Bolt: $0→$40M ARR in 5 months. Lovable: $100M in 8 months. Ship fast, iterate publicly.

9. **Pre-existing audience is critical** — No product in this cohort launched cold. Lovable had 50K GitHub stars, Bolt had StackBlitz users, v0 had Vercel's 6M devs. Pear has LinkedIn (5K) + Substack + waitlist.

### Implications for Pear

- The 60-second demo video is the single most important marketing asset
- 14-day free trial is generous enough for the "aha moment" and conversion (full experience to get comfortable before paying)
- The proficiency mapping / team L&D angle is the path to enterprise revenue
- Don't try to be viral — be excellent. Quality product + developer word-of-mouth scales
- Ship in 5 weeks, iterate based on second-session-rate, not feature completeness

---

## Target Audience (Refined)

**Decision:** Three-tier audience with clear priority.

| Priority | Audience | Why pear |
|----------|----------|----------|
| **Primary** | Developers who ship with AI but want to understand their code | Self-taught engineers, vibe-coders, career switchers, junior-intermediate devs. People who self-select and self-pay. Already spending on Cursor/ChatGPT/courses — pear is the missing layer. |
| **Secondary** | Self-improvers who already pay for dev education | Individuals DO buy tools. Udemy does $700M+. Pluralsight was $400M. The difference: pear is embedded in workflow, not a separate "learning time" activity. |
| **Tertiary** | Engineering managers | Want to track team L&D, identify knowledge gaps, evaluate engineer growth. Proficiency mapping unlocks this (12+ months out). |

**Key reframe:** "Junior devs" describes who needs it, not who buys it. Junior devs have no budget. The buyer is a developer who's *insecure about their gaps and willing to invest in themselves.* Message to the self-selecting learner, not the job title.

---

## Wow Moment Definition

**Decision:** Context injection is the undeniable demo.

The 5-second clip: user asks a question → pear visibly pulls git diff, file tree, error logs from the real codebase → responds with a teaching explanation grounded in their actual code.

**Why this wins:** It's *visual*. You can see pear doing something ChatGPT can't. "It knows my code" is immediately obvious. Follow-up depth ("wait, why?") is the second wow moment that hooks retention.

Voice is UX sugar — impressive but not the differentiator. The context-aware teaching is.

---

## MVP Scope (Round 3 — 2-Week Target)

**Decision:** Ship the smallest thing that makes someone message a friend.

**In scope (2 weeks):**
- CLI in Go
- Teach mode only (one mode, not three)
- Context injection (git, tree, errors, files — token budget to be refined during development)
- Four LLM adapters: Claude, OpenAI, Gemini, **OpenRouter** (BYOK only, no hosted)
- **OpenRouter is key:** gives users access to dozens of cheaper, performant models (Llama, Mistral, DeepSeek, Qwen) at $1-5/M tokens instead of $10-15/M for frontier models. Teaching concepts doesn't need heavy-duty inference — a $3/M-token model teaches goroutines and RBAC patterns just as well as Claude Opus.
- Setup wizard (name, languages, level, experience)
- Local concept tracking — LLM tags concepts per response, stored in `~/.pear/learning.json`
- `pear progress` — concepts with frequency + session streak
- Local user profile (`~/.pear/profile.json`)
- `pear doctor` system health check

**Out of scope (ship later):**
- Voice (feature-flagged, not a launch blocker)
- Mentor and pair modes (teach mode must prove value first)
- Billing / Stripe / dashboard (invoice first 50 users manually)
- Server-side auth / account creation (local-first for beta)
- Hosted LLM mode (can't subsidize pre-revenue)
- MCP (ships v1.5.1 after teaching quality is validated)

**Implication:** The first version must nail context + pedagogy + memory. Voice, billing, and additional modes are growth features, not launch features.

---

## Distribution Strategy (Refined)

**Decision:** Content + ecosystem partnerships + selective open source.

1. **Content-first:** Daily LinkedIn posts documenting the build. Quality audience (5K devs) > large audience. Ship content every day until launch.
2. **Ecosystem partnerships:** Get pear featured in Cursor/Claude Code ecosystems via MCP integration. Piggyback on their distribution. MCP is the Trojan horse into existing workflows.
3. **Open source (TBD):** Some components will be open-sourced for GitHub credibility, HN goodwill, and community contributions. The teaching engine (prompts, context collection strategy, concept tracking, prompt assembly) stays proprietary. Exact scope of what goes open source is a pending decision — see Open Source Strategy below.
4. **Multi-channel community presence:** LinkedIn (primary), Reddit (r/golang, r/programming, r/cursor), Discord (Cursor, Claude), Dev.to/Hashnode (SEO-optimised cross-posts), Hacker News (Show HN launch).

**The MCP distribution play:** If pear works inside Claude Code and Cursor via MCP, every user of those tools is a potential pear user. The install friction drops to one command: `claude mcp add pear -- pear mcp`. This is potentially more powerful than any content strategy.

---

## Open Source Strategy

**Decision: PENDING.** The scope of what goes open source has not been decided yet. This section captures the considerations for when the decision is made.

**What is clear:**
- The teaching engine is proprietary. This includes: teaching prompts and mode system, context collection implementation (token budgeting, parallel gathering, truncation strategy), prompt assembly pipeline, concept tracking and extraction algorithms, and role frame definitions. This is the IP. It stays closed.
- Some components will be open-sourced to earn GitHub credibility, HN goodwill, and developer trust.

**Candidates for open source (to be evaluated):**
- CLI framework (Cobra scaffolding, command registration)
- TUI shell (Bubble Tea terminal UI)
- MCP server transport layer (JSON-RPC stdio)
- LLM adapter interfaces (provider abstraction)
- Config system (TOML parsing, global + per-repo)
- Doctor command (system health check)
- Install scripts and build tooling (Goreleaser, Homebrew tap, install.sh)
- Extension framework and example plugins

**Why open source matters for distribution:**
- GitHub stars drive organic developer discovery. 100+ stars in week 1 puts pear on GitHub trending.
- HN heavily rewards open source. The Show HN post should link to a repo.
- Homebrew tap requires a public repo for `brew install`.
- Fork risk is containable: without the teaching engine, a fork produces generic ChatGPT-quality responses.
- Community contributions on infrastructure (new LLM adapters, TUI enhancements) don't touch the IP.

**Open questions (to resolve before launch):**
- How much of the codebase needs to be public to get meaningful GitHub traction?
- Single repo with closed-source modules vs. separate repos vs. binary-only distribution with open-source tooling?
- What's the minimum viable open source footprint for HN credibility?
- Does the Homebrew tap requirement force a specific repo structure?

---

## Response UX (The Teaching Format)

**Decision:** All three differentiation layers combined.

Every pear response has:
1. **Structured format** — Diagnosis → concept explanation → why it matters in production → offer to go deeper. Visually distinct sections, not a wall of text.
2. **Visible enrichment** — Show context injection happening live in the terminal (files being pulled, diff loading, prompt assembling). Make the magic visible before the response arrives.
3. **Concept tagging** — Each response tags concepts taught (e.g., `[RBAC]` `[Go slices]` `[error handling]`). Builds a visible learning log over time. Shows pear is tracking your growth.

**Why this matters:** Cursor and Claude Code already read your codebase. The teaching format is what makes pear visually and functionally different. A screenshot of a pear response should be immediately distinguishable from a ChatGPT response.

---

## Tech Stack

**Decision:** Go is non-negotiable.

- Single binary distribution, no runtime deps
- Best CLI ecosystem (Cobra, Bubble Tea)
- The Go learning curve is part of the product story (building a learning tool while learning)
- Worth the time investment for distribution simplicity

Voice pipeline will be built in Go. If specific components need subprocess calls (e.g., sox for mic capture), that's fine — but the core is Go.

---

## Data Flywheel (Pulled Forward)

**Decision:** Local tracking + opt-in server sync from v1.5.

- Track concepts locally (on-device) from v1.5 launch
- Show "concepts encountered this session" at session end
- Offer opt-in server sync for users who want cross-device concept history
- Respects privacy by default, starts the flywheel for those who opt in
- Zero server cost for local-only users
- Server-side data collection begins building the proprietary dataset from day one for opted-in users

**Implication:** This means v1.5 needs basic concept extraction from LLM responses (can be simple keyword/pattern matching or ask the LLM to tag its response). The local storage is a JSON file in `~/.pear/learning.json`. Lightweight, no database needed.

---

## Positioning

**Decision:** "The only learning tool that teaches while you work."

This is the one-liner. It does three things:
1. **Differentiates from courses/tutorials** — you don't stop working to learn
2. **Differentiates from AI coding tools** — they work for you, pear teaches you
3. **Implies zero workflow disruption** — the objection "I don't have time to learn" is dead on arrival

Use it as: hero subheadline, PH tagline, LinkedIn bio, first line of cold outreach. Don't overuse — let the rest of the site/product support the claim.

---

## MVP Scope (Round 3 — Ruthless Cut)

**Decision:** Ship the smallest thing that makes someone message a friend.

- CLI + BYOK + teach mode + context injection. That's it.
- No billing at launch. No dashboard. No hosted mode. Free for first beta users.
- Collect feedback, prove second-session rate, bolt on payments once retention is validated.
- You can invoice the first 50 customers manually if needed.
- Build MCP before billing — MCP distribution inside Cursor/Claude Code drives more adoption than Stripe ever will.

**The 60% product that ships in 2 weeks beats the 95% product that ships in 6 weeks.**

---

## 14-Day Free Trial

**Decision:** Full Pro access for 14 days, no credit card required.

- 14-day free trial: Full Pro features (unlimited questions, all modes, concept tracking, learning profile, voice, BYOK)
- After trial: upgrade to Pro at $30/mo or cancel (no autocharge surprise)
- No credit card required removes friction for trial conversion
- 14 days is enough to experience the full "aha moment" and get comfortable with the workflow
- Scarcity of time (vs. scarcity of questions) is more effective for conversion than daily caps

---

## MCP Interjection Model (v2+)

**Decision:** User-controlled interjection levels.

The MCP vision isn't just "ask pear inside Cursor" — it's pear actively watching your coding and interjecting with teaching moments when it detects patterns you don't understand.

**Interjection levels (user configurable):**
- `off` — pear only responds when called
- `subtle` — status bar indicator when a learning opportunity is detected
- `active` — inline teaching cards when pear detects unfamiliar patterns in accepted AI code

Default: `subtle`. The user dials it up when they want to learn, turns it off when they're shipping.

**Why this matters:** This is pear as a *teaching co-pilot*, not a Q&A tool. No competitor is building this.

---

## Progress UX

**Decision:** Concept list with frequency + session streak.

`pear progress` shows:
- Concepts encountered with frequency: `goroutines (3x) | error handling (7x) | channels (1x)`
- Session streak: `This week: 12 concepts across 5 sessions. Streak: 4 days.`
- Clean, scannable, terminal-native
- Gamification-lite — encourages daily use without being obnoxious
- Future: add confidence tiers (mastered / learning / new) based on re-ask frequency

---

## Prompt Iteration / Feedback Loop

**Decision:** Re-ask rate as primary signal + end-of-session summary + AI-powered usage analysis.

- **Re-ask rate:** If a user asks the same concept again within a week, the first explanation didn't stick. Track automatically, no user action.
- **End-of-session:** "Did you learn something new this session? [yes/no]" — one question, once.
- **Usage analytics:** Analyze what users are asking to determine if they're getting value and what they expect. Use AI to surface patterns in usage data and guide product iteration.
- No per-response thumbs up/down — too much friction, most users won't bother.

---

## Churn Strategy

**Decision:** Accept individual churn, design for expansion.

Four churn defenses, each covering a different lifecycle stage:
1. **Infinite depth** — security, performance, architecture, system design. Nobody "finishes" learning engineering.
2. **Career transitions** — new job, new language, new codebase. Pear resets to beginner on new context. Lifetime value extends across career changes.
3. **Habit replacement** — pear replaces Stack Overflow / docs as the default way to understand code. Teaching is the UX, even for experienced devs.
4. **Graduation → org expansion** — individuals churn but recommend pear for their team's juniors. Org seats replace individual retention. (The Slack playbook.)

---

## Trust & Hallucination

**Decision:** Link to sources, accept the risk.

- Every teaching response includes relevant doc links / authoritative sources where applicable
- LLMs are rarely wrong on well-established programming concepts — the risk is overstated
- Don't over-engineer hallucination detection at launch
- Fix edge cases as they surface from user feedback

---

## Data & Privacy

**Decision:** Layered approach.

1. **BYOK = their problem** — code goes directly from user's machine to their own API key. Pear never sees or stores code.
2. **Privacy policy** — pear processes code in-memory only. Nothing stored server-side. No telemetry on code content. Made explicit.
3. **SOC 2 roadmap** — not needed at launch but have a clear "month 12" answer for enterprise conversations.

---

## Onboarding Failure Recovery

**Decision:** Graceful fallback + transparency.

- `pear doctor` validates environment before first question (git, file access, API key)
- If context injection produces bad results, fall back to high-quality generic teaching (no context)
- Every response shows collapsed "Context used: main.go, go.mod, git diff (3 files)" section
- Users can see WHY an answer was bad and adjust — also doubles as debugging tool for you during beta

---

## Community

**Decision:** Slack, support-first.

- Paid users get a private Slack with direct founder access, bug reports, feature requests
- Support, community, AND retention in one channel
- Direct access to founder justifies the upgrade for some users
- Free users get docs and self-serve
- No Discord community yet — community platforms are a time sink to moderate pre-scale

---

## Analytics & Observability

**Decision:** Concept tracking IS the analytics + opt-in error reporting.

- Product data = business data. Concept tracking tells you what users learn, how often they return, what topics they ask about.
- Opt-in error reporting (Sentry or similar) for crash reporting
- Usage analytics (question count, session duration, feature adoption) via PostHog or similar, privacy-respecting
- No telemetry on code content — ever

---

## Updates & Versioning

**Decision:** Auto-update check on launch.

- Pear checks for updates every time it starts
- If new version exists: "Update available (v1.6). Update now? [y/n]"
- Gentle but persistent — users stay current without manual effort
- Homebrew users get `brew upgrade pear`
- Breaking data changes include automatic migration on first launch (learning.json v1 → v2)

---

## i18n

**Decision:** English only. Don't think about it.

- Dev tools are English-first globally. Most code is in English.
- Non-English comments/variables are edge cases
- LLMs handle multilingual input naturally if it comes up
- Revisit only if demand appears

---

## Biggest Risk

**Decision:** Speed to market.

Not fear of failure — fear of being outpaced. The competitive window is real:
- Cursor could ship a teach mode
- Claude Code could add pedagogical prompts
- A well-funded competitor could clone the concept

**Mitigation:** Ship v1.5 in 5 weeks. Voice is feature-flagged (launch without if needed). Text teaching with context injection is the minimum viable amazing. Don't let perfect be the enemy of shipped.

The moat isn't being first — it's being best at teaching. But being first AND best is the play.
