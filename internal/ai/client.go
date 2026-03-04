package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/SShogun/ClawCLI/internal/types"
)

const (
	ClaudeAPIURL     = "https://api.anthropic.com/v1/messages"
	AnthropicVersion = "2023-06-01"
	OllamaChatPath   = "/api/chat"
)

// Client handles communication with the AI provider
type Client struct {
	Provider    string
	APIKey      string
	BaseURL     string
	Model       string
	HTTPClient  *http.Client
	MaxTokens   int
	Temperature float64
}

// NewClient creates a new AI client from config
func NewClient(cfg *types.Config) *Client {
	baseURL := cfg.BaseURL
	if cfg.Provider == "anthropic" {
		baseURL = ClaudeAPIURL
	}

	return &Client{
		Provider: cfg.Provider,
		APIKey:   cfg.APIKey,
		BaseURL:  baseURL,
		Model:    cfg.Model,
		HTTPClient: &http.Client{
			Timeout: 120 * time.Second,
		},
		MaxTokens:   cfg.MaxTokens,
		Temperature: cfg.Temperature,
	}
}

// SetOptions allows setting optional parameters
func (c *Client) SetOptions(maxTokens int, temperature float64) {
	if maxTokens > 0 {
		c.MaxTokens = maxTokens
	}
	if temperature >= 0 && temperature <= 1 {
		c.Temperature = temperature
	}
}

// SendMessage sends a message and returns the response (routes to correct provider)
func (c *Client) SendMessage(messages []types.Message) (string, error) {
	switch c.Provider {
	case "ollama":
		return c.sendOllamaMessage(messages)
	case "anthropic":
		return c.sendAnthropicMessage(messages)
	default:
		return "", fmt.Errorf("unsupported provider: %s", c.Provider)
	}
}

// sendOllamaMessage sends a message to a local Ollama instance
func (c *Client) sendOllamaMessage(messages []types.Message) (string, error) {
	// Build Ollama request
	reqBody := types.OllamaChatRequest{
		Model:    c.Model,
		Messages: messages,
		Stream:   false,
		Options: types.OllamaOptions{
			Temperature: c.Temperature,
			NumPredict:  c.MaxTokens,
		},
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Build URL: baseURL + /api/chat
	url := c.BaseURL + OllamaChatPath

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error connecting to Ollama at %s: %w\nMake sure Ollama is running (ollama serve)", c.BaseURL, err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Handle errors
	if resp.StatusCode != http.StatusOK {
		var errResp types.OllamaErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return "", fmt.Errorf("Ollama error (status %d): %s", resp.StatusCode, string(body))
		}
		return "", fmt.Errorf("Ollama error: %s", errResp.Error)
	}

	// Parse response
	var chatResp types.OllamaChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("error parsing Ollama response: %w", err)
	}

	if chatResp.Message.Content == "" {
		return "", fmt.Errorf("empty response from Ollama")
	}

	return chatResp.Message.Content, nil
}

// sendAnthropicMessage sends a message to the Anthropic Claude API
func (c *Client) sendAnthropicMessage(messages []types.Message) (string, error) {
	// Build request
	reqBody := types.ChatRequest{
		Model:       c.Model,
		Messages:    messages,
		MaxTokens:   c.MaxTokens,
		Temperature: c.Temperature,
		Stream:      false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("error marshaling request: %w", err)
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("anthropic-version", AnthropicVersion)

	// Send request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %w", err)
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	// Handle errors
	if resp.StatusCode != http.StatusOK {
		var errResp types.ErrorResponse
		if err := json.Unmarshal(body, &errResp); err != nil {
			return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
		}
		return "", fmt.Errorf("API error: %s", errResp.Error.Message)
	}

	// Parse successful response
	var chatResp types.ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	// Extract text from content
	if len(chatResp.Content) == 0 {
		return "", fmt.Errorf("empty response from API")
	}

	return chatResp.Content[0].Text, nil
}

// ValidateConnection checks if the provider is reachable
func (c *Client) ValidateConnection() error {
	testMessages := []types.Message{
		{
			Role:    "user",
			Content: "Hi",
		},
	}

	_, err := c.SendMessage(testMessages)
	return err
}
