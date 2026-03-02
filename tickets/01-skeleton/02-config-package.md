# 01-02: Config Package

## Summary
Config read/write for `~/.pear/config.toml` with provider sections and codebase override resolution.

## Event Model Refs
- E1: Config.Load, Config.ResolveCodebase
- E9: /settings, /provider, /model, /key all write config

## Files to Create
- `cli/config/config.go`

## Structs

```go
type Config struct {
    Name      string         `toml:"name"`
    Languages string         `toml:"languages"`
    Level     string         `toml:"level"`
    Provider  ProviderConfig `toml:"provider"`
    Watch     WatchConfig    `toml:"watch"`
}

type ProviderConfig struct {
    Active     string              `toml:"active"`
    Anthropic  ProviderDetail      `toml:"anthropic"`
    OpenAI     ProviderDetail      `toml:"openai"`
    OpenRouter ProviderDetail      `toml:"openrouter"`
}

type ProviderDetail struct {
    APIKey string `toml:"api_key"`
    Model  string `toml:"model"`
}

type WatchConfig struct {
    SettleTime   int `toml:"settle_time"`   // seconds, default 30
    Interval     int `toml:"interval"`      // seconds, default 5
    MinDiffLines int `toml:"min_diff_lines"` // default 5
    Cooldown     int `toml:"cooldown"`      // seconds, default 120
}
```

## Functions
- `Load() (*Config, error)` — read from `~/.pear/config.toml`
- `Save(cfg *Config) error` — write to `~/.pear/config.toml`
- `ResolveCodebase(cfg *Config) error` — detect git root, check `~/.pear/codebases/<slug>.toml`, merge overrides
- `Exists() bool` — check if config file exists
- `Dir() string` — returns `~/.pear/`
- `ActiveProvider(cfg *Config) *ProviderDetail` — returns the active provider's detail
- `SetModel(cfg *Config, model string)` / `SetKey(cfg *Config, key string)`

## Codebase Slug
Full path slugified: `/Users/mitch/Documents/Pear-v0` → `Users-mitch-Documents-Pear-v0.toml`

## Acceptance Criteria
- Round-trip test: write config → read config → assert equal
- `ResolveCodebase` merges overrides correctly (test with fixture)
- `ActiveProvider` returns correct provider detail
- Creates `~/.pear/` and `~/.pear/codebases/` directories if missing
- Defaults applied for WatchConfig when fields are zero-value

## Dependencies
- go.mod must exist (01-01)
- Dep: `github.com/BurntSushi/toml`
