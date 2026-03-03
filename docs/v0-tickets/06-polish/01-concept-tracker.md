# 06-01: Concept Tracker

## Summary
Extracts concepts and relationships from LLM responses via regex. Persists to `~/.pear/learning.json`.

## Event Model Refs
- E11: Concept Tracking (extract, save, display)
- E6c, E7, E8: Tracker.Extract called after every LLM response

## Files to Create
- `cli/learning/tracker.go`

## Types

```go
type ConceptStore struct {
    Concepts map[string]*Concept `json:"concepts"`
}

type Concept struct {
    Count    int      `json:"count"`
    Sessions []string `json:"sessions"` // ISO timestamps
    Related  []string `json:"related"`
}
```

## Functions
- `Load(path string) (*ConceptStore, error)` — read learning.json (create if missing)
- `Save(path string) error` — atomic write (tmp file + rename)
- `Extract(responseText string) ([]string, map[string][]string)` — regex extract concepts + relationships
- `Record(concepts []string, relationships map[string][]string)` — update counts, sessions, merge related
- `Display(w io.Writer)` — render progress bars to writer

## Regex Patterns
- Concepts: `📚 Concepts: \[(.+?)\]` → split on `, `
- Relationships: `🔗 Related: \[(.+?)\]` → parse `x → y` pairs

## Acceptance Criteria
- Extracts concepts from varied LLM output formats (brackets, no brackets, multiline)
- Counts increment correctly across multiple calls
- Relationships merge without duplicates
- Atomic save (no corruption on crash)
- Display shows sorted bars with relationship edges
- Handles missing/empty learning.json gracefully

## Dependencies
- None (pure data package)

## Notes
- Regex extraction is known-fragile (see PRD risk flags). Build tests with varied LLM output samples.
- Progress bar: `████████░░` using block characters, width proportional to count
