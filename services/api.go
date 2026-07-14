package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"01Edu-Groupie-Project/models"
)

const baseURL = "https://groupietrackers.herokuapp.com/api"

// fetchJSON sends an HTTP GET request to the given URL,
// decodes the JSON response into target,
// and returns an error if anything goes wrong.
func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

// GetArtists fetches all artists from the API.
func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist

	err := fetchJSON(baseURL+"/artists", &artists)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

// GetLocations fetches all artist locations.
func GetLocations() ([]models.Location, error) {
	var locations []models.Location

	err := fetchJSON(baseURL+"/locations", &locations)
	if err != nil {
		return nil, err
	}

	return locations, nil
}

// GetDates fetches all concert dates.
func GetDates() ([]models.Date, error) {
	var dates []models.Date

	err := fetchJSON(baseURL+"/dates", &dates)
	if err != nil {
		return nil, err
	}

	return dates, nil
}

// GetRelations fetches all artist relations.
func GetRelations() ([]models.Relation, error) {
	var relations []models.Relation

	err := fetchJSON(baseURL+"/relation", &relations)
	if err != nil {
		return nil, err
	}

	return relations, nil
}
