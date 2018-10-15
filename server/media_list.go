package server

import (
	"github.com/GaruGaru/ciak/discovery"
	"html/template"
	"net/http"
)

type MediaListPage struct {
	Title         string
	MediaCount    int
	MediaList     []discovery.Media
	NoMediasFound bool
}

func (s CiakServer) MediaListHandler(w http.ResponseWriter, r *http.Request) {
	mediaList, err := s.MediaDiscovery.Discover()

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	pageTemplate, err := template.ParseFiles("static/base.html", "static/media-list.html")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	page := MediaListPage{
		Title:         "Home",
		MediaCount:    len(mediaList),
		MediaList:     mediaList,
		NoMediasFound: len(mediaList) == 0,
	}

	pageTemplate.Execute(w, page)

}
