package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDBUrl() string {
	return os.Getenv("DATABASE_URL")
}

func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}

	return "8080"
}
