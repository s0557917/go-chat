package main

import (
	"log"
	"net/http"
	"web-sockets/internal/handlers"
)

func main() {
	routes := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	log.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", routes)
}
