# Pear — Positioning & Copy Guide

> Source of truth for all website copy, messaging hierarchy, and positioning decisions.
> Every word on pearcode.dev should trace back to this document.

---

## Core Positioning Statement

**Pear is the voice-first AI tutor that teaches software engineering while you code.**

It's not a coding assistant. It's not a dictation tool. It's a dedicated learning and development tool that lives in the developer's terminal, uses their actual codebase as the teaching surface, and makes education explicit and intentional — not accidental.

Developers still use Claude Code to code. They use Pear to *learn*.

---

## Messaging Hierarchy

### Level 1: Tagline (7 words or fewer)

**"AI makes you fast. Pear makes you good."**

### Level 2: Subheadline (one sentence)

**"The voice-first AI tutor that teaches software engineering in your terminal — using your actual codebase, not abstract examples."**

### Level 3: Elevator pitch (30 seconds)

AI coding tools are making developers faster but not better. Juniors are shipping code they don't understand. Seniors are drowning in mentoring. Pear is a CLI tool that sits alongside your existing AI tools and teaches you *why* — why that pattern exists, why the LLM suggested that approach, why the alternative is dangerous. Voice in, teaching out, grounded in your real code.

### Level 4: Full positioning (website hero → problem → solution flow)

See page structure below.

---

## Tone

**Direct and confident.** No hedging, no fluff. Statements, not questions. Respect the reader's intelligence — they're developers, they can handle a direct thesis.

- Say "Pear teaches you why." Not "Pear can help you understand."
- Say "AI tools are making developers faster but not better." Not "Some people think AI tools might be..."
- Say "The fix is on line 47." Not "It appears the issue might be on line 47."

Inspired by: Linear's copy (precise, opinionated), Raycast's tone (developer-native, no marketing fluff), Stripe's documentation (clear, authoritative).

---

## Page Structure & Copy

### Section 1: Hero

**Layout:** Full viewport height. Tagline center-left, typewriter animation on load. Subheadline fades in after tagline finishes typing. Waitlist email capture directly below. Terminal mockup or subtle code texture in background.

**Tagline (typewriter):**
> AI makes you fast. Pear makes you good.

**Subheadline (fade in):**
> The voice-first AI tutor that teaches software engineering in your terminal.

**CTA:**
> [email field] **Join the waitlist** →

**Micro-copy under CTA:**
> Early access. macOS. Free to try with your own API key.

---

### Section 2: The Problem

**Section header:**
> The AI-assisted learning crisis is real.

**Three problem cards (stacked or side-by-side):**

**Problem 1: Developers are shipping code they don't understand.**
> 40% of junior developers deploy AI-generated code they can't fully explain. AI tools optimize for output, not understanding. The gap between "code that works" and "code you understand" is widening.

**Problem 2: Senior engineers can't scale mentoring.**
> Every team has a 10:1 ratio of juniors who need guidance to seniors who can provide it. AI tools make juniors faster without making them better — which means more code to review, more questions to answer, more fires to fight.

**Problem 3: Existing learning is disconnected from real work.**
> Udemy courses, documentation, and tutorials exist in a vacuum. They don't activate in the moment when a developer encounters a real pattern in their real codebase. Learning should happen at the point of work, not in a separate tab.

---

### Section 3: The Solution

**Section header:**
> Pear teaches you while you code.

**Body:**
> Pear is a CLI companion that sits alongside your existing AI coding tools. When you accept a Claude suggestion, Pear explains *why* it works. When your build breaks, Pear teaches you the pattern, not just the fix. Voice in, teaching out — grounded in your actual codebase.

**Three capability highlights:**

| Capability | Description |
|---|---|
| **Voice-first** | Hold space, ask a question, hear the answer. Your hands stay on the keyboard. |
| **Context-aware** | Pear reads your git diff, file tree, and error logs automatically. Every response is grounded in your real code. |
| **Teaching-first** | Three modes: teach (explain the concept), mentor (answer + one insight), pair (just solve it). You choose how deep to go. |

