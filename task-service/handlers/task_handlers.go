package handlers

import (
	"encoding/json"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	// Define routes
	mux.HandleFunc("/", rootHandler)

	// Chat-related routes
	mux.HandleFunc("/chat/ai", aiChatGetHandler)
	mux.HandleFunc("/chat/psychologist", psychologistChatGetHandler)

	// Tracker-related routes
	mux.HandleFunc("/tracker/", trackersHandler)
	mux.HandleFunc("/tracker/emotion", emotionTrackerHandler)
	mux.HandleFunc("/tracker/stress", stressTrackerHandler)
	mux.HandleFunc("/tracker/diary", diaryTrackerHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the main page"})
}

func aiChatGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "AI chat page"})
}

func psychologistChatGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Psychologist chat page"})
}

func trackersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "All trackers"})
}

func emotionTrackerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Emotion tracker"})
}

func stressTrackerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Stress tracker"})
}

func diaryTrackerHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Diary tracker"})
}
