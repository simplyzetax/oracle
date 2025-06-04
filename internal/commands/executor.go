package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/simplyzetax/oracle/internal/ui"
)

// CommandPattern defines patterns to detect shell commands in AI responses
var CommandPatterns = []*regexp.Regexp{
	regexp.MustCompile(`(?m)^\$ (.+)$`),                                        // $ command
	regexp.MustCompile(`(?m)^> (.+)$`),                                         // > command
	regexp.MustCompile("```(?:bash|sh|shell)?\n((?:[^`]|`[^`]|``[^`])*)\n```"), // ```bash code blocks
	regexp.MustCompile("`([^`]+)`"),                                            // `inline commands`
}

// ExtractCommands finds potential shell commands in AI response text
func ExtractCommands(text string) []string {
	var commands []string
	seen := make(map[string]bool)

	for _, pattern := range CommandPatterns {
		matches := pattern.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if len(match) > 1 {
				cmd := strings.TrimSpace(match[1])
				if cmd != "" && !seen[cmd] && isValidCommand(cmd) {
					commands = append(commands, cmd)
					seen[cmd] = true
				}
			}
		}
	}

	return commands
}

// isValidCommand checks if a command looks safe to execute
func isValidCommand(cmd string) bool {
	// Skip obviously unsafe commands
	dangerousCommands := []string{
		"rm -rf", "sudo rm", "dd if=", ":(){ :|:& };:",
		"chmod 777", "chown", "mkfs", "fdisk",
		"shutdown", "reboot", "halt", "poweroff",
	}

	cmdLower := strings.ToLower(cmd)
	for _, dangerous := range dangerousCommands {
		if strings.Contains(cmdLower, dangerous) {
			return false
		}
	}

	// Skip very long commands (might be code, not commands)
	if len(cmd) > 200 {
		return false
	}

	// Skip commands with suspicious patterns
	if strings.Contains(cmd, "&&") && strings.Contains(cmd, "rm") {
		return false
	}

	return true
}

// PromptToExecute asks the user if they want to execute detected commands
func PromptToExecute(commands []string) []string {
	if len(commands) == 0 {
		return nil
	}

	ui.ShowCommandsTable(commands) // Updated from ShowCommandsDetected

	var toExecute []string

	for _, cmd := range commands {
		if ui.ConfirmExecution(cmd) {
			toExecute = append(toExecute, cmd)
		}
	}

	return toExecute
}

// ExecuteCommand runs a shell command with proper output handling
func ExecuteCommand(command string) error {
	ui.ShowExecutionStatus(fmt.Sprintf("Executing: %s", command), "executing") // Updated

	// Use the user's default shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		shell = "/bin/sh"
	}

	cmd := exec.Command(shell, "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		ui.ShowExecutionStatus(fmt.Sprintf("Command failed: %s - %v", command, err), "error") // Updated
		return err
	}

	ui.ShowExecutionStatus(fmt.Sprintf("Command completed: %s", command), "success") // Updated
	return nil
}

// ExecuteCommands runs multiple commands in sequence
func ExecuteCommands(commands []string) {
	if len(commands) == 0 {
		return
	}

	ui.ShowExecutionStatus(fmt.Sprintf("Starting execution of %d commands...", len(commands)), "info") // Updated

	for i, cmd := range commands {
		ui.ShowExecutionStatus(fmt.Sprintf("Command %d of %d:", i+1, len(commands)), "info") // Updated

		if err := ExecuteCommand(cmd); err != nil {
			if !ui.ConfirmContinueOnError() {
				ui.ShowExecutionStatus("Execution stopped by user.", "warning") // Updated
				return
			}
		}
	}

	ui.ShowExecutionStatus("All commands executed.", "success") // Updated
}
