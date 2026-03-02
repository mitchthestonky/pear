package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/learning"
	"github.com/pearcode/pear/llm"
	"github.com/pearcode/pear/prompt"
	"github.com/pearcode/pear/repocontext"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask [question]",
	Short: "Ask Pear a question",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		question := args[0]

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

		profile := prompt.UserProfile{
			Name:      cfg.Name,
			Languages: cfg.Languages,
			Level:     cfg.Level,
		}

		rctx, _ := repocontext.Build(repocontext.CollectOpts{
			TriggerType: "user",
			TriggerInfo: "pear ask",
		})
		if rctx == nil {
			rctx = &repocontext.RepoContext{}
		}

		systemPrompt, messages := prompt.Reactive(rctx, profile, []llm.Message{
			{Role: "user", Content: question},
		})

		opts := llm.StreamOptions{
			SystemPrompt: systemPrompt,
			MaxTokens:    4096,
			Temperature:  0.7,
		}

		printSeparator()

		resp, err := llm.StreamWithRetry(context.Background(), client, messages, opts, func(chunk string) {
			fmt.Print(chunk)
		})
		if err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
			os.Exit(1)
		}

		if resp != nil {
			lpath := filepath.Join(config.Dir(), "learning.json")
			store, _ := learning.Load(lpath)
			concepts, relationships := learning.Extract(resp.Content)
			if len(concepts) > 0 {
				store.Record(concepts, relationships)
				_ = store.Save(lpath)
			}
		}

		fmt.Println()
	},
}

func init() {
	rootCmd.AddCommand(askCmd)
}
