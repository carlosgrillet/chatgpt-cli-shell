package handlers

import (
	"chatgpt-cli-shell/utils"
	"fmt"
)

// HandleChat sends the user's input to the ChatGPT API and retrieves the response.
func HandleChat(input string) (string, error) {
	response, err := utils.SendMessageToChatGPT(input)
	if err != nil {
		return "", fmt.Errorf("failed to get response: %w", err)
	}
	return response, nil
}
