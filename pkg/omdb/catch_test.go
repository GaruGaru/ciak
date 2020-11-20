package omdb

import (
	"testing"
)

func TestBasic(t *testing.T) {
	memCache := NewMemoryCache()
	key := "The Queen's Gambit"
	movie := Movie{
		Title:    "The Queen's Gambit",
		Year:     "2020–",
		Rated:    "TV-MA",
		Released: "23 Oct 2020",
		Runtime:  "N/A",
		Genre:    "Drama",
		Director: "N/A",
		Writer:   "N/A",
		Actors:   "Anya Taylor-Joy, Chloe Pirrie, Bill Camp, Matthew Dennis Lewis",
		Plot:     "Eight year-old orphan Beth Harmon is quiet, sullen, and by all appearances unremarkable. That is, until she plays her first game of chess. Her senses grow sharper, her thinking clearer, and...",
		Language: "English",
		Country:  "USA",
		Awards:   "N/A",
		Poster:   "https://m.media-amazon.com/images/M/MV5BM2EwMmRhMmUtMzBmMS00ZDQ3LTg4OGEtNjlkODk3ZTMxMmJlXkEyXkFqcGdeQXVyMjM5ODk1NDU@._V1_SX300.jpg",
		Ratings: []Rating{
			{
				Source: "Internet Movie Database",
				Value:  "8.9/10",
			},
		},
		Metascore:  "N/A",
		ImdbRating: "8.9",
		ImdbVotes:  "6,077",
		ImdbID:     "tt10048342",
		Type:       "series",
		Response:   "True",
	}
	memCache.Put(key, movie)
	movie, exist := memCache.Get(key)
	if !exist {
		t.Error("error, we store a movie in Cache, but it doesn't exist when we get it")
	} else {
		t.Log("movie stored and get successful")
	}
	memCache.Del(key)
	movie, exist = memCache.Get(key)
	if exist {
		t.Error("we delete a movie from cache, but when we get it, it still exists, wired!")
	} else {
		t.Log("delete successful")
	}
}
