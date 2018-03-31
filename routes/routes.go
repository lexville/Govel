package routes

import (
	"log"
	"net/http"

	"Govel/app/controllers/homeController"

	"github.com/gorilla/mux"
)

// InitRoutes stores all the
// routes used in the application
func InitRoutes() {
	r := mux.NewRouter()
	r.HandleFunc("/", homeController.Home)
	r.NotFoundHandler = http.HandlerFunc(homeController.NotFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
