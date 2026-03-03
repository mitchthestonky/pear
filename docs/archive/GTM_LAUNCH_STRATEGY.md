# Pear — Go-To-Market Launch Strategy

> Version: 1.0
> Last updated: February 2026
> Author: Mitch

---

## Strategic Positioning

### The Pitch (One Sentence)

**Pear is the voice-first AI tutor that teaches software engineering while you code — the L&D tool for the AI-assisted era.**

### Why This Matters Now

AI coding tools created a problem nobody's solving. Copilot, Claude Code, Cursor, and Aider make developers faster — but not better. Junior and mid-level engineers are shipping code they don't understand, and the gap between "code that works" and "code you understand" is widening. Senior engineers are overwhelmed by mentoring demands. Online courses are disconnected from real work.

Pear makes learning explicit and intentional. It doesn't interrupt your workflow — it sits alongside your existing tools and teaches you *in the moment*, using your actual codebase. This is a dedicated L&D tool in a non-invasive workflow. Developers still use Claude Code to code. They use Pear to *learn*.

### The Category

Pear isn't competing in the "AI coding assistant" category. It's creating a new one: **AI developer education tools.** This is the intersection of developer tools ($50B+ TAM) and developer education ($15B+ TAM), and it has zero funded entrants.

### Core Differentiator

Every AI coding tool answers: *"How do I fix this?"*
Pear answers: *"Why did this break, what's the pattern, and how do you avoid it next time?"*

The education and upskilling lag introduced by AI is a real, growing problem. Pear brings it to the forefront and solves it at the point of work, not in a separate learning environment.

---

## Target Audiences

### Primary: Individual Developers (1-5 years experience)

- Junior and mid-level developers using AI tools daily
- Already paying for Cursor ($20/mo), Claude API, OpenAI API, online courses
- Want to genuinely understand what they're building, not just ship it
- Feel the gap between what AI helps them produce and what they actually know

**Conversion trigger:** The moment they realize Pear explains the *why* behind code, not just the *what*.

### Secondary: Engineering Teams & L&D Buyers

- Engineering managers frustrated by junior engineer mentoring overhead
- L&D teams looking for developer-specific education tools
- Companies with 10:1 junior-to-senior ratios where mentoring doesn't scale

**Conversion trigger:** "This could save our senior engineers 5+ hours/week of mentoring."

### Tertiary: Startup/Founder Audience (Mitch's Personal Brand)

- 5,000 LinkedIn followers in the startup ecosystem
- People interested in how to build products and startups
- Will follow the "building Pear in public" story
- A subset are developers who become users; the rest become evangelists and course customers

---

## The Flywheel: Product + Personal Brand + Course

The GTM strategy runs on a single insight: **building Pear in public creates content that markets both the product and the founder's expertise.** These aren't separate activities — they're the same activity producing three outputs.

```
Build Pear → Document decisions → Content for LinkedIn/Substack/X
     ↑                                          ↓
     │                              Audience grows (devs + startup people)
     │                                          ↓
     │                              Devs try Pear → usage data → better product
     │                              Startup people follow the story → course audience
     │                                          ↓
     └──────────── Revenue (Pear subs + future course revenue + credibility)
```

### How the Pieces Reinforce Each Other

| Activity | Markets Pear | Builds Personal Brand | Feeds Future Course |
|---|---|---|---|
| LinkedIn build-in-public posts | Direct product awareness | Founder transparency + expertise | Case study content |
| Substack essays on AI learning crisis | Thought leadership → Pear as solution | Authority in dev education space | Course preview content |
| X demo clips + dev community engagement | Viral potential for product | Developer credibility | N/A |
| Technical blog posts (latency, prompts, Go) | Developer trust + SEO | Engineering credibility | Technical decision-making case studies |
| User stories + feedback loops | Social proof → conversions | "Founder who listens" narrative | Customer development lessons |
| PH/HN launches | Install spikes | Founder visibility | Launch strategy case study |

---

## Channel Strategy

### 1. LinkedIn — Primary Channel (5,000 followers)

**Why it's the lead channel:** Highest existing audience, strong engagement, professional context where developers and managers both live. Posts reach beyond followers via algorithm.

**Content cadence:** 3x per week

**Content pillars:**

| Pillar | Example Post | Frequency |
|---|---|---|
| **Build-in-public** | "Week 2 of building Pear. Shipped multi-turn conversation in Go. Here's what broke and what I learned about managing token budgets..." | 1x/week |
| **AI learning crisis hot takes** | "AI tools are creating developers who can ship but can't reason. I asked 10 engineering managers if they've noticed. Every. Single. One. Said yes." | 1x/week |
| **Technical deep-dives** | "How I got voice-to-LLM latency under 2 seconds in Go. The trick: start collecting context while the user is still talking." | 1x/week |

