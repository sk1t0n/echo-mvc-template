package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/sk1t0n/fiber-mvc-template/config"
	"github.com/sk1t0n/fiber-mvc-template/internal/controller/http/errorhandler"
	"github.com/sk1t0n/fiber-mvc-template/internal/controller/http/router"
)

func Run(cfg *config.Config) {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorhandler.New(),
	})

	router.New(app).Setup()

	err := app.Listen("127.0.0.1:" + cfg.Http.Port)
	if err != nil {
		fmt.Printf("Server error: %s", err)
	}
}
