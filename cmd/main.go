package main

import (
	"context"
	"fmt"
	"log"
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

	fmt.Println()
	defer pool.Close()

}
