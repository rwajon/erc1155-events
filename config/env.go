package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	GoEnv           string
	Port            string
	RpcURl          string
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
		AppName: getEnv("APP_NAME", "app"),
		Port:    getEnv("PORT", "3000"),
	}
	log.Println("envs.GoEnv---------", envs.GoEnv)
	switch envs.GoEnv {
	case "production":
		envs.RPCWebSocketURL, envs.DatabaseURL = getEnv("RPC_WS_URL"), getEnv("DB_URL")
	case "test":
		envs.RPCWebSocketURL = getEnv("TEST_RPC_WS_URL", "ws://localhost:8545")
		envs.DatabaseURL = getEnv("TEST_DB_URL", "mongodb://127.0.0.1:27017/erc1155_events_test")
	default:
		envs.RPCWebSocketURL, envs.DatabaseURL = getEnv("DEV_RPC_WS_URL"), getEnv("DEV_DB_URL")
	}

	return envs

}
