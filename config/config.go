package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv reads the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using system environment variables")
	}
}

// GetEnv gets the value of the environment variable
func GetEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
