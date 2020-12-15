package details

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/GaruGaru/ciak/pkg/media/models"
	"github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	omdbEndpoint = "http://www.omdbapi.com"
)

var (
	ErrRatingProviderNotFound = errors.New("no ratings provider parser found")
)

type OmdbClient struct {
	apiKey     string
	httpClient *http.Client
	endpoint   string
}

func Omdb(apiKey string) *OmdbClient {
	return &OmdbClient{
		endpoint: omdbEndpoint,
		apiKey:   apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (o OmdbClient) Details(request Request) (models.Details, error) {
	apiUrl := fmt.Sprintf("%s/?apikey=%s&t=%s", o.endpoint, o.apiKey, url.QueryEscape(normalizeTitle(request.Title)))

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

	releaseDate, err := time.Parse("02 Jan 2006", movie.Released)
	if err != nil {
		logrus.Warnf("unable to parse media release date: %s", err)
	}

	ratingValue, ratingMax, err := parseRating(movie)
	if err != nil {
		logrus.Warnf("unable to parse media release date: %s", err)
	}

	return models.Details{
		Name:        request.Title, // at the moment we use the name to associate request -> metadata
		Director:    movie.Director,
		Genre:       movie.Genre,
		Rating:      ratingValue,
		MaxRating:   ratingMax,
		ReleaseDate: releaseDate,
		ImagePoster: movie.Poster,
	}, nil
}

func parseRating(movie OmdbMovie) (float64, float64, error) {
	for _, r := range movie.Ratings {
		val, max, err := omdbParseProviderRating(r.Source, r.Value)
		if err != nil {
			if err != ErrRatingProviderNotFound {
				logrus.Warnf("error parsing rating: %s", err)
			}
			continue
		}

		return val, max, nil
	}

	return 0, 0, ErrRatingProviderNotFound
}

func omdbParseProviderRating(provider string, value string) (float64, float64, error) {
	switch provider {
	case "Rotten Tomatoes":
		return omdbParseRottenTomatoesRatings(value)
	default:
		return 0, 0, ErrRatingProviderNotFound
	}
}

func omdbParseRottenTomatoesRatings(value string) (float64, float64, error) {
	rawNum := strings.Replace(value, "%", "", 1)
	val, err := strconv.ParseFloat(rawNum, 64)
	if err != nil {
		return 0, 0, err
	}
	return val, 100, nil
}

func normalizeTitle(title string) string {
	removeYear := regexp.MustCompile(`(.*) (.*)(\\d\\d\\d\\d)`)
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
