package routes

import (
	"fmt"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	// GET
	mux.HandleFunc("GET /", rootHandler)

	mux.HandleFunc("GET /login/", loginHandler)

	mux.HandleFunc("GET /chat/ai", aiChatGetHandler)
	mux.HandleFunc("GET /chat/psychologist", psychologistChatGetHandler)

	mux.HandleFunc("GET /tracker/", trackersHandler)
	mux.HandleFunc("GET /tracker/emotion", emotionTrackerHandler)
	mux.HandleFunc("GET /tracker/stress", stressTrackerHandler)
	mux.HandleFunc("GET /tracker/diary", diaryTrackerHandler)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "main page")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login page")
}

func aiChatGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ai chat page")
}

func psychologistChatGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "psyhologist chat page")
}

func trackersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All trackers")
}

func emotionTrackerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "emotion tracker")
}

func stressTrackerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "stress tracker")
}

func diaryTrackerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "diary tracker")
}
