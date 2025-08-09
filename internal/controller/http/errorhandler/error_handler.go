package errorhandler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/open2b/scriggo/native"

	"github.com/sk1t0n/fiber-mvc-template/web/templates"
)

func New() func(*fiber.Ctx, error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		slog.Error("ErrorHandler", "code", code, "message", err)
		c.Status(code)
		c.Response().ResetBody()

		switch code {
		case fiber.StatusForbidden:
			globals := native.Declarations{"title": "403 - Forbidden"}
			return templates.RenderTemplate(c, "web/templates/errors/403.html", globals, nil)
		case fiber.StatusNotFound:
			globals := native.Declarations{"title": "404 - Not Found"}
			return templates.RenderTemplate(c, "web/templates/errors/404.html", globals, nil)
		default:
			globals := native.Declarations{"title": "500 - Internal Server Error"}
			return templates.RenderTemplate(c, "web/templates/errors/500.html", globals, nil)
		}
	}
}
