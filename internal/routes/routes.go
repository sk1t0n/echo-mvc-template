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

func registerResource(e *echo.Echo, name string, controller controllers.Controller) {
	e.GET("/"+name, controller.Index)
	e.GET("/"+name+"/create", controller.Create)
	e.POST("/"+name, controller.Store)
	e.GET("/"+name+"/:id", controller.Show)
	e.GET("/"+name+"/:id/edit", controller.Edit)
	e.PUT("/"+name+"/:id", controller.Update)
	e.DELETE("/"+name+"/:id", controller.Destroy)
}
