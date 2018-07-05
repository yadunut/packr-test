package http

import (
	"html/template"
	"io"

	"fmt"

	"github.com/gobuffalo/packr"
)

var (
	assets    packr.Box
	templates *template.Template
)

func init() {
	assets = packr.NewBox("./assets")
	templateBox := packr.NewBox("./templates")
	templates = template.New("")

	if len(templateBox.List()) == 0 {
		fmt.Println("0 files in ./templates")
	}

	templateBox.Walk(func(s string, _ packr.File) error {
		template.Must(templates.New(s).Parse(templateBox.String(s)))
		return nil
	})
}

func RenderTemplate(w io.Writer, template string, data interface{}) {
	if err := templates.ExecuteTemplate(w, template, data); err != nil {
		fmt.Println(fmt.Errorf("error parsing template: %v", err))
	}
}
