package handlers

import (
	"net/http"
	/"strconv"

	//"01Edu-Groupie-Project/models"
	//"01Edu-Groupie-Project/services"
)

func ArtistHandler(w http.ResponseWriter, r *http.Request) {

	id, err := parseArtistID(r)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	page, err := loadArtistData(id)
	if err != nil {
		InternalServerError(w)
		return
	}

	RenderTemplate(w, "artist.html", page)
}