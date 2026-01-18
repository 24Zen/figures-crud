package main

type Figure struct {
	ID          int      `json:"id"`
	Anime       string   `json:"anime"`
	Character   string   `json:"character"`
	ImageURLs   []string `json:"imageUrls"`
	Description string   `json:"description"`
}
