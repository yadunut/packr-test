package http

import (
	"fmt"

	"github.com/gobuffalo/packr"
)

var (
	templateBox packr.Box
)

func init() {
	templateBox = packr.NewBox("./templates")
	fmt.Println("Init Function")
	fmt.Printf("%d files in ./templates\n", len(templateBox.List()))
}

func ListFiles() string {
	fmt.Println("ListFiles function")
	s := "Files in box:\n"

	templateBox.Walk(func(name string, f packr.File) error {
		s += fmt.Sprintln(name)
		return nil
	})

	return s
}
