package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"chatgpt-cli-shell/handlers"
)

func main() {
	fmt.Println("Welcome to ChatGPT CLI! Type your question below:")
	fmt.Println("Type 'exit' to quit.")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$ ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "exit" {
			fmt.Println("Goodbye!")
			break
		}
		response, err := handlers.HandleChat(input)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		fmt.Println(response)
	}
}
