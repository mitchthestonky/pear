package cmd

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/llm"
	"github.com/MitchTheStonky/pear/cli/logging"
	"github.com/MitchTheStonky/pear/cli/repocontext"
	"github.com/MitchTheStonky/pear/cli/tui"
	"github.com/MitchTheStonky/pear/cli/watcher"
	"github.com/spf13/cobra"
)

var watchCmd = &cobra.Command{
	Use:   "watch",
	Short: "Watch for code changes and teach proactively",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		_ = config.ResolveCodebase(cfg)

		repoRoot, err := repocontext.RepoRoot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "⚠ Not a git repository. Run from a project with git initialized.\n")
			os.Exit(1)
		}

		provider := config.ActiveProvider(cfg)
		client, err := llm.NewClient(cfg.Provider.Active, llm.ProviderDetail{
			APIKey: provider.APIKey,
			Model:  provider.Model,
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating LLM client: %v\n", err)
			os.Exit(1)
		}

		// Init logger
		logsDir := filepath.Join(config.Dir(), "logs")
		logger, _ := logging.NewLogger(logsDir)
		if logger != nil {
			defer logger.Close()
		}

		// Init watcher
		w, err := watcher.New(cfg.Watch, repoRoot, logger)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error initializing watcher: %v\n", err)
			os.Exit(1)
		}

		// Check dirty diff at startup
		baselineDiff := w.BaselineDiff()
		var initialTrigger *tui.ReviewTrigger
		if strings.TrimSpace(baselineDiff) != "" {
			lines := strings.Count(baselineDiff, "\n")
			fmt.Printf("🍐 You have uncommitted changes (%d lines). Review them now? [y/N] ", lines)
			reader := bufio.NewReader(os.Stdin)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(strings.ToLower(answer))
			if answer == "y" || answer == "yes" {
				initialTrigger = &tui.ReviewTrigger{
					Diff:        baselineDiff,
					TriggerType: "settle",
					Info:        fmt.Sprintf("made changes (%d lines)", lines),
				}
			}
		}

		// Start watcher goroutine
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		watcherTriggers := w.Start(ctx)

		// Bridge watcher triggers to TUI triggers channel
		tuiTriggers := make(chan tui.ReviewTrigger, 1)
		go func() {
			// If there's an initial trigger, send it first
			if initialTrigger != nil {
				tuiTriggers <- *initialTrigger
			}
			for wt := range watcherTriggers {
				tuiTriggers <- tui.ReviewTrigger{
					Diff:        wt.Diff,
					TriggerType: wt.Type,
					Info:        formatTriggerInfo(wt),
				}
			}
			close(tuiTriggers)
		}()

		// Launch TUI
		m := tui.NewModel(cfg, client, "watch", tuiTriggers)
		p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithMouseCellMotion(), tea.WithFilter(tui.FilterUnknown))
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		// Cancel context to stop watcher
		cancel()
	},
}

func formatTriggerInfo(wt watcher.ReviewTrigger) string {
	switch wt.Type {
	case "commit":
		return fmt.Sprintf("committed: %s", wt.Summary)
	case "settle":
		return fmt.Sprintf("made changes (%s)", wt.Summary)
	default:
		return wt.Summary
	}
}

func init() {
	rootCmd.AddCommand(watchCmd)
}
