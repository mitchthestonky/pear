package cmd

import (
	"fmt"
	"os"

	"github.com/pearcode/pear/hooks"
	"github.com/pearcode/pear/repocontext"
	"github.com/spf13/cobra"
)

var hooksCmd = &cobra.Command{
	Use:   "hooks",
	Short: "Manage git hooks",
}

var hooksInstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install post-commit git hook",
	Run: func(cmd *cobra.Command, args []string) {
		root, err := repocontext.RepoRoot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: not a git repository\n")
			os.Exit(1)
		}
		if err := hooks.Install(root); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

var hooksUninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall post-commit git hook",
	Run: func(cmd *cobra.Command, args []string) {
		root, err := repocontext.RepoRoot()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: not a git repository\n")
			os.Exit(1)
		}
		if err := hooks.Uninstall(root); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	hooksCmd.AddCommand(hooksInstallCmd)
	hooksCmd.AddCommand(hooksUninstallCmd)
	rootCmd.AddCommand(hooksCmd)
}
