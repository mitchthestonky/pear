# Pear — UX & Product Value Recommendations

> Date: 2026-03-06
> Source: Analysis of dogfood session export (pear-session-2026-03-06-144730.md)
> Perspective: Staff/Principal Product Design Engineering review

---

## Summary

A 14-review dogfood session revealed that Pear's core teaching loop has structural issues: duplicate reviews, flat depth, formulaic output, and a review tone rather than a teaching tone. The Socratic questions specified in the PRD are absent entirely. These issues compound — a user who receives the same insight three times in the same format stops reading, and there's no mechanism to detect that disengagement.

The fixes below are ordered by implementation priority. Items 1-3 address the core value gap and should ship in the next iteration.

---

## P0: Duplicate Detection — Stop Repeating Yourself

**Problem:** The same insight appears 2-3x in a single session. Schema Template Matching appeared 3 times, bigint/symmetric indexes 2x, SSE monotonic ID 2x, idempotency 2x, rollback strategies 2x. The concept graph in `learning.json` is not being consulted before generating reviews.

**Recommendation: Session-level review fingerprinting**

1. After each review, extract a short summary (concept tags + first sentence of each bullet). Store in-memory on a `SessionMemory` struct attached to `tui.Model`.

2. In `prompt/assembler.go`, when building the proactive prompt, append a "previously covered" block:

   ```
   ## Already covered this session (do not repeat):
   - Idempotency: validate-and-skip vs destructive DROP TABLE
   - SQS FIFO: deduplication window and re-upload edge case
   - Secret rotation: rotate-then-scrub sequencing
   ```

3. Add a system prompt instruction: "If your review would primarily cover topics listed above, either skip this review entirely by responding with `[SKIP]`, or focus only on what is genuinely new."

4. In the TUI update loop, if the LLM response starts with `[SKIP]`, silently discard it and re-issue the watcher listener. The user never sees it.

**Why not just use `learning.json`?** The JSON tracks cross-session concept frequency, but duplicate detection is a within-session problem. A 30-minute session where the user iterates on the same file triggers multiple settle events on overlapping diffs. Session memory is ephemeral — lives on the Model struct, dies when the session ends.

**Touch points:** `prompt/assembler.go` (inject previously-covered), `tui/app.go` (store review summaries, handle `[SKIP]`), system prompt templates.

---

## P0: Progressive Depth — Build on Prior Teaching

**Problem:** When a concept reappears, Pear restates instead of deepening. A good teacher would introduce, then challenge, then connect to broader architecture.

**Recommendation: Tiered teaching instructions keyed to concept frequency**

Builds on the session memory from the duplicate detection fix. When the prompt assembler detects a revisited concept (within-session from memory, or cross-session from `learning.json`), modify the teaching instruction:

```go
// prompt/assembler.go
func teachingTier(concept string, sessionCount int, totalCount int) string {
    switch {
    case sessionCount == 0 && totalCount == 0:
        return "introduce"  // First encounter: explain the what and why
    case sessionCount > 0:
        return "deepen"     // Seen this session: challenge or connect, don't restate
    case totalCount >= 3:
        return "connect"    // Well-trodden: only mention if there's a novel angle
    default:
        return "reinforce"  // Seen before, not recently: brief callback
    }
}
```

Inject into the system prompt as behavioral guidance:

```
## Teaching depth for known concepts:
- "Idempotency" -> DEEPEN (covered earlier this session). Do not restate.
  Challenge an assumption or connect to a new part of the architecture.
- "Secret Rotation" -> CONNECT (4 prior sessions). Only mention if there's
  a genuinely new angle the user hasn't seen.
- "Functional Indexes" -> INTRODUCE (first encounter). Explain fully.
```

**Touch points:** `learning/tracker.go` (expose lookup by concept), `prompt/assembler.go` (generate tier instructions), system prompt templates.

---

## P1: Output Variety — Break the Formula

**Problem:** Every review follows the exact same skeleton: intro paragraph, 3 bullets (praise + pattern name), one "watch out" paragraph, tags. After 3-4 reviews this becomes wallpaper.

**Recommendation: Multiple proactive prompt variants, randomly selected**

In `prompt/assembler.go`, maintain 4-5 structural variants instead of one proactive template:

