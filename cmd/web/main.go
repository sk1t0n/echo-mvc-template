package main

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/sk1t0n/echo-mvc-template/internal/routes"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	routes.RegisterRoutes(e)

	err := e.Start(":8080")
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		slog.Error("failed to start", "error", err)
	}
}
