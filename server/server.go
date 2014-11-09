// Palantír Web Server
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	Host string
	Port uint

	router *mux.Router
}

func (s Server) Bind() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func (s *Server) Setup() {
	r := mux.NewRouter()
	ns := r.PathPrefix("/api/v1").Subrouter()

	ns.HandleFunc("/{subject}", getNotification).Methods("GET")
	ns.HandleFunc("/{subject}", createNotification).Methods("POST")

	s.router = r
}

func (s Server) Start() {
	log.Printf("server start listening at: %s", s.Bind())
	log.Fatal(http.ListenAndServe(s.Bind(), s.router))
}

func main() {
	var server Server

	flag.StringVar(&server.Host, "host", "127.0.0.1", "bind host")
	flag.UintVar(&server.Port, "port", 9092, "bind port")
	flag.Parse()

	setupLogger()
	setupStorage()

	server.Setup()
	server.Start()
}

func setupStorage() {
	GetStorageInstance()
}

func setupLogger() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.SetPrefix("Palantír ")
}
