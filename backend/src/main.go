package main

import (
	"log"
	"net/http"

	"main/src/routes"
)

func main() {
	mux := http.NewServeMux()

	routes.SetupRoutes(mux)
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
