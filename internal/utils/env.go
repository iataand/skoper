package utils

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

func LoadEnvVariables() (string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("HACKERONE_USERNAME")
	apiKey := os.Getenv("HACKERONE_API_KEY")

	if user == "" || apiKey == "" {
		log.Fatal("Missing HACKERONE_USERNAME or HACKERONE_API_KEY in environment")
	}

	return user, apiKey
}
