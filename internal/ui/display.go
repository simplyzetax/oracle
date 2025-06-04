package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
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

// ConfirmExecution asks user to confirm command execution with a simple y/n prompt
func ConfirmExecution(command string) bool {
	var confirm bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title(fmt.Sprintf("Run `%s`?", command)).
				Affirmative("Yes").
				Negative("No").
				Value(&confirm),
		),
	)
	err := form.Run()
	if err != nil {
		return false
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

// ShowCommandSuggestion displays a simple command suggestion like GitHub Copilot CLI
func ShowCommandSuggestion(command string) {
	fmt.Printf("\n%s %s\n",
		lipgloss.NewStyle().Foreground(green).Bold(true).Render("→"),
		lipgloss.NewStyle().Foreground(pearl).Render(command))
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

// StartResponseStream initializes the response display
func StartResponseStream() {
	// Simple prefix for the AI's response stream
	fmt.Print(lipgloss.NewStyle().Foreground(pearl).Bold(true).Render("A: "))
}

// EndResponseStream finalizes the response display
func EndResponseStream() {
	fmt.Println() // Just a newline for spacing
}

// ShowFirstRunWelcome displays a welcome message for first-time users
func ShowFirstRunWelcome() {
	welcome := lipgloss.NewStyle().
		Foreground(blue).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(gold).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1).
		Render("🔮 Welcome to Oracle CLI!\n\nThis is your first time using Oracle.\nWould you like to set up a convenient alias 'oa' for quick access?")

	fmt.Println(welcome)
}

// ConfirmAliasSetup asks if user wants to set up the alias
func ConfirmAliasSetup() bool {
	var setupAlias bool
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Set up 'oa' alias for Oracle?").
				Description("This will allow you to run 'oa \"question\"' instead of 'oracle ask \"question\"'").
				Affirmative("Yes, set up alias").
				Negative("No, thanks").
				Value(&setupAlias),
		),
	)
	err := form.Run()
	if err != nil {
		return false
	}
	return setupAlias
}

// ShowAliasInstructions displays manual alias setup instructions
func ShowAliasInstructions() {
	shell := os.Getenv("SHELL")
	var configFile string

	switch {
	case strings.Contains(shell, "zsh"):
		configFile = "~/.zshrc"
	case strings.Contains(shell, "bash"):
		configFile = "~/.bashrc or ~/.bash_profile"
	case strings.Contains(shell, "fish"):
		configFile = "~/.config/fish/config.fish"
	default:
		configFile = "your shell's config file"
	}

	instructions := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(blue).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1).
		Render(fmt.Sprintf(`To set up the alias manually, add this line to %s:

alias oa='oracle ask'

Then run: source %s

After that, you can use: oa "your question here"`, configFile, configFile))

	fmt.Println(instructions)
}

// ShowAliasSetupSuccess displays success message for alias setup
func ShowAliasSetupSuccess() {
	success := lipgloss.NewStyle().
		Foreground(green).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(green).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1).
		Render("✅ Alias setup complete!\n\nYou can now use 'oa \"your question\"' for quick access.\nRestart your terminal or run 'source ~/.zshrc' to activate.")

	fmt.Println(success)
}
