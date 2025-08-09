package main

import (
	"log"

	"github.com/sk1t0n/fiber-mvc-template/config"
	"github.com/sk1t0n/fiber-mvc-template/internal/app"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(config)
}