| Variant | Structure | When to use |
|---|---|---|
| Standard | 2-3 observations + question | Default |
| Deep single | One concept explored thoroughly | Diff touches one complex area |
| Gotcha-first | Lead with the bug/risk, then explain | Concrete issue exists |
| Pattern recognition | "You've been doing X — here's the name for it" | Repeated patterns across reviews |
| Contrast | "You did X here but Y there — here's why" | Diff spans multiple files with different approaches |

Selection logic: track the last 2 variants used in session memory and avoid repeating. Weight toward "deep single" when the diff is focused (1-2 files) and "standard" when it's broad (4+ files).

Add per-variant format instructions to the system prompt:

```
## Response format for THIS review: Deep Single
Focus on the single most interesting pattern in this diff. Go deep —
explain the underlying concept, show where it appears in industry,
and identify one edge case the developer hasn't handled. Do NOT use
a bulleted list of observations. Write in flowing paragraphs.
```

**Touch points:** `prompt/assembler.go` (variant selection + templates), `tui/app.go` (track last variant in session state).

---

## P1: Teaching Tone vs. Review Tone

**Problem:** Pear says "great catch" and "smart move" instead of teaching why something works. The PRD says "Teach the concept, not just the fix" but the output reads like a senior architect code review, not a teaching session.

**Recommendation: Revise the system prompt identity and add anti-patterns**

Primarily a prompt engineering fix in `prompt/assembler.go`:

```
## Who you are
You are a teaching pair programmer. You do NOT evaluate whether decisions
are good or bad. You explain WHY things work, WHERE patterns come from,
and WHAT would break if the approach changed.

## Anti-patterns (never do these):
- "This is a smart move" / "Great catch" / "The right call" -> Instead:
  explain the mechanism that makes it work
- "This is a sophisticated pattern" -> Instead: name the pattern,
  explain where it originated, show a case where it fails
- Validating the developer's choices -> Instead: deepen their
  understanding of the choices they've already made

## Examples:
BAD: "Using LEAST/GREATEST for symmetric relationships is a clean approach."
GOOD: "LEAST/GREATEST normalizes the pair order at the index level. Without
it, you'd need a CHECK constraint or application-side sorting — and either
can drift. This is the same trick social networks use for mutual-friendship
tables."
```

Key insight: praise feels good for one review, but across a session it becomes noise. Explanation compounds — each review makes the user more capable.

**Touch points:** `prompt/assembler.go` (system prompt rewrite). No code changes, just prompt content.

---

## P1: Enforce Socratic Questions

**Problem:** The PRD specifies questions ending proactive reviews. Zero appear in the session export.

**Recommendation: Two-layer enforcement**

**Layer 1 — Prompt hardening:**

Move the question instruction from a general principle to a structural requirement:

```
Your response MUST end with exactly one line starting with a question
that asks something specific and answerable about the code the developer
just wrote. The question should require them to think, not just look
something up.

Good: "What happens to your SQS deduplication if the user deletes and
re-uploads within the 5-minute window?"
Bad: "What do you think about this approach?"
Bad: "Have you considered error handling?" (too vague)
```

**Layer 2 — Post-processing validation in `tui/app.go`:**

After the stream completes, check if the response contains a question. If not, make a lightweight follow-up call with a focused prompt to generate one, and append it to the rendered output. This is a safety net — Layer 1 should handle 95% of cases.

**Touch points:** `prompt/assembler.go` (prompt hardening), `tui/app.go` (optional post-processing check).

---

## P2: Export Format with Metadata

**Problem:** The export is a wall of `## Pear` sections with no timestamps, triggers, diffs, or user messages. Can't tell when something happened, what triggered it, or whether the user engaged.

**Recommendation: Structured session export with full context**

The `/export` slash command should produce:

```markdown
# Pear Session - 2026-03-06 14:47-16:12
Provider: openrouter/claude-3.5-sonnet - 14 reviews - 23 concepts

---

## Review 1 - settle - 14:52 - 3 files, +47 lines
Branch: feat/data-service

> Diff summary: Modified `doc/api-contracts.md`, `doc/decisions.md`

[Pear's review content]

---

## Review 2 - commit - 15:03 - "feat: stateless query service"
Branch: feat/data-service

**User follow-up:** "why did you suggest..."

[Pear's follow-up response]
```

Implementation: extend `tui.Model` to store a `[]SessionReview` slice:

```go
type SessionReview struct {
    Timestamp   time.Time
    TriggerType string        // "settle", "commit", "user"
    TriggerInfo string        // summary or commit message
    Branch      string
    DiffSummary string        // "3 files, +47 lines"
    Response    string        // full LLM response
    FollowUps   []FollowUp   // user questions + responses
    Concepts    []string
}
```

