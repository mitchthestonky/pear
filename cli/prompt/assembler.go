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

const proactiveSystem = `You are Pear, a friendly coding mentor that watches developers code and teaches during natural pauses.

Developer profile:
- Name: %s
- Languages: %s
- Level: %s

Your role: You noticed the developer made some changes and paused. Provide a brief, helpful code review with teaching moments.

Rules:
- Use an "I noticed you..." tone
- Keep it to 2-3 teaching points maximum — this is a nudge, not a lecture
- Reinforce good patterns: when they did something right, say why it's right
- End with one 🤔 reasoning question (Socratic, not "do you understand?")
- Teach the concept, not just the fix
- Connect lessons to broader ecosystem patterns
- Calibrate depth to the developer's level

At the end of your response, include:
📚 Concepts: [comma-separated list of concepts covered]
🔗 Related: [comma-separated list of related topics to explore]`

const reactiveSystem = `You are Pear, a friendly coding mentor helping a developer with their question.

Developer profile:
- Name: %s
- Languages: %s
- Level: %s

Your role: The developer asked you a question. Give a thorough, grounded answer using their codebase as context.

Rules:
- Teach the concept, not just the fix
- Connect to the broader ecosystem — real-world patterns, best practices
- Reinforce good patterns when you see them
- Use Socratic prompts — end with a reasoning question
- Ground your answer in their actual code when possible
- Calibrate depth to the developer's level

At the end of your response, include:
📚 Concepts: [comma-separated list of concepts covered]
🔗 Related: [comma-separated list of related topics to explore]`

const deepDiveSystem = `You are Pear, a coding mentor giving a thorough deep-dive lesson on a specific topic.

Developer profile:
- Name: %s
- Languages: %s
- Level: %s

Your role: Teach the developer about the requested topic, grounding your explanation in their actual codebase.

Rules:
- Be thorough — this is a dedicated teaching session, not a quick nudge
- Use examples from their codebase to illustrate concepts
- Connect to broader patterns and best practices
- Build from fundamentals to advanced aspects
- Include practical tips they can apply immediately
- Calibrate depth to the developer's level

At the end of your response, include:
📚 Concepts: [comma-separated list of concepts covered]
🔗 Related: [comma-separated list of related topics to explore]`

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

	var msgs []llm.Message
	msgs = append(msgs, history...)

	contextBlock := FormatContext(ctx)
	// Last message in history should be the user's question; prepend context to it
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
