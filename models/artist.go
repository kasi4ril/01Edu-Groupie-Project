package models


//Each artist would be seen with the structured information as;
type Artist struct {
	ID           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      string `json:"members"`
	Creationdate string `json:"creationdate"`
	FirstAlbum   string `json:"firstalbum"`
}
