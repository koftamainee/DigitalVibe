package main

import (
	"io"
	"log"
	"net/http"
)

const (
	userServiceURL = "http://localhost:8081"
	taskServiceURL = "http://localhost:8082"
)

func main() {
	mux := http.NewServeMux()

	// Task Service routes
	mux.HandleFunc("/", forwardToTaskService)
	mux.HandleFunc("/chat/ai", forwardToTaskService)
	mux.HandleFunc("/chat/psychologist", forwardToTaskService)
	mux.HandleFunc("/tracker/", forwardToTaskService)
	mux.HandleFunc("/tracker/emotion", forwardToTaskService)
	mux.HandleFunc("/tracker/stress", forwardToTaskService)
	mux.HandleFunc("/tracker/diary", forwardToTaskService)

	// User Service routes
	mux.HandleFunc("/login", forwardToUserService)
	mux.HandleFunc("/register", forwardToUserService)

	// Start the API Gateway
	log.Println("API Gateway running on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

// Forward requests to the Task Service
func forwardToTaskService(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, taskServiceURL)
}

// Forward requests to the User Service
func forwardToUserService(w http.ResponseWriter, r *http.Request) {
	proxyRequest(w, r, userServiceURL)
}

// Helper function to forward requests to appropriate microservices
func proxyRequest(w http.ResponseWriter, r *http.Request, serviceURL string) {
	// Create the new request to forward
	req, err := http.NewRequest(r.Method, serviceURL+r.URL.Path, r.Body)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	// Copy headers from the original request to the new request
	req.Header = r.Header

	// Create an HTTP client and forward the request
	client := &http.Client{}
	log.Println(req)
	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, "Service unavailable", http.StatusServiceUnavailable)
		log.Println(err)
		return
	}
	defer resp.Body.Close()

	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Write the response status and body
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
