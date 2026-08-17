package main

import (
	"context"
	"pet-project/internal/config"
	"pet-project/internal/storage"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	db, err := storage.NewRepository(context.Background(), cfg.Postgres)
	if err != nil {
		panic(err)
	}
}
