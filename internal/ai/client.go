package ai

import (
	"context"
	"os"
	"strings"

	"github.com/simplyzetax/oracle/internal/commands"
	"github.com/simplyzetax/oracle/internal/ui"
	"google.golang.org/genai"
)

// AskQuestion handles the AI interaction with streaming response and optional command execution
func AskQuestion(question, apiKey, model string, enableCommands bool) {
	// Get API key from parameter or environment
	if apiKey == "" {
		apiKey = os.Getenv("GOOGLE_AI_API_KEY")
	}

	if apiKey == "" {
		ui.ShowError("API key is required. Set GOOGLE_AI_API_KEY environment variable or use --api-key flag")
		return
	}

	ctx := context.Background()

	// Create client
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})
	if err != nil {
		ui.ShowError("Failed to create AI client: " + err.Error())
		return
	}

	// Enhanced system prompt to make AI provide executable commands
	systemPrompt := `You are Oracle, an AI assistant that can provide both answers and executable shell commands. 
When providing commands, format them clearly using:
- Code blocks with triple backticks and bash for multi-line commands
- Inline backticks for single commands  
- Or prefix with $ like: $ command

Be helpful and provide working commands when appropriate. Always explain what commands do before suggesting them. Make sure to avoid dangerous commands like "rm -rf" or "sudo rm" unless absolutely necessary. Always provide a brief and short response to avoid spamming the terminal.`

	// Show the enhanced question header
	ui.ShowQuestionHeader(model, question)

	// Stream the response and collect full text for command detection
	ui.StartResponseStream()

	var fullResponse strings.Builder

	for result, err := range client.Models.GenerateContentStream(
		ctx,
		model,
		genai.Text(systemPrompt+"\n\nUser question: "+question),
		&genai.GenerateContentConfig{
			Temperature: genai.Ptr(float32(0.7)),
		},
	) {
		if err != nil {
			ui.ShowError("Error generating content: " + err.Error())
			return
		}

		// Stream each chunk of text as it arrives
		text := result.Text()
		ui.StreamMarkdownText(text)
		fullResponse.WriteString(text)
	}

	ui.EndResponseStream()

	// Render the complete response with full markdown support
	ui.RenderFinalResponse(fullResponse.String())

	// Check for executable commands in the response (only if enabled)
	if enableCommands {
		detectedCommands := commands.ExtractCommands(fullResponse.String())
		if len(detectedCommands) > 0 {
			ui.ShowCommandsDetected(detectedCommands)
			commandsToExecute := commands.PromptToExecute(detectedCommands)
			if len(commandsToExecute) > 0 {
				commands.ExecuteCommands(commandsToExecute)
			}
		}
	}
}
