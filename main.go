package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/firebase/genkit/go/ai"
	"github.com/firebase/genkit/go/genkit"
	"github.com/firebase/genkit/go/plugins/vertexai"
)

// menuSuggestioFlow is a genkit gen ai call
func menuSuggestionFlow(ctx context.Context, input string) (string, error) {
	m := vertexai.Model("gemini-1.5-flash")
	if m == nil {
		return "", errors.New("menuSuggestionFlow: failed to find model")
	}

	// Construct a request and send it to the model API (Google Vertex AI).
	resp, err := ai.Generate(ctx, m,
		ai.WithConfig(&ai.GenerationCommonConfig{Temperature: 1}),
		ai.WithTextPrompt(fmt.Sprintf(`Suggest an item for the menu of a %s themed restaurant`, input)))
	if err != nil {
		return "", err
	}

	text := resp.Text()
	return text, nil

}

// generateCoverageFlow creates coverage statement for a prompt statement
func generateCoverageFlow(ctx context.Context, input string) (string, error) {
	return input, nil
}

// generateQuestionsFlow creates questions and answers for a coverage statement
func generateQuestionsFlow(ctx context.Context, input string) (string, error) {
	return input, nil
}

func main() {
	ctx := context.Background()

	if err := vertexai.Init(ctx, nil); err != nil {
		log.Fatal(err)
	}

	// Define a simple flow that prompts an LLM to generate menu suggestions.
	genkit.DefineFlow("coverage", generateCoverageFlow)
	genkit.DefineFlow("questions", generateQuestionsFlow)
	genkit.DefineFlow("menuSuggestionFlow", menuSuggestionFlow)

	// Initialize Genkit and start a flow server.
	if err := genkit.Init(ctx, &genkit.Options{FlowAddr: ":3400"}); err != nil {
		log.Fatal(err)
	}
}
