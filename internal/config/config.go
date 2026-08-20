package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ServerConfig
	PostgresConfig
}

func Load() (Config, error) {
	var cfg Config

	if err := godotenv.Load("../.env"); err != nil {
		return Config{}, fmt.Errorf("error godotenv load env: %w", err)
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("error load config: %w", err)
	}

	return cfg, nil
}
