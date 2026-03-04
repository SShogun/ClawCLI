package ai

import (
	"testing"

	"github.com/SShogun/ClawCLI/internal/types"
)

func TestNewClientOllama(t *testing.T) {
	cfg := &types.Config{
		Provider:    "ollama",
		BaseURL:     "http://localhost:11434",
		Model:       "qwen2.5-coder",
		MaxTokens:   4096,
		Temperature: 0.7,
	}

	client := NewClient(cfg)

	if client.Provider != "ollama" {
		t.Errorf("NewClient() Provider = %v, want ollama", client.Provider)
	}

	if client.Model != "qwen2.5-coder" {
		t.Errorf("NewClient() Model = %v, want qwen2.5-coder", client.Model)
	}

	if client.BaseURL != "http://localhost:11434" {
		t.Errorf("NewClient() BaseURL = %v, want http://localhost:11434", client.BaseURL)
	}

	if client.MaxTokens != 4096 {
		t.Errorf("NewClient() MaxTokens = %v, want 4096", client.MaxTokens)
	}

	if client.Temperature != 0.7 {
		t.Errorf("NewClient() Temperature = %v, want 0.7", client.Temperature)
	}

	if client.HTTPClient == nil {
		t.Error("NewClient() HTTPClient is nil, want non-nil")
	}
}

func TestNewClientAnthropic(t *testing.T) {
	cfg := &types.Config{
		Provider:    "anthropic",
		APIKey:      "sk-ant-test-key",
		Model:       "claude-3-5-haiku-20241022",
		MaxTokens:   4096,
		Temperature: 0.7,
	}

	client := NewClient(cfg)

	if client.Provider != "anthropic" {
		t.Errorf("NewClient() Provider = %v, want anthropic", client.Provider)
	}

	if client.BaseURL != ClaudeAPIURL {
		t.Errorf("NewClient() BaseURL = %v, want %v", client.BaseURL, ClaudeAPIURL)
	}

	if client.APIKey != "sk-ant-test-key" {
		t.Errorf("NewClient() APIKey = %v, want sk-ant-test-key", client.APIKey)
	}
}

func TestSetOptions(t *testing.T) {
	type args struct {
		maxTokens   int
		temperature float64
	}

	tests := []struct {
		name         string
		args         args
		expectTokens int
		expectTemp   float64
	}{
		{
			name:         "valid options",
			args:         args{maxTokens: 2048, temperature: 0.5},
			expectTokens: 2048,
			expectTemp:   0.5,
		},
		{
			name:         "zero max tokens (should not update)",
			args:         args{maxTokens: 0, temperature: 0.5},
			expectTokens: 4096, // Default value
			expectTemp:   0.5,
		},
		{
			name:         "negative temperature (should not update)",
			args:         args{maxTokens: 1024, temperature: -0.5},
			expectTokens: 1024,
			expectTemp:   0.7, // Default value
		},
		{
			name:         "temperature above 1 (should not update)",
			args:         args{maxTokens: 1024, temperature: 1.5},
			expectTokens: 1024,
			expectTemp:   0.7, // Default value
		},
		{
			name:         "boundary temperature values",
			args:         args{maxTokens: 1024, temperature: 0.0},
			expectTokens: 1024,
			expectTemp:   0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg := &types.Config{
				Provider:    "ollama",
				BaseURL:     "http://localhost:11434",
				Model:       "qwen2.5-coder",
				MaxTokens:   4096,
				Temperature: 0.7,
			}
			client := NewClient(cfg)
			client.SetOptions(tt.args.maxTokens, tt.args.temperature)

			if client.MaxTokens != tt.expectTokens {
				t.Errorf("SetOptions() MaxTokens = %v, want %v", client.MaxTokens, tt.expectTokens)
			}

			if client.Temperature != tt.expectTemp {
				t.Errorf("SetOptions() Temperature = %v, want %v", client.Temperature, tt.expectTemp)
			}
		})
	}
}

func TestOllamaRequestMarshaling(t *testing.T) {
	messages := []types.Message{
		{
			Role:    "user",
			Content: "Hello",
		},
	}

	request := types.OllamaChatRequest{
		Model:    "qwen2.5-coder",
		Messages: messages,
		Stream:   false,
		Options: types.OllamaOptions{
			Temperature: 0.7,
			NumPredict:  4096,
		},
	}

	if request.Model != "qwen2.5-coder" {
		t.Error("OllamaChatRequest Model mismatch")
	}

	if len(request.Messages) != 1 {
		t.Errorf("OllamaChatRequest Messages length = %d, want 1", len(request.Messages))
	}

	if request.Stream != false {
		t.Error("OllamaChatRequest Stream should be false")
	}

	if request.Options.Temperature != 0.7 {
		t.Errorf("OllamaChatRequest Temperature = %v, want 0.7", request.Options.Temperature)
	}

	if request.Options.NumPredict != 4096 {
		t.Errorf("OllamaChatRequest NumPredict = %v, want 4096", request.Options.NumPredict)
	}
}

func TestSendMessageUnsupportedProvider(t *testing.T) {
	cfg := &types.Config{
		Provider:    "invalid",
		BaseURL:     "http://localhost:11434",
		Model:       "test",
		MaxTokens:   4096,
		Temperature: 0.7,
	}
	client := NewClient(cfg)

	_, err := client.SendMessage([]types.Message{{Role: "user", Content: "test"}})
	if err == nil {
		t.Error("SendMessage() should return error for unsupported provider")
	}
}

func TestMessageCreation(t *testing.T) {
	tests := []struct {
		name    string
		role    string
		content string
	}{
		{
			name:    "user message",
			role:    "user",
			content: "What is Go?",
		},
		{
			name:    "assistant message",
			role:    "assistant",
			content: "Go is a programming language...",
		},
		{
			name:    "system message",
			role:    "system",
			content: "You are a helpful assistant",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			msg := types.Message{
				Role:    tt.role,
				Content: tt.content,
			}

			if msg.Role != tt.role {
				t.Errorf("Message Role = %s, want %s", msg.Role, tt.role)
			}

			if msg.Content != tt.content {
				t.Errorf("Message Content = %s, want %s", msg.Content, tt.content)
			}
		})
	}
}