---

### Section 4: How It Works

**Section header:**
> Three steps. Under three seconds.

**Step-by-step (with terminal mockup animation):**

**Step 1: Talk**
> Hold Space, ask your question in plain English.
> *"I just accepted a bunch of Claude suggestions. Walk me through what changed."*

**Step 2: Enrich**
> Pear automatically injects your git diff, file tree, and error logs into the prompt.
> Context injection is visible — you see what's being sent.

**Step 3: Learn**
> Pear responds with a structured teaching breakdown. What changed, why it matters, what to watch for. Offers to go deeper.

---

### Section 5: The Teaching Difference

**Section header:**
> Not just the fix. The lesson.

**Side-by-side comparison:**

**Left column: "A coding assistant says:"**
> "The bug is on line 47 — you're using `==` against a slice. Use `slices.Contains()`."

**Right column: "Pear says:"**
> "The bug is on line 47. You're comparing a string to a slice with `==`, which always returns false in Go — slices aren't comparable with `==`. The fix is `slices.Contains()`. **This is a common Go gotcha.** Unlike Python where `in` works on lists, Go requires explicit membership checks. This pattern matters because RBAC logic like this is a security boundary — a silent `false` here means unauthorized access could slip through."

---

### Section 6: Pricing

**Section header:**
> Simple pricing. Bring your own AI.

**Pricing philosophy intro:**
> Pear is a teaching tool, not an LLM reseller. Bring your own API key (Claude, OpenAI, Gemini) and pay for the intelligence layer — the voice UX, automatic context injection, and pedagogical prompt engine.

**Two pricing cards:**

**Free**
- $0 / forever
- Bring your own API key
- 10 voice-minutes / day
- Text output only
- Teach mode only
- *Try it. See the difference.*

**Pro (Early Access)**
- ~~$30~~ **$20 / month**
- Everything in Free, plus:
- Unlimited voice
- Audio responses (TTS)
- All 3 modes (teach / mentor / pair)
- 50 hosted requests / month included
- Annual billing: $200/yr (2 months free)
- **Lock in early access pricing →**

**Teams — Coming Soon**
> Pear for your engineering team. Reduce senior mentoring overhead. Team learning analytics. Shared billing.
> **Join the team waitlist →**

---

### Section 7: Final CTA

**Section header:**
> AI makes you fast. Pear makes you good.

**Body:**
> Join the waitlist for early access. macOS. Free tier available day one.

**CTA:**
> [email field] **Join the waitlist** →

**Footer:**
> Built by Mitch. Follow the build on [LinkedIn] · [Substack] · [X]

---

## Key Phrases Bank

Use these consistently across the site, social, and marketing:

| Phrase | Use for |
|---|---|
| "AI makes you fast. Pear makes you good." | Tagline, hero, social bios, PH listing |
| "The AI tutor that lives in your terminal." | Subheadline, meta description, social |
| "Not just the fix. The lesson." | Teaching difference section |
| "Voice in. Teaching out." | Feature descriptions |
| "Learn while you code." | CTA contexts, pricing page |
| "Your codebase is the classroom." | Problem/solution sections |
| "The mentor those tools will never be." | Positioning against Claude Code / Cursor |
| "Teaching at the point of work." | L&D / team tier messaging |

---

## SEO Targets

### Primary keywords
- voice AI tutor for developers
- AI coding tutor CLI
- learn software engineering with AI
- voice-first developer education tool

### Secondary keywords
- AI teaching tool for programmers
- voice coding assistant terminal
- developer L&D tool
- AI pair programming tutor

### Long-tail
- "how to understand AI-generated code"
- "learn coding patterns from AI"
- "voice AI for software engineering education"
- "alternative to Udemy for working developers"

### Meta description (homepage)
> Pear is the voice-first AI tutor that teaches software engineering while you code. Talk to your codebase, learn the patterns, understand the why. macOS CLI.

### Title tag
> Pear — AI makes you fast. Pear makes you good.
