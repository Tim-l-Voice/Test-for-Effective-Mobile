package models

type Song struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
}

type SongDetail struct {
	Album       string `json:"album"`
	ReleaseDate string `json:"release_date"`
	Genre       string `json:"genre"`
}

type SongRequest struct {
	Title  string `json:"title"`
	Artist string `json:"artist"`
}
