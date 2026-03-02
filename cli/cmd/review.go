package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/llm"
	"github.com/pearcode/pear/prompt"
	"github.com/pearcode/pear/repocontext"
	"github.com/spf13/cobra"
)

var reviewCommit string

var reviewCmd = &cobra.Command{
	Use:   "review",
	Short: "Review recent code changes",
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := config.Load()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error loading config: %v\n", err)
			os.Exit(1)
		}

		client, err := newLLMClient(cfg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating LLM client: %v\n", err)
			os.Exit(1)
		}

		opts := repocontext.CollectOpts{
			TriggerType: "user",
			TriggerInfo: "pear review",
		}
		if reviewCommit != "" {
			opts.DiffFrom = reviewCommit + "~1"
			opts.DiffTo = reviewCommit
		}

		repoCtx, err := repocontext.Build(opts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error collecting context: %v\n", err)
			os.Exit(1)
		}

		if strings.TrimSpace(repoCtx.Diff) == "" {
			fmt.Println("No changes to review.")
			return
		}

		diffLines := len(strings.Split(repoCtx.Diff, "\n"))
		contextParts := []string{fmt.Sprintf("git diff (%d lines)", diffLines)}
		if repoCtx.Branch != "" {
			contextParts = append(contextParts, fmt.Sprintf("branch: %s", repoCtx.Branch))
		}
		printContextLine(contextParts...)

		profile := prompt.UserProfile{
			Name:      cfg.Name,
			Languages: cfg.Languages,
			Level:     cfg.Level,
		}

		systemPrompt, messages := prompt.Proactive(repoCtx, profile, nil)

		streamOpts := llm.StreamOptions{
			SystemPrompt: systemPrompt,
			MaxTokens:    4096,
			Temperature:  0.7,
		}

		printSeparator()

		_, err = client.Stream(context.Background(), messages, streamOpts, func(chunk string) {
			fmt.Print(chunk)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			os.Exit(1)
		}

		fmt.Println()
	},
}

func init() {
	reviewCmd.Flags().StringVar(&reviewCommit, "commit", "", "commit to review (default HEAD)")
	rootCmd.AddCommand(reviewCmd)
}
