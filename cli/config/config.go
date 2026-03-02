package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Name      string         `toml:"name"`
	Languages string         `toml:"languages"`
	Level     string         `toml:"level"`
	Provider  ProviderConfig `toml:"provider"`
	Watch     WatchConfig    `toml:"watch"`
}

type ProviderConfig struct {
	Active     string         `toml:"active"`
	Anthropic  ProviderDetail `toml:"anthropic"`
	OpenAI     ProviderDetail `toml:"openai"`
	OpenRouter ProviderDetail `toml:"openrouter"`
}

type ProviderDetail struct {
	APIKey string `toml:"api_key"`
	Model  string `toml:"model"`
}

type WatchConfig struct {
	SettleTime   int `toml:"settle_time"`
	Interval     int `toml:"interval"`
	MinDiffLines int `toml:"min_diff_lines"`
	Cooldown     int `toml:"cooldown"`
}

func Dir() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".pear")
}

func configPath() string {
	return filepath.Join(Dir(), "config.toml")
}

func Exists() bool {
	_, err := os.Stat(configPath())
	return err == nil
}

func Load() (*Config, error) {
	cfg := &Config{}
	_, err := toml.DecodeFile(configPath(), cfg)
	if err != nil {
		return nil, fmt.Errorf("loading config: %w", err)
	}
	applyDefaults(cfg)
	return cfg, nil
}

func Save(cfg *Config) error {
	dir := Dir()
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating config dir: %w", err)
	}

	tmp := configPath() + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return fmt.Errorf("creating temp config: %w", err)
	}

	if err := toml.NewEncoder(f).Encode(cfg); err != nil {
		f.Close()
		os.Remove(tmp)
		return fmt.Errorf("encoding config: %w", err)
	}
	f.Close()

	if err := os.Rename(tmp, configPath()); err != nil {
		return fmt.Errorf("writing config: %w", err)
	}
	return nil
}

func ResolveCodebase(cfg *Config) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	slug := pathSlug(wd)
	codebasesDir := filepath.Join(Dir(), "codebases")
	if err := os.MkdirAll(codebasesDir, 0755); err != nil {
		return err
	}

	cbPath := filepath.Join(codebasesDir, slug+".toml")
	if _, err := os.Stat(cbPath); os.IsNotExist(err) {
		return nil
	}

	override := &Config{}
	if _, err := toml.DecodeFile(cbPath, override); err != nil {
		return fmt.Errorf("loading codebase config: %w", err)
	}

	mergeOverrides(cfg, override)
	return nil
}

func ActiveProvider(cfg *Config) *ProviderDetail {
	switch cfg.Provider.Active {
	case "anthropic":
		return &cfg.Provider.Anthropic
	case "openai":
		return &cfg.Provider.OpenAI
	case "openrouter":
		return &cfg.Provider.OpenRouter
	default:
		return &cfg.Provider.Anthropic
	}
}

func SetModel(cfg *Config, model string) {
	p := ActiveProvider(cfg)
	p.Model = model
}

func SetKey(cfg *Config, key string) {
	p := ActiveProvider(cfg)
	p.APIKey = key
}

func applyDefaults(cfg *Config) {
	if cfg.Watch.SettleTime == 0 {
		cfg.Watch.SettleTime = 30
	}
	if cfg.Watch.Interval == 0 {
		cfg.Watch.Interval = 5
	}
	if cfg.Watch.MinDiffLines == 0 {
		cfg.Watch.MinDiffLines = 5
	}
	if cfg.Watch.Cooldown == 0 {
		cfg.Watch.Cooldown = 120
	}
}

func pathSlug(path string) string {
	path = strings.TrimPrefix(path, "/")
	return strings.ReplaceAll(path, "/", "-")
}

func mergeOverrides(base, override *Config) {
	if override.Name != "" {
		base.Name = override.Name
	}
	if override.Languages != "" {
		base.Languages = override.Languages
	}
	if override.Level != "" {
		base.Level = override.Level
	}
	if override.Provider.Active != "" {
		base.Provider.Active = override.Provider.Active
	}
	if override.Watch.SettleTime != 0 {
		base.Watch.SettleTime = override.Watch.SettleTime
	}
	if override.Watch.Interval != 0 {
		base.Watch.Interval = override.Watch.Interval
	}
	if override.Watch.MinDiffLines != 0 {
		base.Watch.MinDiffLines = override.Watch.MinDiffLines
	}
	if override.Watch.Cooldown != 0 {
		base.Watch.Cooldown = override.Watch.Cooldown
	}
}
