package main

import (
	"fmt"
	"log"

	"github.com/perryz0/auto-shorts-pipeline/internal/handlers"
)

func main() {
	// Phase 1: Handle user input
	descriptions := handlers.CaptureInput()

	// Phase 2: Generate a prompt from user input
	prompt := handlers.GeneratePrompt(descriptions)

	// Phase 3: Call GPT API with the generated prompt
	response, err := handlers.CallChatGPT(prompt)
	if err != nil {
		log.Fatalf("Error calling ChatGPT: %v", err)
	}

	fmt.Printf("Generated script: %s\n", response)
}
