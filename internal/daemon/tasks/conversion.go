package tasks

import (
	"fmt"
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"github.com/GaruGaru/ciak/internal/media/encoding"
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

func (mt MediaConvertTask) Type() string {
	panic("MediaConvertTask")
}

func (mt MediaConvertTask) Run() error {

	if mt.OutputFormat == "" {
		return fmt.Errorf("empty output format provided")
	}

	log.Info("Converting media ", mt.Media.Name)

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
		Extension: mt.OutputFormat,
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
