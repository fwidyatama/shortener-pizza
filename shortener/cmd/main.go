package main

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"shortener/internal/api"
)

func main() {
	log.SetFormatter(&log.JSONFormatter{})

	route := api.InitRoute()

	serverAddr := ":8080"
	log.Printf("Server running on port %s", serverAddr)

	if err := http.ListenAndServe(serverAddr, route); err != nil {
		log.Printf("Error starting the server: %v", err.Error())
	}
}
