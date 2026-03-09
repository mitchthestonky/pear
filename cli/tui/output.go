package tui

import (
	"strings"

	"github.com/MitchTheStonky/pear/cli/learning"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/glamour/ansi"
	"github.com/charmbracelet/glamour/styles"
	"github.com/charmbracelet/lipgloss"
)

// OutputModel is a scrollable viewport that renders streaming markdown.
type OutputModel struct {
	viewport      viewport.Model
	content       *strings.Builder // pre-styled UI content (banner, headers, separators)
	stream        *strings.Builder // raw LLM markdown accumulated during streaming
	streaming     bool             // whether we're currently in a stream block
	renderer      *glamour.TermRenderer
	autoScroll    bool
	width         int
	thinkingShown bool   // whether "Thinking..." is currently displayed
	bannerOnly    bool   // true until first non-banner content is added
	bannerFunc    func(int) string // generates banner at given width
}

// NewOutputModel creates a new output component.
func NewOutputModel(width, height int) OutputModel {
	vp := viewport.New(width, height)
	vp.SetContent("")

	r, _ := glamour.NewTermRenderer(
		glamour.WithStyles(glamourStyle()),
		glamour.WithWordWrap(width-4),
	)

	return OutputModel{
		viewport:   vp,
		content:    &strings.Builder{},
		stream:     &strings.Builder{},
		renderer:   r,
		autoScroll: true,
		width:      width,
	}
}

// SetSize updates the viewport dimensions.
func (m *OutputModel) SetSize(width, height int) {
	m.width = width
	m.viewport.Width = width
	m.viewport.Height = height
	if m.renderer != nil {
		m.renderer, _ = glamour.NewTermRenderer(
			glamour.WithStyles(glamourStyle()),
			glamour.WithWordWrap(width-4),
		)
	}
	// Re-render banner at new width if nothing else has been added yet
	if m.bannerOnly && m.bannerFunc != nil {
		m.content.Reset()
		m.content.WriteString(m.bannerFunc(width))
		m.refreshViewport()
	}
}

