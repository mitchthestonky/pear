package cmd

import (
	"fmt"
	"strings"

	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/llm"
)

const separator = "━━━ Pear ━━━"

func printSeparator() {
	fmt.Println(separator)
	fmt.Println()
}

func printContextLine(parts ...string) {
	fmt.Printf("📎 Context: %s\n\n", strings.Join(parts, ", "))
}

func newLLMClient(cfg *config.Config) (llm.LLMClient, error) {
	provider := cfg.Provider.Active
	detail := config.ActiveProvider(cfg)
	return llm.NewClient(provider, llm.ProviderDetail{
		APIKey: detail.APIKey,
		Model:  detail.Model,
	})
}
