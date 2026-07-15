package handlers

import (
	"01Edu-Groupie-Project/services"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	artists, err := services.GetArtists()
	if err != nil {
		InternalServerError(w)
		return
	}

	err = RenderTemplate(w,	"index.html", artists)

	if err != nil {
		InternalServerError(w)
		return
	}
}
