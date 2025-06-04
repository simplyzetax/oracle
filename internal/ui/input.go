package ui

import (
	"os"
	"os/exec"
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
