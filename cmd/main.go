package main

import (
	"context"
	"pet-project/internal/config"
	"pet-project/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	pool, err := database.NewPool(context.Background(), cfg.Postgres)
	if err != nil {
		panic(err)
	}
}
