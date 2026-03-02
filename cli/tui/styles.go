package tui

import (
	"github.com/charmbracelet/lipgloss"
)

// Brand colors
var (
	colorGreen  = lipgloss.AdaptiveColor{Light: "#2D6A2E", Dark: "#5FD75F"}
	colorDim    = lipgloss.AdaptiveColor{Light: "#888888", Dark: "#777777"}
	colorRed    = lipgloss.AdaptiveColor{Light: "#CC3333", Dark: "#FF5555"}
	colorBlue = lipgloss.AdaptiveColor{Light: "#2255CC", Dark: "#5599FF"}
	colorCyan   = lipgloss.AdaptiveColor{Light: "#117777", Dark: "#55DDDD"}
)

// HeaderStyle for session header line (🍐 Pear v0 · watching · ...)
var HeaderStyle = lipgloss.NewStyle().
	Foreground(colorGreen).
	Bold(true)

// TriggerStyle for proactive trigger notifications (🍐 Pear noticed...)
var TriggerStyle = lipgloss.NewStyle().
	Foreground(colorGreen)

// ContextStyle for context lines (📎 Context: ...) — dimmed
var ContextStyle = lipgloss.NewStyle().
	Foreground(colorDim)

// ConceptStyle for concept tags (📚 Concepts: [...])
var ConceptStyle = lipgloss.NewStyle().
	Foreground(colorBlue)

// RelatedStyle for relationship tags (🔗 Related: [...])
var RelatedStyle = lipgloss.NewStyle().
	Foreground(colorCyan)

// ErrorStyle for inline errors (⚠ ...) — red/warning
var ErrorStyle = lipgloss.NewStyle().
	Foreground(colorRed).
	Bold(true)

// SystemStyle for system messages (slash command output) — dimmed
var SystemStyle = lipgloss.NewStyle().
	Foreground(colorDim).
	Italic(true)

// AutocompleteStyle for the @file dropdown box
var AutocompleteStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(colorDim).
	Padding(0, 1)

// AutocompleteSelectedStyle for the highlighted item in dropdown
var AutocompleteSelectedStyle = lipgloss.NewStyle().
	Foreground(colorGreen).
	Bold(true)

// ProgressBarFull for filled progress bar segments
var ProgressBarFull = lipgloss.NewStyle().
	Foreground(colorGreen)

// ProgressBarEmpty for empty progress bar segments
var ProgressBarEmpty = lipgloss.NewStyle().
	Foreground(colorDim)

// UserMessageStyle for displaying the user's input in the log
var UserMessageStyle = lipgloss.NewStyle().
	Foreground(colorDim).
	Bold(true)

// ThinkingStyle for the "Thinking..." indicator
var ThinkingStyle = lipgloss.NewStyle().
	Foreground(colorDim).
	Italic(true)


