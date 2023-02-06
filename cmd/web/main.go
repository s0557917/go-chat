package main

import (
	"log"
	"net/http"
)

func main() {
	routes := routes()

	log.Println("Starting server on port 8080")

	http.ListenAndServe(":8080", routes)
}
