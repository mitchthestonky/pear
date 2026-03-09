package learning

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ConceptStore struct {
	Concepts map[string]*Concept `json:"concepts"`
}

type Concept struct {
	Count    int      `json:"count"`
	Sessions []string `json:"sessions"`
	Related  []string `json:"related"`
}

// CoveredEntry tracks a concept and the angle covered this session.
type CoveredEntry struct {
	Concept string
	Summary string
}

// SessionMemory holds ephemeral state for the current session.
type SessionMemory struct {
	Covered []CoveredEntry
}

// AddCovered records a covered concept, deduplicating by concept name (updates summary if seen again).
func (sm *SessionMemory) AddCovered(concept, summary string) {
	for i, e := range sm.Covered {
		if strings.EqualFold(e.Concept, concept) {
			sm.Covered[i].Summary = summary
			return
		}
	}
	sm.Covered = append(sm.Covered, CoveredEntry{Concept: concept, Summary: summary})
}

// FormatCovered returns the "already covered" block for prompt injection. Empty string if nothing covered.
func (sm *SessionMemory) FormatCovered() string {
	if sm == nil || len(sm.Covered) == 0 {
		return ""
	}
	var b strings.Builder
	b.WriteString("Already covered this session (do not repeat the same angle):\n")
	for _, e := range sm.Covered {
		fmt.Fprintf(&b, "- %s: %s\n", e.Concept, e.Summary)
	}
	return b.String()
}

var titleCaser = cases.Title(language.English)

var (
	// Bracketed format: 📚 Concepts: [A, B, C]
	conceptsBracketRe = regexp.MustCompile(`📚\s*Concepts:\s*\[(.+?)\]`)
	relBracketRe      = regexp.MustCompile(`🔗\s*Related:\s*\[(.+?)\]`)
	coveredRe         = regexp.MustCompile(`📝\s*Covered:\s*\[(.+?)\]`)

	// Inline format without brackets: 📚 Concepts: A, B, C (on same line)
	conceptsInlineRe = regexp.MustCompile(`(?m)^📚\s*Concepts:\s*(.+)$`)
	relInlineRe      = regexp.MustCompile(`(?m)^🔗\s*Related:\s*(.+)$`)

	tagLineRe = regexp.MustCompile(`(?m)^.*(?:📚\s*Concepts:|🔗\s*Related:|📝\s*Covered:).*$\n?`)
)

// StripTags removes concept and related tag lines from response text.
func StripTags(text string) string {
	return strings.TrimRight(tagLineRe.ReplaceAllString(text, ""), "\n")
}

func Load(path string) (*ConceptStore, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return &ConceptStore{Concepts: make(map[string]*Concept)}, nil
		}
		return nil, fmt.Errorf("read learning.json: %w", err)
	}

	if len(data) == 0 {
		return &ConceptStore{Concepts: make(map[string]*Concept)}, nil
	}

	var store ConceptStore
	if err := json.Unmarshal(data, &store); err != nil {
		return nil, fmt.Errorf("parse learning.json: %w", err)
	}
	if store.Concepts == nil {
		store.Concepts = make(map[string]*Concept)
	}
	return &store, nil
}

func (s *ConceptStore) Save(path string) error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal learning.json: %w", err)
	}

	dir := filepath.Dir(path)
	tmp, err := os.CreateTemp(dir, "learning-*.tmp")
	if err != nil {
		return fmt.Errorf("create temp file: %w", err)
	}
	tmpName := tmp.Name()

	if _, err := tmp.Write(data); err != nil {
		tmp.Close()
		os.Remove(tmpName)
		return fmt.Errorf("write temp file: %w", err)
	}
	if err := tmp.Close(); err != nil {
		os.Remove(tmpName)
		return err
	}

	return os.Rename(tmpName, path)
}

func Extract(responseText string) ([]string, map[string][]string, []CoveredEntry) {
	var concepts []string
	relationships := make(map[string][]string)
	var covered []CoveredEntry

	concepts = extractConceptList(responseText, conceptsBracketRe, conceptsInlineRe)
	relationships = extractRelationships(responseText, relBracketRe, relInlineRe)

	for _, match := range coveredRe.FindAllStringSubmatch(responseText, -1) {
		for _, entry := range strings.Split(match[1], "; ") {
			parts := strings.SplitN(entry, ": ", 2)
			if len(parts) == 2 {
				concept := strings.TrimSpace(parts[0])
				summary := strings.TrimSpace(parts[1])
				if concept != "" && summary != "" {
					covered = append(covered, CoveredEntry{Concept: normalizeConcept(concept), Summary: summary})
				}
			}
		}
	}

	return concepts, relationships, covered
}

func (s *ConceptStore) Record(concepts []string, relationships map[string][]string) {
	session := time.Now().UTC().Format(time.RFC3339)

	for _, name := range concepts {
		c, ok := s.Concepts[name]
		if !ok {
			c = &Concept{}
			s.Concepts[name] = c
		}
		c.Count++
		c.Sessions = append(c.Sessions, session)
	}

	for from, tos := range relationships {
		c, ok := s.Concepts[from]
		if !ok {
			c = &Concept{}
			s.Concepts[from] = c
		}
		for _, to := range tos {
			if !contains(c.Related, to) {
				c.Related = append(c.Related, to)
			}
		}
	}
}

