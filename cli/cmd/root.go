package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/llm"
	"github.com/pearcode/pear/logging"
	"github.com/pearcode/pear/tui"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pear",
	Short: "Pear — a coding teacher that watches you code",
	Long:  "Pear is a CLI teaching tool that watches you code and proactively teaches during natural pauses.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Name() == "init" {
			return
		}
		if !config.Exists() {
			runInitWizard()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
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

		logsDir := filepath.Join(config.Dir(), "logs")
		logger, _ := logging.NewLogger(logsDir)
		if logger != nil {
			defer logger.Close()
			logger.Log("session_start", map[string]any{"mode": "interactive", "provider": cfg.Provider.Active})
		}

		m := tui.NewModel(cfg, client, "interactive", nil)
		p := tea.NewProgram(m, tea.WithAltScreen())
		if _, err := p.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
