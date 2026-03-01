package config

// ModelOptions contains available Claude models
var ModelOptions = []string{
	"claude-haiku-4-5-20251001",
	"claude-3-5-haiku-20241022",
	"claude-3-5-sonnet-20241022",
	"claude-3-opus-20250219",
}

// DefaultConfig returns a config struct with default values
func DefaultConfig() map[string]interface{} {
	return map[string]interface{}{
		"model":       DefaultModel,
		"temperature": DefaultTemperature,
		"max-tokens":  DefaultMaxTokens,
		"verbose":     false,
	}
}

// IsValidModel checks if the provided model name is valid
func IsValidModel(model string) bool {
	for _, m := range ModelOptions {
		if m == model {
			return true
		}
	}
	return false
}
