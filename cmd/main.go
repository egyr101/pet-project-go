package main

import (
	"pet-project/internal/app"
	"pet-project/internal/config"
	"pet-project/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal(err.Error())
	}

	app, err := app.New(&cfg)
	err = app.Run()
	if err != nil {
		logger.Error(err.Error())
	}

	// err = app.CreateDb(context.Background(), pool, database.CreateTableUser, database.CreateTableCategory, database.CreateTableProduct, database.CreateTableOrder, database.CreateTableOrderItem)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
