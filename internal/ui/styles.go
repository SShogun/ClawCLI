package ui

import "github.com/charmbracelet/lipgloss"

// Color palette
var (
	ColorPrimary   = lipgloss.Color("12") // Blue
	ColorSuccess   = lipgloss.Color("10") // Green
	ColorWarning   = lipgloss.Color("11") // Yellow
	ColorError     = lipgloss.Color("9")  // Red
	ColorMuted     = lipgloss.Color("8")  // Gray
	ColorUser      = lipgloss.Color("10") // Green
	ColorAssistant = lipgloss.Color("12") // Blue
)

// Text styles
var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(ColorPrimary).
			MarginBottom(1)

	SuccessStyle = lipgloss.NewStyle().
			Foreground(ColorSuccess).
			Bold(true)

	ErrorStyle = lipgloss.NewStyle().
			Foreground(ColorError).
			Bold(true)

	WarningStyle = lipgloss.NewStyle().
			Foreground(ColorWarning)

	InfoStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)

	UserStyle = lipgloss.NewStyle().
			Foreground(ColorUser).
			Bold(true)

	AssistantStyle = lipgloss.NewStyle().
			Foreground(ColorAssistant).
			Bold(true)

	DividerStyle = lipgloss.NewStyle().
			Foreground(ColorMuted)
)

// Box styles
var (
	BoxStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(ColorPrimary).
			Padding(1, 2)

	CodeBlockStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderForeground(ColorMuted).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1)
)

// Helper functions

// RenderBox renders content in a styled box
func RenderBox(content string) string {
	return BoxStyle.Render(content)
}

// RenderCodeBlock renders code in a styled block
func RenderCodeBlock(code string) string {
	return CodeBlockStyle.Render(code)
}

// RenderDivider renders a horizontal divider
func RenderDivider() string {
	return DividerStyle.Render("────────────────────────────────────────")
}

// RenderTitle renders a title
func RenderTitle(title string) string {
	return TitleStyle.Render(title)
}

// RenderSuccess renders a success message
func RenderSuccess(msg string) string {
	return SuccessStyle.Render("✓ " + msg)
}

// RenderError renders an error message
func RenderError(msg string) string {
	return ErrorStyle.Render("✗ " + msg)
}

// RenderWarning renders a warning message
func RenderWarning(msg string) string {
	return WarningStyle.Render("⚠ " + msg)
}

// RenderInfo renders an info message
func RenderInfo(msg string) string {
	return InfoStyle.Render(msg)
}
