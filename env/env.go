package env

import (
	"log"

	"github.com/joho/godotenv"
)

func Load() {
	err := godotenv.Load()

	if err != nil {
		log.Panic("Error loading .env file")
	}
}
