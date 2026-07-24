package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"01Edu-Groupie-Project/models"
)

// nominatimResponse represents one location returned by the Nominatim API.
type nominatimResponse struct {
	DisplayName string `json:"display_name"`
	Lat         string `json:"lat"`
	Lon         string `json:"lon"`
}

// GetCoordinate converts a location name into geographic coordinates.
func GetCoordinate(location string) (models.Coordinate, error) {

	// Convert "london-uk" -> "london uk"
	location = strings.ReplaceAll(location, "-", " ")

	// Escape special characters for use in a URL.
	query := url.QueryEscape(location)

	// Build the Nominatim API URL.
	apiURL := fmt.Sprintf(
		"https://nominatim.openstreetmap.org/search?q=%s&format=json&limit=1",
		query,
	)

	// Create the HTTP request.
	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		return models.Coordinate{}, err
	}

	// Nominatim requires a User-Agent header.
	req.Header.Set("User-Agent", "01Edu-Groupie-Tracker")

	// Send the request.
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return models.Coordinate{}, err
	}
	defer resp.Body.Close()

	// Decode the JSON response.
	var results []nominatimResponse

	err = json.NewDecoder(resp.Body).Decode(&results)
	if err != nil {
		return models.Coordinate{}, err
	}

	if len(results) == 0 {
		return models.Coordinate{}, fmt.Errorf("location not found")
	}

	// Convert latitude and longitude from strings to float64.
	lat, err := strconv.ParseFloat(results[0].Lat, 64)
	if err != nil {
		return models.Coordinate{}, err
	}

	lon, err := strconv.ParseFloat(results[0].Lon, 64)
	if err != nil {
		return models.Coordinate{}, err
	}

	// Return our application's Coordinate model.
	return models.Coordinate{
		Name:      location,
		Latitude:  lat,
		Longitude: lon,
	}, nil
}
