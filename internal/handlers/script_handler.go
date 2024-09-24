package handlers

import (
	"fmt"
	"strings"
)

// GeneratePrompt creates a prompt based on the first few input descriptions.
func GeneratePrompt(descriptions []string) string {
	mainTopics := strings.Join(descriptions[:min(3, len(descriptions))], ", ")
	return fmt.Sprintf("Generate a short video script based on the following keywords: %s.", mainTopics)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
