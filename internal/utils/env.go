package utils

import (
	"github.com/lpernett/godotenv"
	"log"
	"os"
	"strconv"
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

func LoadDbEnvVariables() (string, int, string, string, string) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		log.Fatal("Missing one or more required POSTGRESS environment variables (DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)")
	}

	var portInt int
	portInt, _ = strconv.Atoi(port)

	return host, portInt, user, password, dbname
}
