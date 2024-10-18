package main

import (
	"log"
	"net/http"

	"main/handlers"
)

func main() {
	mux := http.NewServeMux()

	handlers.SetupRoutes(mux)

	log.Println("User Service running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
