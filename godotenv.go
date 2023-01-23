package main

import (
	"github.com/joho/godotenv"
	"os"
)

func main() {
	// Create a map of environment variables
	envVars := map[string]string{
		"API_KEY": "abc123",
		"DEBUG": "true",
	}

	// Write the environment variables to a file named .env
	err := godotenv.Write(envVars, ".env")
	if err != nil {
		panic(err)
	}
}


func main2() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	fmt.Println("API Key:", apiKey)
}
