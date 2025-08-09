package templates

import (
	"cmp"
	"log/slog"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/native"
)

func RenderTemplate(
	c *fiber.Ctx,
	file string,
	globals native.Declarations,
	vars map[string]any,
) error {
	c.Response().Header.Set("Content-Type", "text/html; charset=utf-8")
	w := c.Response().BodyWriter()

	contentLayout, err1 := os.ReadFile("web/templates/layouts/base.html")
	content, err2 := os.ReadFile(file)
	if err := cmp.Or(err1, err2); err != nil {
		slog.Error("RenderTemplate: ReadFile:", "message", err)
		return err
	}

	fsys := scriggo.Files{
		"web/templates/layouts/base.html": contentLayout,
		file:                              content,
	}
	opts := &scriggo.BuildOptions{Globals: globals}
	template, err := scriggo.BuildTemplate(fsys, file, opts)
	if err != nil {
		slog.Error("RenderTemplate: BuildTemplate:", "message", err)
		return err
	}

	err = template.Run(w, vars, nil)
	if err != nil {
		slog.Error("RenderTemplate: Run:", "message", err)
		return err
	}

	return nil
}
