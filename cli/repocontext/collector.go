package repocontext

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

// RepoContext holds gathered repository state for prompt assembly.
type RepoContext struct {
	Diff         string
	ChangedFiles []string
	FileTree     string
	Branch       string
	Files        map[string]string // @file path → contents
	TriggerType  string            // "settle", "commit", "user"
	TriggerInfo  string            // summary string
}

// CollectOpts configures what Build collects.
type CollectOpts struct {
	DiffFrom    string            // if set, use GitDiffRange(DiffFrom, DiffTo)
	DiffTo      string
	Files       map[string]string // @file paths to read (path → "")
	TriggerType string
	TriggerInfo string
	RepoDir     string // override repo root detection
}

const (
	maxDiffLines = 300
	maxFileLines = 200
	maxTreeLines = 100
)

// Build constructs a RepoContext from git state.
func Build(opts CollectOpts) (*RepoContext, error) {
	root := opts.RepoDir
	if root == "" {
		var err error
		root, err = RepoRoot()
		if err != nil {
			return nil, err
		}
	}

	ctx := &RepoContext{
		TriggerType: opts.TriggerType,
		TriggerInfo: opts.TriggerInfo,
		Files:       make(map[string]string),
	}

	// Get branch
	branch, err := GitBranch(root)
	if err == nil {
		ctx.Branch = branch
	}

	// Get diff
	if opts.DiffFrom != "" {
		diff, err := GitDiffRange(root, opts.DiffFrom, opts.DiffTo)
		if err == nil {
			ctx.Diff = diff
		}
	} else {
		diff, err := GitDiff(root)
		if err == nil {
			ctx.Diff = diff
		}
	}

	ctx.ChangedFiles = ParseChangedFiles(ctx.Diff)

	// Get file tree
	tree, err := GitFileTree(root)
	if err == nil {
		ctx.FileTree = tree
	}

	// Read @files
	for path := range opts.Files {
		content, err := ReadFile(root, path)
		if err == nil {
			ctx.Files[path] = content
		}
	}

	return ctx, nil
}

// RepoRoot returns the git repository root directory.
func RepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("not in a git repository: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// GitDiff returns `git diff HEAD`, truncated to maxDiffLines.
func GitDiff(repoDir string) (string, error) {
	cmd := exec.Command("git", "diff", "HEAD")
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git diff failed: %w", err)
	}
	return truncateLines(string(out), maxDiffLines), nil
}

// GitDiffRange returns `git diff from..to`, truncated to maxDiffLines.
func GitDiffRange(repoDir, from, to string) (string, error) {
	cmd := exec.Command("git", "diff", from+".."+to)
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git diff range failed: %w", err)
	}
	return truncateLines(string(out), maxDiffLines), nil
}

// GitFileTree returns `git ls-files` with depth 2, max 100 entries.
func GitFileTree(repoDir string) (string, error) {
	cmd := exec.Command("git", "ls-files")
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git ls-files failed: %w", err)
	}

	lines := strings.Split(strings.TrimSpace(string(out)), "\n")
	var filtered []string
	for _, line := range lines {
		if line == "" {
			continue
		}
		// depth 2: at most 2 slashes (file is at depth ≤ 2)
		if strings.Count(line, string(os.PathSeparator)) <= 2 && strings.Count(line, "/") <= 2 {
			filtered = append(filtered, line)
		}
		if len(filtered) >= maxTreeLines {
			break
		}
	}
	return strings.Join(filtered, "\n"), nil
}

// GitBranch returns the current git branch name.
func GitBranch(repoDir string) (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	cmd.Dir = repoDir
	out, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git branch failed: %w", err)
	}
	return strings.TrimSpace(string(out)), nil
}

// ReadFile reads a file relative to repoDir, truncated to maxFileLines.
func ReadFile(repoDir, path string) (string, error) {
	fullPath := path
	if !strings.HasPrefix(path, "/") {
		fullPath = repoDir + "/" + path
	}
	data, err := os.ReadFile(fullPath)
	if err != nil {
		return "", fmt.Errorf("read file %s: %w", path, err)
	}
	return truncateLines(string(data), maxFileLines), nil
}

var diffFilePattern = regexp.MustCompile(`^diff --git a/(.+) b/`)

// ParseChangedFiles extracts file paths from unified diff headers.
func ParseChangedFiles(diff string) []string {
	var files []string
	seen := make(map[string]bool)
	for _, line := range strings.Split(diff, "\n") {
		matches := diffFilePattern.FindStringSubmatch(line)
		if len(matches) >= 2 && !seen[matches[1]] {
			files = append(files, matches[1])
			seen[matches[1]] = true
		}
	}
	return files
}

// DiffSummary returns a human-readable summary like "3 files, +47 lines".
func DiffSummary(diff string) string {
	files := ParseChangedFiles(diff)
	added, removed := 0, 0
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			added++
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			removed++
		}
	}
	return fmt.Sprintf("%d files, +%d/-%d lines", len(files), added, removed)
}

func truncateLines(s string, max int) string {
	lines := strings.Split(s, "\n")
	if len(lines) <= max {
		return s
	}
	truncated := strings.Join(lines[:max], "\n")
	return truncated + fmt.Sprintf("\n... (%d lines truncated)", len(lines)-max)
}
