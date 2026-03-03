# 04-02: Input Component

## Summary
Text input with @file live dropdown autocomplete and slash command detection.

## Event Model Refs
- E7: User Input — parse input, resolve @files
- E7 @file Autocomplete Sub-Flow

## Files to Create
- `cli/tui/input.go`

## Model

```go
type InputModel struct {
    textinput    textinput.Model   // charmbracelet/bubbles
    autocomplete AutocompleteModel
    enabled      bool
    fileCache    []string          // git ls-files, cached per session
}

type AutocompleteModel struct {
    active   bool
    prefix   string     // text after @
    matches  []string   // filtered file list
    selected int        // cursor index
}
```

## Behavior
- Always visible at bottom of screen with `> ` prompt
- Disabled during streaming (keystrokes buffered by Bubble Tea, not sent to input)
- `@` keystroke: activate autocomplete, load file cache if empty (`git ls-files`)
- Each keystroke after `@`: filter matches, render dropdown above input
- Arrow keys: navigate dropdown
- Enter in autocomplete: insert selected path, deactivate autocomplete
- Esc in autocomplete: cancel, deactivate
- Enter (no autocomplete): submit input → return `SubmitMsg{Text, Files}`
- Slash detection: if input starts with `/`, return `SlashMsg{Command, Args}`

## Output Messages
- `SubmitMsg{Text string, Files map[string]string}` — user submitted input with resolved @files
- `SlashMsg{Command string, Args string}` — slash command detected

## Acceptance Criteria
- Text input renders with `> ` prompt
- @file typing shows live dropdown of matching files
- Arrow keys navigate, Enter selects, Esc cancels
- Submitted input has @file paths resolved to contents
- Slash commands detected and routed separately
- Input disables/enables correctly

## Dependencies
- Dep: `github.com/charmbracelet/bubbles/textinput`
