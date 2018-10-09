package server

import (
	"fmt"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path/filepath"
)

func (s CiakServer) MediaStreamingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	media, err := s.MediaDiscovery.Resolve(params["media"])

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	media, err = s.tryEncodeMedia(media)

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

func (s CiakServer) tryEncodeMedia(media discovery.Media) (discovery.Media, error) {

	if !s.MediaEncoder.CanEncode(media.Extension) {
		return media, nil
	}

	srcPath, srcName := filepath.Split(media.FilePath)

	outFile := fmt.Sprintf("%s.%s", srcName, "mp4")

	outPath := filepath.Join(srcPath, outFile)

	os.Remove(outPath)

	err := s.MediaEncoder.Encode(media.FilePath, outPath)

	if err != nil {
		return discovery.Media{}, err
	}

	return discovery.Media{
		Name:      srcName,
		Extension: "mp4",
		FilePath:  outPath,
		Size:      0,
	}, nil

}
