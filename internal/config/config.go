package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	HTTPPort string
	DatabaseURL string
}

func Load() (*Config, error){
	err := godotenv.Load()
	if err != nil{
		return nil, ErrorEnvNotFound
	}

	port := os.Getenv("HTTP_Port")
	if port == ""{
		return nil, ErrorPortNotFound
	}

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl == ""{
		return nil, ErrorDatabaseUrlNotFound
	}

	return &Config{
		HTTPPort: port,
		DatabaseURL: databaseUrl,
	}, nil
}