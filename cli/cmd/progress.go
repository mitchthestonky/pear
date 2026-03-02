package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/learning"
	"github.com/spf13/cobra"
)

var progressCmd = &cobra.Command{
	Use:   "progress",
	Short: "Show learning progress",
	Run: func(cmd *cobra.Command, args []string) {
		path := filepath.Join(config.Dir(), "learning.json")
		store, err := learning.Load(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading learning data: %v\n", err)
			os.Exit(1)
		}
		store.Display(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(progressCmd)
}
