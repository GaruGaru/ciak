package server

import (
	"github.com/GaruGaru/ciak/pkg/media/details"
	"github.com/GaruGaru/ciak/pkg/media/models"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type PageMediaRating struct {
	Value   float64
	Max     float64
	Present bool
}

type PageMedia struct {
	Media    models.Media
	Cover    string
	Playable bool
	Rating   PageMediaRating
}

type MediaListPage struct {
	Title         string
	MediaCount    int
	PageMedia     []PageMedia
	NoMediasFound bool
}

func mediaToTitlesList(media []models.Media) []details.Request {
	titles := make([]details.Request, 0)
	for _, item := range media {
		titles = append(titles, details.Request{
			Title: item.Name,
		})
	}
	return titles
}

func (s CiakServer) MediaListHandler(w http.ResponseWriter, r *http.Request) {
	mediaList, err := s.MediaDiscovery.Discover()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	mediaMetadata, err := s.DetailsRetriever.DetailsByTitleBulk(mediaToTitlesList(mediaList)...)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	pageMediaList := make([]PageMedia, 0)

	for _, media := range mediaList {

		metadata := mediaMetadata[media.Name]

		if metadata.ImagePoster == "" {
			metadata.ImagePoster = "https://via.placeholder.com/300"
		}

		pageMediaList = append(pageMediaList, PageMedia{
			Media:    media,
			Cover:    metadata.ImagePoster,
			Playable: isExtensionPlayable(media.Format),
			Rating: PageMediaRating{
				Value:   metadata.Rating,
				Max:     metadata.MaxRating,
				Present: metadata.MaxRating != 0,
			},
		})

	}

	mediaListPage := MediaListPage{
		Title:         "Home",
		MediaCount:    len(pageMediaList),
		PageMedia:     pageMediaList,
		NoMediasFound: len(pageMediaList) == 0,
	}

	var mediaListTemplate = template.Must(template.ParseFiles("ui/base.html", "ui/media-list.html"))

	if err := mediaListTemplate.Execute(w, mediaListPage); err != nil {
		logrus.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
