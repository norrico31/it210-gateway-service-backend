package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	JWTSecret   string
	BaseURL     string
	AppVersion1 string
	AuthPath    string
	CorePath    string
	AuthPort    string
	CorePort    string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		BaseURL:     getEnv("BASE_URL", "http://127.0.0.1"),
		AppVersion1: getEnv("APP_VERSION_1", "v1"),
		AuthPath:    getEnv("AUTH_PATH", "auth"),
		CorePath:    getEnv("CORE_PATH", "core"),
		AuthPort:    getEnv("AUTH_PORT", "8081"),
		CorePort:    getEnv("CORE_PORT", "8082"),
		JWTSecret:   getEnv("JWT_SECRET", "IS-IT_REALL-A_SECRET-?~JWT-NOT_SO-SURE"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
