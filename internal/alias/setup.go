package alias

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/simplyzetax/oracle/internal/ui"
)

// SetupAlias attempts to automatically set up the 'oa' alias
func SetupAlias() error {
	shell := os.Getenv("SHELL")
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	var configFile string
	var aliasLine string

	switch {
	case strings.Contains(shell, "zsh"):
		configFile = filepath.Join(homeDir, ".zshrc")
		aliasLine = "alias oa='oracle ask'"
	case strings.Contains(shell, "bash"):
		// Try .bashrc first, then .bash_profile
		configFile = filepath.Join(homeDir, ".bashrc")
		if _, err := os.Stat(configFile); os.IsNotExist(err) {
			configFile = filepath.Join(homeDir, ".bash_profile")
		}
		aliasLine = "alias oa='oracle ask'"
	case strings.Contains(shell, "fish"):
		configFile = filepath.Join(homeDir, ".config", "fish", "config.fish")
		aliasLine = "alias oa='oracle ask'"
	default:
		return fmt.Errorf("unsupported shell: %s", shell)
	}

	// Check if alias already exists
	if aliasExists(configFile, aliasLine) {
		ui.ShowSuccess("Alias 'oa' already exists in your shell config!")
		return nil
	}

	// Create config directory if it doesn't exist (for fish)
	configDir := filepath.Dir(configFile)
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return fmt.Errorf("failed to create config directory: %w", err)
	}

	// Append alias to config file
	file, err := os.OpenFile(configFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer file.Close()

	// Add a comment and the alias
	_, err = file.WriteString(fmt.Sprintf("\n# Oracle CLI alias\n%s\n", aliasLine))
	if err != nil {
		return fmt.Errorf("failed to write alias to config file: %w", err)
	}

	return nil
}

// aliasExists checks if the alias already exists in the config file
func aliasExists(configFile, aliasLine string) bool {
	file, err := os.Open(configFile)
	if err != nil {
		return false // File doesn't exist, so alias doesn't exist
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "alias oa=") {
			return true
		}
	}

	return false
}
