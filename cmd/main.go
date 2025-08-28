package main

import (
	"context"
	"log"
	"test/internal/api"
	"test/internal/config"
	"test/internal/service"
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

	api.RegisterRoutes(service.New(db), conf.AppPort)
}
