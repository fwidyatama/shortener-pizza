package main

import (
	"log"
	"net/http"
	"pizza-hub/internal/api"
)

func main() {

	route := api.InitRoute()

	serverAddr := ":8080"
	log.Printf("Server running on port %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, route); err != nil {
		log.Printf("Error starting the server: %v", err.Error())
	}
}
