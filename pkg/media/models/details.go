package models

import "time"

type Details struct {
	Name        string
	Director    string
	Genre       string
	Rating      float64
	MaxRating   float64
	ReleaseDate time.Time
	ImagePoster string
}
