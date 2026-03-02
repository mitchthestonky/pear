package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// FilterUnknown drops terminal escape responses that leak into the TUI.
// These include OSC 11 background color responses and cursor position reports
// from termenv/lipgloss/glamour.
func FilterUnknown(m tea.Model, msg tea.Msg) tea.Msg {
	// Drop unexported bubbletea message types
	typeName := fmt.Sprintf("%T", msg)
	if strings.Contains(typeName, "unknown") {
		return nil
	}

	// Drop KeyMsg that contain escape sequence fragments
	if key, ok := msg.(tea.KeyMsg); ok {
		s := key.String()
		if strings.Contains(s, ";rgb:") ||
			strings.Contains(s, "\x1b[") ||
			strings.Contains(s, "\x1b]") ||
			strings.HasPrefix(s, "]") ||
			(len(s) > 1 && strings.HasSuffix(s, "R") && strings.Contains(s, ";")) {
			return nil
		}
	}

	return msg
}
