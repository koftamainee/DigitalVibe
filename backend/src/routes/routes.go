package routes

import (
	"fmt"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	// GET
	mux.HandleFunc("GET /", rootHandler)
	mux.HandleFunc("GET /c", allChatsGetHandler)
	mux.HandleFunc("GET /c/{id}", chatGetHandler)

	// POST
	mux.HandleFunc("POST /c}", chatPostHandler)
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Hello, World!</title>
	</head>
	<body>
		<h1>Hello, World!</h1>
		<p>Welcome to my web server!</p>
	</body>
	</html>
`)
}

func allChatsGetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from all chats")
}

func chatGetHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "Hello from chat, id = %s", id)
}

func chatPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "created new chat")
}
