package tui

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/learning"
	"github.com/MitchTheStonky/pear/cli/llm"
	"github.com/MitchTheStonky/pear/cli/repocontext"
	"github.com/MitchTheStonky/pear/cli/watcher"
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
type listenTickMsg struct{}
type conceptPickerDismissMsg struct{ gen int }

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
	cancelFn     context.CancelFunc
	chunkCh      <-chan string
	conceptStore  *learning.ConceptStore
	learningPath  string
	sessionMemory  *learning.SessionMemory
	newConcepts      []string // concepts available in the picker
	conceptIdx       int      // currently highlighted picker index
	conceptPickerGen int      // generation counter to ignore stale timeouts
	lastReviewCtx    *repocontext.RepoContext
	settings     settingsState
	watcher      *watcher.Watcher
	watchCancel  context.CancelFunc
	listenDots        int // 0-3, cycles for "Pear is listening" animation
	consecutiveErrors int // tracks consecutive LLM errors for graceful degradation
}

// NewModel creates a new TUI model.
func NewModel(cfg *config.Config, client llm.LLMClient, mode string, triggers <-chan ReviewTrigger) Model {
	lpath := filepath.Join(config.Dir(), "learning.json")
	store, _ := learning.Load(lpath)
	output := NewOutputModel(80, 20)
	bannerFn := func(w int) string { return WelcomeBanner(cfg, w) }
	output.bannerFunc = bannerFn
	output.bannerOnly = true
	output.content.WriteString(bannerFn(80))
	output.refreshViewport()
	return Model{
		input:        NewInputModel(),
		output:       output,
		mode:         mode,
		state:        "idle",
		config:       cfg,
		llmClient:    client,
		triggers:     triggers,
		stats:        SessionStats{StartTime: time.Now()},
		conceptStore:  store,
		learningPath:  lpath,
		sessionMemory: &learning.SessionMemory{},
		width:        80,
		height:       24,
	}
}

