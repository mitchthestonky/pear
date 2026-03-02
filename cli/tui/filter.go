package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// FilterUnknown drops unexported bubbletea messages (like unknownCSISequenceMsg
// and unknownInputByteMsg) to prevent OSC terminal responses from leaking into the UI.
func FilterUnknown(m tea.Model, msg tea.Msg) tea.Msg {
	typeName := fmt.Sprintf("%T", msg)
	if strings.Contains(typeName, "unknown") {
		return nil
	}
	return msg
}
