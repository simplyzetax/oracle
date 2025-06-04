package ui

import (
	"fmt"
	"os"
	"strings"

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

	// Command execution UI functions

	commandStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFA500")).
			Bold(true).
			SetString("⚡ ")

	warningStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFD700")).
			Bold(true).
			SetString("⚠️  ")
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

// ShowCommandsDetected displays the detected commands
func ShowCommandsDetected(commands []string) {
	fmt.Println(warningStyle.Render() + "Detected executable commands:")
	for i, cmd := range commands {
		fmt.Printf("  %d. %s\n", i+1, cmd)
	}
	fmt.Println()
}

// ConfirmExecution asks user to confirm command execution
func ConfirmExecution(command string) bool {
	prompt := fmt.Sprintf("Execute: %s", command)
	fmt.Print(warningStyle.Render() + prompt + " (y/N): ")

	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// ShowExecutingCommand displays that a command is being executed
func ShowExecutingCommand(command string) {
	fmt.Println(commandStyle.Render() + "Executing: " + command)
}

// ShowCommandError displays command execution error
func ShowCommandError(command string, err error) {
	fmt.Println(errorStyle.Render() + fmt.Sprintf("Command failed: %s - %v", command, err))
}

// ShowCommandSuccess displays successful command execution
func ShowCommandSuccess(command string) {
	fmt.Println(successStyle.Render() + "Command completed successfully")
}

// ShowStartingExecution displays start of batch execution
func ShowStartingExecution(count int) {
	fmt.Printf("\n%sStarting execution of %d commands...\n\n", commandStyle.Render(), count)
}

// ShowCommandProgress displays progress during batch execution
func ShowCommandProgress(current, total int) {
	fmt.Printf("%sCommand %d of %d:\n", commandStyle.Render(), current, total)
}

// ConfirmContinueOnError asks if execution should continue after error
func ConfirmContinueOnError() bool {
	fmt.Print(warningStyle.Render() + "Continue with remaining commands? (y/N): ")

	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))
	return response == "y" || response == "yes"
}

// ShowExecutionStopped displays that execution was stopped
func ShowExecutionStopped() {
	fmt.Println(warningStyle.Render() + "Execution stopped by user")
}

// ShowExecutionComplete displays completion of all commands
func ShowExecutionComplete() {
	fmt.Println(successStyle.Render() + "All commands executed successfully!")
}
