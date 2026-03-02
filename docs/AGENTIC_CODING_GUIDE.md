# Agentic Coding Guide

> How to orchestrate AI agents to build software projects in hours instead of days.

---

## The Core Principle

Your job as orchestrator is **managing interfaces and dependencies, not writing code**. Agents write code. You define the contracts between them, sequence the work, and verify the output.

---

## Phase 0: Architecture & Decomposition (20% of total time)

This is the highest-leverage phase. Every minute here saves ten minutes of agent confusion later.

### Write Four Documents

**1. Project Instructions (`CLAUDE.md` or equivalent)**
- What the project is, in one paragraph
- Directory structure and module layout
- Critical rules: things agents MUST do, things they must NEVER do
- Key interfaces and type signatures
- Dependencies and their versions
- Common mistakes specific to your stack

The "Do NOT" section is the most important part. Agents will make the same mistakes repeatedly without explicit guardrails.

**2. Architecture Doc**
- Package/module responsibilities
- Data flow between packages
- Concurrency model (if applicable)
- State machine definitions
- Key design decisions and *why*

**3. Event/Behavior Model**
- Every user-facing flow described as pseudocode
- Input → processing → output for each feature
- Edge cases called out explicitly
- Error handling expectations

**4. Tickets**
- One file per unit of work
- Each ticket specifies: files to create/modify, function signatures, acceptance criteria, dependencies on other tickets
- Define exact interfaces at package boundaries *in the ticket itself* — both producer and consumer tickets should contain the same type definitions

### Decomposition Rules

- **Right-size tickets:** 1-5 files, 50-300 lines each. Too small = overhead. Too large = agents lose focus.
- **Explicit dependencies:** Every ticket lists which other tickets must complete first.
- **Interface-first:** Define the types and function signatures before implementation. Put them in the tickets.
- **One owner per file:** No two parallel agents should modify the same file. If unavoidable, designate one agent as owner and have others work against the interface.

---

## Phase 1: Wave Planning

Organize tickets into waves based on the dependency graph.

### Wave Structure

```
Wave 0:  You write the docs and tickets (manual)
Wave 1:  Foundation — establish interfaces, types, scaffolding
         1 agent, sequential. This creates the contract everything else builds on.
         → Smoke test: does it build? Do types exist?

Wave 2:  Core packages — implement against the interfaces
         2-4 agents in parallel. Each owns a distinct package/module.
         → Quick audit agent in parallel (P0/P1 scan only)

Wave 3:  Integration — wire packages together, build features
         2-3 agents in parallel. This is where cross-package code lives.
         → Fix audit findings as a dedicated agent

Wave 4:  Final integration — touch shared files, polish
         1 agent, sequential. Resolves conflicts, handles shared state.
```

### Agent Sizing

- **Sweet spot: 3-5 tickets per agent.** Fewer than 3 isn't worth the overhead. More than 5 and the agent loses coherence.
- **Prefer fewer, larger agents** over many small ones. Each agent has startup cost (reading docs, understanding context) and coordination cost (permissions, resuming).
- **Keep one agent sequential for shared files.** Files like the main entry point, shared types, or the primary UI component will be touched by multiple features. Assign these to a single integration agent.

---

## Phase 2: Agent Prompts

### Prompt Template

```
cd /path/to/project

Read [project instructions file] and [architecture docs] first.

Implement these tickets IN ORDER:
1. path/to/ticket-1.md
2. path/to/ticket-2.md
3. path/to/ticket-3.md

[Stack-specific context or warnings]

After each ticket:
- Run [build command] and [lint command]
- Create a completion file in [completion directory]
- Commit with [commit convention]
```

### Prompt Principles

- **Start with "read X first"** — agents need the same context you have
- **Specify the order** — even if tickets seem independent, ordering prevents conflicts
- **Include stack-specific gotchas** — these are the things agents get wrong repeatedly
- **Define the completion ritual** — build, lint, completion file, commit. Every time. This gives you a verification checkpoint.
- **Commit your agent prompts** to the repo in an `AGENTS.md` file. When you need to re-run or resume, you're not reconstructing from memory.

