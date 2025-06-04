package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	ApiKey string
	Model  string
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
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&ApiKey, "api-key", "k", "", "Google AI API key (can also use GOOGLE_AI_API_KEY env var)")
	RootCmd.PersistentFlags().StringVarP(&Model, "model", "m", "gemini-2.0-flash-exp", "AI model to use")
}
