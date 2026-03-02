package tui

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// SubmitMsg is sent when the user submits input.
type SubmitMsg struct {
	Text  string
	Files map[string]string
}

// SlashMsg is sent when the user enters a slash command.
type SlashMsg struct {
	Command string
	Args    string
}

// AutocompleteModel tracks @file autocomplete state.
type AutocompleteModel struct {
	active   bool
	prefix   string
	matches  []string
	selected int
}

// InputModel wraps a text input with @file autocomplete and slash detection.
type InputModel struct {
	textinput    textinput.Model
	autocomplete AutocompleteModel
	enabled      bool
	fileCache    []string
}

// NewInputModel creates a new input component.
func NewInputModel() InputModel {
	ti := textinput.New()
	ti.Placeholder = "Ask a question or type /help..."
	ti.Prompt = "> "
	ti.Focus()

	return InputModel{
		textinput: ti,
		enabled:   true,
	}
}

// SetEnabled enables/disables input.
func (m *InputModel) SetEnabled(enabled bool) {
	m.enabled = enabled
	if enabled {
		m.textinput.Focus()
	} else {
		m.textinput.Blur()
	}
}

// Update handles input messages.
func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	if !m.enabled {
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.autocomplete.active {
			switch msg.String() {
			case "up":
				if m.autocomplete.selected > 0 {
					m.autocomplete.selected--
				}
				return m, nil
			case "down":
				if m.autocomplete.selected < len(m.autocomplete.matches)-1 {
					m.autocomplete.selected++
				}
				return m, nil
			case "enter":
				if len(m.autocomplete.matches) > 0 {
					selected := m.autocomplete.matches[m.autocomplete.selected]
					// Replace @prefix with @selected
					val := m.textinput.Value()
					atIdx := strings.LastIndex(val, "@")
					if atIdx >= 0 {
						m.textinput.SetValue(val[:atIdx] + "@" + selected + " ")
						m.textinput.SetCursor(len(m.textinput.Value()))
					}
				}
				m.autocomplete = AutocompleteModel{}
				return m, nil
			case "esc":
				m.autocomplete = AutocompleteModel{}
				return m, nil
			}
		}

		switch msg.String() {
		case "enter":
			text := strings.TrimSpace(m.textinput.Value())
			if text == "" {
				return m, nil
			}
			m.textinput.SetValue("")
			m.autocomplete = AutocompleteModel{}

			// Slash command detection
			if strings.HasPrefix(text, "/") {
				parts := strings.SplitN(text, " ", 2)
				cmd := strings.TrimPrefix(parts[0], "/")
				if cmd == "" {
					cmd = "help"
				}
				args := ""
				if len(parts) > 1 {
					args = parts[1]
				}
				return m, func() tea.Msg {
					return SlashMsg{Command: cmd, Args: args}
				}
			}

			// Resolve @files
			files := resolveAtFiles(text)
			return m, func() tea.Msg {
				return SubmitMsg{Text: text, Files: files}
			}
		}
	}

	var cmd tea.Cmd
	m.textinput, cmd = m.textinput.Update(msg)

	// Check for @ autocomplete trigger
	val := m.textinput.Value()
	atIdx := strings.LastIndex(val, "@")
	if atIdx >= 0 {
		prefix := val[atIdx+1:]
		if !strings.Contains(prefix, " ") {
			if m.fileCache == nil {
				m.fileCache = loadFileCache()
			}
			m.autocomplete.active = true
			m.autocomplete.prefix = prefix
			m.autocomplete.matches = filterFiles(m.fileCache, prefix)
			if m.autocomplete.selected >= len(m.autocomplete.matches) {
				m.autocomplete.selected = 0
			}
		} else {
			m.autocomplete = AutocompleteModel{}
		}
	} else {
		m.autocomplete = AutocompleteModel{}
	}

	return m, cmd
}

// View renders the input component.
func (m InputModel) View() string {
	var b strings.Builder

	if m.autocomplete.active && len(m.autocomplete.matches) > 0 {
		var items []string
		max := 5
		if len(m.autocomplete.matches) < max {
			max = len(m.autocomplete.matches)
		}
		for i := 0; i < max; i++ {
			if i == m.autocomplete.selected {
				items = append(items, AutocompleteSelectedStyle.Render(m.autocomplete.matches[i]))
			} else {
				items = append(items, m.autocomplete.matches[i])
			}
		}
		b.WriteString(AutocompleteStyle.Render(strings.Join(items, "\n")))
		b.WriteString("\n")
	}

	b.WriteString(m.textinput.View())
	return b.String()
}

func loadFileCache() []string {
	cmd := exec.Command("git", "ls-files")
	if root, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err == nil {
		cmd.Dir = strings.TrimSpace(string(root))
	}
	out, err := cmd.Output()
	if err != nil {
		return nil
	}
	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var result []string
	for _, l := range lines {
		if l != "" {
			result = append(result, l)
		}
	}
	return result
}

func filterFiles(files []string, prefix string) []string {
	if prefix == "" {
		if len(files) > 10 {
			return files[:10]
		}
		return files
	}
	prefix = strings.ToLower(prefix)
	var matches []string
	for _, f := range files {
		if strings.Contains(strings.ToLower(f), prefix) {
			matches = append(matches, f)
			if len(matches) >= 10 {
				break
			}
		}
	}
	return matches
}

func resolveAtFiles(text string) map[string]string {
	files := make(map[string]string)
	words := strings.Fields(text)
	for _, w := range words {
		if strings.HasPrefix(w, "@") && len(w) > 1 {
			path := strings.TrimPrefix(w, "@")
			files[path] = ""
		}
	}
	return files
}
