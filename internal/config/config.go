package config

import (
	"fmt"
	"go/types"
	"os"

	"github.com/SShogun/ClawCLI/internal/types"

	"github.com/spf13/viper"
)

// Default configuration values
const (
	DefaultModel       = "claude-sonnet-4.5"
	DefaultTemperature = 0.7
	DefaultMaxTokens   = 4096
)

// Load reads configuration from viper and returns a Config struct
func Load() (*types.Config, error) {
	cfg := &types.Config{
		APIKey:      viper.GetString("api-key"),
		Model:       viper.GetString("model"),
		Temperature: viper.GetFloat64("temperature"),
		MaxTokens:   viper.GetInt("max-tokens"),
		Verbose:     viper.GetBool("verbose"),
	}

	// Apply defaults if not set
	if cfg.Model == "" {
		cfg.Model = DefaultModel
	}

	if cfg.Temperature == 0 {
		cfg.Temperature = DefaultTemperature
	}

	if cfg.MaxTokens == 0 {
		cfg.MaxTokens = DefaultMaxTokens
	}

	// Validate API key
	if cfg.APIKey == "" {
		return nil, fmt.Errorf(
			"API key not set. Please set it using:\n" +
				"  1. clawcli config set api-key YOUR_KEY\n" +
				"  2. --api-key flag\n" +
				"  3. CLAW_CLI_API_KEY environment variable",
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

	configPath := home + "/.claw-cli.yaml"

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

	return home + "/.claw-cli.yaml", nil
}

// Validate checks if the configuration is valid
func Validate(cfg *types.Config) error {
	if cfg.APIKey == "" {
		return fmt.Errorf("API key is required")
	}

	if cfg.Temperature < 0 || cfg.Temperature > 1 {
		return fmt.Errorf("temperature must be between 0 and 1")
	}

	if cfg.MaxTokens < 1 || cfg.MaxTokens > 100000 {
		return fmt.Errorf("max-tokens must be between 1 and 100000")
	}

	return nil
}
