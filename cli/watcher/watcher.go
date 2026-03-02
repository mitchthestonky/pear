package watcher

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/pearcode/pear/config"
	"github.com/pearcode/pear/logging"
)

// ReviewTrigger represents a watcher-generated review trigger.
type ReviewTrigger struct {
	Type    string // "settle" or "commit"
	Diff    string
	Summary string
}

// Watcher implements hybrid file watching: fsnotify for instant change
// detection, git polling for diff content and commit detection.
type Watcher struct {
	settleTime  time.Duration
	minDiffSize int
	cooldown    time.Duration

	lastChangeTime time.Time
	lastReviewTime time.Time
	lastHEAD       string
	lastReviewDiff string
	settled        bool

	fsWatcher *fsnotify.Watcher
	triggers  chan ReviewTrigger
	repoRoot  string
	logger    *logging.Logger
}

// New creates a new Watcher. It initializes fsnotify, sets baseline HEAD and diff.
func New(cfg config.WatchConfig, repoRoot string, logger *logging.Logger) (*Watcher, error) {
	fsw, err := fsnotify.NewWatcher()
	if err != nil {
		// Fallback: log warning, continue without fsnotify
		if logger != nil {
			logger.Log("watcher.fsnotify_fail", map[string]any{"error": err.Error()})
		}
	}

	w := &Watcher{
		settleTime:  time.Duration(cfg.SettleTime) * time.Second,
		minDiffSize: cfg.MinDiffLines,
		cooldown:    time.Duration(cfg.Cooldown) * time.Second,
		triggers:    make(chan ReviewTrigger, 1),
		repoRoot:    repoRoot,
		fsWatcher:   fsw,
		settled:     true, // start settled, no pending changes
		logger:      logger,
	}

	// Set baseline HEAD
	head, err := w.gitRevParseHEAD()
	if err == nil {
		w.lastHEAD = head
	}

	// Set baseline diff
	diff, err := w.gitDiffHEAD()
	if err == nil {
		w.lastReviewDiff = diff
	}

	// Add repo root directories recursively to fsnotify
	if fsw != nil {
		w.addWatchDirs(repoRoot)
	}

	return w, nil
}

// BaselineDiff returns the diff at watcher init time (for dirty diff prompt).
func (w *Watcher) BaselineDiff() string {
	return w.lastReviewDiff
}

// Start launches the watcher goroutine and returns the triggers channel.
func (w *Watcher) Start(ctx context.Context) <-chan ReviewTrigger {
	go w.run(ctx)
	return w.triggers
}

func (w *Watcher) run(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	defer close(w.triggers)

	if w.fsWatcher != nil {
		defer w.fsWatcher.Close()
	}

	for {
		select {
		case <-ctx.Done():
			return

		case event, ok := <-w.fsEvents():
			if !ok {
				return
			}
			if w.shouldIgnorePath(event.Name) {
				continue
			}
			if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) {
				w.lastChangeTime = time.Now()
				w.settled = false
			}
			// Watch new directories
			if event.Has(fsnotify.Create) {
				if info, err := os.Stat(event.Name); err == nil && info.IsDir() {
					if !w.shouldIgnorePath(event.Name) {
						w.fsWatcher.Add(event.Name)
					}
				}
			}

		case _, ok := <-w.fsErrors():
			if !ok {
				return
			}

		case <-ticker.C:
			w.checkCommit(ctx)
			w.checkSettle(ctx)
		}
	}
}

// fsEvents returns the fsnotify events channel, or a nil channel if no watcher.
func (w *Watcher) fsEvents() <-chan fsnotify.Event {
	if w.fsWatcher == nil {
		return nil
	}
	return w.fsWatcher.Events
}

func (w *Watcher) fsErrors() <-chan error {
	if w.fsWatcher == nil {
		return nil
	}
	return w.fsWatcher.Errors
}

func (w *Watcher) checkCommit(ctx context.Context) {
	head, err := w.gitRevParseHEAD()
	if err != nil {
		return
	}
	if head == w.lastHEAD {
		return
	}

	oldHEAD := w.lastHEAD
	w.lastHEAD = head
	w.lastReviewDiff = "" // reset baseline

	commitDiff, err := w.gitDiffRange(oldHEAD, head)
	if err != nil {
		return
	}

	commitMsg, err := w.gitLogOneline(head)
	if err != nil {
		commitMsg = head[:8]
	}

	if w.logger != nil {
		w.logger.Log("watcher.trigger", map[string]any{
			"type":       "commit",
			"diff_lines": countLines(commitDiff),
		})
	}

	select {
	case w.triggers <- ReviewTrigger{Type: "commit", Diff: commitDiff, Summary: commitMsg}:
	default:
		// channel full, drop
	}
}

