package main

import (
	"context"
	"fmt"
	"log"
	"test/internal/api"
	"test/internal/config"
	"test/internal/store"
)

func main() {
	ctx := context.Background()

	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	_, err = store.NewDbConn(ctx, conf)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	r := api.RegisterRoutes()

	if err := r.Run(conf.AppPort); err != nil {
		fmt.Println("error:", err)
	}
}
