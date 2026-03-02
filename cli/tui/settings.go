package tui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/llm"
)

// settingsState tracks the /settings numbered editor flow.
type settingsState struct {
	active       bool
	awaitingEdit int  // which field number (1-9), 0 = showing menu
	providerPick bool // awaiting provider choice (1-3)
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
