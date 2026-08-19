package main

import (
	"context"
	"log"
	"pet-project/internal/app"
	"pet-project/internal/config"
	"pet-project/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	pool, err := database.NewPool(context.Background(), cfg.PostgresConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	err = app.CreateDb(context.Background(), pool, database.CreateTableUser, database.CreateTableCategory, database.CreateTableProduct, database.CreateTableOrder, database.CreateTableOrderItem)
	if err != nil {
		log.Fatal(err)
	}

}
