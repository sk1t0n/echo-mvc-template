package templates

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/osteele/liquid.v1"
)

func LoadTemplate(w http.ResponseWriter, path string, bindings map[string]any) {
	rawTemplate, err := os.ReadFile(path)
	if err != nil {
		log.Fatalln(err)
	}

	engine := liquid.NewEngine()
	engine.RegisterFilter("asset_url", func(text string) string {
		// local or S3
		return text
	})

	out, err := engine.ParseAndRenderString(string(rawTemplate), bindings)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Fprint(w, out)
}
