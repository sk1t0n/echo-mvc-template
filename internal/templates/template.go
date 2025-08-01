package templates

import (
	"cmp"
	"log"
	"net/http"
	"os"

	"github.com/open2b/scriggo"
	"github.com/open2b/scriggo/native"
)

func RenderTemplate(
	w http.ResponseWriter,
	file string,
	globals native.Declarations,
	vars map[string]any,
) error {
	contentLayout, err1 := os.ReadFile("internal/templates/layouts/default.html")
	content, err2 := os.ReadFile(file)
	if err := cmp.Or(err1, err2); err != nil {
		log.Printf("RenderTemplate: ReadFile: %v", err)
		return err
	}

	fsys := scriggo.Files{
		"internal/templates/layouts/default.html": contentLayout,
		file: content,
	}
	opts := &scriggo.BuildOptions{Globals: globals}
	template, err := scriggo.BuildTemplate(fsys, file, opts)
	if err != nil {
		log.Printf("RenderTemplate: BuildTemplate: %v", err)
		return err
	}

	err = template.Run(w, vars, nil)
	if err != nil {
		log.Printf("RenderTemplate: Run: %v", err)
		return err
	}

	return nil
}
