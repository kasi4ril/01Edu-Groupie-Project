package handlers

import (
	"log"
	"net/http"
	"strconv"

	"01Edu-Groupie-Project/models"
	"01Edu-Groupie-Project/services"
)

// buildFilter reads the query parameters from the URL
// and builds a Filter struct.
func buildFilter(r *http.Request) models.Filter {

	filter := models.Filter{}

	// -------------------------------
	// Creation Year
	// -------------------------------

	if value := r.URL.Query().Get("creationFrom"); value != "" {
		if year, err := strconv.Atoi(value); err == nil {
			filter.CreationFrom = year
		}
	}

	if value := r.URL.Query().Get("creationTo"); value != "" {
		if year, err := strconv.Atoi(value); err == nil {
			filter.CreationTo = year
		}
	}

	// -------------------------------
	// First Album Year
	// -------------------------------

	if value := r.URL.Query().Get("albumFrom"); value != "" {
		if year, err := strconv.Atoi(value); err == nil {
			filter.AlbumFrom = year
		}
	}

	if value := r.URL.Query().Get("albumTo"); value != "" {
		if year, err := strconv.Atoi(value); err == nil {
			filter.AlbumTo = year
		}
	}

	// -------------------------------
	// Members
	// -------------------------------

	members := r.URL.Query()["members"]

	for _, m := range members {
		if value, err := strconv.Atoi(m); err == nil {
			filter.Members = append(filter.Members, value)
		}
	}

	// -------------------------------
	// Locations
	// -------------------------------

	location := r.URL.Query().Get("location")

	if location != "" {
		filter.Locations = append(filter.Locations, location)
	}

	return filter
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		NotFound(w)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Build filter from query parameters
	filter := buildFilter(r)

	// Print filter to terminal (for testing)
	log.Printf("Current Filter: %+v\n", filter)

	// Fetch all artists
	artists, err := services.GetArtists()
	if err != nil {
		InternalServerError(w)
		return
	}

	// Apply creation year filter
	artists = services.FilterCreation(artists, filter)
	artists = services.FilterAlbum(artists, filter)
	artists = services.FilterMembers(artists, filter)

	// Filtering will be added in the next step

	if err := RenderTemplate(w, "index.html", artists); err != nil {
		InternalServerError(w)
		return
	}

	locations, err := services.GetLocations()
	if err != nil {
		InternalServerError(w)
		return
	}

	artists = services.FilterCreation(artists, filter)
	artists = services.FilterAlbum(artists, filter)
	artists = services.FilterMembers(artists, filter)
	artists = services.FilterLocations(artists, locations, filter)
}
