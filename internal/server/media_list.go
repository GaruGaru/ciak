package server

import (
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"github.com/GaruGaru/duty/task"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type PageMedia struct {
	Media          discovery.Media
	TransferStatus task.ScheduledTask
}

type MediaListPage struct {
	Title           string
	MediaCount      int
	PageMedia       []PageMedia
	NoMediasFound   bool
	TransferEnabled bool
}

func (p PageMedia) TButtonClass() string {
	switch p.TransferStatus.Status.State {
	case task.StateSuccess:
		return "btn-success"
	case task.StateError:
		return "btn-danger"
	case task.StateRunning:
		return "btn-secondary"
	case task.StatePending:
		return "btn-warning"
	default:
		return "btn-primary"
	}

}

var mediaListTemplate = template.Must(template.ParseFiles("static/base.html", "static/media-list.html"))

func (s CiakServer) MediaListHandler(w http.ResponseWriter, r *http.Request) {

	mediaList, err := s.MediaDiscovery.Discover()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	pageMediaList := make([]PageMedia, 0)

	for _, media := range mediaList {
		transferResult, err := s.Daemon.Task(media.Hash())
		if err != nil {
			logrus.Error(err)
		}

		pageMediaList = append(pageMediaList, PageMedia{
			Media:          media,
			TransferStatus: transferResult,
		})

	}

	mediaListPage := MediaListPage{
		Title:           "Home",
		MediaCount:      len(pageMediaList),
		PageMedia:       pageMediaList,
		NoMediasFound:   len(pageMediaList) == 0,
		TransferEnabled: s.Daemon.Conf.TransferDestination != "",
	}

	_ = mediaListTemplate.Execute(w, mediaListPage)

}
