package models

// Filter stores the user's filter selections.
type Filter struct {
	CreationFrom int
	CreationTo   int

	AlbumFrom int
	AlbumTo   int

	Members []int

	Locations []string
}
