package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CaptureInput captures user input from the terminal and processes it.
func CaptureInput() []string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter comma-separated descriptions:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	descriptions := strings.Split(input, ",")
	return descriptions
}
