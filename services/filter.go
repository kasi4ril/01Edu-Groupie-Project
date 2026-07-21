package services

import (
	"01Edu-Groupie-Project/models"
	"strconv"
	"strings"
)

// FilterCreation filters artists by creation year.
func FilterCreation(artists []models.Artist, filter models.Filter) []models.Artist {

	// If no creation year filter is selected,
	// return all artists.
	if filter.CreationFrom == 0 && filter.CreationTo == 0 {
		return artists
	}

	var filtered []models.Artist

	for _, artist := range artists {

		year := artist.CreationDate

		// Check minimum year
		if filter.CreationFrom != 0 && year < filter.CreationFrom {
			continue
		}

		// Check maximum year
		if filter.CreationTo != 0 && year > filter.CreationTo {
			continue
		}

		filtered = append(filtered, artist)
	}

	return filtered
}

// FilterAlbum filters artists by first album year.
func FilterAlbum(artists []models.Artist, filter models.Filter) []models.Artist {

	// No filter selected
	if filter.AlbumFrom == 0 && filter.AlbumTo == 0 {
		return artists
	}

	var filtered []models.Artist

	for _, artist := range artists {

		// Example: "26-03-2001"
		parts := strings.Split(artist.FirstAlbum, "-")

		if len(parts) != 3 {
			continue
		}

		year, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		if filter.AlbumFrom != 0 && year < filter.AlbumFrom {
			continue
		}

		if filter.AlbumTo != 0 && year > filter.AlbumTo {
			continue
		}

		filtered = append(filtered, artist)
	}

	return filtered
}

// FilterMembers filters artists by number of members.
func FilterMembers(artists []models.Artist, filter models.Filter) []models.Artist {

	// No member filter selected
	if len(filter.Members) == 0 {
		return artists
	}

	var filtered []models.Artist

	for _, artist := range artists {

		memberCount := len(artist.Members)

		for _, selected := range filter.Members {

			if memberCount == selected {
				filtered = append(filtered, artist)
				break
			}
		}
	}

	return filtered
}

// FilterLocations filters artists by concert locations.
func FilterLocations(
	artists []models.Artist,
	locations []models.Location,
	filter models.Filter,
) []models.Artist {

	// No location selected
	if len(filter.Locations) == 0 {
		return artists
	}

	var filtered []models.Artist

	for _, artist := range artists {

		// Find this artist's locations
		for _, loc := range locations {

			if loc.ID != artist.ID {
				continue
			}

			found := false

			// Compare every artist location
			for _, artistLocation := range loc.Locations {

				// Compare with every selected location
				for _, selected := range filter.Locations {

					if artistLocation == selected {
						found = true
						break
					}
				}

				if found {
					break
				}
			}

			if found {
				filtered = append(filtered, artist)
			}

			break
		}
	}

	return filtered
}
