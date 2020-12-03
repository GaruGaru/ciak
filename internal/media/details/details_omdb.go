package details

import (
	"encoding/json"
	"fmt"
	"github.com/GaruGaru/ciak/internal/media/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"regexp"
	"time"
)

type OmdbClient struct {
	apiKey     string
	httpClient *http.Client
}

func Omdb(apiKey string) *OmdbClient {
	return &OmdbClient{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (o OmdbClient) Details(request Request) (models.Details, error) {
	apiUrl := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", o.apiKey, url.QueryEscape(normalizeTitle(request.Title)))

	resp, err := o.httpClient.Get(apiUrl)

	if err != nil {
		return models.Details{}, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logrus.Warn("unable to close response body: " + err.Error())
		}
	}()

	var movie OmdbMovie

	err = json.NewDecoder(resp.Body).Decode(&movie)

	if err != nil {
		return models.Details{}, err
	}

	if movie.Response == "False" {
		return models.Details{}, ErrDetailsNotFound
	}

	return models.Details{
		Name:        request.Title, // at the moment we use the name to associate request -> metadata
		Director:    movie.Director,
		Genre:       movie.Genre,
		Rating:      0,
		MaxRating:   0,
		ReleaseDate: time.Time{},
		ImagePoster: movie.Poster,
	}, nil
}

func normalizeTitle(title string) string {
	removeYear := regexp.MustCompile("(.*) (.*)(\\d\\d\\d\\d)")
	matches := removeYear.FindAllString(title, -1)
	if len(matches) == 0 {
		return title
	}
	return matches[0]
}

type OmdbMovie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}
