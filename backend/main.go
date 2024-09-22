package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter comma-separated descriptions:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Split input into separate descriptions
	descriptions := strings.Split(input, ",")
	prompt := generatePrompt(descriptions)

	// GPT api call
	response, err := callChatGPT(prompt)
	if err != nil {
		log.Fatalf("Error calling ChatGPT: %v", err)
	}

	fmt.Printf("Generated script: %s\n", response)
}

func generatePrompt(descriptions []string) string {
	// Use the first few descriptions to build a prompt
	mainTopics := strings.Join(descriptions[:min(3, len(descriptions))], ", ")
	return fmt.Sprintf("Generate a short video script based on the following keywords: %s.", mainTopics)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func callChatGPT(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY") // Ensure your API key is set in the environment

	reqBody := ChatGPTRequest{
		Model: "gpt-4",
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
	}

	jsonReq, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonReq))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Process the response (pseudo-code; replace with actual response handling)
	// Assume response processing extracts and returns the text completion
	return "Sample generated script text", nil // Replace with the actual script from the response
}
