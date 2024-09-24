package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

type ChatGPTRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// CallChatGPT sends a request to the ChatGPT API and returns the response.
func CallChatGPT(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")

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

	// Here you would parse the actual response body to get the generated script
	return "Sample generated script text", nil // Placeholder, replace with actual response handling
}
