package initializers

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnvVar() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}
}
