package cmd

import (
	"strings"

	"github.com/simplyzetax/oracle/internal/ai"
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
