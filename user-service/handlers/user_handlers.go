package handlers

import (
	"encoding/json"
	"net/http"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/register", registerHandler)
}

// Handle user registration
func registerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// For simplicity, we assume the user is always successfully registered
	response := map[string]string{"message": "User registered successfully"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Handle user login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST is allowed", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// For simplicity, we assume the login is always successful
	response := map[string]string{"message": "Login successful"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
