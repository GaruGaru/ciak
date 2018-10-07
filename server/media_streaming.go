package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

func (s CiakServer) MediaStreamingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	media, err := s.MediaDiscovery.Resolve(params["media"])

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	if IsExtensionPlayable(media.Extension) {
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Type", "video/"+media.Extension)
	} else {
		_, fileName := filepath.Split(media.FilePath)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	}

	http.ServeFile(w, r, media.FilePath)
}
