package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	userStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("10")).Bold(true)
	assistantStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("12")).Bold(true)
	errorStyle     = lipgloss.NewStyle().Foreground(lipgloss.Color("9")).Bold(true)
	infoStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	dividerStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
)

var chatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Start an interactive chat session with the assistant",
	Long: `Start an interactive chat session with the AI assistant.
Type your messages and get responses in real-time.

Special commands:
  exit, quit, q  - Exit the chat
  clear           - Clear conversation history
  history         - Show conversation history`,
	Run: runChat,
}

func init() {
	rootCmd.AddCommand(chatCmd)
}

func runChat(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
	}

	client := ai.NewClient(cfg.APIKey, cfg.Model)
	messages := []types.Message{}

	printWelcome()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print(userStyle.Render("You: "))
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(errorStyle.Render("Error reading input: " + err.Error()))
			continue
		}

		input = strings.TrimSpace(input)

		switch strings.ToLower(input) {
		case "exit", "quit", "q":
			fmt.Println(infoStyle.Render("\nGoodbye! 👋"))
			return
		case "clear":
			messages = []types.Message{}
			fmt.Println(infoStyle.Render("Conversation cleared."))
			continue
		case "history":
			if len(messages) == 0 {
				fmt.Println(infoStyle.Render("No conversation history."))
			} else {
				fmt.Println(infoStyle.Render("Conversation history:"))
				for _, msg := range messages {
					fmt.Println(infoStyle.Render(fmt.Sprintf("%s: %s", msg.Role, msg.Content)))
				}
			}
		case "":
			continue
		}

		messages = append(messages, types.Message{
			Role:    "user",
			Content: input,
		})

		fmt.Print(assistantStyle.Render("Assistant: "))
		response, err := client.SendMessage(messages)
		if err != nil {
			fmt.Println(errorStyle.Render("Error: " + err.Error()))
			continue
		}
		fmt.Println(response)
		fmt.Println(dividerStyle.Render(strings.Repeat("─", 50)))

		messages = append(messages, types.Message{
			Role:    "assistant",
			Content: response,
		})
	}
}

func printWelcome() {
	border := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("12")).
		Padding(1, 2)

	welcome := border.Render(
		"🤖 CLAW CLI Assistant\n\n" +
			"Type your message to chat with the AI.\n" +
			"Commands: exit | clear | history",
	)

	fmt.Println(welcome)
	fmt.Println()
}
