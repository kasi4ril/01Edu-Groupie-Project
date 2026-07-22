package services

import (
	"sort"
	"strconv"
	"strings"

	"01Edu-Groupie-Project/models"
)

// FilterCreation filters artists by creation year.
func FilterCreation(artists []models.Artist, filter models.Filter) []models.Artist {

	// No creation filter selected
	if filter.CreationFrom == 0 && filter.CreationTo == 0 {
		return artists
	}

	var filtered []models.Artist

	for _, artist := range artists {

		year := artist.CreationDate

		if filter.CreationFrom != 0 && year < filter.CreationFrom {
			continue
		}

		if filter.CreationTo != 0 && year > filter.CreationTo {
			continue
		}

		filtered = append(filtered, artist)
	}

	return filtered
}

// FilterAlbum filters artists by first album year.
func FilterAlbum(artists []models.Artist, filter models.Filter) []models.Artist {

	// No album filter selected
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

	// No location filter selected
	if len(filter.Locations) == 0 {
		return artists
	}

	var filtered []models.Artist

	for _, artist := range artists {

		// Find the matching location record for this artist
		for _, loc := range locations {

			if loc.ID != artist.ID {
				continue
			}

			found := false

			// Compare artist locations with selected locations
			for _, artistLocation := range loc.Locations {

				for _, selected := range filter.Locations {

					if strings.Contains(
						strings.ToLower(artistLocation),
						strings.ToLower(selected),
					) {
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

			// Only one Location record exists per artist
			break
		}
	}

	return filtered
}

// GetUniqueLocations returns all unique concert locations.
func GetUniqueLocations(locations []models.Location) []string {

	unique := make(map[string]bool)

	for _, loc := range locations {
		for _, city := range loc.Locations {
			unique[city] = true
		}
	}

	var result []string

	for city := range unique {
		result = append(result, city)
	}

	sort.Strings(result)

	return result
}
// GetCreationYears returns all unique creation years.
func GetCreationYears(artists []models.Artist) []int {

	unique := make(map[int]bool)

	for _, artist := range artists {
		unique[artist.CreationDate] = true
	}

	var years []int

	for year := range unique {
		years = append(years, year)
	}

	sort.Ints(years)

	return years
}

// GetAlbumYears returns all unique album years.
func GetAlbumYears(artists []models.Artist) []int {

	unique := make(map[int]bool)

	for _, artist := range artists {

		parts := strings.Split(artist.FirstAlbum, "-")

		if len(parts) != 3 {
			continue
		}

		year, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		unique[year] = true
	}

	var years []int

	for year := range unique {
		years = append(years, year)
	}

	sort.Ints(years)

	return years
}