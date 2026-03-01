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

var explainCmd = &cobra.Command{
	Use:   "explain [file]",
	Short: "Explain code from a file",
	Long: `Read a source file and get an AI explanation of what the code does.

Examples:
  ai-cli explain main.go
  ai-cli explain src/handler.go`,
	Args: cobra.ExactArgs(1),
	Run:  runExplain,
}

func init() {
	rootCmd.AddCommand(explainCmd)
}

func runExplain(cmd *cobra.Command, args []string) {
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

	// Build prompt
	prompt := fmt.Sprintf(
		"Please explain the following code in detail. "+
			"Break down what each part does and the overall purpose.\n\n"+
			"File: %s\n```\n%s\n```",
		filePath,
		content,
	)

	client := ai.NewClient(cfg.APIKey, cfg.Model)
	client.SetOptions(cfg.MaxTokens, cfg.Temperature)

	messages := []types.Message{
		{
			Role:    "user",
			Content: prompt,
		},
	}

	fmt.Println(infoStyle.Render("Analyzing " + filePath + "..."))

	response, err := client.SendMessage(messages)
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
		os.Exit(1)
	}

	// Header
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("12")).
		MarginBottom(1)

	fmt.Println(header.Render("\n📄 Explanation for: " + filePath))
	fmt.Println(dividerStyle.Render("────────────────────────────────────────"))
	fmt.Println(response)
}