func (s *ConceptStore) Display(w io.Writer) {
	if len(s.Concepts) == 0 {
		fmt.Fprintln(w, "No concepts tracked yet. Start a session with `pear watch`.")
		return
	}

	type entry struct {
		name    string
		concept *Concept
	}

	var entries []entry
	for name, c := range s.Concepts {
		entries = append(entries, entry{name, c})
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].concept.Count > entries[j].concept.Count
	})

	maxCount := entries[0].concept.Count
	totalSessions := uniqueSessions(s)

	fmt.Fprintln(w, "🍐 Concepts Pear has taught you:")
	fmt.Fprintln(w)

	for _, e := range entries {
		bar := progressBar(e.concept.Count, maxCount, 10)
		fmt.Fprintf(w, "  %-20s %s  %d sessions\n", e.name, bar, e.concept.Count)
		if len(e.concept.Related) > 0 {
			fmt.Fprintf(w, "    → %s\n", strings.Join(e.concept.Related, ", "))
		}
	}

	fmt.Fprintf(w, "\n  %d concepts across %d sessions\n", len(s.Concepts), totalSessions)
}

func progressBar(count, max, width int) string {
	filled := (count * width) / max
	if filled == 0 && count > 0 {
		filled = 1
	}
	return strings.Repeat("█", filled) + strings.Repeat("░", width-filled)
}

func uniqueSessions(s *ConceptStore) int {
	seen := make(map[string]struct{})
	for _, c := range s.Concepts {
		for _, sess := range c.Sessions {
			seen[sess] = struct{}{}
		}
	}
	return len(seen)
}

// extractConceptList parses concepts from bracketed, inline, or newline-separated formats.
func extractConceptList(text string, bracketRe, inlineRe *regexp.Regexp) []string {
	var concepts []string

	// Try bracketed: 📚 Concepts: [A, B, C]
	if matches := bracketRe.FindAllStringSubmatch(text, -1); len(matches) > 0 {
		for _, match := range matches {
			for _, c := range strings.Split(match[1], ", ") {
				c = strings.TrimSpace(c)
				if c != "" {
					concepts = append(concepts, normalizeConcept(c))
				}
			}
		}
		return concepts
	}

	// Try inline (same line, with or without brackets): 📚 Concepts: A, B, C
	if matches := inlineRe.FindAllStringSubmatch(text, -1); len(matches) > 0 {
		for _, match := range matches {
			inner := strings.TrimSpace(match[1])
			inner = strings.TrimPrefix(inner, "[")
			inner = strings.TrimSuffix(inner, "]")
			if inner == "" {
				continue
			}
			for _, c := range strings.Split(inner, ", ") {
				c = strings.TrimSpace(c)
				if c != "" {
					concepts = append(concepts, normalizeConcept(c))
				}
			}
		}
		return concepts
	}

	// Try newline-separated: 📚 Concepts:\nA\nB\nC
	lines := strings.Split(text, "\n")
	inBlock := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "📚") && strings.Contains(trimmed, "Concepts") {
			// Check if concepts are on the same line after the colon
			if idx := strings.Index(trimmed, ":"); idx >= 0 {
				after := strings.TrimSpace(trimmed[idx+1:])
				if after != "" {
					// Same line, already handled above
					break
				}
			}
			inBlock = true
			continue
		}
		if inBlock {
			if trimmed == "" || strings.Contains(trimmed, "🔗") || strings.Contains(trimmed, "📝") {
				break
			}
			c := strings.TrimLeft(trimmed, "- •*")
			c = strings.TrimSpace(c)
			if c != "" {
				concepts = append(concepts, normalizeConcept(c))
			}
		}
	}

	return concepts
}

// extractRelationships parses relationships from bracketed, inline, or newline-separated formats.
func extractRelationships(text string, bracketRe, inlineRe *regexp.Regexp) map[string][]string {
	relationships := make(map[string][]string)

	parseRelPairs := func(inner string) {
		for _, pair := range strings.Split(inner, ", ") {
			parts := strings.SplitN(pair, " → ", 2)
			if len(parts) == 2 {
				from := strings.TrimSpace(parts[0])
				to := strings.TrimSpace(parts[1])
				if from != "" && to != "" {
					relationships[normalizeConcept(from)] = append(relationships[normalizeConcept(from)], normalizeConcept(to))
				}
			}
		}
	}

	// Try bracketed
	if matches := bracketRe.FindAllStringSubmatch(text, -1); len(matches) > 0 {
		for _, match := range matches {
			parseRelPairs(match[1])
		}
		return relationships
	}

	// Try inline
	if matches := inlineRe.FindAllStringSubmatch(text, -1); len(matches) > 0 {
		for _, match := range matches {
			inner := strings.TrimSpace(match[1])
			inner = strings.TrimPrefix(inner, "[")
			inner = strings.TrimSuffix(inner, "]")
			parseRelPairs(inner)
		}
		return relationships
	}

	// Try newline-separated
	lines := strings.Split(text, "\n")
	inBlock := false
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.Contains(trimmed, "🔗") && strings.Contains(trimmed, "Related") {
			if idx := strings.Index(trimmed, ":"); idx >= 0 {
				after := strings.TrimSpace(trimmed[idx+1:])
				if after != "" {
					break
				}
			}
			inBlock = true
			continue
		}
		if inBlock {
			if trimmed == "" || strings.Contains(trimmed, "📝") || strings.Contains(trimmed, "📚") {
				break
			}
			pair := strings.TrimLeft(trimmed, "- •*")
			pair = strings.TrimSpace(pair)
			parts := strings.SplitN(pair, " → ", 2)
			if len(parts) == 2 {
				from := strings.TrimSpace(parts[0])
				to := strings.TrimSpace(parts[1])
				if from != "" && to != "" {
					relationships[normalizeConcept(from)] = append(relationships[normalizeConcept(from)], normalizeConcept(to))
				}
			}
		}
	}

	return relationships
}

// normalizeConcept normalizes concept names to title case for consistent lookups.
func normalizeConcept(s string) string {
	return titleCaser.String(strings.ToLower(s))
}

func contains(slice []string, s string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
