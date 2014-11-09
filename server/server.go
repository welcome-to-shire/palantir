// Palantír Web Server
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	setupStorage()

	log.Fatal(http.ListenAndServe(":8080", setupServer()))
}

func setupStorage() {
	GetStorageInstance()
}

func setupServer() *mux.Router {
	r := mux.NewRouter()
	ns := r.PathPrefix("/api/v1").Subrouter()

	ns.HandleFunc("/{subject}", getNotification).Methods("GET")
	ns.HandleFunc("/{subject}", createNotification).Methods("POST")

	return r
}
