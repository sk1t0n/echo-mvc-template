package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sk1t0n/fiber-mvc-template/internal/controller/http"
	"github.com/sk1t0n/fiber-mvc-template/internal/controller/http/middleware"
)

type Router struct {
	app *fiber.App
}

func New(app *fiber.App) *Router {
	return &Router{app: app}
}

func (r *Router) Setup() {
	r.registerMiddleware()
	r.registerRoutes()
}

func (r *Router) registerMiddleware() {
	r.app.Use(middleware.Recovery())
}

func (r *Router) registerRoutes() {
	homeController := http.NewHomeController()

	r.app.Static("/", "web/build")
	r.app.Get("/", homeController.Index)
}

func (r *Router) registerResource(name string, c http.Controller) {
	r.app.Get("/"+name, c.Index)
	r.app.Get("/"+name+"/create", c.Create)
	r.app.Post("/"+name, c.Store)
	r.app.Get("/"+name+"/:id", c.Show)
	r.app.Get("/"+name+"/:id/edit", c.Edit)
	r.app.Put("/"+name+"/:id", c.Update)
	r.app.Delete("/"+name+"/:id", c.Destroy)
}
