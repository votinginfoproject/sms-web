package env

import "github.com/joho/godotenv"

func Load() {
	ok := godotenv.Load()

	if ok != nil {
		panic("Error loading .env file")
	}
}
