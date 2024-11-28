package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const apiUrl = "https://api.openai.com/v1/chat/completions"

type ChatGPTRequest struct {
	Model    string      `json:"model"`
	Messages []Message   `json:"messages"`
	MaxTokens int        `json:"max_tokens"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatGPTResponse struct {
	Choices []struct {
		Message Message `json:"message"`
	} `json:"choices"`
}

// SendMessageToChatGPT sends a message to the ChatGPT API and retrieves the response.
func SendMessageToChatGPT(prompt string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("OPENAI_API_KEY is not set")
	}

	requestBody := ChatGPTRequest{
		Model: "gpt-3.5-turbo",
		Messages: []Message{
			{Role: "user", Content: prompt},
		},
		MaxTokens: 150, // Adjust as needed
	}

	body, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf("failed to serialize request body: %w", err)
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := ioutil.ReadAll(resp.Body)
		return "", fmt.Errorf("non-OK HTTP status: %s - %s", resp.Status, string(respBody))
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var chatResponse ChatGPTResponse
	if err := json.Unmarshal(respBody, &chatResponse); err != nil {
		return "", fmt.Errorf("failed to parse response body: %w", err)
	}

	if len(chatResponse.Choices) == 0 {
		return "", fmt.Errorf("no response from ChatGPT")
	}

	return chatResponse.Choices[0].Message.Content, nil
}
