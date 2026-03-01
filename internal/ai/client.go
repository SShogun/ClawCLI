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
)

// Client handles communication with the Claude API
type Client struct {
	APIKey      string
	BaseURL     string
	Model       string
	HTTPClient  *http.Client
	MaxTokens   int
	Temperature float64
}

// NewClient creates a new AI client
func NewClient(apiKey, model string) *Client {
	return &Client{
		APIKey:  apiKey,
		BaseURL: ClaudeAPIURL,
		Model:   model,
		HTTPClient: &http.Client{
			Timeout: 120 * time.Second,
		},
		MaxTokens:   4096,
		Temperature: 0.7,
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

// SendMessage sends a message to Claude and returns the response
func (c *Client) SendMessage(messages []types.Message) (string, error) {
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

// ValidateAPIKey checks if the API key is valid by making a test request
func (c *Client) ValidateAPIKey() error {
	testMessages := []types.Message{
		{
			Role:    "user",
			Content: "Hi",
		},
	}

	_, err := c.SendMessage(testMessages)
	return err
}
