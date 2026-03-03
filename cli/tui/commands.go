package tui

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/repocontext"
	"github.com/MitchTheStonky/pear/cli/watcher"
)

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
		if m.watchCancel != nil {
			m.watchCancel()
		}
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
