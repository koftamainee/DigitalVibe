package main

import (
	"log"
	"net/http"

	"main/handlers"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	mux := http.NewServeMux()

	handlers.SetupRoutes(mux)

	http.Handle("/metrics", promhttp.Handler())

	log.Println("Task Service running on port 8082")
	log.Fatal(http.ListenAndServe(":8082", mux))
}
