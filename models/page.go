package models

// Modelled into one so that the templates receive one object
type ArtistPage struct {
	Artist    Artist
	Locations Location
	Dates     Date
	Relation  Relation
}
