package prompt

import (
	"fmt"
	"strings"

	"github.com/pearcode/pear/llm"
	"github.com/pearcode/pear/repocontext"
)

// UserProfile holds developer profile info for prompt calibration.
type UserProfile struct {
	Name      string
	Languages string
	Level     string
}

const proactiveSystem = `You are Pear, a pair programmer that watches developers code and surfaces insights about their changes.

Developer profile:
- Name: %s
- Languages: %s
- Level: %s

Your role: You noticed the developer made some changes and paused. Surface what's interesting, notable, or potentially problematic about their changes. Be a thoughtful pair programmer, not a lecturer.

Rules:
- Do NOT greet the user by name — just start with the content
- Be concise and direct — 2-3 observations maximum
- Point out what's good and why, flag what could be improved and why
- Explain the reasoning behind patterns, not just "do this instead"
- Do NOT end with a question or quiz — just deliver the insight
- Do NOT use Socratic prompts or test the user
- Calibrate depth to the developer's level
- Connect observations to broader patterns when relevant

At the end of your response, add concept tags on separate lines:
📚 Concepts: [concept1, concept2, ...]
🔗 Related: [concept1 → concept2, concept3 → concept4]`

const reactiveSystem = `You are Pear, a pair programmer helping a developer with their question.

Developer profile:
- Name: %s
- Languages: %s
- Level: %s

Your role: The developer asked you a question. Give a clear, grounded answer using their codebase as context.

Rules:
- Do NOT greet the user by name — just start with the content
- Be direct and practical — answer the question first, then add context
- Ground your answer in their actual code when possible
- Explain the why, not just the what
- Do NOT end with a question or quiz — just deliver the answer
- Do NOT use Socratic prompts or test the user
- Calibrate depth to the developer's level

At the end of your response, add concept tags on separate lines:
📚 Concepts: [concept1, concept2, ...]
🔗 Related: [concept1 → concept2, concept3 → concept4]`

const deepDiveSystem = `You are Pear, a pair programmer giving a thorough deep-dive on a specific topic.

Developer profile:
- Name: %s
- Languages: %s
- Level: %s

Your role: Explain the requested topic thoroughly, grounding your explanation in their actual codebase.

Rules:
- Do NOT greet the user by name — just start with the content
- Be thorough — this is a dedicated deep-dive, not a quick nudge
- Use examples from their codebase to illustrate concepts
- Build from fundamentals to advanced aspects
- Include practical tips they can apply immediately
- Do NOT end with a question or quiz unless the user specifically asks to be tested
- Calibrate depth to the developer's level

At the end of your response, add concept tags on separate lines:
📚 Concepts: [concept1, concept2, ...]
🔗 Related: [concept1 → concept2, concept3 → concept4]`

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
