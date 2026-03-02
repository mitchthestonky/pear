package logging

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"
	"time"
)

type Logger struct {
	mu   sync.Mutex
	file *os.File
}

func NewLogger(logsDir string) (*Logger, error) {
	if err := os.MkdirAll(logsDir, 0755); err != nil {
		return nil, fmt.Errorf("create logs dir: %w", err)
	}

	rotate(logsDir, 5)

	name := time.Now().UTC().Format("2006-01-02T15-04-05Z") + ".log"
	f, err := os.Create(filepath.Join(logsDir, name))
	if err != nil {
		return nil, fmt.Errorf("create log file: %w", err)
	}

	return &Logger{file: f}, nil
}

func (l *Logger) Log(event string, fields map[string]any) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.file == nil {
		return
	}

	entry := make(map[string]any, len(fields)+2)
	entry["ts"] = time.Now().UTC().Format(time.RFC3339Nano)
	entry["event"] = event
	for k, v := range fields {
		entry[k] = v
	}

	data, err := json.Marshal(entry)
	if err != nil {
		return
	}
	l.file.Write(append(data, '\n'))
}

func (l *Logger) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}

func rotate(dir string, keep int) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	var logs []string
	for _, e := range entries {
		if !e.IsDir() && filepath.Ext(e.Name()) == ".log" {
			logs = append(logs, e.Name())
		}
	}

	if len(logs) < keep {
		return
	}

	sort.Strings(logs)
	for _, name := range logs[:len(logs)-keep+1] {
		os.Remove(filepath.Join(dir, name))
	}
}
