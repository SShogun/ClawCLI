package ai

import (
	"testing"

	"github.com/SShogun/ClawCLI/internal/types"
)

func TestNewClient(t *testing.T) {
	apiKey := "test-api-key"
	model := "claude-3-5-haiku-20241022"

	client := NewClient(apiKey, model)

	if client.APIKey != apiKey {
		t.Errorf("NewClient() APIKey = %v, want %v", client.APIKey, apiKey)
	}

	if client.Model != model {
		t.Errorf("NewClient() Model = %v, want %v", client.Model, model)
	}

	if client.BaseURL != ClaudeAPIURL {
		t.Errorf("NewClient() BaseURL = %v, want %v", client.BaseURL, ClaudeAPIURL)
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
			client := NewClient("test-key", "claude-3-5-haiku-20241022")
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

func TestChatRequestMarshaling(t *testing.T) {
	messages := []types.Message{
		{
			Role:    "user",
			Content: "Hello",
		},
	}

	request := types.ChatRequest{
		Model:       "claude-3-5-haiku-20241022",
		Messages:    messages,
		MaxTokens:   2048,
		Temperature: 0.7,
		Stream:      false,
	}

	// This test just verifies the struct can be created and has valid fields
	if request.Model != "claude-3-5-haiku-20241022" {
		t.Error("ChatRequest Model mismatch")
	}

	if len(request.Messages) != 1 {
		t.Errorf("ChatRequest Messages length = %d, want 1", len(request.Messages))
	}

	if request.Messages[0].Role != "user" {
		t.Errorf("Message Role = %s, want user", request.Messages[0].Role)
	}

	if request.Messages[0].Content != "Hello" {
		t.Errorf("Message Content = %s, want Hello", request.Messages[0].Content)
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
