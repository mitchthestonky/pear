package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var reviewCommit string

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review recent code changes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("not implemented yet")
	},
}

func init() {
	reviewCmd.Flags().StringVar(&reviewCommit, "commit", "", "commit to review (default HEAD)")
	rootCmd.AddCommand(reviewCmd)
}
