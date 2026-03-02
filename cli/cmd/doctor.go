package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pearcode/pear/config"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check system health",
	Run: func(cmd *cobra.Command, args []string) {
		ok := RunDoctor()
		if !ok {
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}

// RunDoctor runs all health checks and returns true if all pass.
func RunDoctor() bool {
	allPassed := true

	// Check 1: git installed
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Println("✗ git: not found")
		allPassed = false
	} else {
		fmt.Println("✓ git: installed")
	}

	// Check 2: config valid
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("✗ config: invalid")
		allPassed = false
	} else {
		fmt.Println("✓ config: valid")
	}

	// Check 3: API key set
	if cfg != nil {
		p := config.ActiveProvider(cfg)
		if p.APIKey == "" {
			fmt.Println("✗ API key: missing")
			allPassed = false
		} else {
			fmt.Println("✓ API key: set")
		}
	} else {
		fmt.Println("✗ API key: missing (no config)")
		allPassed = false
	}

	// Check 4: LLM test request (stubbed until LLM package is ready)
	// TODO: implement when llm package is available (ticket 02-01)
	fmt.Println("- LLM connection: skipped (not yet implemented)")

	return allPassed
}
