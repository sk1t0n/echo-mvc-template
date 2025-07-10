package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/sk1t0n/echo-mvc-template/internal/app/http/controllers"
)

func RegisterRoutes(e *echo.Echo) {
	homeController := controllers.NewHomeController()

	e.Static("/", "web")
	e.GET("/", homeController.Index)
}
