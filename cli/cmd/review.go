package cmd

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/learning"
	"github.com/MitchTheStonky/pear/cli/llm"
	"github.com/MitchTheStonky/pear/cli/prompt"
	"github.com/MitchTheStonky/pear/cli/repocontext"
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

		lpath := config.LearningPath()
		store, _ := learning.Load(lpath)
		systemPrompt, messages := prompt.Proactive(repoCtx, profile, nil, nil, store)

		streamOpts := llm.StreamOptions{
			SystemPrompt: systemPrompt,
			MaxTokens:    4096,
			Temperature:  0.7,
		}

		printSeparator()

		resp, err := llm.StreamWithRetry(context.Background(), client, messages, streamOpts, func(chunk string) {
			fmt.Print(chunk)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			os.Exit(1)
		}

		if resp != nil {
			lpath := config.LearningPath()
			store, _ := learning.Load(lpath)
			concepts, relationships, _ := learning.Extract(resp.Content)
			if len(concepts) > 0 {
				store.Record(concepts, relationships)
				_ = store.Save(lpath)
			}
		}

		fmt.Println()
	},
}

func init() {
	reviewCmd.Flags().StringVar(&reviewCommit, "commit", "", "commit to review (default HEAD)")
	rootCmd.AddCommand(reviewCmd)
}