---

## Phase 3: Execution

### Before Launching Agents

1. **Verify tooling permissions.** Run a test agent with a trivial command (build, commit) to confirm it can actually execute. Permission blocks were the #1 time sink in practice.
2. **Verify the foundation builds.** Wave 1 must produce a compiling project with all shared types defined before Wave 2 starts.
3. **Check completion markers between waves.** `ls tickets/completed/` (or equivalent) confirms dependencies are met.

### During Execution

- **Launch parallel agents in a single message** — don't send them one at a time
- **Don't duplicate agent work** — if an agent is researching something, don't search for the same thing yourself
- **Monitor for stuck agents** — the most common failure mode is permission blocks or missing dependencies. Resume quickly.
- **When an agent gets stuck, finish its work directly** rather than resuming it repeatedly. Three failed resumes = just do it yourself.

### Between Waves

- **Run the build** — every file from every agent must compile together
- **Run a quick audit** — not a full review, just "does the core loop actually work?" Check that the primary user flow executes end-to-end.
- **Check for interface mismatches** — parallel agents may have slightly different assumptions about shared types

---

## Phase 4: Audit

### When to Audit

- **Quick audit after Wave 2** — catch architectural issues before integration agents build on them
- **Full audit after all waves** — comprehensive review before polish

### What to Audit

Run an audit agent with a specific persona and checklist:

1. **Architecture & Code Quality** — package separation, error handling, concurrency bugs
2. **Product & UX** — does the primary flow work? Edge cases?
3. **Robustness** — no internet, bad input, corruption, large data
4. **Security** — credential handling, injection risks, input sanitization
5. **Performance** — memory growth, responsiveness, efficiency
6. **Gaps vs Spec** — what's implemented vs what was promised
7. **Top 10 Issues** — ranked P0-P3 with file paths and line numbers

### Acting on the Audit

Write findings to a file (e.g., `AUDIT.md`). Then either:
- Launch a new agent session with: "Read AUDIT.md. Fix all P0 and P1 issues in priority order."
- Or fix P0s yourself (they're usually 1-2 critical bugs) and delegate P1/P2 to agents.

---

## Anti-Patterns

**Don't skip the docs phase.** "Just start coding" with agents produces incoherent architecture that costs more to fix than to plan.

**Don't launch too many parallel agents.** 4 is usually the practical max. Beyond that, file conflicts and coordination overhead eat the parallelism gains.

**Don't let agents touch shared files in parallel.** This creates merge conflicts and subtle bugs that are hard to diagnose.

**Don't audit only at the end.** The most expensive bugs are architectural — catch them between Wave 2 and Wave 3, not after everything is built.

**Don't retry stuck agents indefinitely.** If an agent fails twice on the same issue, take over that piece and move on.

**Don't assume "implemented" means "working."** Agents are great at making things compile. They're less reliable at making things work end-to-end. Verify the core user flow manually between waves.

---

## Checklist

```
□ Project instructions doc with critical rules and "Do NOT" section
□ Architecture doc with package responsibilities and data flow
□ Event/behavior model with pseudocode for every flow
□ Tickets with explicit deps, interfaces, and acceptance criteria
□ Wave plan with dependency graph
□ Agent prompts committed to repo
□ Tooling permissions verified before launch
□ Wave 1 builds and types exist before Wave 2
□ Quick audit after Wave 2
□ Build check between every wave
□ Full audit after final wave
□ P0 issues fixed before shipping
```

---

## Expected Results

A well-orchestrated agent swarm can build a **2,000-5,000 line project in 1-2 hours** that would take a senior engineer 2-5 days solo. The tradeoff: you spend 20% of the time on architecture docs and tickets, and the output needs a targeted fix pass for the 2-3 critical bugs agents reliably miss (usually around integration points and end-to-end flows).

The ROI is highest for projects where:
- The architecture is well-understood upfront
- Packages have clean interfaces between them
- The work can be parallelized across 3-4 independent streams
- You have clear acceptance criteria for each piece
