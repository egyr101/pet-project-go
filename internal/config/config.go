package config

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
}

func Load() (Config, error) {
	var cfg Config

	godotenv.Load()

	if err := envconfig.Process("", &cfg); err != nil {
		return Config{}, fmt.Errorf("error load config: %w", err)
	}

	return cfg, nil
}
