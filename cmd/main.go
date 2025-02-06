package main

import (
	"leadgen/internal/adapter/http"
	"log"
)

func main() {
	server := http.NewServer()
	defer server.Close()
	log.Println("Starting server on :8080")
	server.Run(":8080")
}
