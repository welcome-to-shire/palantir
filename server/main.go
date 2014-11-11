package main

import (
	"flag"
	"log"
)

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
