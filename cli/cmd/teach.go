package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var teachCmd = &cobra.Command{
	Use:   "teach [topic]",
	Short: "Deep-dive teaching on a topic",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented yet")
	},
}

func init() {
	rootCmd.AddCommand(teachCmd)
}
