package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GoEnv           string
	Port            string
	RPCWebSocketURL string
	AppName         string
	DatabaseURL     string
}

func getEnv(key string, defaultValue ...string) string {
	val := os.Getenv(key)
	if val == "" && len(defaultValue) > 0 && defaultValue[0] != "" {
		return defaultValue[0]
	}
	return val
}

func GetEnvs() Env {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	envs := Env{
		GoEnv:   getEnv("GO_ENV", "development"),
		AppName: getEnv("APP_NAME", "erc1155-events"),
		Port:    getEnv("PORT", "3000"),
	}

	switch envs.GoEnv {
	case "development":
		envs.RPCWebSocketURL, envs.DatabaseURL = getEnv("DEV_RPC_WS_URL"), getEnv("DEV_DB_URL")
	case "test":
		envs.RPCWebSocketURL = getEnv("TEST_RPC_WS_URL", "ws://localhost:8545")
		envs.DatabaseURL = getEnv("TEST_DB_URL", "mongodb://localhost:27017/erc1155_events_test")
	case "production":
		envs.RPCWebSocketURL, envs.DatabaseURL = getEnv("RPC_WS_URL"), getEnv("DB_URL")
	}

	return envs

}
