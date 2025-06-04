package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
)

var (
	// Color and style definitions
	primaryColor = lipgloss.Color("#00D9FF")
	errorColor   = lipgloss.Color("#FF6B6B")
	successColor = lipgloss.Color("#51CF66")
	mutedColor   = lipgloss.Color("#8B949E")

	// Styles
	headerStyle = lipgloss.NewStyle().
			Foreground(primaryColor).
			Bold(true).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(primaryColor).
			MarginBottom(1)

	errorStyle = lipgloss.NewStyle().
			Foreground(errorColor).
			Bold(true).
			SetString("❌ ")

	successStyle = lipgloss.NewStyle().
			Foreground(successColor).
			Bold(true).
			SetString("✅ ")

	questionStyle = lipgloss.NewStyle().
			Foreground(mutedColor).
			Italic(true).
			SetString("❓ ")

	responseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFFFF")).
			MarginLeft(2)
)

// ShowError displays an error message with styling
func ShowError(message string) {
	fmt.Println(errorStyle.Render() + message)
	os.Exit(1)
}

// ShowSuccess displays a success message with styling
func ShowSuccess(message string) {
	fmt.Println(successStyle.Render() + message)
}

// ShowQuestionHeader displays the question being asked with model info
func ShowQuestionHeader(model, question string) {
	header := fmt.Sprintf("🤖 Oracle (%s)", model)
	fmt.Println(headerStyle.Render(header))
	fmt.Println(questionStyle.Render() + question)
	fmt.Println()
}

// StartResponseStream initializes the response display
func StartResponseStream() {
	fmt.Print(responseStyle.Render("🔮 "))
}

// StreamText outputs streaming text with proper styling
func StreamText(text string) {
	fmt.Print(text)
}

// EndResponseStream finalizes the response display
func EndResponseStream() {
	fmt.Println()
	fmt.Println()
}
