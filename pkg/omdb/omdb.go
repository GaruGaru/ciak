package omdb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

type Omdb struct {
	ApiKey     string
	Cache      Cache
	httpClient *http.Client
}

type MovieByTitle struct {
	Title string
	Movie Movie
}

func New(apiKey string) Client {
	if apiKey == "" {
		return &NoOpClient{}
	}
	return &Omdb{
		ApiKey: apiKey,
		Cache:  &MemoryCache{},
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (o *Omdb) ByTitleBulk(titles ...string) (map[string]Movie, error) {

	var wg sync.WaitGroup

	wg.Add(len(titles))

	results := make(chan MovieByTitle, len(titles))

	for _, title := range titles {
		go func(t string) {
			defer wg.Done()
			movie, _, err := o.ByTitle(t)
			if err != nil {
				logrus.Warnf("unable to get title metadata for %s: %s", t, err.Error())
			}

			results <- MovieByTitle{
				Title: t,
				Movie: movie,
			}
		}(title)
	}

	wg.Wait()

	close(results)

	out := make(map[string]Movie)

	for res := range results {
		out[res.Title] = res.Movie
	}

	return out, nil
}

func (o *Omdb) ByTitle(title string) (Movie, bool, error) {

	cached, present := o.Cache.Get(title)

	if present {
		return cached, true, nil
	}

	apiUrl := fmt.Sprintf("http://www.omdbapi.com/?apikey=%s&t=%s", o.ApiKey, url.QueryEscape(normalizeTitle(title)))

	resp, err := o.httpClient.Get(apiUrl)

	if err != nil {
		return Movie{}, false, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			logrus.Warn("unable to close response body: " + err.Error())
		}
	}()

	var movie Movie

	err = json.NewDecoder(resp.Body).Decode(&movie)

	o.Cache.Put(title, movie)

	if err != nil {
		return Movie{}, false, err
	}

	if movie.Response == "False" {
		return Movie{}, false, nil
	}

	return movie, true, nil
}

func normalizeTitle(title string) string {

	removeYear := regexp.MustCompile("(.*) (.*)(\\d\\d\\d\\d)")

	matches := removeYear.FindAllString(title, -1)

	if len(matches) == 0 {
		return title
	}

	return matches[0]
}
