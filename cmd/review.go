package cmd

import (
	"fmt"
	"os"

	"github.com/SShogun/ClawCLI/internal/ai"
	"github.com/SShogun/ClawCLI/internal/config"
	"github.com/SShogun/ClawCLI/internal/types"
	"github.com/SShogun/ClawCLI/internal/utils"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var reviewCmd = &cobra.Command{
	Use:   "review [file]",
	Short: "Review code from a file",
	Long: `Read a source file and get an AI code review with suggestions.

Examples:
  ai-cli review main.go
  ai-cli review src/handler.go`,
	Args: cobra.ExactArgs(1),
	Run:  runReview,
}

func init() {
	rootCmd.AddCommand(reviewCmd)
}

func runReview(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
		os.Exit(1)
	}

	filePath := args[0]

	// Read file content
	content, err := utils.ReadFile(filePath)
	if err != nil {
		fmt.Println(errorStyle.Render("Error reading file: " + err.Error()))
		os.Exit(1)
	}

	// Build review prompt
	prompt := fmt.Sprintf(
		"Please review the following code. Provide:\n"+
			"1. A brief summary of what the code does\n"+
			"2. Potential bugs or issues\n"+
			"3. Performance improvements\n"+
			"4. Code style and best practice suggestions\n"+
			"5. Security concerns (if any)\n\n"+
			"File: %s\n```\n%s\n```",
		filePath,
		content,
	)

	client := ai.NewClient(cfg)

	messages := []types.Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	fmt.Println(infoStyle.Render("Reviewing " + filePath + "..."))

	response, err := client.SendMessage(messages)
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
		os.Exit(1)
	}

	// Header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("11")).
		MarginBottom(1)

	fmt.Println(header.Render("\n🔍 Code Review: " + filePath))
	fmt.Println(dividerStyle.Render("────────────────────────────────────────"))
	fmt.Println(response)
}
