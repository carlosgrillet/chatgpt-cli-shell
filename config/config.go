package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from a .env file.
func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)

	return nil
}
