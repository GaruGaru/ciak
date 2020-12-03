package models

import "time"

type Details struct {
	Name        string
	Director    string
	Genre       string
	Rating      uint8
	MaxRating   uint8
	ReleaseDate time.Time
	ImagePoster string
}
