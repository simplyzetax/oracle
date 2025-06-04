package ui

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// PromptForQuestion prompts the user to enter a question interactively using gum
func PromptForQuestion() string {
	// Use gum input via command execution for compatibility
	cmd := exec.Command("gum", "input", "--placeholder", "Enter your question here...", "--prompt", "What would you like to ask? ")
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		ShowError("Failed to get input: " + err.Error())
		return ""
	}

	// Remove trailing newline
	question := string(output)
	if len(question) > 0 && question[len(question)-1] == '\n' {
		question = question[:len(question)-1]
	}

	return question
}

// PromptForAPIKey prompts the user to enter their Google AI API key
func PromptForAPIKey() string {
	fmt.Print("Enter your Google AI API Key (get it from https://ai.google.dev/gemini-api/docs/api-key): ")

	var apiKey string
	_, err := fmt.Scanln(&apiKey)
	if err != nil {
		ShowError("Failed to read API key: " + err.Error())
		return ""
	}

	apiKey = strings.TrimSpace(apiKey)
	if apiKey == "" {
		ShowError("API key cannot be empty")
		return ""
	}

	return apiKey
}

// ShowAPIKeySetupSuccess displays success message for API key setup
func ShowAPIKeySetupSuccess() {
	success := lipgloss.NewStyle().
		Foreground(green).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(green).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1).
		Render("🔑 API Key configured successfully!\n\nYour Google AI API key has been saved securely to ~/.oracle/config.json\nYou're now ready to start asking questions!")

	fmt.Println(success)
}

// ShowAPIKeyPrompt displays a message about needing to set up the API key
func ShowAPIKeyPrompt() {
	prompt := lipgloss.NewStyle().
		Foreground(gold).
		Bold(true).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(blue).
		Padding(1, 2).
		MarginTop(1).
		MarginBottom(1).
		Render("🔑 API Key Required\n\nTo use Oracle, you need a Google AI API key.\nThis will be saved securely in your local configuration.")

	fmt.Println(prompt)
}
