# 03-02: Prompt Assembler

## Summary
Builds system prompts and message arrays from context, user profile, and conversation history. Three prompt variants.

## Event Model Refs
- E6c: Assembler.Proactive(repoContext, last3History)
- E7: Assembler.Reactive(repoContext, fullHistory)
- E8 [teach with topic]: Assembler.DeepDive(repoContext, topic)

## Files to Create
- `cli/prompt/assembler.go`

## Functions
- `Proactive(ctx *repocontext.RepoContext, profile UserProfile, history []llm.Message) (string, []llm.Message)` — returns system prompt + messages for proactive review
- `Reactive(ctx *repocontext.RepoContext, profile UserProfile, history []llm.Message) (string, []llm.Message)` — returns system prompt + messages for user-driven question
- `DeepDive(ctx *repocontext.RepoContext, profile UserProfile, topic string) (string, []llm.Message)` — returns system prompt + messages for topic deep dive
- `FormatContext(ctx *repocontext.RepoContext) string` — builds the context block injected into user message

```go
type UserProfile struct {
    Name      string
    Languages string
    Level     string
}
```

## Prompt Templates
See PRD.md "Teaching Prompt Principles" and ARCHITECTURE.md prompt section. Key rules:
- Proactive: "I noticed you..." tone, 2-3 points max, one 🤔 question, request `📚 Concepts: [...]` and `🔗 Related: [...]`
- Reactive: full teaching prompt, same concept tagging
- Deep dive: thorough, codebase-grounded, request concept tagging

## Context Block Format
```
Here are the recent changes in the codebase:

**Branch:** {branch}
**Changed files:** {file list}

```diff
{diff content}
```

**File tree:**
{tree}

{@file contents if any}
```

## Acceptance Criteria
- Golden-file tests: fixed profile + fixed RepoContext → expected system prompt (compare with testdata/*.golden)
- History is included correctly: proactive gets last 3, reactive gets full
- Context block is well-formatted and under token budget
- All three variants request concept + relationship tagging

## Dependencies
- 03-01 (RepoContext type)
- 02-01 (llm.Message type)
