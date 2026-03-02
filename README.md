# pear

**A pair programmer that watches your code and tells you what matters.**

pear runs in your terminal alongside your AI coding tools. It watches your diffs, detects what you don't understand, and teaches you at the right moment — without breaking your flow.

AI makes you faster. pear makes you smarter.

## Features

- **Watch mode** — Monitors file changes and git diffs, surfaces insights during natural pauses
- **Interactive Q&A** — Ask questions with full codebase context (`pear ask`)
- **Code review** — On-demand review of your recent changes (`pear review`)
- **Guided teaching** — Deep-dive explanations adapted to your level (`pear teach`)
- **Concept tracking** — Tags and tracks engineering concepts as you encounter them
- **Learning state memory** *(Pro)* — Remembers what you understand across sessions
- **Adaptive pedagogy** *(Pro)* — Changes how it teaches based on your behavior
- **Skill progression** *(Pro)* — Shows growth over time across languages and frameworks
- **BYO any LLM** — Bring your own API key for Anthropic, OpenAI, or OpenRouter

## Quick Start

### Prerequisites

- Go 1.24+
- macOS (Linux planned)
- An API key from Anthropic, OpenAI, or OpenRouter

### Install from source

```bash
git clone https://github.com/pearcode/pear.git
cd pear/cli
go build -o pear .
```

### Set up

```bash
# Initialize config and add your API key
./pear init

# Start watching your project
cd /path/to/your/project
pear watch
```

## Usage

```bash
# Watch mode — monitors changes, teaches during pauses
pear watch

# Ask a question with codebase context
pear ask "what does this middleware do?"

# Review recent changes
pear review

# Deep-dive teaching on a topic
pear teach

# See your learning progress
pear progress

# Check your setup
pear doctor

# Install git hooks for post-commit reviews
pear hooks install
```

## How It Works

1. **You code.** Use whatever editor and AI tools you want — Cursor, Claude Code, OpenCode, Copilot.
2. **Pear watches.** It reads your diffs, file tree, and recent changes in the background.
3. **You understand.** During natural pauses — while your agent thinks, after a diff lands — pear surfaces what matters.

Pear doesn't write code. It makes sure you understand the code being written.

## Configuration

Config lives in `~/.pear/`:

```
~/.pear/
├── config.toml          # API keys, provider, preferences
├── learning.json        # Your learning state (Pro)
├── codebases/<slug>.toml  # Per-project context
└── logs/                # Session logs
```

### Supported providers

| Provider | Models | Config key |
|----------|--------|------------|
| Anthropic | Claude 4.5 Sonnet, Claude 4 Opus | `anthropic` |
| OpenAI | GPT-4o, o3 | `openai` |
| OpenRouter | Any supported model | `openrouter` |

## Architecture

Written in Go. Built with:

- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) — Terminal UI
- [Glamour](https://github.com/charmbracelet/glamour) — Markdown rendering
- [Lipgloss](https://github.com/charmbracelet/lipgloss) — TUI styling
- [fsnotify](https://github.com/fsnotify/fsnotify) — File system watching

All LLM clients are hand-rolled with `net/http` for streaming control. No external SDKs.

```
cli/
├── cmd/           # Cobra commands
├── watcher/       # fsnotify + git polling with pause detection
├── repocontext/   # Git diff, file tree, @file reading
├── prompt/        # System prompt assembly
├── llm/           # LLMClient interface + provider implementations
├── config/        # ~/.pear/ config management
├── learning/      # Concept extraction & learning state
├── hooks/         # Git hook install/uninstall
├── tui/           # Bubble Tea app
└── logging/       # Structured JSON logging
```

## Contributing

pear is open source. Contributions welcome.

```bash
# Clone and build
git clone https://github.com/pearcode/pear.git
cd pear/cli
go build -o pear .

# Run
./pear doctor  # Verify setup
./pear watch   # Start watching
```

Please open an issue before submitting large PRs.

## License

MIT

## Links

- **Website:** [pearcode.dev](https://pearcode.dev)
- **Email:** mitch@pearcode.dev