// AppendHeader adds a styled header line.
func (m *OutputModel) AppendHeader(text string) {
	m.bannerOnly = false
	m.content.WriteString("\n")
	m.content.WriteString(TriggerStyle.Render(text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendUserMessage displays the user's input in the log.
func (m *OutputModel) AppendUserMessage(text string) {
	m.bannerOnly = false
	m.content.WriteString(UserMessageStyle.Render("❯ " + text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// StartStream begins a new response block.
func (m *OutputModel) StartStream(width int) {
	m.bannerOnly = false
	m.content.WriteString("\n")
	m.content.WriteString(ThinkingStyle.Render("Thinking..."))
	m.content.WriteString("\n")
	m.thinkingShown = true
	m.stream.Reset()
	m.streaming = true
	m.refreshViewport()
}

// AppendChunk adds a streaming chunk and re-renders.
func (m *OutputModel) AppendChunk(text string) {
	if m.thinkingShown {
		// Remove the "Thinking..." line from content
		s := m.content.String()
		thinkingLine := ThinkingStyle.Render("Thinking...") + "\n"
		if idx := strings.LastIndex(s, thinkingLine); idx >= 0 {
			m.content.Reset()
			m.content.WriteString(s[:idx] + s[idx+len(thinkingLine):])
		}
		m.thinkingShown = false
	}
	m.stream.WriteString(text)
	m.refreshViewport()
}

// EndStream finalizes the stream block.
func (m *OutputModel) EndStream(width int) {
	// Remove "Thinking..." if it was never cleared by a chunk
	if m.thinkingShown {
		s := m.content.String()
		thinkingLine := ThinkingStyle.Render("Thinking...") + "\n"
		if idx := strings.LastIndex(s, thinkingLine); idx >= 0 {
			m.content.Reset()
			m.content.WriteString(s[:idx] + s[idx+len(thinkingLine):])
		}
		m.thinkingShown = false
	}

	// Strip concept/related tag lines from stream before rendering
	streamText := learning.StripTags(m.stream.String())

	// Render final markdown and bake it into content
	if m.renderer != nil {
		if r, err := m.renderer.Render(streamText); err == nil {
			m.content.WriteString(r)
		} else {
			m.content.WriteString(streamText)
		}
	} else {
		m.content.WriteString(streamText)
	}
	m.stream.Reset()
	m.streaming = false
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendConcepts adds concept tags after a response.
func (m *OutputModel) AppendConcepts(concepts []string) {
	if len(concepts) == 0 {
		return
	}
	m.bannerOnly = false
	m.content.WriteString(ConceptStyle.Render("📚 Concepts: [" + strings.Join(concepts, ", ") + "]"))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendRelationships adds relationship tags after a response.
func (m *OutputModel) AppendRelationships(relationships map[string][]string) {
	if len(relationships) == 0 {
		return
	}
	m.bannerOnly = false
	var pairs []string
	for from, tos := range relationships {
		for _, to := range tos {
			pairs = append(pairs, from+" → "+to)
		}
	}
	m.content.WriteString(RelatedStyle.Render("🔗 Related: [" + strings.Join(pairs, ", ") + "]"))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendError adds an error message.
func (m *OutputModel) AppendError(text string) {
	m.bannerOnly = false
	m.content.WriteString(ErrorStyle.Render("⚠ " + text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendSystem adds a system message.
func (m *OutputModel) AppendSystem(text string) {
	m.bannerOnly = false
	m.content.WriteString(SystemStyle.Render(text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendContext adds a context line.
func (m *OutputModel) AppendContext(text string) {
	m.bannerOnly = false
	m.content.WriteString(ContextStyle.Render("📎 Context: " + text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// Clear resets the output content.
func (m *OutputModel) Clear() {
	m.content.Reset()
	m.stream.Reset()
	m.streaming = false
	m.refreshViewport()
}

func (m *OutputModel) refreshViewport() {
	if m.streaming && m.stream.Len() > 0 {
		// During streaming, show raw text for smooth character-by-character output.
		// Glamour re-renders the full buffer on every chunk which causes visible jumps.
		// Final glamour render happens in EndStream.
		// Strip tags so they don't flash before being replaced by styled UI elements.
		m.viewport.SetContent(m.content.String() + learning.StripTags(m.stream.String()))
	} else {
		m.viewport.SetContent(m.content.String())
	}
	if m.autoScroll {
		if m.viewport.TotalLineCount() > m.viewport.Height {
			m.viewport.GotoBottom()
		} else {
			m.viewport.SetYOffset(0)
		}
	}
}

// Update handles viewport messages.
func (m OutputModel) Update(msg tea.Msg) (OutputModel, tea.Cmd) {
	// Only track user-initiated scrolling (keys/mouse), not content changes
	isScrollInput := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "down", "pgup", "pgdown", "home", "end":
			isScrollInput = true
		}
	case tea.MouseMsg:
		isScrollInput = true
	}

	var cmd tea.Cmd
	prevOffset := m.viewport.YOffset
	m.viewport, cmd = m.viewport.Update(msg)

	if isScrollInput {
		if m.viewport.YOffset < prevOffset {
			m.autoScroll = false
		}
		if m.viewport.AtBottom() {
			m.autoScroll = true
		}
	}

	return m, cmd
}

// View renders the viewport.
func (m OutputModel) View() string {
	return m.viewport.View()
}

// glamourStyle returns a glamour style config that uses the terminal's default
// text color instead of hardcoded grays. This ensures black text on light
// backgrounds and white text on dark backgrounds.
func glamourStyle() ansi.StyleConfig {
	if lipgloss.HasDarkBackground() {
		s := styles.DarkStyleConfig
		s.Document.Color = nil // use terminal default
		return s
	}
	s := styles.LightStyleConfig
	s.Document.Color = nil // use terminal default (black on light bg)
	return s
}
