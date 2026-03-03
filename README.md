```
        ██
       ████
      ██████
     ████████
    ██████████
    ██████████
     ████████
      ██████
```

# pear

**AI makes you fast. Pear makes you sharp.**

A CLI teaching tool that watches you code and proactively teaches during natural pauses. Pear doesn't write code for you — it helps you understand what you're writing and why.

## Install

**Go:**

```sh
go install github.com/MitchTheStonky/pear/cli@latest
```

**Homebrew:**

```sh
brew install MitchTheStonky/pear/pear
```

**curl:**

```sh
curl -fsSL https://raw.githubusercontent.com/MitchTheStonky/pear/main/install.sh | sh
```

## Quick Start

```sh
# Set up your config and API key
pear init

# Start watching — Pear teaches as you code
pear watch
```

## How It Works

1. **You code.** Use whatever editor and AI tools you want — Cursor, Claude Code, Copilot.
2. **Pear watches.** It reads your diffs, file tree, and recent changes in the background.
3. **You understand.** During natural pauses, Pear surfaces what matters.

Pear doesn't write code. It makes sure you understand the code being written.

## Usage

### `pear watch`

Interactive TUI that watches your files. When you pause, Pear reviews your recent changes and teaches you something relevant — patterns, gotchas, better approaches.

### `pear ask "question"`

Ask a question about your codebase. Pear reads your current context (git diff, file tree) and answers with teaching intent.

### `pear review`

Review your recent code changes. Pear analyzes your uncommitted diff and gives feedback focused on learning, not just linting.

### `pear teach [topic]`

Deep-dive teaching on a topic, grounded in your actual code. Optionally specify a topic or let Pear pick based on what you've been working on.

### `pear doctor`

Check system health — verifies your config, API keys, and provider connectivity.

### `pear hooks install|uninstall`

Install a post-commit git hook that triggers Pear to review each commit.

### `pear progress`

Show your learning progress across sessions.

## Commands

| Command | Description |
|---------|-------------|
| `pear init` | Initialize configuration |
| `pear watch` | Watch files and teach proactively |
| `pear ask "q"` | Ask a question |
| `pear review` | Review recent changes |
| `pear teach` | Deep-dive on a topic |
| `pear doctor` | Check system health |
| `pear hooks install` | Install git hook |
| `pear hooks uninstall` | Remove git hook |
| `pear progress` | Show learning progress |

### TUI Slash Commands

| Command | Description |
|---------|-------------|
| `/help` | Show all commands |
| `/watch` | Start file watcher |
| `/review` | Review current changes |
| `/settings` | Configure provider & model |
| `/status` | Session info |
| `/copy` | Copy last response |
| `/export` | Export conversation |
| `/clear` | Reset conversation |
| `/quit` | Exit |

## Configuration

All config lives in `~/.pear/`:

```
~/.pear/
├── config.toml          # Provider, model, preferences
├── learning.json         # Learning progress
├── codebases/<slug>.toml # Per-repo overrides
└── logs/<timestamp>.log  # Session logs
```

## Providers

| Provider | Default Model | API Key Env |
|----------|--------------|-------------|
| Anthropic | `claude-sonnet-4-20250514` | `ANTHROPIC_API_KEY` |
| OpenAI | `gpt-4o` | `OPENAI_API_KEY` |
| OpenRouter | `anthropic/claude-sonnet-4-20250514` | `OPENROUTER_API_KEY` |

Set your provider during `pear init` or edit `~/.pear/config.toml` directly.

## Built With

- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — TUI framework
- [Glamour](https://github.com/charmbracelet/glamour) — Markdown rendering
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — TUI styling
- [fsnotify](https://github.com/fsnotify/fsnotify) — File watching

All LLM clients are hand-rolled with `net/http` for streaming control. No external SDKs.

## Contributing

Contributions welcome. Please open an issue first to discuss what you'd like to change.

## License

MIT

## Links

- **Website:** [pearcode.dev](https://pearcode.dev)
- **Docs:** [pearcode.dev/docs](https://pearcode.dev/docs)
- **GitHub:** [github.com/MitchTheStonky/pear](https://github.com/MitchTheStonky/pear)
