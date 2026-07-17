package handlers

import (
	"log"
	"net/http"
)

// ArtistHandler handles requests for a single artist page.
func ArtistHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests.
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse artist ID.
	id, err := parseArtistID(r)
	if err != nil {
		log.Printf("parseArtistID: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Load artist data.
	page, err := loadArtistData(id)
	if err != nil {
		log.Printf("loadArtistData: %v", err)
		InternalServerError(w)
		return
	}

	// Render template.
	if err := RenderTemplate(w, "artist.html", page); err != nil {
		log.Printf("RenderTemplate: %v", err)
		InternalServerError(w)
		return
	}
}