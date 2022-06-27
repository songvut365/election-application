package config

import (
	"log"

	"github.com/joho/godotenv"
)

func SetupEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Panic("Error loading .env file")
	}
}
