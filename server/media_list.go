package server

import (
	"github.com/GaruGaru/ciak/media/discovery"
	"html/template"
	"net/http"
)

type MediaListPage struct {
	Title         string
	MediaCount    int
	MediaList     []discovery.Media
	NoMediasFound bool
}

var mediaListTemplate = template.Must( template.ParseFiles("static/base.html", "static/media-list.html"))

func (s CiakServer) MediaListHandler(w http.ResponseWriter, r *http.Request) {
	mediaList, err := s.MediaDiscovery.Discover()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	mediaListPage := MediaListPage{
		Title:         "Home",
		MediaCount:    len(mediaList),
		MediaList:     mediaList,
		NoMediasFound: len(mediaList) == 0,
	}

	mediaListTemplate.Execute(w, mediaListPage)

}
