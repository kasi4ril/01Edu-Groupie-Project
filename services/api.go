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

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return err
	}

	return nil
}

// GetArtists fetches all artists.
func GetArtists() ([]models.Artist, error) {
	var artists []models.Artist

	if err := fetchJSON(baseURL+"/artists", &artists); err != nil {
		return nil, err
	}

	return artists, nil
}

// GetLocations fetches all artist locations.
func GetLocations() ([]models.Location, error) {
	var response models.LocationsResponse

	if err := fetchJSON(baseURL+"/locations", &response); err != nil {
		return nil, err
	}

	return response.Index, nil
}

// GetDates fetches all concert dates.
func GetDates() ([]models.Date, error) {
	var response models.DatesResponse

	if err := fetchJSON(baseURL+"/dates", &response); err != nil {
		return nil, err
	}

	return response.Index, nil
}

// GetRelations fetches all artist relations.
func GetRelations() ([]models.Relation, error) {
	var response models.RelationsResponse

	if err := fetchJSON(baseURL+"/relation", &response); err != nil {
		return nil, err
	}

	return response.Index, nil
}