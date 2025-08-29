package main

import (
	"fmt"
	"log"
	"test/internal/api"
	"test/internal/config"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(conf)

	r := api.RegisterRoutes()

	if err := r.Run(":8080"); err != nil {
		fmt.Println("error:", err)
	}
}
