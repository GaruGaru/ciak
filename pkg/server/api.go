package server

import (
	"net/http"
)

func (s CiakServer) MediaTransferApi(w http.ResponseWriter, r *http.Request) {

	mediaID := r.URL.Query().Get("media")

	if mediaID == "" {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("media id parameter not provided"))
		return
	}

	media, err := s.MediaDiscovery.Resolve(mediaID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("media with id " + mediaID + " not found"))
		return
	}

	err = s.Daemon.ScheduleMediaTransfer(media)

	if err != nil {
		w.WriteHeader(http.StatusTooManyRequests)
		_, _ = w.Write([]byte("worker pool full, unable to accept other tasks"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write([]byte{})
}
