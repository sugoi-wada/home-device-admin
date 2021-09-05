package env

import (
	"os"

	"github.com/joho/godotenv"
)

func Get(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		// maybe production
	}

	return os.Getenv(key)
}
