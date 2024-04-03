package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Serve static files (like your HTML and Tailwind CSS)
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	r.HandleFunc("/", serveIndex)
	r.HandleFunc("/conf", getConfigurations).Methods("GET")
	r.HandleFunc("/conf", addConfiguration).Methods("POST")
	r.HandleFunc("/conf/{id}", removeConfiguration).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
