package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

// Color palette
var (
	// Brand colors
	blue   = lipgloss.Color("#00D9FF")
	gold   = lipgloss.Color("#F59E0B")
	green  = lipgloss.Color("#10B981")
	orange = lipgloss.Color("#F97316")
	yellow = lipgloss.Color("#FBBF24")

	// Neutral colors
	slate = lipgloss.Color("#64748B")
	pearl = lipgloss.Color("#F8FAFC")

	// Status colors
	statusErrorColor   = lipgloss.Color("#FF6B6B")
	statusSuccessColor = lipgloss.Color("#51CF66")
)

// Styles
var (
	// BannerStyle for the main application banner
	BannerStyle = lipgloss.NewStyle().
			Foreground(blue).
			Bold(true).
			Align(lipgloss.Center).
			MarginBottom(1)

	// ResponseStyle for AI responses
	ResponseStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(slate).
			Padding(1, 2).
			MarginTop(1).
			MarginBottom(1)

	// HeaderStyle for question headers
	HeaderStyle = lipgloss.NewStyle().
			Foreground(yellow).
			Bold(true).
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(blue).
			MarginBottom(1)

	// ErrorStyle for error messages
	ErrorStyle = lipgloss.NewStyle().
			Foreground(statusErrorColor).
			Bold(true)

	// SuccessStyle for success messages
	SuccessStyle = lipgloss.NewStyle().
			Foreground(statusSuccessColor).
			Bold(true)

	// QuestionStyle for displaying the user's question
	QuestionStyle = lipgloss.NewStyle().
			Foreground(slate).
			Italic(true)
)

// Markdown renderer
var markdownRenderer *glamour.TermRenderer

func init() {
	var err error
	markdownRenderer, err = glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(80),
	)
	if err != nil {
		markdownRenderer = nil // Fallback to nil if glamour fails
	}
}

// StreamMarkdownText outputs streaming text
func StreamMarkdownText(text string) {
	fmt.Print(text)
}

// RenderFinalResponse renders the complete response with markdown support
func RenderFinalResponse(fullText string) {
	if markdownRenderer == nil {
		fmt.Println(ResponseStyle.Render(fullText))
		return
	}
	rendered, err := markdownRenderer.Render(fullText)
	if err != nil {
		fmt.Println(ResponseStyle.Render(fullText))
		return
	}
	fmt.Println(ResponseStyle.Render(rendered))
}

// ConfirmExecution asks user to confirm command execution using huh
func ConfirmExecution(command string) bool {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Execute this command?").
				Description(fmt.Sprintf("Command: %s", command)).
				Affirmative("Yes, execute").
				Negative("No, skip").
				Value(&confirm),
		),
	)
	err := form.Run()
	if err != nil {
		// Fallback to basic text confirmation if huh fails
		fmt.Print(lipgloss.NewStyle().Foreground(gold).Render(fmt.Sprintf("Execute: %s (y/N): ", command)))
		var responseText string
		fmt.Scanln(&responseText)
		responseText = strings.ToLower(strings.TrimSpace(responseText))
		return responseText == "y" || responseText == "yes"
	}
	return confirm
}

// ConfirmContinueOnError asks if execution should continue after error using huh
func ConfirmContinueOnError() bool {
	var continueExec bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Command failed! Continue?").
				Affirmative("Yes, continue").
				Negative("No, stop").
				Value(&continueExec),
		),
	)
	err := form.Run()
	if err != nil {
		// Fallback to basic text confirmation if huh fails
		fmt.Print(lipgloss.NewStyle().Foreground(gold).Render("Continue with remaining commands? (y/N): "))
		var responseText string
		fmt.Scanln(&responseText)
		responseText = strings.ToLower(strings.TrimSpace(responseText))
		return responseText == "y" || responseText == "yes"
	}
	return continueExec
}

// ShowCommandsTable displays commands in a table
func ShowCommandsTable(commands []string) {
	if len(commands) == 0 {
		return
	}

	header := lipgloss.NewStyle().
		Foreground(gold).
		Bold(true).
		SetString("Detected Executable Commands:")

	fmt.Println(header.Render())
	fmt.Println()

	rows := make([][]string, len(commands))
	for i, cmd := range commands {
		rows[i] = []string{
			fmt.Sprintf("%d", i+1),
			cmd,
		}
	}

	t := table.New().
		Border(lipgloss.RoundedBorder()).
		BorderStyle(lipgloss.NewStyle().Foreground(gold)).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == table.HeaderRow {
				return lipgloss.NewStyle().
					Foreground(gold).
					Bold(true).
					Align(lipgloss.Center)
			}
			if col == 0 { // Index column
				return lipgloss.NewStyle().
					Foreground(blue).
					Bold(true).
					Align(lipgloss.Center).
					Width(5)
			}
			return lipgloss.NewStyle().Padding(0, 1) // Command column
		}).
		Headers("#", "Command").
		Rows(rows...)

	fmt.Println(t.Render())
	fmt.Println()
}

// ShowExecutionStatus displays execution status messages
func ShowExecutionStatus(message string, statusType string) {
	var style lipgloss.Style
	var prefix string

	switch statusType {
	case "success":
		style = lipgloss.NewStyle().Foreground(green).Bold(true)
		prefix = "Success:"
	case "error":
		style = lipgloss.NewStyle().Foreground(statusErrorColor).Bold(true)
		prefix = "Error:"
	case "executing":
		style = lipgloss.NewStyle().Foreground(orange).Bold(true)
		prefix = "Executing:"
	case "warning":
		style = lipgloss.NewStyle().Foreground(gold).Bold(true)
		prefix = "Warning:"
	default: // "info" or any other type
		style = lipgloss.NewStyle().Foreground(blue).Bold(true)
		prefix = "Info:"
	}
	fmt.Println(style.Render(fmt.Sprintf("%s %s", prefix, message)))
}

// ShowError displays a general error message and exits
func ShowError(message string) {
	fmt.Println(ErrorStyle.Render("Error: " + message))
	os.Exit(1)
}

// ShowSuccess displays a general success message
func ShowSuccess(message string) {
	fmt.Println(SuccessStyle.Render("Success: " + message))
}

// ShowQuestionHeader displays the question being asked with model info
func ShowQuestionHeader(model, questionText string) {
	headerMsg := fmt.Sprintf("Oracle (%s)", model)
	fmt.Println(HeaderStyle.Render(headerMsg))
	fmt.Println(QuestionStyle.Render("Q: " + questionText))
	fmt.Println()
}

// StartResponseStream initializes the response display
func StartResponseStream() {
	// Simple prefix for the AI's response stream
	fmt.Print(lipgloss.NewStyle().Foreground(pearl).Bold(true).Render("A: "))
}

// EndResponseStream finalizes the response display
func EndResponseStream() {
	fmt.Println() // Just a newline for spacing
}
