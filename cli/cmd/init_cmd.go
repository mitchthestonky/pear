package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize Pear configuration",
	Run: func(cmd *cobra.Command, args []string) {
		runInitWizard()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInitWizard() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("🍐 Welcome to Pear! Let's get you set up.")
	fmt.Println()

	name := promptRequired(scanner, "Your name: ")
	languages := promptRequired(scanner, "Languages you're learning (e.g. Go, Python): ")
	level := promptSelect(scanner, "Your level", []string{"junior", "mid", "senior"})
	provider := promptSelect(scanner, "LLM provider", []string{"anthropic", "openai", "openrouter"})
	apiKey := promptRequired(scanner, "API key: ")

	defaultModel := providerDefaultModel(provider)
	fmt.Printf("Model (default: %s): ", defaultModel)
	scanner.Scan()
	model := strings.TrimSpace(scanner.Text())
	if model == "" {
		model = defaultModel
	}

	cfg := &config.Config{
		Name:      name,
		Languages: languages,
		Level:     level,
		Provider: config.ProviderConfig{
			Active: provider,
		},
	}

	detail := config.ActiveProvider(cfg)
	detail.APIKey = apiKey
	detail.Model = model

	if err := config.Save(cfg); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()
	fmt.Println("Config saved! Running doctor...")
	fmt.Println()

	ok := RunDoctor()
	if ok {
		fmt.Println()
		fmt.Println("🍐 You're all set! Pear is ready to help you learn.")
	} else {
		fmt.Println()
		fmt.Println("Some checks failed. Run `pear doctor` after fixing the issues.")
	}
}

func promptRequired(scanner *bufio.Scanner, prompt string) string {
	for {
		fmt.Print(prompt)
		scanner.Scan()
		val := strings.TrimSpace(scanner.Text())
		if val != "" {
			return val
		}
		fmt.Println("  This field is required.")
	}
}

func promptSelect(scanner *bufio.Scanner, label string, options []string) string {
	for {
		fmt.Printf("%s:\n", label)
		for i, opt := range options {
			fmt.Printf("  %d. %s\n", i+1, opt)
		}
		fmt.Print("Choose (1-", len(options), "): ")
		scanner.Scan()
		val := strings.TrimSpace(scanner.Text())
		for i, opt := range options {
			if val == fmt.Sprintf("%d", i+1) || val == opt {
				return opt
			}
		}
		fmt.Println("  Invalid choice.")
	}
}

func providerDefaultModel(provider string) string {
	switch provider {
	case "anthropic":
		return "claude-haiku-4-5"
	case "openai":
		return "gpt-4o"
	case "openrouter":
		return "anthropic/claude-3.5-sonnet"
	default:
		return ""
	}
}
