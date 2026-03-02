package hooks

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	startMarker = "# pear-hook-start"
	endMarker   = "# pear-hook-end"
	pearBlock   = "\n# pear-hook-start\npear review --commit HEAD\n# pear-hook-end\n"
	shebang     = "#!/bin/sh\n"
)

func hookPath(repoRoot string) string {
	return filepath.Join(repoRoot, ".git", "hooks", "post-commit")
}

// Install installs a post-commit git hook that runs pear review.
func Install(repoRoot string) error {
	p := hookPath(repoRoot)

	// Ensure hooks directory exists
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return fmt.Errorf("creating hooks directory: %w", err)
	}

	data, err := os.ReadFile(p)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("reading hook file: %w", err)
	}

	content := string(data)

	// Already installed
	if strings.Contains(content, startMarker) {
		fmt.Println("Pear hook already installed.")
		return nil
	}

	if os.IsNotExist(err) {
		// Create new file with shebang
		content = shebang + pearBlock
	} else {
		// Append to existing
		content += pearBlock
	}

	if err := os.WriteFile(p, []byte(content), 0755); err != nil {
		return fmt.Errorf("writing hook file: %w", err)
	}

	if err := os.Chmod(p, 0755); err != nil {
		return fmt.Errorf("setting hook permissions: %w", err)
	}

	fmt.Println("Pear post-commit hook installed.")
	return nil
}

// Uninstall removes the pear block from the post-commit hook.
func Uninstall(repoRoot string) error {
	p := hookPath(repoRoot)

	data, err := os.ReadFile(p)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("No post-commit hook found.")
			return nil
		}
		return fmt.Errorf("reading hook file: %w", err)
	}

	content := string(data)
	if !strings.Contains(content, startMarker) {
		fmt.Println("Pear hook not installed.")
		return nil
	}

	// Remove everything between markers (inclusive)
	startIdx := strings.Index(content, startMarker)
	endIdx := strings.Index(content, endMarker)
	if startIdx == -1 || endIdx == -1 {
		return nil
	}

	// Include the trailing newline after end marker
	endIdx += len(endMarker)
	if endIdx < len(content) && content[endIdx] == '\n' {
		endIdx++
	}

	// Also remove the leading newline before start marker if present
	if startIdx > 0 && content[startIdx-1] == '\n' {
		startIdx--
	}

	remaining := content[:startIdx] + content[endIdx:]

	// If only shebang or whitespace remains, delete the file
	trimmed := strings.TrimSpace(remaining)
	if trimmed == "" || trimmed == strings.TrimSpace(shebang) {
		if err := os.Remove(p); err != nil {
			return fmt.Errorf("removing hook file: %w", err)
		}
		fmt.Println("Pear post-commit hook removed (file deleted).")
		return nil
	}

	if err := os.WriteFile(p, []byte(remaining), 0755); err != nil {
		return fmt.Errorf("writing hook file: %w", err)
	}

	fmt.Println("Pear post-commit hook removed.")
	return nil
}

// IsInstalled returns true if the pear hook is installed.
func IsInstalled(repoRoot string) bool {
	data, err := os.ReadFile(hookPath(repoRoot))
	if err != nil {
		return false
	}
	return strings.Contains(string(data), startMarker)
}
