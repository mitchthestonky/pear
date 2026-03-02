package tui

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/learning"
	"github.com/pearcode/pear/llm"
	"github.com/pearcode/pear/prompt"
	"github.com/pearcode/pear/repocontext"
	"github.com/pearcode/pear/watcher"
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

// SessionStats tracks session metrics.
type SessionStats struct {
	StartTime time.Time
	Reviews   int
	Concepts  int
}

// settingsState tracks the /settings numbered editor flow.
type settingsState struct {
	active       bool
	awaitingEdit int  // which field number (1-9), 0 = showing menu
	providerPick bool // awaiting provider choice (1-3)
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
	conceptStore *learning.ConceptStore
	learningPath string
	settings     settingsState
	watcher      *watcher.Watcher
	watchCancel  context.CancelFunc
	listenDots   int // 0-3, cycles for "Pear is listening" animation
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
		conceptStore: store,
		learningPath: lpath,
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
				concepts, relationships := learning.Extract(msg.Response.Content)
				if len(concepts) > 0 {
					m.output.AppendConcepts(concepts)
					m.output.AppendRelationships(relationships)
					m.conceptStore.Record(concepts, relationships)
					m.stats.Concepts += len(concepts)
					_ = m.conceptStore.Save(m.learningPath)
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
		m.output.AppendError(msg.Err.Error())
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
		key.Render("/") + dim.Render(" commands • ") +
		key.Render("/copy") + dim.Render(" latest response")
	b.WriteString("\n")
	b.WriteString(hints)

	// Status line below input
	var status string
	if m.mode == "watch" && m.state == "idle" && !m.paused {
		dots := strings.Repeat(".", m.listenDots)
		pad := strings.Repeat(" ", 3-m.listenDots)
		status = lipgloss.NewStyle().Foreground(colorGreen).Bold(true).Render(
			fmt.Sprintf(" Pear is watching%s%s", dots, pad))
	} else if m.state == "streaming" {
		status = ThinkingStyle.Render(" Pear is thinking...")
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
// Measures a placeholder View to get the exact chrome lines, avoiding magic numbers.
func (m *Model) syncViewportHeight() {
	bottom := m.renderBottom()
	bottomLines := strings.Count(bottom, "\n") + 1
	vpHeight := m.height - 1 - bottomLines // 1 = header line
	if vpHeight < 1 {
		vpHeight = 1
	}
	m.output.SetSize(m.width, vpHeight)
}

// View renders the full layout. All state mutations happen in Update;
// this is a pure render function.
func (m Model) View() string {
	var b strings.Builder
	b.WriteString(m.renderHeader())
	b.WriteString("\n")
	b.WriteString(m.output.View())
	b.WriteString("\n")
	b.WriteString(m.renderBottom())
	return b.String()
}

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

	systemPrompt, messages := prompt.Proactive(ctx, profile, m.history)

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

func (m *Model) handleSlash(msg SlashMsg) tea.Cmd {
	switch msg.Command {
	case "help":
		help := `Available commands:
  /help      — Show this help
  /clear     — Clear chat history
  /exit      — End session
  /watch     — Start file watcher from interactive mode
  /review    — One-shot review of current git diff
  /pause     — Pause proactive reviews (watch mode)
  /resume    — Resume proactive reviews (watch mode)
  /status    — Show session status
  /settings  — Edit configuration
  /provider  — Change LLM provider
  /model <n> — Change model
  /copy      — Copy last response to clipboard
  /export    — Export chat history to file
  /key       — Update API key`
		m.output.AppendSystem(help)

	case "clear":
		m.history = nil
		m.output.Clear()
		m.output.AppendSystem("🍐 History cleared.")

	case "exit", "quit", "q":
		return tea.Quit

	case "pause":
		if m.mode != "watch" {
			m.output.AppendSystem("/pause is only available in watch mode.")
			return nil
		}
		m.paused = true
		m.output.AppendSystem("🍐 Proactive reviews paused. Type /resume to restart.")

	case "resume":
		if m.mode != "watch" {
			m.output.AppendSystem("/resume is only available in watch mode.")
			return nil
		}
		m.paused = false
		m.output.AppendSystem("🍐 Proactive reviews resumed.")
		return listenTick()

	case "status":
		uptime := time.Since(m.stats.StartTime).Truncate(time.Second)
		p := config.ActiveProvider(m.config)
		status := fmt.Sprintf(`🍐 Session status:
  Uptime:    %s
  Reviews:   %d
  Concepts:  %d
  Provider:  %s
  Model:     %s`, uptime, m.stats.Reviews, m.stats.Concepts, m.config.Provider.Active, p.Model)
		m.output.AppendSystem(status)

	case "settings":
		m.showSettingsMenu()
		m.settings = settingsState{active: true}

	case "provider":
		m.showProviderPicker()
		m.settings = settingsState{active: true, awaitingEdit: 4, providerPick: true}

	case "model":
		if msg.Args == "" {
			m.output.AppendSystem("Usage: /model <model-name>")
			return nil
		}
		config.SetModel(m.config, msg.Args)
		if err := config.Save(m.config); err != nil {
			m.output.AppendSystem(fmt.Sprintf("⚠ Error saving config: %s", err))
			return nil
		}
		m.reinitLLM()
		m.output.AppendSystem(fmt.Sprintf("✓ Model changed to %s (%s)", msg.Args, m.config.Provider.Active))

	case "key":
		m.output.AppendSystem(fmt.Sprintf("New API key for %s?", m.config.Provider.Active))
		m.settings = settingsState{active: true, awaitingEdit: 6}

	case "watch":
		// Clean up existing watcher if it exists
		if m.watchCancel != nil {
			m.watchCancel()
			m.watchCancel = nil
			m.watcher = nil
		}
		repoRoot, err := repocontext.RepoRoot()
		if err != nil {
			m.output.AppendSystem("⚠ Not a git repository. Cannot start watcher.")
			return nil
		}
		w, err := watcher.New(m.config.Watch, repoRoot, nil)
		if err != nil {
			m.output.AppendSystem(fmt.Sprintf("⚠ Error starting watcher: %s", err))
			return nil
		}
		ctx, cancel := context.WithCancel(context.Background())
		m.watcher = w
		m.watchCancel = cancel
		watcherTriggers := w.Start(ctx)
		tuiTriggers := make(chan ReviewTrigger, 1)
		go func() {
			for wt := range watcherTriggers {
				tuiTriggers <- ReviewTrigger{
					Diff:        wt.Diff,
					TriggerType: wt.Type,
					Info:        formatWatchTriggerInfo(wt),
				}
			}
			close(tuiTriggers)
		}()
		m.triggers = tuiTriggers
		m.mode = "watch"
		m.output.AppendSystem("🍐 File watcher started. Pear will review your changes automatically.")
		return waitForTrigger(m.triggers)

	case "review":
		if m.state == "streaming" {
			return nil
		}
		repoRoot, _ := repocontext.RepoRoot()
		diff, err := repocontext.GitDiff(repoRoot)
		if err != nil || strings.TrimSpace(diff) == "" {
			m.output.AppendSystem("No changes to review.")
			return nil
		}
		lines := strings.Count(diff, "\n")
		trigger := ReviewTrigger{
			Diff:        diff,
			TriggerType: "settle",
			Info:        fmt.Sprintf("made changes (%d lines)", lines),
		}
		return m.handleTrigger(trigger)

	case "copy":
		// Find the last assistant message
		var last string
		for i := len(m.history) - 1; i >= 0; i-- {
			if m.history[i].Role == "assistant" {
				last = m.history[i].Content
				break
			}
		}
		if last == "" {
			m.output.AppendSystem("Nothing to copy yet.")
			return nil
		}
		clipCmd := clipboardCmd()
		if clipCmd == nil {
			m.output.AppendSystem("⚠ No clipboard tool found (need pbcopy, xclip, or wl-copy).")
			return nil
		}
		clipCmd.Stdin = strings.NewReader(last)
		if err := clipCmd.Run(); err != nil {
			m.output.AppendSystem(fmt.Sprintf("⚠ Copy failed: %s", err))
			return nil
		}
		m.output.AppendSystem("✓ Last response copied to clipboard.")

	case "export":
		filename := fmt.Sprintf("pear-session-%s.md", time.Now().Format("2006-01-02-150405"))
		var b strings.Builder
		b.WriteString("# Pear Session Export\n\n")
		for _, msg := range m.history {
			if msg.Role == "user" {
				b.WriteString("## You\n\n")
			} else {
				b.WriteString("## Pear\n\n")
			}
			b.WriteString(msg.Content)
			b.WriteString("\n\n---\n\n")
		}
		if err := os.WriteFile(filename, []byte(b.String()), 0644); err != nil {
			m.output.AppendSystem(fmt.Sprintf("⚠ Export failed: %s", err))
			return nil
		}
		m.output.AppendSystem(fmt.Sprintf("✓ Chat exported to %s", filename))

	default:
		m.output.AppendSystem("Unknown command. Type /help for available commands.")
	}
	return nil
}

func clipboardCmd() *exec.Cmd {
	switch runtime.GOOS {
	case "darwin":
		return exec.Command("pbcopy")
	case "linux":
		if _, err := exec.LookPath("wl-copy"); err == nil {
			return exec.Command("wl-copy")
		}
		if _, err := exec.LookPath("xclip"); err == nil {
			return exec.Command("xclip", "-selection", "clipboard")
		}
		return nil
	default:
		return nil
	}
}

func (m *Model) showSettingsMenu() {
	p := config.ActiveProvider(m.config)
	maskedKey := maskKey(p.APIKey)
	s := fmt.Sprintf(`🍐 Current configuration:

  1. Name:       %s
  2. Languages:  %s
  3. Level:      %s
  4. Provider:   %s
  5. Model:      %s
  6. API Key:    %s

  Watch settings:
  7. Settle time:    %ds
  8. Min diff lines: %d
  9. Cooldown:       %ds

  Enter a number to edit, or press Esc to close.`,
		m.config.Name, m.config.Languages, m.config.Level,
		m.config.Provider.Active, p.Model, maskedKey,
		m.config.Watch.SettleTime, m.config.Watch.MinDiffLines, m.config.Watch.Cooldown)
	m.output.AppendSystem(s)
}

func (m *Model) showProviderPicker() {
	current := m.config.Provider.Active
	lines := fmt.Sprintf(`  Choose your LLM provider:
    1. Anthropic (Claude)%s
    2. OpenAI%s
    3. OpenRouter%s`,
		currentMarker("anthropic", current),
		currentMarker("openai", current),
		currentMarker("openrouter", current))
	m.output.AppendSystem(lines)
}

func currentMarker(name, current string) string {
	if name == current {
		return " ← current"
	}
	return ""
}

func maskKey(key string) string {
	if len(key) <= 8 {
		return "****"
	}
	return key[:6] + "...****"
}

func (m *Model) handleSettingsInput(text string) {
	text = strings.TrimSpace(text)

	// If awaiting provider pick
	if m.settings.providerPick {
		providers := map[string]string{"1": "anthropic", "2": "openai", "3": "openrouter"}
		prov, ok := providers[text]
		if !ok {
			m.output.AppendSystem("Invalid choice. Enter 1, 2, or 3.")
			return
		}
		m.config.Provider.Active = prov
		p := config.ActiveProvider(m.config)
		if p.APIKey == "" {
			m.output.AppendSystem(fmt.Sprintf("%s API key?", prov))
			m.settings = settingsState{active: true, awaitingEdit: 6}
			return
		}
		if err := config.Save(m.config); err != nil {
			m.output.AppendSystem(fmt.Sprintf("⚠ Error saving config: %s", err))
		}
		m.reinitLLM()
		m.output.AppendSystem(fmt.Sprintf("✓ Switched to %s (%s)", prov, p.Model))
		m.settings = settingsState{}
		return
	}

	// If awaiting a field edit value
	if m.settings.awaitingEdit > 0 {
		m.applySettingsEdit(m.settings.awaitingEdit, text)
		return
	}

	// Otherwise, parse number selection from settings menu
	num, err := strconv.Atoi(text)
	if err != nil || num < 1 || num > 9 {
		m.output.AppendSystem("Enter a number 1-9, or press Esc to close.")
		return
	}

	switch num {
	case 4:
		m.showProviderPicker()
		m.settings = settingsState{active: true, awaitingEdit: 4, providerPick: true}
	case 6:
		m.output.AppendSystem(fmt.Sprintf("New API key for %s?", m.config.Provider.Active))
		m.settings = settingsState{active: true, awaitingEdit: 6}
	default:
		labels := map[int]string{
			1: "Name", 2: "Languages", 3: "Level", 5: "Model",
			7: "Settle time (seconds)", 8: "Min diff lines", 9: "Cooldown (seconds)",
		}
		m.output.AppendSystem(fmt.Sprintf("New value for %s?", labels[num]))
		m.settings = settingsState{active: true, awaitingEdit: num}
	}
}

func (m *Model) applySettingsEdit(field int, value string) {
	switch field {
	case 1:
		m.config.Name = value
	case 2:
		m.config.Languages = value
	case 3:
		m.config.Level = value
	case 5:
		config.SetModel(m.config, value)
	case 6:
		config.SetKey(m.config, value)
	case 7:
		n, err := strconv.Atoi(value)
		if err != nil || n <= 0 {
			m.output.AppendSystem("Invalid number.")
			m.settings = settingsState{}
			return
		}
		m.config.Watch.SettleTime = n
	case 8:
		n, err := strconv.Atoi(value)
		if err != nil || n <= 0 {
			m.output.AppendSystem("Invalid number.")
			m.settings = settingsState{}
			return
		}
		m.config.Watch.MinDiffLines = n
	case 9:
		n, err := strconv.Atoi(value)
		if err != nil || n <= 0 {
			m.output.AppendSystem("Invalid number.")
			m.settings = settingsState{}
			return
		}
		m.config.Watch.Cooldown = n
	}

	if err := config.Save(m.config); err != nil {
		m.output.AppendSystem(fmt.Sprintf("⚠ Error saving config: %s", err))
		m.settings = settingsState{}
		return
	}

	// Reinit LLM if provider/model/key changed
	if field == 5 || field == 6 {
		m.reinitLLM()
	}

	m.output.AppendSystem("✓ Setting updated.")
	m.settings = settingsState{}
}

func (m *Model) reinitLLM() {
	p := config.ActiveProvider(m.config)
	client, err := llm.NewClient(m.config.Provider.Active, llm.ProviderDetail{
		APIKey: p.APIKey,
		Model:  p.Model,
	})
	if err != nil {
		m.output.AppendSystem(fmt.Sprintf("⚠ Error reinitializing LLM: %s", err))
		return
	}
	m.llmClient = client
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

func waitForChunk(ch <-chan string) tea.Cmd {
	return func() tea.Msg {
		chunk, ok := <-ch
		if !ok {
			return nil
		}
		return ChunkMsg{Text: chunk}
	}
}

func formatWatchTriggerInfo(wt watcher.ReviewTrigger) string {
	switch wt.Type {
	case "commit":
		return fmt.Sprintf("committed: %s", wt.Summary)
	case "settle":
		return fmt.Sprintf("made changes (%s)", wt.Summary)
	default:
		return wt.Summary
	}
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
	return tea.Tick(500*time.Millisecond, func(time.Time) tea.Msg {
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
