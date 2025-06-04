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

	// Simplified system prompt for commands
	systemPrompt := `You are Oracle, an AI assistant that provides answers and executable shell commands.
Format commands clearly using:
- Code blocks with triple backticks for multi-line commands
- Inline backticks for single commands
- Prefix with $ for commands

Explain what commands do before suggesting them. Avoid dangerous commands and keep responses concise. Again, keep the response length to a maximum of 3 sentences.`

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
		fullResponse.WriteString(text)
	}

	// Only render final response if it would look different from the streamed text
	// This prevents duplicate text in the terminal
	if strings.Contains(fullResponse.String(), "```") ||
		strings.Contains(fullResponse.String(), "#") ||
		strings.Contains(fullResponse.String(), "*") {
		// Response likely contains markdown, so render with formatting
		ui.RenderFinalResponse(fullResponse.String())
	}

	// Check for executable commands in the response (only if enabled)
	if enableCommands {
		detectedCommands := commands.ExtractCommands(fullResponse.String())
		if len(detectedCommands) > 0 {
			ui.ShowCommandsTable(detectedCommands) // Changed from ShowCommandsDetected
			commandsToExecute := commands.PromptToExecute(detectedCommands)
			if len(commandsToExecute) > 0 {
				commands.ExecuteCommands(commandsToExecute)
			}
		}
	}
}
