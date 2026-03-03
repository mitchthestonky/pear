# 04-04: Styles & Theme

## Summary
Lipgloss styles for all TUI elements. Single source of truth for colors, borders, separators.

## Files to Create
- `cli/tui/styles.go`

## Styles to Define
- `HeaderStyle` — session header (`🍐 Pear v0 · watching · ...`)
- `TriggerStyle` — proactive trigger notification (`🍐 Pear noticed...`)
- `ContextStyle` — context line (`📎 Context: ...`) — dimmed/muted
- `SeparatorOpen` — `━━━ Pear ━━━`
- `SeparatorClose` — `━━━━━━━━━━━━`
- `ConceptStyle` — concept tags (`📚 Concepts: [...]`)
- `RelatedStyle` — relationship tags (`🔗 Related: [...]`)
- `QuestionStyle` — Socratic question (`🤔 ...`)
- `ErrorStyle` — inline errors (`⚠ ...`) — red/warning
- `SystemStyle` — system messages (slash command output) — dimmed
- `InputPromptStyle` — `> ` prompt
- `AutocompleteStyle` — dropdown box for @file matches
- `ProgressBarFull` / `ProgressBarEmpty` — for `pear progress` bars

## Acceptance Criteria
- All styles defined as lipgloss.Style constants/vars
- Consistent color palette across all elements
- Separators are full terminal width
- Looks good on both light and dark terminal backgrounds

## Dependencies
- Dep: `lipgloss`

## Notes
- Keep it simple — don't over-design. Pear's brand color is green (🍐).
- Test on at least iTerm2 and default Terminal.app
