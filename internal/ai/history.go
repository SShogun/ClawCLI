package ai

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SShogun/ClawCLI/internal/types"
)

// ConversationHistory manages chat history
type ConversationHistory struct {
	Messages  []types.Message `json:"messages"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}

// NewHistory creates a new conversation history
func NewHistory() *ConversationHistory {
	now := time.Now()
	return &ConversationHistory{
		Messages:  []types.Message{},
		CreatedAt: now,
		UpdatedAt: now,
	}
}

// Add appends a message to the history
func (h *ConversationHistory) Add(role, content string) {
	h.Messages = append(h.Messages, types.Message{
		Role:    role,
		Content: content,
	})
	h.UpdatedAt = time.Now()
}

// Clear removes all messages from history
func (h *ConversationHistory) Clear() {
	h.Messages = []types.Message{}
	h.UpdatedAt = time.Now()
}

// Count returns the number of messages
func (h *ConversationHistory) Count() int {
	return len(h.Messages)
}

// Save writes the conversation history to a file
func (h *ConversationHistory) Save(filename string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	historyDir := filepath.Join(home, ".claw-cli", "history")
	if err := os.MkdirAll(historyDir, 0755); err != nil {
		return fmt.Errorf("error creating history directory: %w", err)
	}

	historyPath := filepath.Join(historyDir, filename)

	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling history: %w", err)
	}

	if err := os.WriteFile(historyPath, data, 0644); err != nil {
		return fmt.Errorf("error writing history file: %w", err)
	}

	return nil
}

// Load reads conversation history from a file
func LoadHistory(filename string) (*ConversationHistory, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	historyPath := filepath.Join(home, ".claw-cli", "history", filename)

	data, err := os.ReadFile(historyPath)
	if err != nil {
		return nil, fmt.Errorf("error reading history file: %w", err)
	}

	var history ConversationHistory
	if err := json.Unmarshal(data, &history); err != nil {
		return nil, fmt.Errorf("error parsing history file: %w", err)
	}

	return &history, nil
}

// ListHistoryFiles returns all saved conversation history files
func ListHistoryFiles() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	historyDir := filepath.Join(home, ".claw-cli", "history")

	entries, err := os.ReadDir(historyDir)
	if err != nil {
		if os.IsNotExist(err) {
			return []string{}, nil
		}
		return nil, err
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() {
			files = append(files, entry.Name())
		}
	}

	return files, nil
}
