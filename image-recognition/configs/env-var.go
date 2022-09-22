package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvS3APIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("S3_API_KEY")
}
