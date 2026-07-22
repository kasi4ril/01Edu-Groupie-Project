package models

type HomePage struct {
	Artists      []Artist
	Locations    []string
	CreationYears []int
	AlbumYears    []int
	Filter       Filter
}