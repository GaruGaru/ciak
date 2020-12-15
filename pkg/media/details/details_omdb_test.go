package details

import (
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path"
	"testing"
	"time"
)

func TestRetrieveDataFromOmdb(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		apikey := r.URL.Query()["apikey"]
		if len(apikey) == 0 || apikey[0] != "test" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		queries := r.URL.Query()["t"]
		if len(queries) == 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		query := queries[0]
		if query == "found" {
			response, err := ioutil.ReadFile(path.Join("testdata/omdb-found.json"))
			require.NoError(t, err)
			_, err = w.Write(response)
			require.NoError(t, err)
		} else {
			response, err := ioutil.ReadFile(path.Join("testdata/omdb-not-found.json"))
			require.NoError(t, err)
			_, err = w.Write(response)
			require.NoError(t, err)
		}

	}))
	defer srv.Close()

	omdbClient := &OmdbClient{
		endpoint: srv.URL,
		apiKey:   "test",
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	details, err := omdbClient.Details(Request{
		Title: "found",
	})

	require.NoError(t, err)
	require.Equal(t, details.Name, "found")
	require.Equal(t, details.Director, "Stanley Kubrick")
	require.Equal(t, details.ImagePoster, "https://m.media-amazon.com/images/M/MV5BZWFlYmY2MGEtZjVkYS00YzU4LTg0YjQtYzY1ZGE3NTA5NGQxXkEyXkFqcGdeQXVyMTQxNzMzNDI@._V1_SX300.jpg")
	require.NotEqual(t, details.ReleaseDate, time.Time{})
	require.Equal(t, details.Genre, "Drama, Horror")

	require.Equal(t, details.MaxRating, 100.)
	require.Equal(t, details.Rating, 84.)

	_, err = omdbClient.Details(Request{
		Title: "notfound",
	})
	require.Error(t, err)

}
