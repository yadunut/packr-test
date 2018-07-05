package main

import (
	"os"

	"github.com/yadunut/packr-test/http"
)

func main() {
	http.RenderTemplate(os.Stdout, "index.tmpl", "World")
}
