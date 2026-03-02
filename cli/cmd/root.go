package cmd

import (
	"fmt"
	"os"

	"github.com/pearcode/pear/config"
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
		fmt.Println("not implemented yet")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
