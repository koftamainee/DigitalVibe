package handlers

import (
	"fmt"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	// GET
	mux.HandleFunc("/login", loginHandler)
	mux.HandleFunc("/register", registerHandler)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login page")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "register page")
}
