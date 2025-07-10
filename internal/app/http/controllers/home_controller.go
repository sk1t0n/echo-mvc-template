package controllers

import (
	"github.com/labstack/echo/v4"

	"github.com/sk1t0n/echo-mvc-template/internal/templates"
)

type HomeController struct {
}

func NewHomeController() HomeController {
	return HomeController{}
}

func (HomeController) Index(c echo.Context) error {
	w := c.Response().Writer
	bindings := map[string]any{
		"page": map[string]string{
			"title": "Home page",
		},
	}
	templates.LoadTemplate(w, "internal/templates/index.html", bindings)

	return nil
}
