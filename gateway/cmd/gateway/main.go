package main

import (
	"log"

	"gateway/internal/server"
)

var Version = "v0.1.0"

func main() {
	log.Printf("Starting API version: %s\n", Version)
	s := server.New()
	s.Init(Version)
	s.Run()
}
