package cmd

import (
	"fmt"
	"os"

	"github.com/simplyzetax/oracle/internal/alias"
	"github.com/simplyzetax/oracle/internal/config"
	"github.com/simplyzetax/oracle/internal/ui"
	"github.com/spf13/cobra"
)

var (
	ApiKey         string
	Model          string
	EnableCommands bool
)

var RootCmd = &cobra.Command{
	Use:   "oracle",
	Short: "Oracle - AI CLI tool for asking questions to AI models",
	Long: `Oracle is a CLI tool that allows you to ask questions to AI models like Gemini with streaming responses.
	
Examples:
  oracle ask "What is the meaning of life?"
  oracle ask "Explain quantum computing" --model gemini-pro
  oracle ask "Write a haiku about coding" --api-key your-key`,
}

func Execute() {
	// Check for first run and offer alias setup
	if config.IsFirstRun() {
		ui.ShowFirstRunWelcome()

		if ui.ConfirmAliasSetup() {
			if err := alias.SetupAlias(); err != nil {
				ui.ShowError("Failed to set up alias automatically: " + err.Error())
				ui.ShowAliasInstructions()
			} else {
				ui.ShowAliasSetupSuccess()
			}
		} else {
			ui.ShowAliasInstructions()
		}

		// Mark first run as complete
		if err := config.MarkFirstRunComplete(); err != nil {
			// Non-fatal error, just continue
			fmt.Printf("Warning: Could not mark first run as complete: %v\n", err)
		}

		fmt.Println() // Add spacing
	}

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&ApiKey, "api-key", "k", "", "Google AI API key (can also use GOOGLE_AI_API_KEY env var)")
	RootCmd.PersistentFlags().StringVarP(&Model, "model", "m", "gemini-2.0-flash-exp", "AI model to use")
	RootCmd.PersistentFlags().BoolVarP(&EnableCommands, "execute", "x", false, "Enable command execution (allows Oracle to run shell commands)")
}
