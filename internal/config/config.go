package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/SShogun/ClawCLI/internal/types"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// Default configuration values
const (
	DefaultProvider    = "ollama"
	DefaultBaseURL     = "http://localhost:11434"
	DefaultModel       = "qwen2.5-coder"
	DefaultTemperature = 0.7
	DefaultMaxTokens   = 4096
)

// Init initializes viper to read from .env file
func Init() error {
	// Load .env file from current working directory
	wd, err := os.Getwd()
	if err == nil {
		envPath := filepath.Join(wd, ".env")
		_ = godotenv.Load(envPath) // Ignore error if file doesn't exist
	}

	// Set viper to read environment variables with CLAW_ prefix
	viper.SetEnvPrefix("CLAW")
	viper.AutomaticEnv()

	return nil
}

// Load reads configuration from viper and returns a Config struct
func Load() (*types.Config, error) {
	cfg := &types.Config{
		Provider:    viper.GetString("PROVIDER"),
		BaseURL:     viper.GetString("BASE_URL"),
		APIKey:      viper.GetString("API_KEY"),
		Model:       viper.GetString("MODEL"),
		Temperature: viper.GetFloat64("TEMPERATURE"),
		MaxTokens:   viper.GetInt("MAX_TOKENS"),
		Verbose:     viper.GetBool("VERBOSE"),
	}

	// Apply defaults if not set
	if cfg.Provider == "" {
		cfg.Provider = DefaultProvider
	}

	if cfg.BaseURL == "" {
		cfg.BaseURL = DefaultBaseURL
	}

	if cfg.Model == "" {
		cfg.Model = DefaultModel
	}

	if cfg.Temperature == 0 {
		cfg.Temperature = DefaultTemperature
	}

	if cfg.MaxTokens == 0 {
		cfg.MaxTokens = DefaultMaxTokens
	}

	// Only require API key for anthropic provider
	if cfg.Provider == "anthropic" && cfg.APIKey == "" {
		return nil, fmt.Errorf(
			"API key not set. Please set CLAW_API_KEY in .env file or as environment variable",
		)
	}

	return cfg, nil
}

// Save writes a configuration key-value pair to the config file
func Save(key, value string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error finding home directory: %w", err)
	}

	configPath := filepath.Join(home, ".claw-cli.yaml")

	// Set the value in viper
	viper.Set(key, value)

	// Write to config file
	if err := viper.WriteConfigAs(configPath); err != nil {
		// If file doesn't exist, create it
		if os.IsNotExist(err) {
			if err := viper.SafeWriteConfigAs(configPath); err != nil {
				return fmt.Errorf("error creating config file: %w", err)
			}
			return nil
		}
		return fmt.Errorf("error writing config file: %w", err)
	}

	return nil
}

// GetConfigPath returns the path to the config file
func GetConfigPath() (string, error) {
	if viper.ConfigFileUsed() != "" {
		return viper.ConfigFileUsed(), nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(home, ".claw-cli.yaml"), nil
}

// Validate checks if the configuration is valid
func Validate(cfg *types.Config) error {
	// Only require API key for anthropic
	if cfg.Provider == "anthropic" && cfg.APIKey == "" {
		return fmt.Errorf("API key is required for anthropic provider")
	}

	if cfg.Provider != "ollama" && cfg.Provider != "anthropic" {
		return fmt.Errorf("invalid provider: %s (must be 'ollama' or 'anthropic')", cfg.Provider)
	}

	if cfg.Temperature < 0 || cfg.Temperature > 1 {
		return fmt.Errorf("temperature must be between 0 and 1")
	}

	if cfg.MaxTokens < 1 || cfg.MaxTokens > 100000 {
		return fmt.Errorf("max-tokens must be between 1 and 100000")
	}

	return nil
}
