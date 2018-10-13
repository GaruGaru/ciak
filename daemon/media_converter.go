package daemon

import (
	"fmt"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type MediaConvertTask struct {
	Encoder        encoding.MediaEncoder
	Media          discovery.Media
	OutputPath     string
	DeleteOriginal bool
	OverrideOld    bool
	OutputFormat   string
}

func (mt MediaConvertTask) Run() error {

	if mt.OutputFormat == "" {
		mt.OutputFormat = "mp4"
	}

	log.Info("Trying to convert ", mt.Media.Name)

	if !mt.Encoder.CanEncode(mt.Media.Extension) {
		log.Warn("Unable to convert ", mt.Media.FilePath, " extension not supported.")
		return nil
	}

	_, srcName := filepath.Split(mt.Media.FilePath)

	outFile := fmt.Sprintf("%s.%s", srcName, mt.OutputFormat)

	outPath := filepath.Join(mt.OutputPath, outFile)

	if mt.OverrideOld {
		if os.Remove(outPath) == nil {
			log.Warn("Remove pre existing converted file")
		}
	} else {
		if _, err := os.Stat(outPath); !os.IsNotExist(err) {
			log.Warn("Media already converted: ", mt.Media.Name)
			return nil
		}
	}

	err := mt.Encoder.Encode(mt.Media.FilePath, outPath)

	if err != nil {
		log.Warn("Unable to convert ", mt.Media.FilePath, " encoding error: ", err)
		return err
	}

	output := discovery.Media{
		Name:      srcName,
		Extension: "mp4",
		FilePath:  outPath,
		Size:      0,
	}

	if mt.DeleteOriginal {
		err = os.Remove(mt.Media.FilePath)
		if err != nil {
			log.Warn("Unable to delete original media: ", err)
		}
	}

	log.Info("Media convert task completed successfully: ", output)

	return nil
}

type ConvertAllMediaTask struct {
	MediaDiscovery discovery.MediaDiscovery
	OutputPath     string
	Encoder        encoding.MediaEncoder
}

func (ct ConvertAllMediaTask) Run() error {

	mediaList, err := ct.MediaDiscovery.Discover()
	if err != nil {
		panic(err)
	}

	for _, media := range mediaList {
		log.Info("Scheduled ", media.Name, " for conversion")
		// TODO Implement conversion
		//workerPool.Schedule(daemon.MediaConvertTask{
		//	Encoder:    encoder,
		//	Media:      media,
		//	OutputPath: "/tmp/",
		//})
	}
	return err
}
