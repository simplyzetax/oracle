package ai

import (
	"context"
	"os"

	"github.com/simplyzetax/oracle/internal/ui"
	"google.golang.org/genai"
)

// AskQuestion handles the AI interaction with streaming response
func AskQuestion(question, apiKey, model string) {
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

	// Show the question being asked
	ui.ShowQuestionHeader(model, question)

	// Stream the response
	ui.StartResponseStream()

	for result, err := range client.Models.GenerateContentStream(
		ctx,
		model,
		genai.Text(question),
		&genai.GenerateContentConfig{
			Temperature: genai.Ptr(float32(0.7)),
		},
	) {
		if err != nil {
			ui.ShowError("Error generating content: " + err.Error())
			return
		}

		// Stream each chunk of text as it arrives
		ui.StreamText(result.Text())
	}

	ui.EndResponseStream()
}
