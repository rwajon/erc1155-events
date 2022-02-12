package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GoEnv   string
	Port    string
	AppName string
}

func getEnv(key string, defaultValue string) string {
	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}
	return val
}

func GetEnvs() Env {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	return Env{
		GoEnv:   getEnv("GO_ENV", "development"),
		Port:    getEnv("PORT", "3000"),
		AppName: getEnv("APP_NAME", ""),
	}
}
