# Pedagogy & Concept Hierarchy Design Notes

> Captured from conversation on 2026-02-16. Reference for v1.7+ adaptive learning system.

---

## Starting Point: What Pear Needs to Teach

Pear doesn't teach a curriculum — it teaches what you just built. The system can't control the order concepts appear. A user might encounter goroutines, channels, and error wrapping all in their first session.

The concept hierarchy isn't a syllabus. It's a **knowledge graph that tracks what the user has encountered vs. what they actually understand.**

---

## Key Insight: It's Not Just "Concepts"

Most of what developers need to understand AI-generated code isn't just declarative knowledge ("what is a channel?"). It's:

- **Skills** (procedural): "trace a goroutine leak", "reason about cancellation propagation"
- **Mental models** (conceptual): how the runtime scheduler works, what "blocking" means
- **Misconceptions** (bugs waiting to happen): "I thought defer runs only on success", "I assumed maps are safe for concurrent reads"

Traditional ITS research (ACT-R based tutors) treats these differently. Pear should too.

**The unit of adaptation shouldn't be "concepts in code." It should be "what misunderstanding will cause the next bug."**

---

## What Research Says Works for Programming

### Effective techniques (relevant to Pear)

- **Worked examples + self-explanation**: Pear's "explain this diff" is a worked example pipeline. Structure it so the user actively processes it (micro questions, prediction prompts) rather than passively reading.
- **Cognitive apprenticeship** (model → coach → scaffold → fade): model the reasoning ("why this change"), scaffold ("here's the invariant"), then fade to short prompts ("what would break if X").
- **Retrieval practice + spacing**: if Pear never asks the user to retrieve, you get "felt clear at the time" but low retention.
- **Parsons problems** (reordering code blocks) and "trace the output" checks: effective and fast. Perfect for CLI-friendly micro-checks.

### Where it's hard

- **Knowledge tracing for programming is noisy** — source code is unstructured evidence of understanding. Richer signals like user questions are more valuable.
- **Misconceptions dominate failures** in code understanding. Adaptive systems break when they treat a misconception as "concept not learned" instead of "wrong mental model."

### References
- ACT-R Intelligent Tutoring Systems: https://act-r.psy.cmu.edu/wordpress/wp-content/uploads/2012/12/122IntTutSys.pdf
- Worked Examples in Programming: https://dl.acm.org/doi/10.1145/3560266
- Cognitive Apprenticeship: https://ocw.metu.edu.tr/pluginfile.php/9107/mod_resource/content/1/Collins%20report.pdf
- Spaced Retrieval Practice: https://dl.acm.org/doi/10.1145/3291279.3339411
- Parsons Problems: https://www2.eecs.berkeley.edu/Pubs/TechRpts/2020/EECS-2020-88.pdf
- Knowledge Tracing in Programming: https://aclanthology.org/2025.acl-long.1343.pdf
- Misconceptions in Programming: https://dl.acm.org/doi/full/10.1145/3702652.3744209

---

## Concept Graph Design

### Node model (start simple, expand later)

Start with one node type: **concept** — with an optional `misconceptions` field.

Add skill/pattern/context node types in v2+ when usage data justifies it.

### Edge types (two only to start)

- `requires` (with strength: strong/weak)
- `related_to`

Don't over-engineer with `is_example_of`, `causes_bug`, etc. until needed.

### Granularity: two-layer model

- **Macro concepts (~50-60):** "concurrency coordination", "error handling", "memory & escape", "API boundaries"
- **Micro concepts (expand over time):** "buffered vs unbuffered channel semantics", "`errors.Is` vs `==`", "context cancellation propagation"

Start with macro + a small high-yield micro set. Expand from real diffs.

### Bootstrap sources

- CS2023 / CS2013 knowledge areas (IEEE): https://ieeecs-media.computer.org/media/education/reports/CS2023.pdf
- SWEBOK: https://ieeecs-media.computer.org/media/education/swebok/swebok-v3.pdf
- ACM Computing Classification System: https://www.acm.org/publications/class-2012
- CS-specific learning taxonomy: https://kar.kent.ac.uk/23997/1/TaxonomyFuller.pdf
- The Go Programming Language (book chapter structure)
- Effective Go

---

## User Knowledge State Model

Per-user, per-concept state progression:

```
unseen → seen_in_code → explained → checked → reinforced → solid
```

### What counts as evidence

| Signal | Strength |
|--------|----------|
| User asked a question about it | Strong |
| User edited relevant lines correctly | Medium |
| User passed a micro-check (predict output, choose invariant) | Strong |
| User read explanation | Weak |
| Repeat exposure across different contexts | Medium |

### Micro-checks (CLI-friendly)

- "Predict the output" — one-liner
- "Which change fixes this bug?" — multiple choice
- "In one sentence, why did the AI add this?" — free text (LLM-graded)
- Parsons problems (reorder code blocks) — works in terminal

---

## Pedagogy Policy (Rules-Based, v1.7)

Function signature: `(detected_concepts, user_knowledge_state, code_context) → (explanation_depth_per_concept, optional_micro_check)`

Rules:
- **Novelty**: unseen → teach fully. seen 3+ times in 7 days → skim.
- **Risk**: concurrency/security/data-loss concepts get deeper treatment regardless of state.
- **Prerequisite gaps**: missing strong prerequisite → teach prerequisite briefly first.
- **Micro-check trigger**: offer exactly one micro-check when there's a high-value new concept.

No ML. No learned policies. Just if/else rules until 100+ paying users.

---

## Data Model (Design Now, Build in v1.7)

### Tables

**ConceptNode**
- id, language, name, description, category
- misconceptions: array of {wrong_belief, correction, risk_level}

**ConceptEdge**
- from_id, to_id, edge_type (requires | related_to), strength (strong | weak)

**UserKnowledge**
- user_id, concept_id, state, last_seen_at, last_checked_at
- evidence_counts: {seen, explained, checked, corrected}

**LearningEvent** (append-only log — start logging in v1.5)
- timestamp, user_id, repo, file, diff_hunk
- detected_concepts[], user_actions (opened, asked, edited, checked, correct?)

---

## Failure Modes to Watch

1. **Misdiagnosis**: reading an explanation ≠ understanding. "Explained" is not evidence; "retrieved correctly later" is.
2. **Concept detection errors**: if the parser maps code to the wrong concept, personalization becomes random.
3. **Over-eager prerequisite nagging**: seniors will churn. Use weak edges and only escalate when a misunderstanding likely causes a real bug.
4. **Cognitive load spikes**: code diffs already have high intrinsic load. Don't dump theory on top.
5. **No misconception model**: if you don't represent misconceptions explicitly, you'll keep "teaching the concept" while the user keeps failing.
6. **Tutor vibe mismatch**: Pear is a learning tool, but micro-checks must still feel fast and useful for today's code — not like homework.

---

## Implementation Sequencing

| Version | What to build |
|---------|---------------|
| v1.5 | No concept tracking. Log every explanation as a `LearningEvent` with LLM-tagged concepts. Append-only logging only. |
| v1.7 | Build 50-concept Go macro graph. Add `UserKnowledge` state tracking. Adjust explanation depth in prompts. |
| v1.8 | Add optional micro-checks. Track evidence quality. |
| v2.0 | Add misconception nodes. Use accumulated data to identify top 20 Go misconceptions. Build targeted corrections. |

---

## Pending: ChatGPT Deliverables

Asked ChatGPT to produce:
1. JSON schema for concept graph + user knowledge state (simplified)
2. Initial Go macro concept set (~50-60 concepts) with misconceptions and prerequisite edges
3. Pedagogy policy spec as simple rules (Go function signature)
