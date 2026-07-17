package handlers

import (
	"net/http"

	"01Edu-Groupie-Project/services"
)

// HomeHandler handles requests to the home page.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow the root path.
	if r.URL.Path != "/" {
		NotFound(w)
		return
	}

	// Only allow GET requests.
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Fetch artists.
	artists, err := services.GetArtists()
	if err != nil {
		InternalServerError(w)
		return
	}

	// Render homepage.
	if err := RenderTemplate(w, "index.html", artists); err != nil {
		InternalServerError(w)
		return
	}
}