// Init initializes the model.
func (m Model) Init() tea.Cmd {
	var cmds []tea.Cmd
	cmds = append(cmds, m.input.textinput.Focus())
	if m.triggers != nil {
		cmds = append(cmds, waitForTrigger(m.triggers))
	}
	if m.mode == "watch" {
		cmds = append(cmds, listenTick())
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
		m.syncViewportHeight()
		return m, nil

	case tea.KeyMsg:
		// Concept picker handles its own keys when active
		if m.state == "concept_pick" {
			switch msg.String() {
			case "ctrl+c":
				if m.watchCancel != nil {
					m.watchCancel()
				}
				return m, tea.Quit
			case "up":
				if m.conceptIdx > 0 {
					m.conceptIdx--
					m.output.ShowConceptPicker(m.newConcepts, m.conceptIdx)
				}
				return m, nil
			case "down":
				if m.conceptIdx < len(m.newConcepts)-1 {
					m.conceptIdx++
					m.output.ShowConceptPicker(m.newConcepts, m.conceptIdx)
				}
				return m, nil
			case "enter":
				concept := m.newConcepts[m.conceptIdx]
				_ = m.dismissConceptPicker() // deep dive overrides the idle state anyway
				return m, m.handleDeepDive(concept)
			case "esc":
				return m, m.dismissConceptPicker()
			}
			return m, nil
		}

		switch msg.String() {
		case "ctrl+c":
			if m.watchCancel != nil {
				m.watchCancel()
			}
			return m, tea.Quit
		case "esc":
			if m.settings.active {
				m.settings = settingsState{}
				m.output.AppendSystem("Settings closed.")
				return m, nil
			}
		case "ctrl+p":
			if m.mode == "watch" {
				m.paused = !m.paused
				if m.paused {
					m.output.AppendSystem("⏸ Proactive reviews paused. Ctrl+P to resume.")
				} else {
					m.output.AppendSystem("▶ Proactive reviews resumed.")
					return m, listenTick()
				}
				return m, nil
			}
		}

	case listenTickMsg:
		if m.mode == "watch" && m.state == "idle" && !m.paused {
			m.listenDots = (m.listenDots + 1) % 4
			return m, listenTick()
		}
		return m, nil

	case conceptPickerDismissMsg:
		if m.state == "concept_pick" && msg.gen == m.conceptPickerGen {
			return m, m.dismissConceptPicker()
		}
		return m, nil

	case ReviewTriggerMsg:
		if m.state == "concept_pick" {
			_ = m.dismissConceptPicker() // trigger handling below takes over
		}
		if m.paused {
			return m, waitForTrigger(m.triggers)
		}
		if m.mode == "watch" && m.consecutiveErrors >= 3 {
			// Silently drop triggers during sustained LLM outage
			return m, waitForTrigger(m.triggers)
		}
		if m.state == "streaming" {
			t := msg.Trigger
			m.queuedTrig = &t
			return m, nil
		}
		return m, m.handleTrigger(msg.Trigger)

	case SubmitMsg:
		m.consecutiveErrors = 0
		if m.state == "streaming" {
			return m, nil
		}
		if m.state == "concept_pick" {
			_ = m.dismissConceptPicker()
		}
		if m.settings.active {
			m.handleSettingsInput(msg.Text)
			return m, nil
		}
		return m, m.handleUserInput(msg)

	case SlashMsg:
		return m, m.handleSlash(msg)

	case ChunkMsg:
		m.output.AppendChunk(msg.Text)
		return m, waitForChunk(m.chunkCh)

	case StreamDoneMsg:
		m.output.EndStream(m.width)
		m.state = "idle"
		m.consecutiveErrors = 0
		if cmd := m.input.SetEnabled(true); cmd != nil {
			cmds = append(cmds, cmd)
		}
		m.stats.Reviews++

		if msg.Response != nil {
			m.history = append(m.history, llm.Message{Role: "assistant", Content: msg.Response.Content})

			const maxHistory = 50
			if len(m.history) > maxHistory {
				m.history = m.history[len(m.history)-maxHistory:]
			}

			if m.conceptStore != nil {
				concepts, relationships, covered := learning.Extract(msg.Response.Content)
				if len(concepts) > 0 {
					// Check for new concepts before recording (so counts are still 0)
					var newConcepts []string
					for _, c := range concepts {
						if existing, ok := m.conceptStore.Concepts[c]; !ok || existing.Count == 0 {
							newConcepts = append(newConcepts, c)
						}
					}

					m.output.AppendConcepts(concepts)
					m.output.AppendRelationships(relationships)
					m.conceptStore.Record(concepts, relationships)
					m.stats.Concepts += len(concepts)
					_ = m.conceptStore.Save(m.learningPath)

					// Show concept picker if there are new concepts (non-blocking: input stays enabled)
					if len(newConcepts) > 0 {
						m.conceptPickerGen++
						m.newConcepts = newConcepts
						m.conceptIdx = 0
						m.state = "concept_pick"
						m.output.ShowConceptPicker(newConcepts, 0)
						cmds = append(cmds, conceptPickerTimeout(m.conceptPickerGen))
					}
				}
				for _, entry := range covered {
					m.sessionMemory.AddCovered(entry.Concept, entry.Summary)
				}
			}
		}

		if m.queuedTrig != nil {
			t := *m.queuedTrig
			m.queuedTrig = nil
			cmds = append(cmds, m.handleTrigger(t))
		}
		if m.triggers != nil {
			cmds = append(cmds, waitForTrigger(m.triggers))
		}
		if m.mode == "watch" {
			cmds = append(cmds, listenTick())
		}
		return m, tea.Batch(cmds...)

	case StreamErrorMsg:
		m.output.EndStream(m.width)
		m.consecutiveErrors++
		if m.consecutiveErrors >= 3 {
			m.output.AppendError("⚠ LLM unavailable — Pear will retry when you make changes. Type /status to check.")
		} else {
			m.output.AppendError(msg.Err.Error())
		}
		m.state = "idle"
		if cmd := m.input.SetEnabled(true); cmd != nil {
			cmds = append(cmds, cmd)
		}
		if m.triggers != nil {
			cmds = append(cmds, waitForTrigger(m.triggers))
		}
		if m.mode == "watch" {
			cmds = append(cmds, listenTick())
		}
		return m, tea.Batch(cmds...)
	}

	// Forward to sub-models
	var cmd tea.Cmd
	prevAC := m.input.autocomplete.active
	m.input, cmd = m.input.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	// Recalculate viewport height if autocomplete toggled (changes chrome height)
	if m.input.autocomplete.active != prevAC {
		m.syncViewportHeight()
	}

	// Only forward scroll-relevant messages to the viewport to prevent
	// typing and other keys from causing viewport jumps.
	if isScrollMsg(msg) && !m.input.AutocompleteActive() {
		m.output, cmd = m.output.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

// renderHeader returns the rendered header line.
func (m Model) renderHeader() string {
	provider := config.ActiveProvider(m.config)
	header := HeaderStyle.Render(fmt.Sprintf("🍐 Pear v0 · %s · %s/%s", m.mode, m.config.Provider.Active, provider.Model))
	if m.paused {
		header += " · paused"
	}
	return header
}

// renderBottom returns the rendered input box + status line below the viewport.
func (m Model) renderBottom() string {
	var b strings.Builder

	// Bordered input box
	inputBox := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(colorDim).
		Width(m.width - 2).
		Render(m.input.View())
	b.WriteString(inputBox)

	// Hints bar
	dim := lipgloss.NewStyle().Foreground(colorDim)
	key := lipgloss.NewStyle().Foreground(colorDim).Bold(true)
	hints := dim.Render(" ") +
		key.Render("Enter") + dim.Render(" to send • ") +
		key.Render("@") + dim.Render(" mention data • ") +
		key.Render("/copy") + dim.Render(" latest response • ") +
		key.Render("/") + dim.Render(" commands list • ") +
		key.Render("↑↓") + dim.Render(" scroll")
	b.WriteString("\n")
	b.WriteString(hints)

	// Status line below input
	var status string
	if m.mode == "watch" && m.state == "idle" && !m.paused {
		// Pulse between bright and dim dot
		dot := "●"
		if m.listenDots%2 == 1 {
			dot = "·"
		}
		dots := strings.Repeat(".", m.listenDots)
		pad := strings.Repeat(" ", 3-m.listenDots)
		greenDot := lipgloss.NewStyle().Foreground(colorGreen).Render(dot)
		status = " " + greenDot + lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Render(
			fmt.Sprintf(" Pear is watching%s%s", dots, pad))
	} else if m.state == "concept_pick" {
		status = ConceptPickerStyle.Render(" ↕ Pick a concept to explore")
	} else if m.state == "streaming" {
		status = ThinkingStyle.Render(" ✦ Pear is thinking...")
	} else if m.paused {
		status = ThinkingStyle.Render(" ⏸ Paused")
	}
	if status != "" {
		b.WriteString("\n\n")
		b.WriteString(status)
	}

	return b.String()
}

// syncViewportHeight recalculates the viewport height from the rendered chrome.
func (m *Model) syncViewportHeight() {
	bottom := m.renderBottom()
	bottomLines := strings.Count(bottom, "\n") + 1
	vpHeight := m.height - 1 - bottomLines // 1 = header line
	if vpHeight < 1 {
		vpHeight = 1
	}
	m.output.SetSize(m.width, vpHeight)
}

// View renders the full layout.
func (m Model) View() string {
	var b strings.Builder
	b.WriteString(m.renderHeader())
	b.WriteString("\n")
	b.WriteString(m.output.View())
	b.WriteString("\n")
	b.WriteString(m.renderBottom())
	return b.String()
}

func isScrollMsg(msg tea.Msg) bool {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "down", "pgup", "pgdown", "home", "end":
			return true
		}
	case tea.MouseMsg:
		return true
	}
	return false
}

func listenTick() tea.Cmd {
	return tea.Tick(800*time.Millisecond, func(time.Time) tea.Msg {
		return listenTickMsg{}
	})
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

func waitForChunk(ch <-chan string) tea.Cmd {
	return func() tea.Msg {
		chunk, ok := <-ch
		if !ok {
			return nil
		}
		return ChunkMsg{Text: chunk}
	}
}

func conceptPickerTimeout(gen int) tea.Cmd {
	return tea.Tick(10*time.Second, func(time.Time) tea.Msg {
		return conceptPickerDismissMsg{gen: gen}
	})
}

func (m *Model) dismissConceptPicker() tea.Cmd {
	m.output.RemoveConceptPicker()
	m.newConcepts = nil
	m.conceptIdx = 0
	m.state = "idle"
	if m.mode == "watch" {
		return listenTick()
	}
	return nil
}
