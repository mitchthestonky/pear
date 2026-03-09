package prompt

import (
	"fmt"
	"strings"

	"github.com/MitchTheStonky/pear/cli/llm"
	"github.com/MitchTheStonky/pear/cli/repocontext"
)

// UserProfile holds developer profile info for prompt calibration.
type UserProfile struct {
	Name      string
	Languages string
	Level     string
}

const proactiveSystem = `You are Pear, a pair programmer. Review code changes, then teach the underlying patterns.

Profile: %s | %s | %s

Behavior:
- Observe the diff, then explain the pattern/mechanism behind it
- 2-3 observations max, depth over breadth
- Level calibration: junior=full explanation with context, intermediate=non-obvious parts + edge cases, senior=tradeoffs + failure modes only
- Hard bugs: state clearly what breaks, under what conditions, how to spot this class of problem next time
- Subtle risks: lighter flag, explain the conditions where it becomes a problem
- Tone: soft evaluation ok ("this holds up well — here's why"), no empty praise ("great catch", "smart move", "nice work")
- No greeting by name, no questions, no quizzes, no suggestions to read docs

Tags (end of response):
📚 Concepts: [...]
🔗 Related: [... → ...]`

const reactiveSystem = `You are Pear, a pair programmer answering a developer's question.

Profile: %s | %s | %s

Behavior:
- Answer the question first, then teach the underlying concept
- Ground answers in their actual code when possible
- Level calibration: junior=full explanation, intermediate=direct answer + non-obvious details, senior=concise answer + tradeoffs
- Tone: direct and practical, soft evaluation ok, no empty praise
- No greeting by name, no questions, no quizzes

Tags (end of response):
📚 Concepts: [...]
🔗 Related: [... → ...]`

const deepDiveSystem = `You are Pear, a pair programmer giving a thorough explanation of a topic.

Profile: %s | %s | %s

Behavior:
- Start with what the concept is and why it exists
- Use their codebase to illustrate, then broaden to the general pattern
- Cover when it breaks, common mistakes, how to recognize it in unfamiliar code
- Level calibration: junior=fundamentals up, intermediate=focus on gaps, senior=internals + tradeoffs + history
- Tone: thorough but conversational, not academic. Soft evaluation ok, no empty praise.
- No greeting by name, no questions or quizzes unless user asks to be tested

Tags (end of response):
📚 Concepts: [...]
🔗 Related: [... → ...]`

// Proactive builds system prompt and messages for a proactive code review.
// History is capped to last 3 messages.
func Proactive(ctx *repocontext.RepoContext, profile UserProfile, history []llm.Message) (string, []llm.Message) {
	system := fmt.Sprintf(proactiveSystem, profile.Name, profile.Languages, profile.Level)

	var msgs []llm.Message
	// Include last 3 history messages
	if len(history) > 3 {
		msgs = append(msgs, history[len(history)-3:]...)
	} else {
		msgs = append(msgs, history...)
	}

	contextBlock := FormatContext(ctx)
	msgs = append(msgs, llm.Message{
		Role:    "user",
		Content: contextBlock,
	})

	return system, msgs
}

// Reactive builds system prompt and messages for a user-driven question.
// Full history is included.
func Reactive(ctx *repocontext.RepoContext, profile UserProfile, history []llm.Message) (string, []llm.Message) {
	system := fmt.Sprintf(reactiveSystem, profile.Name, profile.Languages, profile.Level)

	msgs := make([]llm.Message, len(history))
	copy(msgs, history)

	contextBlock := FormatContext(ctx)
	if len(msgs) > 0 && msgs[len(msgs)-1].Role == "user" {
		msgs[len(msgs)-1].Content = contextBlock + "\n\n" + msgs[len(msgs)-1].Content
	} else {
		msgs = append(msgs, llm.Message{
			Role:    "user",
			Content: contextBlock,
		})
	}

	return system, msgs
}

// DeepDive builds system prompt and messages for a topic deep-dive.
func DeepDive(ctx *repocontext.RepoContext, profile UserProfile, topic string) (string, []llm.Message) {
	system := fmt.Sprintf(deepDiveSystem, profile.Name, profile.Languages, profile.Level)

	contextBlock := FormatContext(ctx)
	msgs := []llm.Message{
		{
			Role:    "user",
			Content: contextBlock + "\n\nPlease teach me about: " + topic,
		},
	}

	return system, msgs
}

// FormatContext builds the context block injected into user messages.
func FormatContext(ctx *repocontext.RepoContext) string {
	var b strings.Builder

	b.WriteString("Here are the recent changes in the codebase:\n\n")

	if ctx.Branch != "" {
		fmt.Fprintf(&b, "**Branch:** %s\n", ctx.Branch)
	}

	if len(ctx.ChangedFiles) > 0 {
		fmt.Fprintf(&b, "**Changed files:** %s\n", strings.Join(ctx.ChangedFiles, ", "))
	}

	if ctx.Diff != "" {
		b.WriteString("\n```diff\n")
		b.WriteString(ctx.Diff)
		b.WriteString("\n```\n")
	}

	if ctx.FileTree != "" {
		fmt.Fprintf(&b, "\n**File tree:**\n%s\n", ctx.FileTree)
	}

	if len(ctx.Files) > 0 {
		b.WriteString("\n")
		for path, content := range ctx.Files {
			fmt.Fprintf(&b, "**@%s:**\n```\n%s\n```\n\n", path, content)
		}
	}

	return b.String()
}
