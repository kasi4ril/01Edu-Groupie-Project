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

	locations := r.URL.Query()["location"]

	filter.Locations = append(filter.Locations, locations...)

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

	// ----------------------------------------
	// Auto-correct invalid ranges
	// ----------------------------------------

	// Creation year
	if filter.CreationFrom > filter.CreationTo &&
		filter.CreationTo != 0 {

		filter.CreationFrom, filter.CreationTo =
			filter.CreationTo, filter.CreationFrom
	}

	// Album year
	if filter.AlbumFrom > filter.AlbumTo &&
		filter.AlbumTo != 0 {

		filter.AlbumFrom, filter.AlbumTo =
			filter.AlbumTo, filter.AlbumFrom
	}

	log.Printf("Current Filter: %+v\n", filter)

	// ----------------------------------------
	// Fetch data
	// ----------------------------------------

	artists, err := services.GetArtists()
	if err != nil {
		InternalServerError(w)
		return
	}

	// Save a copy BEFORE filtering
	allArtists := artists

	locations, err := services.GetLocations()
	if err != nil {
		InternalServerError(w)
		return
	}

	// ----------------------------------------
	// Apply filters
	// ----------------------------------------

	artists = services.FilterCreation(artists, filter)
	artists = services.FilterAlbum(artists, filter)
	artists = services.FilterMembers(artists, filter)
	artists = services.FilterLocations(artists, locations, filter)

	// ----------------------------------------
	// Build page
	// ----------------------------------------

	page := models.HomePage{
		Artists:       artists,
		Locations:     services.GetUniqueLocations(locations),
		CreationYears: services.GetCreationYears(allArtists),
		AlbumYears:    services.GetAlbumYears(allArtists),
		Filter:        filter,
	}

	if err := RenderTemplate(w, "index.html", page); err != nil {
		InternalServerError(w)
		return
	}
}