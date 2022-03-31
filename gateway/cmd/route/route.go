package main

import (
	"fmt"
	"log"

	"gateway/internal/server"
)

var Version = "v0.1.0"

func main() {
	log.Printf("starting API version: %s\n", Version)
	s := server.New()
	s.InitDomains()
	fmt.Println("Registered Routes:")
	s.PrintAllRegisteredRoutes()
}
