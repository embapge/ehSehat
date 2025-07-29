package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	GRPCPort string
}

// LoadEnv loads .env file and returns EnvConfig struct
func LoadEnv() *EnvConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return &EnvConfig{
		DBHost:   os.Getenv("DB_HOST"),
		DBPort:   os.Getenv("DB_PORT"),
		DBUser:   os.Getenv("DB_USER"),
		DBPass:   os.Getenv("DB_PASS"),
		DBName:   os.Getenv("DB_NAME"),
		GRPCPort: os.Getenv("GRPC_PORT"),
	}
}
