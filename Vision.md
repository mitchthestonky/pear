# Pear — Vision

> Last updated: March 2026

---

## One-Liner

**The pair programmer that makes sure you understand your code.**

---

## The Problem

AI agents are writing more of your code every day. Cursor, Claude Code, Copilot — they make you fast. But they don't make you better.

1. **You're shipping code you don't understand.** AI tools optimize for output, not comprehension. The gap between "code that works" and "code you understand" is growing.
2. **Nobody has time to teach you.** Senior engineers are overwhelmed. Mentorship is broken. The 10:1 junior-to-senior ratio means most developers learn by guessing.
3. **Learning happens somewhere else.** Courses, docs, and tutorials exist in a vacuum. They don't activate in the moment you encounter a real pattern in your real codebase.

### The Data

- 19% slower with AI on familiar codebases (METR, 2025)
- 67% drop in junior developer roles (Stanford)
- 66% of developers spend time fixing flawed AI-generated code (Stack Overflow)
- 4x increase in code duplication (GitClear)

---

## The Solution

Pear is a CLI pair programmer that sits alongside your AI coding tools. It watches your code as you work — including code written by agents — and teaches you what changed and why it matters.

- **Learns with you, not for you.** Pear helps you understand what AI agents are doing to your codebase so you grow with the tools, not behind them.
- **Your code, explained.** Pear watches your coding sessions and teaches you what changed, why it works, and what to watch for.
- **The developer who stays sharp.** AI agents write your code. Pear makes sure you still understand it.

---

## How It Works

Pear runs in a terminal alongside your editor. It monitors your repo via file watching and git, and teaches during natural pauses in your workflow.

1. **Watch** — Pear detects file changes and pauses in your coding. On pause (~30s) or commit, it reviews what changed.
2. **Teach** — Pear streams a contextual teaching response: what changed, why it matters, what patterns to learn.
3. **Ask** — You can ask follow-up questions, request deep dives on topics, or reference specific files with `@file`.

Three interaction modes:
- **Watch mode** (`pear watch`) — proactive, always-on pair programming
- **Interactive mode** (`pear`) — on-demand REPL for questions and conversation
- **One-shot commands** (`pear ask`, `pear review`, `pear teach`) — quick answers

---

## Positioning

### What Pear Is

A **complementary layer** for developers using AI coding tools. Cursor writes your code. Claude Code builds your features. Pear makes sure you understand what just happened and why.

### What Pear Is Not

- Not a replacement for Claude Code, Cursor, or Copilot
- Not a linter or static analysis tool
- Not an online course or tutorial platform
- Not an LLM wrapper — Pear's value is the teaching engine, context injection, and concept tracking

### Tagline

**"AI makes you fast. Pear makes you good."**

---

## Target Users

1. **Junior-to-mid developers (0-5 years)** — Primary. Learning on the job, using AI tools daily, worried about skill gaps. Want to understand, not just ship.
2. **Vibe coders** — Non-traditional developers building with AI who want to actually learn what they're shipping.
3. **Any developer using AI agents** — Experienced developers who want to stay sharp as agents write more of their code.

---

## Pricing

- **Free (BYOK)** — $0 forever. Bring your own API key. Watch mode, interactive Q&A, context injection, concept tags.
- **Pro** — $20/month ($130/year). Learning state memory, adaptive teaching, concept tracking, progress visibility, cross-machine sync.
- **Teams** — $30/seat/month. Team-wide learning metrics, shared billing.

BYOK means near-zero infrastructure costs. Pear handles the teaching system — you control your LLM costs.

---

## Competitive Landscape

No funded company is building "learning at point of execution." The space is wide open.

| Tool | What it does | Pear's relationship |
|------|-------------|---------------------|
| Claude Code | AI coding agent | Pear teaches you what Claude Code built |
| Cursor | AI-powered editor | Pear explains the suggestions you accepted |
| Copilot | Inline completions | Pear helps you understand the code you kept |
| Aider | CLI coding agent | Pear is the teaching layer Aider doesn't have |

---

## Long-Term Vision

The v0 CLI is structured LLM prompts with context injection. Over time, Pear builds a real learning engine:

- **Concept graphs** — Track which concepts a developer understands and which have gaps
- **Proficiency mapping** — Know not just what you've seen, but what you've mastered
- **Misconception modeling** — Identify and correct specific misunderstandings
- **Spaced retrieval** — Surface concepts at intervals that reinforce long-term retention
- **Team analytics** — Help engineering managers understand where their team's knowledge gaps are

The moat is the learning data flywheel: more users → better teaching → more retention → more data → better personalization.
