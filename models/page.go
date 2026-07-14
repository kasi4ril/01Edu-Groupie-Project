package models


//Modelled into one so that the templates receive one object
type ArtistPage struct {
	Artist    Artist
	Locations []string
	Dates     []string
	Relation  map[string][]string
}
