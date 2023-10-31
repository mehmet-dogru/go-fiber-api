package config

import (
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("There is error in .env file")
	}

	return os.Getenv(key)
}
