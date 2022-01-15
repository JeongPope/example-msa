package model

type Movie struct {
	ID      int    `json:"id,omitempty"`
	Title   string `json:"title"`
	Image   string `json:"image"`
	AltText string `json:"alt"`
	Desc    string `json:"desc"`
}
