package main

import (
	"log"
	"net/http"

	"main/handlers"
)

func main() {
	mux := http.NewServeMux()

	handlers.SetupRoutes(mux)

	log.Println("Task Service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
