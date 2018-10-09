package discovery

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"strings"
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

	err := filepath.Walk(d.BasePath, func(path string, file os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if !file.IsDir() {
			mediaList = append(mediaList, fileToMedia(file, path))
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
	return Media{
		Name:      strings.Replace(fileInfo.Name(), extension, "", 1),
		FilePath:  filePath,
		Size:      fileInfo.Size(),
		Extension: strings.TrimLeft(extension, "."),
	}
}