The `/export` handler iterates this slice and renders structured markdown.

**Touch points:** `tui/app.go` (store reviews in session), new or extended `/export` handler, `repocontext/collector.go` (expose diff summary string).

---

## P2: Concept + Insight Tracking

**Problem:** `learning.json` tracks `"Idempotency": { count: 4 }` — useless for recall. It tells you the topic but not what was learned.

**Recommendation: Track concept-insight pairs**

Extend the `Concept` struct:

```go
type Concept struct {
    Count    int              `json:"count"`
    Sessions []string         `json:"sessions"`
    Related  []string         `json:"related"`
    Insights []Insight        `json:"insights"`  // NEW
}

type Insight struct {
    Summary   string `json:"summary"`   // one-line takeaway
    Session   string `json:"session"`   // ISO timestamp
    Codebase  string `json:"codebase"`  // which repo
}
```

Add a new LLM tag to the system prompt alongside the existing concept and related tags:

```
Insights: ["LEAST/GREATEST normalizes pair order at index level",
"Destructive idempotency creates race conditions for concurrent readers"]
```

Parse with regex in `learning/tracker.go`, associate with parent concepts.

Then `pear progress` becomes genuinely useful:

```
$ pear progress

Concepts Pear has taught you:

  Idempotency                    8/10  4 sessions
    - validate-and-skip > destructive DROP TABLE
    - SQS visibility timeout as retry buffer (2x processing time)
    - schema checksum needed for same-schema-new-data case

  Secret Rotation                6/10  3 sessions
    - rotate-then-scrub sequence (validate before history rewrite)
    - BETTER_AUTH_SECRET rotation invalidates all user sessions
```

**Touch points:** `learning/tracker.go` (new struct, extraction regex, dedup insights), `cmd/progress.go` (render insights), system prompt templates.

---

## P3: Engagement Detection / Adaptive Frequency

**Problem:** 14 reviews fired with no evidence the user read any of them.

**Recommendation: Collect signals in v0, act on them in v1.5**

Three signals to start collecting with minimal effort:

**Signal 1 — Follow-up rate:** Already available. Track `reviews_given` vs `follow_ups_asked` in `SessionStats`. If the ratio drops below 1:5, the reviews aren't landing.

**Signal 2 — Scroll behavior:** The Bubble Tea viewport already tracks scroll position for auto-scroll locking. If a proactive review renders and the user doesn't scroll or interact within 30 seconds before the next change event, mark it as `unread`.

**Signal 3 — Lightweight feedback:** After each review completes, briefly show a subtle hint:

```
> helpful? (y/n/enter to skip)
```

Auto-dismisses after 10 seconds or on any other input. Zero-friction but gives a direct signal.

**Adaptive behavior (v1.5, not v0):**

```go
func (m *Model) shouldReview() bool {
    if m.stats.UnreadReviews >= 3 {
        return false  // User isn't reading. Back off.
    }
    if m.stats.NegativeFeedback >= 2 && m.stats.PositiveFeedback == 0 {
        m.promptUserAboutFrequency()
        return false  // Reviews aren't helpful. Pause and ask.
    }
    return true
}
```

**For v0:** Just collect the data. Log follow-up rate and scroll-after-review to `~/.pear/logs/`. Analyze manually during dogfooding. Don't build adaptive logic until there's real signal.

**Touch points (v0 only):** `tui/app.go` (track follow-up count, optional scroll tracking), `logging/` (emit engagement events).

---

## Implementation Order

| Step | Items | Rationale |
|---|---|---|
| 1 | System prompt rewrite | Covers P1 tone + P1 Socratic + half of P1 variety. Highest ROI, zero code changes. |
| 2 | Session memory + previously-covered injection | P0 duplication. Most user-visible bug. |
| 3 | Teaching tiers | P0 progressive depth. Builds on session memory. |
| 4 | Prompt variants | P1 variety. Builds on rewritten prompt. |
| 5 | Export format | P2 quality-of-life for dogfooding. |
| 6 | Insight tracking | P2 extends existing learning infra. |
| 7 | Engagement signals | P3 collect passively, act on later. |

Steps 1-3 address the core value gap and should ship in the next iteration. They fix the fundamental problem: Pear talks too much, says the same things, and reviews instead of teaches.
