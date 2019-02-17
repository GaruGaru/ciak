package discovery

import (
	"fmt"
	"github.com/GaruGaru/ciak/internal/media/translator"
	"github.com/GaruGaru/ciak/internal/utils"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
)

var (
	FormatsWhitelist = []string{
		".avi", ".mkv", ".flac", ".mp4", ".m4a", ".mp3", ".ogv",
		".ogm", ".ogg", ".oga", ".opus", ".webm", ".wav",
	}
)

type FileSystemMediaDiscovery struct {
	BasePath string
}

func (d FileSystemMediaDiscovery) Resolve(hash string) (Media, error) {
	mediaList, err := d.Discover()

	if err != nil {
		return Media{}, nil
	}

	for _, m := range mediaList {
		if m.Hash() == hash {
			return m, nil
		}
	}

	return Media{}, fmt.Errorf("no media found with Hash %s", hash)
}

func (d FileSystemMediaDiscovery) Discover() ([]Media, error) {

	mediaList := make([]Media, 0)

	err := filepath.Walk(d.BasePath, func(filePath string, file os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !file.IsDir() && utils.StringIn(path.Ext(filePath), FormatsWhitelist) {
			mediaList = append(mediaList, fileToMedia(file, filePath))
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	log.Info("Found ", len(mediaList), " media after discovery")

	return mediaList, nil
}

func fileToMedia(fileInfo os.FileInfo, filePath string) Media {
	extension := path.Ext(filePath)

	name := strings.Replace(fileInfo.Name(), extension, "", 1)
	return Media{
		Name:      translator.Translate(name),
		FilePath:  filePath,
		Size:      fileInfo.Size() / 1024 / 1024,
		Extension: strings.TrimLeft(extension, "."),
	}
}
