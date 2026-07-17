package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"01Edu-Groupie-Project/models"
	"01Edu-Groupie-Project/services"
)

// parseArtistID extracts and validates the artist ID from the query string.
func parseArtistID(r *http.Request) (int, error) {
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		return 0, fmt.Errorf("missing artist id")
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid artist id")
	}

	return id, nil
}

// loadArtistData fetches and combines all data needed for artist.html.
func loadArtistData(id int) (models.ArtistPage, error) {
	artists, err := services.GetArtists()
	if err != nil {
		log.Println("GetArtists:", err)
		return models.ArtistPage{}, err
	}

	locations, err := services.GetLocations()
	if err != nil {
		log.Println("GetLocations:", err)
		return models.ArtistPage{}, err
	}

	dates, err := services.GetDates()
	if err != nil {
		log.Println("GetDates:", err)
		return models.ArtistPage{}, err
	}

	relations, err := services.GetRelations()
	if err != nil {
		log.Println("GetRelations:", err)
		return models.ArtistPage{}, err
	}

	var artist models.Artist
	found := false

	for _, a := range artists {
		if a.ID == id {
			artist = a
			found = true
			break
		}
	}

	if !found {
		return models.ArtistPage{}, fmt.Errorf("artist not found")
	}

	var location models.Location
	for _, l := range locations {
		if l.ID == id {
			location = l
			break
		}
	}

	var date models.Date
	for _, d := range dates {
		if d.ID == id {
			date = d
			break
		}
	}

	var relation models.Relation
	for _, rel := range relations {
		if rel.ID == id {
			relation = rel
			break
		}
	}

	return models.ArtistPage{
		Artist:   artist,
		Location: location,
		Date:     date,
		Relation: relation,
	}, nil
}