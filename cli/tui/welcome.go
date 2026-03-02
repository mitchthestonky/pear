package tui

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/pearcode/pear/config"
)

// WelcomeBanner renders the startup welcome screen.
func WelcomeBanner(cfg *config.Config, width int) string {
	if width < 40 {
		width = 80
	}

	green := lipgloss.NewStyle().Foreground(colorGreen)
	greenBold := lipgloss.NewStyle().Foreground(colorGreen).Bold(true)
	dim := lipgloss.NewStyle().Foreground(colorDim)

	// ASCII pear
	pear := green.Render(strings.Join([]string{
		`    ,`,
		`   / \`,
		`  /   \`,
		` (  🍐 )`,
		`  \   /`,
		`   \_/`,
	}, "\n"))

	provider := config.ActiveProvider(cfg)
	modelInfo := fmt.Sprintf("%s/%s", cfg.Provider.Active, provider.Model)

	// Left column: welcome + pear art + model info
	welcome := greenBold.Render(fmt.Sprintf("Welcome back %s!", cfg.Name))
	info := dim.Render(fmt.Sprintf("%s\n%s", modelInfo, cwd()))

	leftContent := lipgloss.JoinVertical(lipgloss.Center,
		welcome,
		"",
		pear,
		"",
		info,
	)

	leftBox := lipgloss.NewStyle().
		Width(28).
		Align(lipgloss.Center).
		Render(leftContent)

	// Right column: tips + commands
	tipsHeader := greenBold.Render("Getting started")
	tips := strings.Join([]string{
		"Ask anything — Pear teaches, not just answers",
		"Use @file to include code context",
		"Run `pear watch` for proactive teaching",
	}, "\n")

	cmdsHeader := greenBold.Render("Commands")
	cmds := strings.Join([]string{
		dim.Render("/settings") + "  configure provider & model",
		dim.Render("/clear") + "     reset conversation",
		dim.Render("/quit") + "      exit",
	}, "\n")

	rightContent := lipgloss.JoinVertical(lipgloss.Left,
		tipsHeader,
		tips,
		"",
		cmdsHeader,
		cmds,
	)

	rightBox := lipgloss.NewStyle().
		Width(width - 32).
		PaddingLeft(2).
		Render(rightContent)

	// Divider
	divider := lipgloss.NewStyle().
		Foreground(colorDim).
		Render("│")

	// Compose inner content
	inner := lipgloss.JoinHorizontal(lipgloss.Top, leftBox, divider, rightBox)

	// Outer box with version label
	versionLabel := green.Render("─ Pear v0 ")
	topBorder := versionLabel + dim.Render(strings.Repeat("─", max(0, width-12)))
	bottomBorder := dim.Render(strings.Repeat("─", width))

	return lipgloss.JoinVertical(lipgloss.Left,
		topBorder,
		inner,
		bottomBorder,
		"",
	)
}

func cwd() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	home, _ := os.UserHomeDir()
	if home != "" && strings.HasPrefix(dir, home) {
		return "~" + dir[len(home):]
	}
	return filepath.Clean(dir)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
