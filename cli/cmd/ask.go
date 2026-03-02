package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/llm"
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

		systemPrompt := "You are Pear, a teaching-first coding companion. Teach concepts, not just answers."

		messages := []llm.Message{
			{Role: "user", Content: question},
		}

		opts := llm.StreamOptions{
			SystemPrompt: systemPrompt,
			MaxTokens:    4096,
			Temperature:  0.7,
		}

		printSeparator()

		_, err = client.Stream(context.Background(), messages, opts, func(chunk string) {
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
	rootCmd.AddCommand(askCmd)
}
