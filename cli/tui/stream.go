package tui

import (
	"context"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/MitchTheStonky/pear/cli/llm"
	"github.com/MitchTheStonky/pear/cli/prompt"
	"github.com/MitchTheStonky/pear/cli/repocontext"
)

func (m *Model) handleTrigger(trigger ReviewTrigger) tea.Cmd {
	m.state = "streaming"
	_ = m.input.SetEnabled(false)

	m.output.AppendHeader(fmt.Sprintf("🍐 Pear noticed you %s", trigger.Info))
	m.output.AppendContext(fmt.Sprintf("git diff, %s", repocontext.DiffSummary(trigger.Diff)))
	m.output.StartStream(m.width)

	profile := prompt.UserProfile{
		Name:      m.config.Name,
		Languages: m.config.Languages,
		Level:     m.config.Level,
	}

	ctx := &repocontext.RepoContext{
		Diff:         trigger.Diff,
		ChangedFiles: repocontext.ParseChangedFiles(trigger.Diff),
		TriggerType:  trigger.TriggerType,
		TriggerInfo:  trigger.Info,
	}

	systemPrompt, messages := prompt.Proactive(ctx, profile, m.history, m.sessionMemory)

	return m.startStream(systemPrompt, messages)
}

func (m *Model) handleUserInput(msg SubmitMsg) tea.Cmd {
	m.state = "streaming"
	_ = m.input.SetEnabled(false)

	m.output.AppendUserMessage(msg.Text)
	m.history = append(m.history, llm.Message{Role: "user", Content: msg.Text})

	// Resolve @files
	if len(msg.Files) > 0 {
		var fileNames []string
		for path := range msg.Files {
			fileNames = append(fileNames, path)
		}
		m.output.AppendContext(strings.Join(fileNames, ", "))
	}

	m.output.StartStream(m.width)

	profile := prompt.UserProfile{
		Name:      m.config.Name,
		Languages: m.config.Languages,
		Level:     m.config.Level,
	}

	rctx, _ := repocontext.Build(repocontext.CollectOpts{
		Files:       msg.Files,
		TriggerType: "user",
	})
	if rctx == nil {
		rctx = &repocontext.RepoContext{Files: msg.Files}
	}

	systemPrompt, messages := prompt.Reactive(rctx, profile, m.history)

	return m.startStream(systemPrompt, messages)
}

func (m *Model) startStream(systemPrompt string, messages []llm.Message) tea.Cmd {
	client := m.llmClient
	opts := llm.StreamOptions{
		SystemPrompt: systemPrompt,
		MaxTokens:    4096,
		Temperature:  0.7,
	}

	ctx, cancel := context.WithCancel(context.Background())
	m.cancelFn = cancel

	ch := make(chan string, 64)
	m.chunkCh = ch

	streamCmd := func() tea.Msg {
		resp, err := llm.StreamWithRetry(ctx, client, messages, opts, func(chunk string) {
			ch <- chunk
		})
		close(ch)
		if err != nil {
			return StreamErrorMsg{Err: err}
		}
		return StreamDoneMsg{Response: resp}
	}

	return tea.Batch(streamCmd, waitForChunk(ch))
}
