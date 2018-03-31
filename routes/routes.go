package routes

import (
	"log"
	"net/http"

	"Govel/app/controllers/home"

	"github.com/gorilla/mux"
)

// Init stores all the
// routes used in the application
func Init() {
	r := mux.NewRouter()
	r.HandleFunc("/", home.Index)
	r.NotFoundHandler = http.HandlerFunc(home.NotFound)
	log.Fatal(http.ListenAndServe(":8080", r))
}
