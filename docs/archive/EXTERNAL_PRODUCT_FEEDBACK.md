# External Product Feedback — Feb 2026

## Source 1: Meeting with Josh & Logan (Memory/Learning Experts)

**Context**: Jacob connected Mitch with Josh and Logan, who spent ~1 year building learning/memory tech and ~2 years researching the psychology of learning before that. Designers/marketers by background, now running an agency with AI backend.

### Resources Recommended
- **Andy Matuschak** — independent researcher on psychology of learning systems. Created **Quantum Country** (learning system embedded directly into content about quantum computing). Key insight: separating learning tools from content doesn't work well; embedding the learning system into the content is more effective but extremely time-intensive.
- **Anki** — spaced repetition flashcard system. Works at high volume but can feel like rote repetition. UX is a known weakness.
- **Scrimba** — interactive coding tutorials with pause-and-code UX. Good reference for "learning while doing" interface design.
- **2008 paper on effective student study techniques** (~50 pages) — sent via LinkedIn. Was the backbone of Logan's learning research.
- **RAFF** — open source agentic coding that recurses on itself. Worth studying for recursive agent architecture.

### Key Insights

**1. Socratic questioning > information delivery**
Logan found Claude far more useful as a "podcast host" asking great questions than as an information provider. Mentors are most valuable for the questions they ask, not the information they provide. Any LLM can provide information — getting it to ask the right questions of the user is the real value.

**2. Don't break flow — teach in the gaps**
Biggest risk is slowing users down. When someone's in flow with Claude, breaking their enjoyment means they won't use the product. Teach during natural gaps: while Claude is thinking, while a plan is being reviewed, while code is compiling. Like loading screen tips in old games.

**3. The plan acceptance moment is the ideal intervention point**
When Claude presents a plan and user is about to hit "accept" — that's when to say "here's what this plan means for your codebase, here's the pattern being used, do you understand what's about to be implemented?"

**4. Go higher than syntax — teach engineering habits**
Logan suggested maybe you don't need to codify a coding language at all. Instead, figure out the "8 habits of a successful programmer" and implement that. LLMs are consistent with high-level craftsmanship; people get lazy. Teaching architectural thinking and system design trade-offs may be more valuable than syntax.

**5. Insight cascades / zoom out as circuit breaker**
When users are frustrated and tunnel-visioning on a specific bug, they need to zoom out — but they don't remember to. Pear can act as the circuit breaker that prompts zooming out to first principles. Scientific papers on "insight cascades" support this.

**6. A/B comparison teaching**
Show two approaches side-by-side with explanations of why one is better. Like Midjourney's variant model — the comparison is what teaches, not just commentary. "Here's why this approach is better than that one" → user learns and can execute the preferred response.

**7. Target no-time-pressure learners first**
People cranking under deadlines won't stop to learn. Hobby coders, side project builders, people learning new stacks with no time pressure are happier to pause. This could be a specific market targeting insight.

**8. Spaced repetition is the backbone**
Track what user knows, reintroduce concepts at optimal intervals. The challenge is making spaced repetition easy enough that unmotivated users still do it. Building it into the workflow (not as a separate activity) is key.

**9. Seamlessness is everything**
For the target audience (not top-level devs), it needs to be download-and-it-works simple. No complex setup.

### Logan & Josh's Business
- Agency-first (done-for-you) marketing/design with AI backend
- Looking for companies needing consistently updated marketing/design collateral
- Mitch offered to plug their services

---

## Source 2: External Product & Market Analysis

### Core Positioning Reframe

**Current**: "AI explains code" (too vague, everyone does this)
**Should be**: "A learning system, not an explainer"

**Killer headline**:
> "AI is making engineers faster. Pear makes engineers better."

**Subhead**:
> Pear is a learning engine for AI-generated code. It tracks what you understand, detects gaps, and teaches you at the moment of execution — not six months later when prod breaks.

### The 5 Real Differentiators (The Moat)

1. **Learning State Memory** — "Pear remembers what you know." No other tool does this.
2. **Concept Graph (Curriculum Engine)** — "Pear knows what concepts exist and how they relate." This is infrastructure, not UX fluff.
3. **Adaptive Pedagogy** — "Pear changes how it teaches you based on your behaviour." This is the actual IP.
4. **Intervention Timing** — "Pear interrupts at the moment learning is maximally useful." The UX wedge Claude won't solve.
5. **Skill Growth Tracking** — "Pear shows you how you're becoming a better engineer over time." The emotional hook.

### Market Segments

The pain is felt by:
- Self-taught devs
- Juniors using Cursor
- Bootcamp grads
- Founders vibe-coding SaaS
- Product managers writing code
- Seniors jumping stacks (Go, Rust, infra)

**New category**: "Learning at point of execution."

### Key Risk & Response

**Risk**: "Why wouldn't I just ask Claude to explain the code?"

**Response** (must be crystal clear on site):
- No memory of user learning state
- No pedagogy
- No proactive circuit breaker
- No spaced repetition
- No detection of conceptual gaps
- No "I know you don't understand channels yet"

### Problem Framing (Sharper)

AI tools let you ship code you don't understand. This creates:
- Brittle systems
- Shallow engineering judgment
- Slower career growth
- Fragile founders
- Dangerous production risk

### MVP Core Loop (Right-Sized)

1. Watch diffs
2. Detect concepts
3. Check knowledge state
4. Explain only what's new or misunderstood
5. Save learning state
6. Ask 1 reflective question max

Everything else (voice, spaced repetition, concept DAG) layers in after.

### Pricing Refinement

| Tier | Price | Features |
|------|-------|----------|
| Free | $0 | Limited explanations, no memory, no concept tracking |
| Pro | $20–40/mo | Learning state memory, concept progression, unlimited explanations, multi-language |
| Team | $10–25/dev/mo | Skill dashboard, knowledge risk detection, onboarding accelerator, junior ramp-up metrics |

**Team tier is the sleeper revenue play.** "Knowledge risk detection" and "junior ramp-up metrics" are things CTOs will pay for before individual devs pay $20/mo for self-improvement.

### Distribution Strategy Validation

Channels confirmed correct: X, HN, Indie Hackers, AI discourse, founder content, shipping logs, demo clips.

**Controversial positioning drives traffic**: "AI coding tools are creating worse engineers."

Landing page must convert attention into: waitlist, email capture, Discord, CLI install.

### Strategic Identity

Pear is not a SaaS first. It's a **movement product**: "Don't let AI make you dumber."

This attracts: thoughtful engineers, founders, dev influencers, learning nerds, serious builders. That's the right early user base.

### Immediate Tests (Pre-Build Validation)

1. Fake CLI demo (recorded GIF)
2. Show diff → explanation → learning state update
3. Post on X + HN
4. Measure: waitlist signups, DM volume, "holy shit I need this" reactions
5. If no emotional pull → reposition harder
