package homeController

import (
	"fmt"
	"net/http"
)

// Home Handler provides the base of the application
func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Welcome Home </h1>")
}

// NotFound Handler is called when a route give
// ia not present in the routes
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1> Sorry page not found </h1>")
}
