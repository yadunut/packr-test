package main

import (
	"os"

	"github.com/yadunut/test-packr/http"
)

func main() {
	http.RenderTemplate(os.Stdout, "index.tmpl", "World")
}
