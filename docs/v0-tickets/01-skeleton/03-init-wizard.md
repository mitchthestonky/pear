# 01-03: Init Wizard

## Summary
First-run setup wizard. Blocks all commands until complete. Collects name, languages, level, provider, API key, model. Auto-runs doctor at end.

## Event Model Refs
- E2: Init Wizard (full flow)
- E1: Config.Load [not found] → BLOCK → RunInitWizard

## Files to Edit
- `cli/cmd/init_cmd.go` — replace stub with wizard
- `cli/cmd/root.go` — add pre-run hook: if `!config.Exists()` → run wizard

## Behavior
1. Print `🍐 Welcome to Pear! Let's get you set up.`
2. Prompt each field (all required, no skipping):
   - Name → string
   - Languages → string
   - Level → select: junior / mid / senior
   - Provider → select: 1. Anthropic  2. OpenAI  3. OpenRouter
   - API key → string (masked input)
   - Model → string with provider-specific default
3. Write config via `config.Save()`
4. Auto-run doctor (ticket 01-04)
5. If doctor passes → print ready message, return to original command
6. If doctor fails → show failure, prompt to fix, re-run doctor

## Pre-Run Hook
Add to root command's `PersistentPreRun`: check `config.Exists()`. If false, run init wizard. This gates ALL commands except `init` itself.

## Acceptance Criteria
- Running any `pear` command without config drops into wizard
- All prompts are required — empty input re-prompts
- Config file written correctly to `~/.pear/config.toml`
- Doctor auto-runs after config saved
- After wizard, original command continues (e.g., `pear watch` starts watch mode)
- `pear init` can be run explicitly to re-configure

## Dependencies
- 01-02 (config package)
- 01-04 (doctor — called at end of wizard)

## Notes
- Use `bufio.Scanner` for input (Bubble Tea not needed here — wizard runs before TUI)
- Provider-specific defaults: anthropic → `claude-haiku-4-5`, openai → `gpt-4o`, openrouter → `anthropic/claude-3.5-sonnet`
