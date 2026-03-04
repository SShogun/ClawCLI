package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/SShogun/ClawCLI/internal/ai"
	"github.com/SShogun/ClawCLI/internal/config"
	"github.com/SShogun/ClawCLI/internal/types"
	"github.com/SShogun/ClawCLI/internal/utils"
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
  exit, quit, q      - Exit the chat
  clear              - Clear conversation history
  history            - Show conversation history
  explain <file>     - Explain code from a file (in-context)
  review <file>      - Review code from a file (in-context)`,
	Run: runChat,
}

func init() {
	rootCmd.AddCommand(chatCmd)
}

func runChat(cmd *cobra.Command, args []string) {
	cfg, err := config.Load()
	if err != nil {
		fmt.Println(errorStyle.Render("Error: " + err.Error()))
		os.Exit(1)
	}

	client := ai.NewClient(cfg)
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
			continue
		case "":
			continue
		}

		// Check for special commands: explain and review
		lowerInput := strings.ToLower(input)
		if strings.HasPrefix(lowerInput, "explain ") {
			filePath := strings.TrimPrefix(input, "explain ")
			filePath = strings.TrimPrefix(filePath, "Explain ")
			filePath = strings.TrimSpace(filePath)

			content, err := utils.ReadFile(filePath)
			if err != nil {
				fmt.Println(errorStyle.Render("Error reading file: " + err.Error()))
				continue
			}

			prompt := fmt.Sprintf(
				"Please explain the following code in detail. "+
					"Break down what each part does and the overall purpose.\n\n"+
					"File: %s\n```\n%s\n```",
				filePath,
				content,
			)

			fmt.Println(infoStyle.Render("Analyzing " + filePath + "..."))

			messages = append(messages, types.Message{
				Role:    "user",
				Content: prompt,
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
			continue
		}

		if strings.HasPrefix(lowerInput, "review ") {
			filePath := strings.TrimPrefix(input, "review ")
			filePath = strings.TrimPrefix(filePath, "Review ")
			filePath = strings.TrimSpace(filePath)

			content, err := utils.ReadFile(filePath)
			if err != nil {
				fmt.Println(errorStyle.Render("Error reading file: " + err.Error()))
				continue
			}

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

			fmt.Println(infoStyle.Render("Reviewing " + filePath + "..."))

			messages = append(messages, types.Message{
				Role:    "user",
				Content: prompt,
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
			continue
		}

		// Regular chat message
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
			"Special Commands:\n" +
			"  exit | quit        - Exit the chat\n" +
			"  clear              - Clear conversation history\n" +
			"  history            - Show conversation history\n" +
			"  explain <file>     - Explain code from a file\n" +
			"  review <file>      - Review code from a file",
	)

	fmt.Println(welcome)
	fmt.Println()
}
