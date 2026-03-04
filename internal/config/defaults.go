package config

// OllamaModelOptions contains popular Ollama models
var OllamaModelOptions = []string{
	"qwen2.5-coder",
	"qwen2.5-coder:7b",
	"codellama:7b",
	"deepseek-coder:6.7b",
	"llama3.1:8b",
	"mistral:7b",
	"phi3:mini",
}

// AnthropicModelOptions contains available Claude models
var AnthropicModelOptions = []string{
	"claude-haiku-4-5-20251001",
	"claude-3-5-haiku-20241022",
	"claude-3-5-sonnet-20241022",
	"claude-3-opus-20250219",
}

// ProviderOptions contains available providers
var ProviderOptions = []string{
	"ollama",
	"anthropic",
}

// DefaultConfig returns a config struct with default values
func DefaultConfig() map[string]interface{} {
	return map[string]interface{}{
		"provider":    DefaultProvider,
		"base-url":    DefaultBaseURL,
		"model":       DefaultModel,
		"temperature": DefaultTemperature,
		"max-tokens":  DefaultMaxTokens,
		"verbose":     false,
	}
}

// IsValidModel checks if the provided model name is valid
// For Ollama, any model name is valid (user can pull any model)
func IsValidModel(model string) bool {
	// Check Ollama models
	for _, m := range OllamaModelOptions {
		if m == model {
			return true
		}
	}
	// Check Anthropic models
	for _, m := range AnthropicModelOptions {
		if m == model {
			return true
		}
	}
	// For Ollama, any model is valid since users can pull custom models
	return true
}

// IsValidProvider checks if the provider is supported
func IsValidProvider(provider string) bool {
	for _, p := range ProviderOptions {
		if p == provider {
			return true
		}
	}
	return false
}
