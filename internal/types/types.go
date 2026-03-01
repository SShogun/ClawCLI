package types

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
	Stream      bool      `json:"stream,omitempty"`
}
type ChatResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

// Content represents content blocks in the response
type Content struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

// Usage tracks token usage
type Usage struct {
	InputTokens  int `json:"input_tokens"`
	OutputTokens int `json:"output_tokens"`
}

// ErrorResponse represents API error responses
type ErrorResponse struct {
	Type  string `json:"type"`
	Error Error  `json:"error"`
}

// Error contains error details
type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// Config represents application configuration
type Config struct {
	APIKey      string  `mapstructure:"api-key"`
	Model       string  `mapstructure:"model"`
	Temperature float64 `mapstructure:"temperature"`
	MaxTokens   int     `mapstructure:"max-tokens"`
	Verbose     bool    `mapstructure:"verbose"`
}
