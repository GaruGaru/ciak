package server

import (
	"fmt"
	"github.com/GaruGaru/ciak/pkg/media/models"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"path/filepath"
)

var playableMediaFormats = []models.MediaFormat{
	models.MediaFormatFlac,
	models.MediaFormatMp4,
	models.MediaFormatMp4a,
	models.MediaFormatMp3,
	models.MediaFormatOgv,
	models.MediaFormatOgm,
	models.MediaFormatOgg,
	models.MediaFormatOga,
	models.MediaFormatOpus,
	models.MediaFormatWebm,
	models.MediaFormatWav,
}

func (s CiakServer) MediaStreamingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	media, err := s.MediaDiscovery.Resolve(params["media"])

	if err != nil {
		if _, err := w.Write([]byte(err.Error())); err != nil {
			logrus.Error(err.Error())
		}
		return
	}

	if isExtensionPlayable(media.Format) {
		w.Header().Set("Accept-Ranges", "bytes")
		w.Header().Set("Content-Type", "video/"+media.Format.Name())
	} else {
		_, fileName := filepath.Split(media.FilePath)
		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	}

	http.ServeFile(w, r, media.FilePath)
}

func isExtensionPlayable(format models.MediaFormat) bool {
	for _, playableFormat := range playableMediaFormats {
		if playableFormat == format {
			return true
		}
	}
	return false
}
