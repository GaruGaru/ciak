package server

import (
	"github.com/GaruGaru/ciak/internal/daemon/tasks"
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"net/http"
)

func (s CiakServer) MediaTransferApi(w http.ResponseWriter, r *http.Request) {

	mediaID := r.URL.Query().Get("media")

	if mediaID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("media id parameter not provided"))
		return
		return
	}

	mediaList, err := s.MediaDiscovery.Discover()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}

	media, found := findMediaById(mediaList, mediaID)

	if !found {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("media with id " + mediaID + " not found"))
		return
	}

	transferTask := tasks.TransferTask{
		Source:      media.FilePath,
		Destination: "/tmp",
	}

	scheduled := s.Daemon.WorkerPool.Schedule(transferTask)

	if !scheduled {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte("worker pool full, unable to accept other tasks"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte{})
}

func findMediaById(list []discovery.Media, hash string) (discovery.Media, bool) {
	for _, item := range list {
		if item.Hash() == hash {
			return item, true
		}
	}
	return discovery.Media{}, false
}
