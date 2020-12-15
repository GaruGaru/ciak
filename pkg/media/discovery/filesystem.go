package discovery

import (
	"fmt"
	"github.com/GaruGaru/ciak/pkg/media/models"
	"github.com/GaruGaru/ciak/pkg/media/translator"
	log "github.com/sirupsen/logrus"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
)

type FileSystemMediaDiscovery struct {
	BasePath string
}

func NewFileSystemDiscovery(basePath string) FileSystemMediaDiscovery {
	return FileSystemMediaDiscovery{BasePath: basePath}
}

func (d FileSystemMediaDiscovery) Resolve(hash string) (models.Media, error) {
	mediaList, err := d.Discover()

	if err != nil {
		return models.Media{}, nil
	}

	for _, m := range mediaList {
		if m.Hash() == hash {
			return m, nil
		}
	}

	return models.Media{}, fmt.Errorf("no media found with Hash %s", hash)
}

func (d FileSystemMediaDiscovery) Discover() ([]models.Media, error) {

	mediaList := make([]models.Media, 0)

	err := filepath.Walk(d.BasePath, func(filePath string, file os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if file.IsDir() {
			return nil
		}

		media, err := fileToMedia(file, filePath)
		if err == models.ErrMediaFormatNotSupported {
			return nil
		}

		if err != nil {
			return err
		}

		mediaList = append(mediaList, media)

		return nil
	})

	if err != nil {
		return nil, err
	}

	log.Info("Found ", len(mediaList), " media after discovery")

	sort.Slice(mediaList, func(i, j int) bool {
		return mediaList[i].Name < mediaList[j].Name
	})
	return mediaList, nil
}

func fileToMedia(fileInfo os.FileInfo, filePath string) (models.Media, error) {
	extension := path.Ext(filePath)
	mediaExt, err := models.MediaFormatFrom(extension)
	if err != nil {
		return models.Media{}, err
	}

	name := strings.Replace(fileInfo.Name(), extension, "", 1)
	return models.Media{
		Name:     translator.Translate(name),
		FilePath: filePath,
		Size:     fileInfo.Size() / 1024 / 1024,
		Format:   mediaExt,
	}, nil
}
