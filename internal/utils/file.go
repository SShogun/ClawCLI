package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// ReadFile reads the contents of a file
func ReadFile(path string) (string, error) {
	// Check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("file not found: %s", path)
	}

	// Read file
	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}

// GetFileExtension returns the file extension
func GetFileExtension(path string) string {
	return filepath.Ext(path)
}

// IsCodeFile checks if a file is a code file based on extension
func IsCodeFile(path string) bool {
	codeExtensions := map[string]bool{
		".go":   true,
		".py":   true,
		".js":   true,
		".ts":   true,
		".java": true,
		".c":    true,
		".cpp":  true,
		".rs":   true,
		".rb":   true,
		".php":  true,
		".html": true,
		".css":  true,
		".sh":   true,
	}

	ext := GetFileExtension(path)
	return codeExtensions[ext]
}
