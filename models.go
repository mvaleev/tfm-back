package main

// Thumb description
type Thumb struct {
	ThumbURL string `json:"url"`
	ThumbID  string `json:"id"`
}

// Video description
type Video struct {
	VideoURL   string `json:"url"`
	Popularity string `json:"popularity"`
}
