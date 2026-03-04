package types

// Message represents a single message in the conversation
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// ============================================================
// Ollama API Types
// ============================================================

// OllamaChatRequest is the request body for Ollama's /api/chat endpoint
type OllamaChatRequest struct {
	Model    string        `json:"model"`
	Messages []Message     `json:"messages"`
	Stream   bool          `json:"stream"`
	Options  OllamaOptions `json:"options,omitempty"`
}

// OllamaOptions contains model parameters for Ollama
type OllamaOptions struct {
	Temperature float64 `json:"temperature,omitempty"`
	NumPredict  int     `json:"num_predict,omitempty"`
}

// OllamaChatResponse is the response from Ollama's /api/chat endpoint
type OllamaChatResponse struct {
	Model           string  `json:"model"`
	CreatedAt       string  `json:"created_at"`
	Message         Message `json:"message"`
	Done            bool    `json:"done"`
	TotalDuration   int64   `json:"total_duration"`
	EvalCount       int     `json:"eval_count"`
	PromptEvalCount int     `json:"prompt_eval_count"`
}

// OllamaErrorResponse represents an error from Ollama
type OllamaErrorResponse struct {
	Error string `json:"error"`
}

// ============================================================
// Anthropic API Types (kept for future use)
// ============================================================

// ChatRequest represents a request to the Claude API
type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	MaxTokens   int       `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
	Stream      bool      `json:"stream,omitempty"`
}

// ChatResponse represents the response from Claude API
type ChatResponse struct {
	Content []ContentBlock `json:"content"`
}

// ContentBlock represents content blocks in the response
type ContentBlock struct {
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

// ============================================================
// Application Config
// ============================================================

// Config represents application configuration
type Config struct {
	Provider    string  `mapstructure:"provider"`
	BaseURL     string  `mapstructure:"base-url"`
	APIKey      string  `mapstructure:"api-key"`
	Model       string  `mapstructure:"model"`
	Temperature float64 `mapstructure:"temperature"`
	MaxTokens   int     `mapstructure:"max-tokens"`
	Verbose     bool    `mapstructure:"verbose"`
}
