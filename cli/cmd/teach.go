package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strings"

	"github.com/MitchTheStonky/pear/cli/config"
	"github.com/MitchTheStonky/pear/cli/learning"
	"github.com/MitchTheStonky/pear/cli/llm"
	"github.com/MitchTheStonky/pear/cli/prompt"
	"github.com/MitchTheStonky/pear/cli/repocontext"
	"github.com/spf13/cobra"
)

var teachCmd = &cobra.Command{
	Use:   "teach [topic]",
	Short: "Deep-dive teaching on a topic",
	Args:  cobra.MaximumNArgs(1),
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

		profile := prompt.UserProfile{
			Name:      cfg.Name,
			Languages: cfg.Languages,
			Level:     cfg.Level,
		}

		var systemPrompt string
		var messages []llm.Message

		if len(args) == 0 {
			repoCtx, err := repocontext.Build(repocontext.CollectOpts{
				TriggerType: "user",
				TriggerInfo: "pear teach",
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error collecting context: %v\n", err)
				os.Exit(1)
			}

			if strings.TrimSpace(repoCtx.Diff) == "" {
				fmt.Println("No changes to teach about.")
				return
			}

			diffLines := len(strings.Split(repoCtx.Diff, "\n"))
			contextParts := []string{fmt.Sprintf("git diff (%d lines)", diffLines)}
			if repoCtx.Branch != "" {
				contextParts = append(contextParts, fmt.Sprintf("branch: %s", repoCtx.Branch))
			}
			printContextLine(contextParts...)
			lpath := config.LearningPath()
			store, _ := learning.Load(lpath)
			systemPrompt, messages = prompt.Proactive(repoCtx, profile, nil, nil, store)
		} else {
			// Topic: deep dive with auto-selected files
			topic := args[0]
			root, err := repocontext.RepoRoot()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: not in a git repo: %v\n", err)
				os.Exit(1)
			}

			autoFiles := grepForTopic(root, topic)

			filesMap := make(map[string]string)
			for _, f := range autoFiles {
				filesMap[f] = ""
			}

			repoCtx, err := repocontext.Build(repocontext.CollectOpts{
				Files:       filesMap,
				TriggerType: "user",
				TriggerInfo: "pear teach " + topic,
			})
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error collecting context: %v\n", err)
				os.Exit(1)
			}

			if len(autoFiles) > 0 {
				printContextLine(fmt.Sprintf("auto-selected %s", strings.Join(autoFiles, ", ")))
			} else {
				printContextLine("no matching files found, teaching from general knowledge")
			}

			systemPrompt, messages = prompt.DeepDive(repoCtx, profile, topic)
		}

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
	rootCmd.AddCommand(teachCmd)
}

// grepForTopic uses git grep -c to find files matching the topic, returns top 3 by match count.
func grepForTopic(repoDir, topic string) []string {
	cmd := exec.Command("git", "grep", "-c", "-i", topic)
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return nil
	}

	type fileCount struct {
		path  string
		count int
	}

	var results []fileCount
	for _, line := range strings.Split(strings.TrimSpace(string(out)), "\n") {
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}
		count := 0
		fmt.Sscanf(parts[1], "%d", &count)
		results = append(results, fileCount{path: parts[0], count: count})
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].count > results[j].count
	})

	var files []string
	for i, r := range results {
		if i >= 3 {
			break
		}
		files = append(files, r.path)
	}
	return files
}
