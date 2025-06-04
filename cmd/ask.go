package cmd

import (
	"fmt"
	"strings"

	"github.com/simplyzetax/oracle/internal/ai"
	"github.com/simplyzetax/oracle/internal/config"
	"github.com/simplyzetax/oracle/internal/ui"
	"github.com/spf13/cobra"
)

var askCmd = &cobra.Command{
	Use:   "ask [question]",
	Short: "Ask a question to the AI model",
	Long: `Ask a question to the AI model and get a streaming response.
	
The question can be provided as arguments or you'll be prompted to enter it interactively.

Examples:
  oracle ask "What is the meaning of life?"
  oracle ask "Explain quantum computing in simple terms"
  oracle ask`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check for API key and prompt if needed
		if err := checkAndSetupAPIKey(); err != nil {
			ui.ShowError("Failed to setup API key: " + err.Error())
			return
		}

		var question string

		if len(args) == 0 {
			// If no question provided, prompt for it
			question = ui.PromptForQuestion()
		} else {
			question = strings.Join(args, " ")
		}

		if question == "" {
			ui.ShowError("No question provided")
			return
		}

		ai.AskQuestion(question, ApiKey, Model, EnableCommands)
	},
}

func init() {
	RootCmd.AddCommand(askCmd)
}

// checkAndSetupAPIKey checks if API key is available and prompts for it if needed
func checkAndSetupAPIKey() error {
	// Check if API key is available from any source
	apiKey, err := config.GetAPIKey(ApiKey)
	if err != nil {
		return fmt.Errorf("failed to check API key: %w", err)
	}

	// If we have an API key, we're good
	if apiKey != "" {
		ApiKey = apiKey
		return nil
	}

	// No API key found anywhere, prompt for it
	ui.ShowAPIKeyPrompt()
	newAPIKey := ui.PromptForAPIKey()
	if newAPIKey == "" {
		return fmt.Errorf("API key is required to use Oracle")
	}

	// Save the API key to config
	if err := config.SetAPIKey(newAPIKey); err != nil {
		return fmt.Errorf("failed to save API key to config: %w", err)
	}

	// Set the global variable
	ApiKey = newAPIKey

	// Show success message
	ui.ShowAPIKeySetupSuccess()

	return nil
}
