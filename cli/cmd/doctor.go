package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/llm"
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

	// Check 4: LLM test request
	if cfg != nil && config.ActiveProvider(cfg).APIKey != "" {
		client, cerr := newLLMClient(cfg)
		if cerr != nil {
			fmt.Printf("✗ LLM connection: %v\n", cerr)
			allPassed = false
		} else {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			_, cerr = client.Stream(ctx, []llm.Message{{Role: "user", Content: "Respond with OK"}}, llm.StreamOptions{MaxTokens: 10}, func(string) {})
			if cerr != nil {
				var llmErr *llm.LLMError
				if errors.As(cerr, &llmErr) && llmErr.Code == "auth" {
					fmt.Println("✗ LLM connection: invalid API key")
				} else {
					fmt.Printf("✗ LLM connection: %v\n", cerr)
				}
				allPassed = false
			} else {
				fmt.Println("✓ LLM connection: valid")
			}
		}
	} else {
		fmt.Println("- LLM connection: skipped (no API key)")
	}

	return allPassed
}
