package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	layoutDirectory = "resources/views/layouts/"
	extension       = ".html"
)

// RenderView is responsible for rendering the
// view needed
func RenderView(w http.ResponseWriter, data interface{}, files ...string) {
	t := createView(files...)
	t.ExecuteTemplate(w, "main", data)
}

// createView is used to parse the files
// needed for the view
func createView(files ...string) *template.Template {
	files = append(
		files,
		getLayoutFiles()...,
	)
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return t
}

// getLayoutFiles returns all the layouts files
// as a slice of strings
func getLayoutFiles() []string {
	files, err := filepath.Glob(layoutDirectory + "*" + extension)
	if err != nil {
		panic(err)
	}
	return files
}
