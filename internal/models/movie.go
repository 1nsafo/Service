package models

type Film struct {
	ID          int64   `json:"id"`
	Title       string  `json:"title"`
	ReleaseYear string  `json:"release_year"`
	Director    string  `json:"director"`
	Genre       string  `json:"genre"`
	Rating      float64 `json:"rating"`
}
