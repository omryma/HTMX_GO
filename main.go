package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Serve static files (like your HTML and Tailwind CSS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Home route
	r.HandleFunc("/", serveIndex)

	log.Fatal(http.ListenAndServe(":8080", r))
}
