package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/open2b/scriggo/native"

	"github.com/sk1t0n/echo-mvc-template/internal/templates"
)

type HomeController struct {
}

func NewHomeController() HomeController {
	return HomeController{}
}

func (HomeController) Index(c echo.Context) error {
	w := c.Response().Writer
	globals := native.Declarations{
		"title": "Home page",
	}
	vars := map[string]any{}

	err := templates.RenderTemplate(w, "internal/templates/index.html", globals, vars)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something unexpected happened")
	}

	return nil
}