func (w *Watcher) checkSettle(ctx context.Context) {
	if w.settled {
		return
	}
	if w.lastChangeTime.IsZero() {
		return
	}
	if time.Since(w.lastChangeTime) < w.settleTime {
		return
	}

	// Get current diff
	currentDiff, err := w.gitDiffHEAD()
	if err != nil {
		return
	}

	// Subtract last reviewed diff
	newDiff := subtractDiff(currentDiff, w.lastReviewDiff)

	lines := countLines(newDiff)
	if lines < w.minDiffSize {
		if w.logger != nil {
			w.logger.Log("watcher.skip", map[string]any{"reason": "too_small", "lines": lines})
		}
		w.settled = true
		return
	}

	if !w.lastReviewTime.IsZero() && time.Since(w.lastReviewTime) < w.cooldown {
		if w.logger != nil {
			w.logger.Log("watcher.skip", map[string]any{"reason": "cooldown"})
		}
		return
	}

	w.settled = true
	w.lastReviewDiff = currentDiff
	w.lastReviewTime = time.Now()

	summary := diffSummary(newDiff)

	if w.logger != nil {
		w.logger.Log("watcher.trigger", map[string]any{
			"type":       "settle",
			"diff_lines": lines,
			"summary":    summary,
		})
	}

	select {
	case w.triggers <- ReviewTrigger{Type: "settle", Diff: newDiff, Summary: summary}:
	default:
	}
}

// addWatchDirs recursively adds directories to fsnotify, skipping ignored dirs.
func (w *Watcher) addWatchDirs(root string) {
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() {
			return nil
		}
		if w.shouldIgnorePath(path) {
			return filepath.SkipDir
		}
		if err := w.fsWatcher.Add(path); err != nil {
			if w.logger != nil {
				w.logger.Log("watcher.watch_add_fail", map[string]any{"path": path, "error": err.Error()})
			}
		}
		return nil
	})
}

func (w *Watcher) shouldIgnorePath(path string) bool {
	base := filepath.Base(path)
	if base == ".git" || base == "node_modules" {
		return true
	}
	if strings.HasPrefix(base, ".") && base != "." {
		return true
	}
	return false
}

// Git helpers

func (w *Watcher) gitRevParseHEAD() (string, error) {
	cmd := exec.Command("git", "rev-parse", "HEAD")
	cmd.Dir = w.repoRoot
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

func (w *Watcher) gitDiffHEAD() (string, error) {
	cmd := exec.Command("git", "diff", "HEAD")
	cmd.Dir = w.repoRoot
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (w *Watcher) gitDiffRange(from, to string) (string, error) {
	cmd := exec.Command("git", "diff", from+".."+to)
	cmd.Dir = w.repoRoot
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}

func (w *Watcher) gitLogOneline(ref string) (string, error) {
	cmd := exec.Command("git", "log", "-1", "--format=%s", ref)
	cmd.Dir = w.repoRoot
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}

// Helpers

func countLines(s string) int {
	if s == "" {
		return 0
	}
	return strings.Count(s, "\n")
}

func subtractDiff(current, previous string) string {
	if previous == "" {
		return current
	}
	if current == previous {
		return ""
	}

	// Split into per-file chunks keyed by diff header
	prevChunks := splitDiffChunks(previous)
	curChunks := splitDiffChunks(current)

	var result []string
	for _, chunk := range curChunks {
		header := chunkHeader(chunk)
		if prev, ok := prevChunks[header]; ok && prev == chunk {
			continue // identical chunk already reviewed
		}
		result = append(result, chunk)
	}

	if len(result) == 0 {
		return ""
	}
	return strings.Join(result, "")
}

func splitDiffChunks(diff string) map[string]string {
	chunks := make(map[string]string)
	parts := strings.Split(diff, "\ndiff --git ")
	for i, part := range parts {
		if i == 0 {
			if strings.HasPrefix(part, "diff --git ") {
				part = part[len("diff --git "):]
			} else {
				continue
			}
		}
		full := "diff --git " + part
		header := chunkHeader(full)
		chunks[header] = full
	}
	return chunks
}

func chunkHeader(chunk string) string {
	if idx := strings.Index(chunk, "\n"); idx != -1 {
		return chunk[:idx]
	}
	return chunk
}

func diffSummary(diff string) string {
	files := make(map[string]bool)
	added := 0
	for _, line := range strings.Split(diff, "\n") {
		if strings.HasPrefix(line, "diff --git a/") {
			parts := strings.Fields(line)
			if len(parts) >= 4 {
				files[strings.TrimPrefix(parts[2], "a/")] = true
			}
		}
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			added++
		}
	}
	return fmt.Sprintf("%d files, +%d lines", len(files), added)
}
