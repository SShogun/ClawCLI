package config

import (
	"os"
	"testing"

	"github.com/SShogun/ClawCLI/internal/types"
)

func TestIsValidModel(t *testing.T) {
	tests := []struct {
		model    string
		expected bool
	}{
		{"claude-haiku-4-5-20251001", true},
		{"claude-3-5-haiku-20241022", true},
		{"claude-3-5-sonnet-20241022", true},
		{"claude-3-opus-20250219", true},
		{"invalid-model", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			result := IsValidModel(tt.model)
			if result != tt.expected {
				t.Errorf("IsValidModel(%q) = %v, want %v", tt.model, result, tt.expected)
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	// Check required keys exist
	if _, ok := cfg["model"]; !ok {
		t.Error("DefaultConfig missing 'model' key")
	}
	if _, ok := cfg["temperature"]; !ok {
		t.Error("DefaultConfig missing 'temperature' key")
	}
	if _, ok := cfg["max-tokens"]; !ok {
		t.Error("DefaultConfig missing 'max-tokens' key")
	}
	if _, ok := cfg["verbose"]; !ok {
		t.Error("DefaultConfig missing 'verbose' key")
	}

	// Check default values
	if model, ok := cfg["model"].(string); !ok || model != DefaultModel {
		t.Errorf("DefaultConfig model = %v, want %s", model, DefaultModel)
	}

	if temp, ok := cfg["temperature"].(float64); !ok || temp != DefaultTemperature {
		t.Errorf("DefaultConfig temperature = %v, want %f", temp, DefaultTemperature)
	}

	if tokens, ok := cfg["max-tokens"].(int); !ok || tokens != DefaultMaxTokens {
		t.Errorf("DefaultConfig max-tokens = %v, want %d", tokens, DefaultMaxTokens)
	}
}

func TestValidate(t *testing.T) {
	type testCase struct {
		name    string
		config  func() *types.Config
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid config",
			config: func() *types.Config {
				return &types.Config{
					APIKey:      "sk-ant-valid-key",
					Model:       "claude-haiku-4-5-20251001",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: false,
		},
		{
			name: "missing API key",
			config: func() *types.Config {
				return &types.Config{
					APIKey:      "",
					Model:       "claude-haiku-4-5-20251001",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: true,
		},
		{
			name: "invalid temperature too high",
			config: func() *types.Config {
				return &types.Config{
					APIKey:      "sk-ant-valid-key",
					Model:       "claude-haiku-4-5-20251001",
					Temperature: 1.5,
					MaxTokens:   4096,
				}
			},
			wantErr: true,
		},
		{
			name: "invalid temperature too low",
			config: func() *types.Config {
				return &types.Config{
					APIKey:      "sk-ant-valid-key",
					Model:       "claude-haiku-4-5-20251001",
					Temperature: -0.5,
					MaxTokens:   4096,
				}
			},
			wantErr: true,
		},
		{
			name: "invalid max tokens",
			config: func() *types.Config {
				return &types.Config{
					APIKey:      "sk-ant-valid-key",
					Model:       "claude-haiku-4-5-20251001",
					Temperature: 0.7,
					MaxTokens:   -1,
				}
			},
			wantErr: true,
		},
		{
			name: "invalid model",
			config: func() *types.Config {
				return &types.Config{
					APIKey:      "sk-ant-valid-key",
					Model:       "invalid-model",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := tt.config()
			err := Validate(cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInit(t *testing.T) {
	// Save current working directory
	oldCwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get working directory: %v", err)
	}
	defer os.Chdir(oldCwd)

	err = Init()
	if err != nil {
		t.Errorf("Init() error = %v, want nil", err)
	}
}

func TestLoadWithoutAPIKey(t *testing.T) {
	// Clear the environment variable if it exists
	os.Unsetenv("CLAW_API_KEY")

	err := Init()
	if err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	_, err = Load()
	if err == nil {
		t.Error("Load() should return error when API key is missing")
	}
}
