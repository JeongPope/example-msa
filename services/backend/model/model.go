package model

type Movie struct {
	Title   string `json:"title"`
	Image   string `json:"image"`
	AltText string `json:"alt"`
	Desc    string `json:"desc"`
}
