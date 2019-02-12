package server

import (
	"fmt"
	"github.com/GaruGaru/ciak/internal/utils"
	"github.com/gorilla/mux"
	"net/http"
	"path/filepath"
)

var SupportedVideoFormats = []string{
	"flac", "mp4", "m4a",
	"mp3", "ogv", "ogm",
	"ogg", "oga", "opus",
	"webm", "wav",
}

func (s CiakServer) MediaStreamingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	media, err := s.MediaDiscovery.Resolve(params["media"])

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

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

func IsExtensionPlayable(ext string) bool {
	return utils.StringIn(ext, SupportedVideoFormats)
}
