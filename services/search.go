package services

import (
	"01Edu-Groupie-Project/models"
	"strconv"
	"strings"
)

// SearchArtists searches artists using the user's search query.
//
// It searches through:
// - artist/band name
// - members
// - concert locations
// - first album date
// - creation year
//
// Search is case-insensitive.
func SearchArtists(
	artists []models.Artist,
	locations []models.Location,
	query string,
) []models.Artist {

	query = strings.TrimSpace(strings.ToLower(query))

	if query == "" {
		return artists
	}

	var results []models.Artist

	for _, artist := range artists {

		found := false

		// -------------------------------
		// Artist / Band name
		// -------------------------------
		if strings.Contains(
			strings.ToLower(artist.Name),
			query,
		) {
			found = true
		}

		// -------------------------------
		// Members
		// -------------------------------
		if !found {
			for _, member := range artist.Members {

				if strings.Contains(
					strings.ToLower(member),
					query,
				) {
					found = true
					break
				}
			}
		}

		// -------------------------------
		// Creation year
		// -------------------------------
		if !found {

			creationYear := strconv.Itoa(artist.CreationDate)

			if strings.Contains(creationYear, query) {
				found = true
			}
		}

		// -------------------------------
		// First album date
		// -------------------------------
		if !found {

			if strings.Contains(
				strings.ToLower(artist.FirstAlbum),
				query,
			) {
				found = true
			}
		}

		// -------------------------------
		// Concert locations
		// -------------------------------
		if !found {

			for _, locationData := range locations {

				if locationData.ID != artist.ID {
					continue
				}

				for _, location := range locationData.Locations {

					if strings.Contains(
						strings.ToLower(location),
						query,
					) {
						found = true
						break
					}
				}

				if found {
					break
				}
			}
		}

		// Add artist only once.
		if found {
			results = append(results, artist)
		}
	}

	return results
}

// GetSearchSuggestions returns live suggestions for the search bar.
func GetSearchSuggestions(
	artists []models.Artist,
	locations []models.Location,
	query string,
) []models.SearchSuggestion {

	query = strings.TrimSpace(strings.ToLower(query))

	if query == "" {
		return nil
	}

	var suggestions []models.SearchSuggestion

	for _, artist := range artists {

		// -------------------------------
		// Artist / Band
		// -------------------------------
		if strings.Contains(
			strings.ToLower(artist.Name),
			query,
		) {
			suggestions = append(
				suggestions,
				models.SearchSuggestion{
					Value: artist.Name,
					Type:  "artist/band",
				},
			)
		}

		// -------------------------------
		// Members
		// -------------------------------
		for _, member := range artist.Members {

			if strings.Contains(
				strings.ToLower(member),
				query,
			) {
				suggestions = append(
					suggestions,
					models.SearchSuggestion{
						Value: member,
						Type:  "member",
					},
				)
			}
		}

		// -------------------------------
		// Creation year
		// -------------------------------
		creationYear := strconv.Itoa(artist.CreationDate)

		if strings.Contains(creationYear, query) {
			suggestions = append(
				suggestions,
				models.SearchSuggestion{
					Value: creationYear,
					Type:  "creation date",
				},
			)
		}

		// -------------------------------
		// First album
		// -------------------------------
		if strings.Contains(
			strings.ToLower(artist.FirstAlbum),
			query,
		) {
			suggestions = append(
				suggestions,
				models.SearchSuggestion{
					Value: artist.FirstAlbum,
					Type:  "first album",
				},
			)
		}
	}

	// -------------------------------
	// Locations
	// -------------------------------
	for _, locationData := range locations {

		for _, place := range locationData.Locations {

			if strings.Contains(
				strings.ToLower(place),
				query,
			) {
				suggestions = append(
					suggestions,
					models.SearchSuggestion{
						Value: place,
						Type:  "location",
					},
				)
			}
		}
	}

	return suggestions
}
