package commands

import (
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/simplyzetax/oracle/internal/ui"
)

// ExtractCommands finds potential shell commands in AI response text
func ExtractCommands(text string) []string {
	var commands []string
	seen := make(map[string]bool)

	// Pattern 1: Commands prefixed with $ (but clean them)
	dollarPattern := regexp.MustCompile(`(?m)^\s*\$\s+(.+)$`)
	matches := dollarPattern.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		if len(match) > 1 {
			cmd := strings.TrimSpace(match[1])
			if cmd != "" && !seen[cmd] && isValidCommand(cmd) {
				commands = append(commands, cmd)
				seen[cmd] = true
			}
		}
	}

	// Pattern 2: Code blocks with bash/shell content (multi-line commands)
	codeBlockPattern := regexp.MustCompile("```(?:bash|sh|shell|zsh)?\n((?:[^`]|`[^`]|``[^`])*?)\n```")
	matches = codeBlockPattern.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		if len(match) > 1 {
			blockContent := strings.TrimSpace(match[1])
			// Split multi-line commands and clean them
			lines := strings.Split(blockContent, "\n")
			for _, line := range lines {
				line = strings.TrimSpace(line)
				// Remove $ prefix if present
				if strings.HasPrefix(line, "$") {
					line = strings.TrimSpace(line[1:])
				}
				if line != "" && !seen[line] && isValidCommand(line) && !strings.HasPrefix(line, "#") {
					commands = append(commands, line)
					seen[line] = true
				}
			}
		}
	}

	// Pattern 3: Single backtick commands (but be more selective)
	inlinePattern := regexp.MustCompile("`([^`\n]+)`")
	matches = inlinePattern.FindAllStringSubmatch(text, -1)
	for _, match := range matches {
		if len(match) > 1 {
			cmd := strings.TrimSpace(match[1])
			// Remove $ prefix if present
			if strings.HasPrefix(cmd, "$") {
				cmd = strings.TrimSpace(cmd[1:])
			}
			// Only include if it looks like a shell command (starts with common command words)
			if cmd != "" && !seen[cmd] && isLikelyShellCommand(cmd) && isValidCommand(cmd) {
				commands = append(commands, cmd)
				seen[cmd] = true
			}
		}
	}

	// Remove similar/redundant commands before returning
	return removeSimilarCommands(commands)
}

// removeSimilarCommands filters out redundant similar commands, keeping the most specific ones
func removeSimilarCommands(commands []string) []string {
	if len(commands) <= 1 {
		return commands
	}

	var filtered []string

	for i, cmd1 := range commands {
		shouldInclude := true

		// Check if this command is a subset of another more specific command
		for j, cmd2 := range commands {
			if i != j && isCommandSubset(cmd1, cmd2) {
				shouldInclude = false
				break
			}
		}

		if shouldInclude {
			filtered = append(filtered, cmd1)
		}
	}

	return filtered
}

// isCommandSubset checks if cmd1 is a subset of cmd2 (cmd2 is more specific)
func isCommandSubset(cmd1, cmd2 string) bool {
	words1 := strings.Fields(cmd1)
	words2 := strings.Fields(cmd2)

	// If cmd1 has fewer words and all its words are at the beginning of cmd2
	if len(words1) < len(words2) {
		for i, word := range words1 {
			if i >= len(words2) || word != words2[i] {
				return false
			}
		}
		return true
	}

	return false
}

// isLikelyShellCommand checks if a string looks like a shell command
func isLikelyShellCommand(cmd string) bool {
	// Common shell command prefixes
	commonCommands := []string{
		"ls", "cd", "pwd", "mkdir", "rmdir", "rm", "cp", "mv", "cat", "grep", "find", "sort",
		"awk", "sed", "head", "tail", "wc", "chmod", "chown", "ps", "kill", "top", "df", "du",
		"tar", "zip", "unzip", "curl", "wget", "ssh", "scp", "git", "docker", "npm", "yarn",
		"pip", "python", "node", "go", "java", "gcc", "make", "cmake", "echo", "printf",
		"which", "whereis", "alias", "export", "source", "history", "jobs", "nohup",
	}

	// Get the first word of the command
	words := strings.Fields(cmd)
	if len(words) == 0 {
		return false
	}

	firstWord := strings.ToLower(words[0])

	// Check if it starts with a common command
	for _, common := range commonCommands {
		if firstWord == common {
			return true
		}
	}

	// Check if it contains shell operators or looks like a path
	if strings.Contains(cmd, "/") || strings.Contains(cmd, "|") || strings.Contains(cmd, ">") || strings.Contains(cmd, "<") {
		return true
	}

	return false
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

// PromptToExecute asks the user if they want to execute detected commands with simple prompts
func PromptToExecute(commands []string) []string {
	if len(commands) == 0 {
		return nil
	}

	var toExecute []string

	for _, cmd := range commands {
		ui.ShowCommandSuggestion(cmd)
		if ui.ConfirmExecution(cmd) {
			toExecute = append(toExecute, cmd)
		}
	}

	return toExecute
}

// ExecuteCommand runs a shell command with minimal output
func ExecuteCommand(command string) error {
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
		fmt.Printf("Command failed: %s\n", command)
		return err
	}

	return nil
}

// ExecuteCommands runs multiple commands in sequence with minimal logging
func ExecuteCommands(commands []string) {
	if len(commands) == 0 {
		return
	}

	for _, cmd := range commands {
		if err := ExecuteCommand(cmd); err != nil {
			if !ui.ConfirmContinueOnError() {
				return
			}
		}
	}
}
