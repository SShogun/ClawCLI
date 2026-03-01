package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"honnef.co/go/tools/config"
)

var askCmd = &cobra.Command{
	Use:   "ask [question]",
	Short: "Ask a one-shot question",
	Long: `Ask a single question and get a response.
No conversation history is maintained.

Examples:
  ai-cli ask "What is a goroutine?"
  ai-cli ask "How to reverse a string in Go?"`,
	Args: cobra.MinimumNArgs(1),
	Run:  runAsk,
}

func init() {
	rootCmd.AddCommand(askCmd)
}

func runAsk(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
		os.Exit(1)
	}

	question := strings.Join(args, " ")

	client := ai.NewClient(cfg.APIKey, cfg.Model)

	messages := []types.Message{
		{
			Role:    "user",
			Content: question,
		},
	}

	// Show loading
	fmt.Println(infoStyle.Render("Thinking..."))

	response, err := client.SendMessage(messages)
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
		os.Exit(1)
	}

	// Format output
	box := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("12")).
		Padding(1, 2).
		MarginTop(1)

	fmt.Println(box.Render(response))
}
