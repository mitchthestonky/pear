package tui

import (
	"os/exec"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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

// slashCommand defines a slash command with its description.
type slashCommand struct {
	Name string
	Desc string
}

var slashCommands = []slashCommand{
	{"help", "Show all commands"},
	{"clear", "Clear chat history"},
	{"exit", "End session"},
	{"watch", "Start file watcher"},
	{"review", "Review current changes"},
	{"pause", "Pause proactive reviews"},
	{"resume", "Resume proactive reviews"},
	{"status", "Session info"},
	{"settings", "Configure provider & model"},
	{"provider", "Change LLM provider"},
	{"model", "Change model"},
	{"key", "Update API key"},
}

// AutocompleteItem is a single autocomplete suggestion.
type AutocompleteItem struct {
	Name string
	Desc string // optional description (slash commands)
}

// AutocompleteModel tracks @file and /command autocomplete state.
type AutocompleteModel struct {
	active   bool
	prefix   string
	items    []AutocompleteItem
	selected int
	isSlash  bool
}

// InputModel wraps a text input with @file autocomplete and slash detection.
type InputModel struct {
	textinput    textinput.Model
	autocomplete AutocompleteModel
	enabled      bool
	fileCache    []string
}

// AutocompleteActive reports whether the autocomplete menu is showing.
func (m InputModel) AutocompleteActive() bool {
	return m.autocomplete.active
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

// SetEnabled enables/disables input and returns a tea.Cmd for focus.
func (m *InputModel) SetEnabled(enabled bool) tea.Cmd {
	m.enabled = enabled
	if enabled {
		return m.textinput.Focus()
	}
	m.textinput.Blur()
	return nil
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
				if m.autocomplete.selected < len(m.autocomplete.items)-1 {
					m.autocomplete.selected++
				}
				return m, nil
			case "tab", "enter":
				if len(m.autocomplete.items) == 0 {
					m.autocomplete = AutocompleteModel{}
					return m, nil
				}
				item := m.autocomplete.items[m.autocomplete.selected]
				isSlash := m.autocomplete.isSlash
				m.autocomplete = AutocompleteModel{}

				if isSlash && msg.String() == "enter" {
					m.textinput.SetValue("")
					cmd := item.Name
					return m, func() tea.Msg {
						return SlashMsg{Command: cmd, Args: ""}
					}
				}
				if isSlash {
					m.textinput.SetValue("/" + item.Name + " ")
				} else {
					v := m.textinput.Value()
					atIdx := strings.LastIndex(v, "@")
					if atIdx >= 0 {
						m.textinput.SetValue(v[:atIdx] + "@" + item.Name + " ")
					}
				}
				m.textinput.SetCursor(len(m.textinput.Value()))
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

	// Check for autocomplete triggers (/ commands and @files).
	// Preserve selected index when the prefix hasn't changed (e.g. on cursor blink).
	val := m.textinput.Value()
	prevPrefix := m.autocomplete.prefix
	prevSelected := m.autocomplete.selected

	if strings.HasPrefix(val, "/") && !strings.Contains(val, " ") {
		prefix := strings.TrimPrefix(val, "/")
		items := filterSlashCommands(prefix)
		if len(items) > 0 {
			sel := 0
			if prefix == prevPrefix && prevSelected < len(items) {
				sel = prevSelected
			}
			m.autocomplete = AutocompleteModel{
				active:   true,
				prefix:   prefix,
				items:    items,
				isSlash:  true,
				selected: sel,
			}
		} else {
			m.autocomplete = AutocompleteModel{}
		}
	} else if atIdx := strings.LastIndex(val, "@"); atIdx >= 0 {
		prefix := val[atIdx+1:]
		if !strings.Contains(prefix, " ") {
			if m.fileCache == nil {
				m.fileCache = loadFileCache()
			}
			files := filterFiles(m.fileCache, prefix)
			items := make([]AutocompleteItem, len(files))
			for i, f := range files {
				items[i] = AutocompleteItem{Name: f}
			}
			sel := 0
			if prefix == prevPrefix && prevSelected < len(items) {
				sel = prevSelected
			}
			m.autocomplete = AutocompleteModel{
				active:   true,
				prefix:   prefix,
				items:    items,
				selected: sel,
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

	if m.autocomplete.active && len(m.autocomplete.items) > 0 {
		var lines []string
		max := 6
		if len(m.autocomplete.items) < max {
			max = len(m.autocomplete.items)
		}
		dim := lipgloss.NewStyle().Foreground(colorDim)
		for i := 0; i < max; i++ {
			item := m.autocomplete.items[i]
			name := item.Name
			if m.autocomplete.isSlash {
				name = "/" + name
			}
			desc := ""
			if item.Desc != "" {
				desc = "  " + dim.Render(item.Desc)
			}
			if i == m.autocomplete.selected {
				lines = append(lines, AutocompleteSelectedStyle.Render(name)+desc)
			} else {
				lines = append(lines, dim.Render(name)+desc)
			}
		}
		b.WriteString(AutocompleteStyle.Render(strings.Join(lines, "\n")))
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

func filterSlashCommands(prefix string) []AutocompleteItem {
	prefix = strings.ToLower(prefix)
	var items []AutocompleteItem
	for _, sc := range slashCommands {
		if strings.HasPrefix(sc.Name, prefix) {
			items = append(items, AutocompleteItem{Name: sc.Name, Desc: sc.Desc})
		}
	}
	return items
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