**Every post ends with a soft CTA:** "Building Pear — a voice AI tutor for developers. [pearcode.dev]"

**Target growth:** 5k → 8k pre-launch, 8k → 12k by month 3, 12k → 20k by month 6.

### 2. Substack — Long-Form Depth

**Launch immediately** with a manifesto: *"AI Makes You Fast. Pear Makes You Good."*

**Content cadence:** Weekly (500-1000 words — you're building a product, not writing a blog)

**Content alternates between:**
- Pear build updates (what shipped, what's hard, honest numbers)
- Startup philosophy (pricing decisions, why BYOK-first, why closed source, choosing Go as a learning project)
- AI + education perspective pieces
- Developer education thought leadership

**This doubles as course preview content.** You're literally teaching how to build a startup by documenting your own.

**Target:** 500 subscribers pre-launch, 1,000 by month 3, 2,500 by month 6.

### 3. X (Twitter) — New Account, Secondary Channel

**Don't try to grow X from zero while building.** Use it for:
- Cross-posting short-form versions of LinkedIn content
- 15-30 second demo clips (these can go viral)
- Engaging in threads about Claude Code, Aider, vibe coding, developer education
- Replying to developer influencers discussing AI coding tools

**Target:** 1k followers by month 3 (organic from cross-posts + engagement).

### 4. Product Hunt — Launch Event

- Category: Developer Tools / Education
- Title: "Pear — AI tutor that lives in your terminal"
- Subtitle: "Voice-first teaching companion for developers. Talk to your codebase, learn while you code."
- Rally LinkedIn audience to upvote
- Reply to every PH comment personally
- Target: Top 5 Product of the Day

### 5. Hacker News — Technical Launch

Two approaches (choose based on the day's news cycle):
- **Show HN:** "Show HN: Pear — voice AI tutor for developers (Go CLI)"
- **Blog post:** "I built a voice AI tutor because Copilot is making developers worse" (the provocative angle HN loves)

Be in the comments for 4+ hours. HN values founder presence and honesty.

### 6. Developer Communities

- r/programming, r/ExperiencedDevs, r/ChatGPTCoding, r/golang
- Discord servers: Claude Code, Cursor, indie hackers, dev tools
- Share genuinely, not spammily. Focus on the problem ("AI learning crisis"), not the product pitch.

---

## Launch Sequence

### Pre-Launch: Weeks 1-4 (Building Phase)

| Week | LinkedIn | Substack | X | Other |
|---|---|---|---|---|
| **Week 1** | "I'm building something." Post about the AI learning crisis problem. Start the narrative. | Launch Substack with manifesto: "AI Makes You Fast. Pear Makes You Good." | Create account. Pin manifesto. | Set up pearcode.dev with email capture ("Get early access") |
| **Week 2** | Build-in-public: voice pipeline in Go. Technical detail post. | Essay: "Why I chose Go for a voice AI product (and what I'm learning)" | Cross-post highlights. Engage in AI coding threads. | Collect early access emails |
| **Week 3** | Hot take on AI learning crisis. Tag relevant thought leaders. | Essay: "The $30/mo question — how to price a developer education tool" | Demo clip: 15s of push-to-talk voice interaction (even if rough) | DM 20 developer friends for alpha testing |
| **Week 4** | Build-in-public: the teaching prompt engine. Show a before/after (raw Claude vs. Pear response). | Essay: "What GitHub learned when they killed Copilot Voice — and why 2026 is different" | Teaser: "Launching next week." | Final testing. Lock pricing page. Prepare PH listing. |

**Pre-launch goals:**
- 200+ early access email signups
- 3 LinkedIn posts with 100+ likes each
- 500+ Substack subscribers
- 10+ alpha testers with qualitative feedback

### Launch Week: Week 5

**Critical: Do NOT launch PH and HN the same day.** Spread the energy across the week.

| Day | Action | Channel |
|---|---|---|
| **Monday** | Final prep. Stage PH listing. Pre-write HN post. Prepare launch email. | — |
| **Tuesday** | **Soft launch.** "After 5 weeks of building in public, Pear is live." Demo video + install link. Email waitlist with launch pricing. | LinkedIn, Substack, X |
| **Wednesday** | **Product Hunt launch.** Rally LinkedIn audience. Reply to every comment. | PH, LinkedIn, X |
| **Thursday** | PH follow-up. Share early user reactions on LinkedIn. Keep engaging PH comments. | PH, LinkedIn |
| **Friday** | **Hacker News.** Post Show HN or blog post. Be in comments for 4+ hours. | HN, X, Reddit |
| **Weekend** | Post on r/programming, r/ChatGPTCoding, r/golang. Share in Discord servers. | Reddit, Discord |

**Launch week goals:**
- 500+ CLI installs
- 5,000+ pearcode.dev visits
- 50+ trial signups
- 10+ Pro conversions ($30/mo)
- PH Top 5, HN front page (stretch)

### Post-Launch: Weeks 6-10

| Activity | Cadence | Goal |
|---|---|---|
| LinkedIn posts | 3x/week | Continue build-in-public + user story content |
| Substack | Weekly | Shift to user insights + technical deep-dives |
| X engagement | Daily | Respond to mentions, engage in threads, post demo clips |
| User feedback collection | Continuous | In-app feedback prompt after session 5, DM power users |
| Iterate prompts server-side | Weekly | Improve teaching quality without CLI releases |
| v1.5.1 MCP release | Week 6 | Distribution play — Pear inside Claude Code/Cursor |

### Months 3-6: Scale & Monetize

| Activity | Timeline | Goal |
|---|---|---|
| **Early adopter pricing remains** | Ongoing (manual change when ready) | No artificial deadline — change when growth and product quality justify it |
| **Team tier outreach** | After 50+ individual subscribers | Personal outreach to users who mentioned teams. "What would your team need?" |
| **Startup course launch** | Month 4-5 | Formalize "Building a Startup" Substack content into paid course. Pear is the case study. |
| **Conference talks / podcast appearances** | Month 3+ | Submit to DevRelCon, local meetups. Pitch: "The AI-assisted learning crisis" |
| **SEO content** | Month 3+ | Blog posts targeting "learn coding with AI", "AI tutor for developers", "voice coding tool" |

---

## The Demo That Sells Everything

**The single most important marketing asset is a 60-second screen recording.** Every channel converges on this one artifact.

### The Script

1. **(0-5s)** Developer opens terminal. `pear` is running. Screen shows: "Hold [Space] to talk."
2. **(5-15s)** Developer holds space: *"I just accepted a bunch of Claude suggestions in this file. Walk me through what changed and tell me if it's solid."*
3. **(15-20s)** Screen shows context injection happening in real-time: "Enriching... git diff (3 files), src/auth/rbac.go, file tree"
4. **(20-50s)** Pear responds with a structured teaching breakdown. Each change is explained — *what* it does, *why* it matters, *what to watch for*. Flags a subtle issue. Offers to go deeper.
5. **(50-55s)** Developer holds space: *"Explain that RWMutex thing more."* — Multi-turn follow-up.
6. **(55-60s)** Pear responds with a deeper explanation grounded in the developer's actual code.

### Why This Works

The viewer *sees* three things:
1. **Context injection is visible.** The prompt isn't just "explain my code" — it's enriched with diff, tree, and active files. This is the "not just a wrapper" proof.
2. **Teaching quality gap.** The response isn't "here's the fix." It's "here's the fix, here's the concept, here's why it matters at scale, want to go deeper?" Visibly better than raw Claude.
3. **Voice is natural.** Talking is faster than typing a prompt. The UX feels effortless.

**Invest a full day making this perfect.** Record 20 takes. Get the script right. This video appears on:
- pearcode.dev hero section
- LinkedIn launch post
- Product Hunt listing
- Substack launch essay
- X pinned tweet
- HN blog post

---

## Content Calendar: First 12 Weeks

### Pre-Launch (Weeks 1-4)

| Week | LinkedIn | Substack |
|---|---|---|
| 1 | "The AI-assisted learning crisis is real. Here's what I'm building about it." | Manifesto: "AI Makes You Fast. Pear Makes You Good." |
| 2 | "Building a voice pipeline in Go as a junior Go dev. Here's what I learned about sox, raw terminal mode, and concurrency." | "Why I chose Go for a voice AI product" |
| 3 | "I asked 10 engineering managers if AI tools are making their juniors better or worse. The answers were unanimous." | "The $30/mo question — pricing a developer education tool" |
| 4 | "Before/after: raw Claude response vs. Pear teaching response. Same question, same code. See the difference." | "What GitHub learned from killing Copilot Voice" |

### Launch (Week 5)

| Day | LinkedIn | Substack | Other |
|---|---|---|---|
| Tue | "Pear is live. 5 weeks, solo founder, Go CLI." + demo video | Launch essay + email to waitlist | X announcement |
| Wed | "We're live on Product Hunt. Here's why I built an AI tutor." | — | PH launch |
| Fri | "On Hacker News today. The responses have been incredible." | — | HN post |

### Post-Launch (Weeks 6-12)

| Week | LinkedIn | Substack |
|---|---|---|
| 6 | User story: "A junior dev told me Pear explained something their senior couldn't." | "Week 1 post-launch: numbers, surprises, what's next" |
| 7 | "How Pear's prompt engine works — the teaching system prompt that makes the difference." | "Iterating on teaching quality: what I learned from 500 sessions" |
| 8 | "MCP server mode is live — Pear now works inside Claude Code." | "The MCP distribution bet" |
| 9 | Hot take: "Vibe coding is creating a generation of developers who can't debug." | "Building a startup in 9 weeks: what I'd do differently" |
| 10 | "First 50 paying users. Here's what they told me." | Revenue transparency post |
| 11 | Technical: "Latency optimization for voice AI in the terminal." | "The build-in-public playbook that actually works" |
| 12 | "What I've learned about developer education in 3 months of building Pear." | Q1 retrospective |

---

## Metrics & KPIs

### Launch Week Targets

| Metric | Target |
|---|---|
| pearcode.dev visits | 5,000+ |
| CLI installs | 500+ |
| Trial signups | 50+ |
| Pro conversions | 10+ |
| PH ranking | Top 5 Product of the Day |
| HN front page | Yes |

### Month 1 Targets

| Metric | Target |
|---|---|
| Weekly active CLI users | 50+ |
| Pro subscribers | 20+ |
| MRR | $600+ |
| Voice sessions/user/week | 3+ |
| "Teach" mode usage | >60% of sessions |
| Substack subscribers | 1,000 |
| LinkedIn followers | 8,000 |

### Month 3 Targets

| Metric | Target |
|---|---|
| Weekly active CLI users | 200+ |
| Pro subscribers | 75+ |
| MRR | $2,250+ |
| Trial → Pro conversion rate | >5% |
| Team tier waitlist | 10+ companies |
| Substack subscribers | 2,500 |
| LinkedIn followers | 12,000 |

### Month 6 Targets

| Metric | Target |
|---|---|
| Weekly active CLI users | 500+ |
| Pro subscribers | 200+ |
| MRR | $6,000+ |
| First team tier customers | 3-5 companies |
| LinkedIn followers | 20,000 |
| Course revenue (if launched) | $2,000+/mo |

---

## The Startup Course Integration

### Timing

Don't launch the course at the same time as Pear. Let Pear establish first. Start formalizing course content around month 4, when you have:
- Real product metrics to share
- Genuine lessons learned
- A growing audience hungry for the "how I built this" narrative

### Course Structure (Preview)

Working title: **"Zero to Shipped: Building a Startup from Scratch"**

The course uses Pear as the running case study. Each module maps to a real decision you made:

| Module | Pear Case Study |
|---|---|
| Idea validation | Competitive landscape analysis, the "AI learning crisis" thesis |
| Product definition | PRD creation, scope decisions, what to build vs. defer |
| Pricing strategy | BYOK-first model, $30/mo flat pricing |
| Technical architecture | Go CLI, monorepo, solo-founder architecture tradeoffs |
| GTM strategy | Build-in-public, LinkedIn-first, PH/HN launch playbook |
| First 50 users | What worked, what didn't, honest metrics |
| Scaling from 50 to 500 | Community, content, team tier validation |

### Brand Separation

The course and Pear are complementary, not competing:
- **Pear** targets developers who want to learn software engineering
- **The course** targets founders who want to learn startup building
- The audiences overlap minimally but cross-pollinate: developers who become founders, founders who try developer tools
- Your personal brand sits at the center: "the person who builds and teaches"

---

## Risk Mitigation

| Risk | Mitigation |
|---|---|
| PH/HN launch falls flat | It's one channel, not the entire strategy. LinkedIn + Substack provide steady growth regardless of launch spikes. |
| "Just a wrapper" perception on HN | Lead with teaching quality, not voice. Show the before/after comparison. Be honest about what it is and isn't. |
| Low conversion from free to Pro | Free tier is deliberately limited (10 voice-minutes/day, teach mode only). Monitor friction points. If users aren't converting, the Pro value prop isn't clear enough — iterate messaging. |
| Audience confusion (tool maker vs. course seller) | Don't overlap the launches. Pear first, course later. Personal brand sits above both: "builder who teaches." |
| Content fatigue (3x/week LinkedIn + weekly Substack while building) | Batch content creation. One afternoon per week writes next week's posts. Build-in-public content writes itself — you're documenting what's already happening. |
| Solo founder bandwidth | Focus on one channel at a time. LinkedIn is the engine. Substack is the depth. Everything else is gravy. Don't try to be everywhere — be consistent where you are. |

---

## Summary: The GTM Thesis

Pear's GTM strategy is founder-led content marketing with a product launch event. The competitive advantage isn't paid ads or a growth team — it's Mitch's existing audience, willingness to build in public, and the narrative that building Pear itself teaches startup lessons worth sharing.

The flywheel:
1. Build Pear → generates authentic content
2. Content grows audience → drives installs
3. Usage data improves product → generates more content
4. Personal brand grows → unlocks course revenue + speaking + advisory
5. Course audience → discovers Pear → back to step 2

**The goal is not just to launch a product. It's to launch a brand, a community, and a category — "AI developer education" — and own it from day one.**
