package tui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
)

// OutputModel is a scrollable viewport that renders streaming markdown.
type OutputModel struct {
	viewport   viewport.Model
	content    *strings.Builder // pre-styled UI content (banner, headers, separators)
	stream     *strings.Builder // raw LLM markdown accumulated during streaming
	streaming  bool             // whether we're currently in a stream block
	renderer   *glamour.TermRenderer
	autoScroll bool
	width      int
}

// NewOutputModel creates a new output component.
func NewOutputModel(width, height int) OutputModel {
	vp := viewport.New(width, height)
	vp.SetContent("")

	r, _ := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
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
			glamour.WithAutoStyle(),
			glamour.WithWordWrap(width-4),
		)
	}
}

// AppendHeader adds a styled header line.
func (m *OutputModel) AppendHeader(text string) {
	m.content.WriteString(TriggerStyle.Render(text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// StartStream begins a new response block with an opening separator.
func (m *OutputModel) StartStream(width int) {
	m.content.WriteString(SeparatorOpen(width))
	m.content.WriteString("\n")
	m.stream.Reset()
	m.streaming = true
	m.refreshViewport()
}

// AppendChunk adds a streaming chunk and re-renders.
func (m *OutputModel) AppendChunk(text string) {
	m.stream.WriteString(text)
	m.refreshViewport()
}

// EndStream adds a closing separator.
func (m *OutputModel) EndStream(width int) {
	// Render final markdown and bake it into content
	if m.renderer != nil {
		if r, err := m.renderer.Render(m.stream.String()); err == nil {
			m.content.WriteString(r)
		} else {
			m.content.WriteString(m.stream.String())
		}
	} else {
		m.content.WriteString(m.stream.String())
	}
	m.stream.Reset()
	m.streaming = false
	m.content.WriteString("\n")
	m.content.WriteString(SeparatorClose(width))
	m.content.WriteString("\n\n")
	m.refreshViewport()
}

// AppendError adds an error message.
func (m *OutputModel) AppendError(text string) {
	m.content.WriteString(ErrorStyle.Render("⚠ " + text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendSystem adds a system message.
func (m *OutputModel) AppendSystem(text string) {
	m.content.WriteString(SystemStyle.Render(text))
	m.content.WriteString("\n")
	m.refreshViewport()
}

// AppendContext adds a context line.
func (m *OutputModel) AppendContext(text string) {
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
		// Render LLM markdown and append to pre-styled content
		raw := m.stream.String()
		rendered := raw
		if m.renderer != nil {
			if r, err := m.renderer.Render(raw); err == nil {
				rendered = r
			}
		}
		m.viewport.SetContent(m.content.String() + rendered)
	} else {
		m.viewport.SetContent(m.content.String())
	}
	if m.autoScroll {
		m.viewport.GotoBottom()
	}
}

// Update handles viewport messages.
func (m OutputModel) Update(msg tea.Msg) (OutputModel, tea.Cmd) {
	var cmd tea.Cmd
	prevOffset := m.viewport.YOffset
	m.viewport, cmd = m.viewport.Update(msg)

	// If user scrolled up, disable auto-scroll
	if m.viewport.YOffset < prevOffset {
		m.autoScroll = false
	}
	// If at bottom, re-enable auto-scroll
	if m.viewport.AtBottom() {
		m.autoScroll = true
	}

	return m, cmd
}

// View renders the viewport.
func (m OutputModel) View() string {
	return m.viewport.View()
}
