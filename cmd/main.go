package main

import (
	"context"
	"log"
	"test/internal/api"
	handlers "test/internal/api/handlers/v2"
	"test/internal/config"
	service "test/internal/service/v2"
	"test/internal/store"
)

func main() {
	ctx := context.Background()

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.NewDbConn(ctx, conf)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	defer db.Close()

	engine := api.Register(handlers.NewCvHandler(service.NewCvService(db)))

	if err = engine.Run(conf.AppPort); err != nil {
		log.Fatal(err)
	}
}
