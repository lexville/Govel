package home

import (
	"fmt"
	"net/http"

	"Govel/view"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	views.RenderView(w, nil, "resources/views/home/home.html")
}

// NotFound Handler is called when a route give
// ia not present in the routes
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Sorry page not found </h1>")
}
