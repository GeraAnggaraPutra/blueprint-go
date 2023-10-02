package main

import (
	"log"

	"github.com/GeraAnggaraPutra/blueprint-go/routes"
)

func main() {
	err := routes.Init()
	if err != nil {
		log.Fatalf("Error start the server with err: %s", err)
	}
}
