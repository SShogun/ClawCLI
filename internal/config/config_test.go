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
		{"qwen2.5-coder", true},
		{"codellama:7b", true},
		{"llama3.1:8b", true},
		{"claude-3-5-haiku-20241022", true},
		{"custom-model-name", true}, // Ollama accepts any model
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

func TestIsValidProvider(t *testing.T) {
	tests := []struct {
		provider string
		expected bool
	}{
		{"ollama", true},
		{"anthropic", true},
		{"invalid", false},
		{"", false},
	}

	for _, tt := range tests {
		t.Run(tt.provider, func(t *testing.T) {
			result := IsValidProvider(tt.provider)
			if result != tt.expected {
				t.Errorf("IsValidProvider(%q) = %v, want %v", tt.provider, result, tt.expected)
			}
		})
	}
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	// Check required keys exist
	if _, ok := cfg["provider"]; !ok {
		t.Error("DefaultConfig missing 'provider' key")
	}
	if _, ok := cfg["base-url"]; !ok {
		t.Error("DefaultConfig missing 'base-url' key")
	}
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
	if provider, ok := cfg["provider"].(string); !ok || provider != DefaultProvider {
		t.Errorf("DefaultConfig provider = %v, want %s", provider, DefaultProvider)
	}

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
			name: "valid ollama config",
			config: func() *types.Config {
				return &types.Config{
					Provider:    "ollama",
					BaseURL:     "http://localhost:11434",
					Model:       "qwen2.5-coder",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: false,
		},
		{
			name: "valid anthropic config",
			config: func() *types.Config {
				return &types.Config{
					Provider:    "anthropic",
					APIKey:      "sk-ant-valid-key",
					Model:       "claude-3-5-haiku-20241022",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: false,
		},
		{
			name: "anthropic missing API key",
			config: func() *types.Config {
				return &types.Config{
					Provider:    "anthropic",
					APIKey:      "",
					Model:       "claude-3-5-haiku-20241022",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: true,
		},
		{
			name: "ollama without API key (should be valid)",
			config: func() *types.Config {
				return &types.Config{
					Provider:    "ollama",
					BaseURL:     "http://localhost:11434",
					APIKey:      "",
					Model:       "qwen2.5-coder",
					Temperature: 0.7,
					MaxTokens:   4096,
				}
			},
			wantErr: false,
		},
		{
			name: "invalid provider",
			config: func() *types.Config {
				return &types.Config{
					Provider:    "openai",
					Model:       "gpt-4",
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
					Provider:    "ollama",
					Model:       "qwen2.5-coder",
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
					Provider:    "ollama",
					Model:       "qwen2.5-coder",
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
					Provider:    "ollama",
					Model:       "qwen2.5-coder",
					Temperature: 0.7,
					MaxTokens:   -1,
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
	// Clear the environment variables
	os.Unsetenv("CLAW_API_KEY")
	os.Setenv("CLAW_PROVIDER", "ollama")
	defer os.Unsetenv("CLAW_PROVIDER")

	err := Init()
	if err != nil {
		t.Fatalf("Init() failed: %v", err)
	}

	// Ollama should load fine without API key
	cfg, err := Load()
	if err != nil {
		t.Errorf("Load() should not return error for ollama without API key: %v", err)
	}

	if cfg != nil && cfg.Provider != "ollama" {
		t.Errorf("Load() provider = %v, want ollama", cfg.Provider)
	}
}
