package tui

import (
	"context"
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/llm"
	"github.com/pearcode/pear/prompt"
	"github.com/pearcode/pear/repocontext"
)

// ReviewTrigger represents a watcher-generated review trigger.
type ReviewTrigger struct {
	Diff        string
	TriggerType string // "settle", "commit"
	Info        string
}

// Messages
type ChunkMsg struct{ Text string }
type StreamDoneMsg struct{ Response *llm.Response }
type StreamErrorMsg struct{ Err error }
type ReviewTriggerMsg struct{ Trigger ReviewTrigger }

// SessionStats tracks session metrics.
type SessionStats struct {
	StartTime time.Time
	Reviews   int
	Concepts  int
}

// Model is the main Bubble Tea model.
type Model struct {
	input      InputModel
	output     OutputModel
	mode       string // "interactive" or "watch"
	state      string // "idle", "streaming"
	paused     bool
	history    []llm.Message
	stats      SessionStats
	config     *config.Config
	llmClient  llm.LLMClient
	triggers   <-chan ReviewTrigger
	queuedTrig *ReviewTrigger
	width      int
	height     int
	cancelFn   context.CancelFunc
}

// NewModel creates a new TUI model.
func NewModel(cfg *config.Config, client llm.LLMClient, mode string, triggers <-chan ReviewTrigger) Model {
	return Model{
		input:     NewInputModel(),
		output:    NewOutputModel(80, 20),
		mode:      mode,
		state:     "idle",
		config:    cfg,
		llmClient: client,
		triggers:  triggers,
		stats:     SessionStats{StartTime: time.Now()},
		width:     80,
		height:    24,
	}
}

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.input.textinput.Focus())
	if m.triggers != nil {
		cmds = append(cmds, waitForTrigger(m.triggers))
	}
	return tea.Batch(cmds...)
}

// Update handles all messages.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		outputHeight := m.height - 3 // reserve space for input
		m.output.SetSize(m.width, outputHeight)
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+p":
			if m.mode == "watch" {
				m.paused = !m.paused
				if m.paused {
					m.output.AppendSystem("⏸ Proactive reviews paused. Ctrl+P to resume.")
				} else {
					m.output.AppendSystem("▶ Proactive reviews resumed.")
				}
				return m, nil
			}
		}

	case ReviewTriggerMsg:
		if m.paused {
			return m, waitForTrigger(m.triggers)
		}
		if m.state == "streaming" {
			t := msg.Trigger
			m.queuedTrig = &t
			return m, nil
		}
		return m, m.handleTrigger(msg.Trigger)

	case SubmitMsg:
		if m.state == "streaming" {
			return m, nil
		}
		return m, m.handleUserInput(msg)

	case SlashMsg:
		return m, m.handleSlash(msg)

	case ChunkMsg:
		m.output.AppendChunk(msg.Text)
		return m, nil

	case StreamDoneMsg:
		m.output.EndStream(m.width)
		m.state = "idle"
		m.input.SetEnabled(true)
		m.stats.Reviews++

		var cmds []tea.Cmd
		if m.queuedTrig != nil {
			t := *m.queuedTrig
			m.queuedTrig = nil
			cmds = append(cmds, m.handleTrigger(t))
		}
		if m.triggers != nil {
			cmds = append(cmds, waitForTrigger(m.triggers))
		}
		return m, tea.Batch(cmds...)

	case StreamErrorMsg:
		m.output.AppendError(msg.Err.Error())
		m.state = "idle"
		m.input.SetEnabled(true)
		if m.triggers != nil {
			return m, waitForTrigger(m.triggers)
		}
		return m, nil
	}

	// Forward to sub-models
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}
	m.output, cmd = m.output.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

// View renders the full layout.
func (m Model) View() string {
	var b strings.Builder

	// Header
	header := HeaderStyle.Render(fmt.Sprintf("🍐 Pear v0 · %s", m.mode))
	if m.paused {
		header += " · paused"
	}
	b.WriteString(header)
	b.WriteString("\n")

	// Output
	b.WriteString(m.output.View())
	b.WriteString("\n")

	// Input
	b.WriteString(m.input.View())

	return b.String()
}

func (m *Model) handleTrigger(trigger ReviewTrigger) tea.Cmd {
	m.state = "streaming"
	m.input.SetEnabled(false)

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

	systemPrompt, messages := prompt.Proactive(ctx, profile, m.history)

	return m.startStream(systemPrompt, messages)
}

func (m *Model) handleUserInput(msg SubmitMsg) tea.Cmd {
	m.state = "streaming"
	m.input.SetEnabled(false)

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

func (m *Model) handleSlash(msg SlashMsg) tea.Cmd {
	switch msg.Command {
	case "help":
		m.output.AppendSystem("Commands: /pause, /resume, /settings, /quit")
	case "pause":
		m.paused = true
		m.output.AppendSystem("⏸ Proactive reviews paused.")
	case "resume":
		m.paused = false
		m.output.AppendSystem("▶ Proactive reviews resumed.")
	case "quit", "q":
		return tea.Quit
	default:
		m.output.AppendSystem(fmt.Sprintf("Unknown command: /%s", msg.Command))
	}
	return nil
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

	return func() tea.Msg {
		resp, err := client.Stream(ctx, messages, opts, func(chunk string) {
			// We need to send chunks via the program; for now we accumulate
		})
		_ = cancel
		if err != nil {
			return StreamErrorMsg{Err: err}
		}
		return StreamDoneMsg{Response: resp}
	}
}

func waitForTrigger(ch <-chan ReviewTrigger) tea.Cmd {
	return func() tea.Msg {
		trigger, ok := <-ch
		if !ok {
			return nil
		}
		return ReviewTriggerMsg{Trigger: trigger}
	}
}
