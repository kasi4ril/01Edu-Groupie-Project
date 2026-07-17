package main

import (
	"log"
	"net/http"

	"01Edu-Groupie-Project/handlers"
)

func main() {
	// Create router
	mux := http.NewServeMux()

	// Serve static files
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	// Register routes
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artist", handlers.ArtistHandler)

	log.Println("Server running on http://localhost:8080")

	// Start server
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}