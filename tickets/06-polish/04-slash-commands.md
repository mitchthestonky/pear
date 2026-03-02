# 06-04: Slash Commands

## Summary
Implement all slash commands in the TUI.

## Event Model Refs
- E9: Slash Commands (full table)

## Files to Edit
- `cli/tui/app.go` — add slash command handler

## Commands

| Command | Implementation |
|---|---|
| `/help` | Render command table in output viewport |
| `/clear` | Clear history + viewport, print confirmation |
| `/exit` | Trigger session end (E12) |
| `/pause` | Set `paused = true`, render confirmation (watch only) |
| `/resume` | Set `paused = false`, render confirmation (watch only) |
| `/status` | Render: uptime, reviews, concepts, provider, model, branch |
| `/settings` | Render numbered config → user types number → edit field → config.Save → LLM.Reinit if needed |
| `/provider` | Render provider picker → select → prompt API key if new → config.Save → LLM.Reinit |
| `/model <name>` | Update model in config → config.Save → LLM.Reinit → confirm |
| `/key` | Prompt new API key → config.Save → LLM.Reinit → confirm |

## /settings Flow (Display + Quick Actions)
```
🍐 Current configuration:

  1. Name:       Mitch
  2. Languages:  Go, TypeScript, React
  3. Level:      mid
  4. Provider:   openrouter
  5. Model:      anthropic/claude-3.5-sonnet
  6. API Key:    sk-or-...****

  Watch settings:
  7. Settle time:    30s
  8. Min diff lines: 5
  9. Cooldown:       120s

  Enter a number to edit, or press Esc to close.
```

User types a number → inline edit for that field → save.

## Acceptance Criteria
- All 10 commands work in both interactive and watch mode
- `/pause` and `/resume` only functional in watch mode (show "only available in watch mode" in interactive)
- `/settings` numbered editor works end-to-end
- Provider/model/key changes take effect immediately (LLM client re-initialized)
- Unknown commands: "Unknown command. Type /help for available commands."

## Dependencies
- 04-01 (TUI app shell)
- 01-02 (config read/write)
- 02-01 (LLM NewClient for reinit)
