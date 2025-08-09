package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/open2b/scriggo/native"

	"github.com/sk1t0n/fiber-mvc-template/web/templates"
)

type HomeController struct {
}

func NewHomeController() HomeController {
	return HomeController{}
}

func (HomeController) Index(c *fiber.Ctx) error {
	globals := native.Declarations{
		"title": "Home page",
	}
	vars := map[string]any{}

	template := "web/templates/home/index.html"
	err := templates.RenderTemplate(c, template, globals, vars)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"failed to render template "+template,
		)
	}

	return nil
}
