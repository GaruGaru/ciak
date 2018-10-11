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
	Encoder    encoding.MediaEncoder
	Media      discovery.Media
	OutputPath string
}

func (mt MediaConvertTask) Run() error {

	log.Info("Trying to convert ", mt.Media.Name)

	if !mt.Encoder.CanEncode(mt.Media.Extension) {
		log.Warn("Unable to convert ", mt.Media.FilePath, " extension not supported.")
		return nil
	}

	_, srcName := filepath.Split(mt.Media.FilePath)

	outFile := fmt.Sprintf("%s.%s", srcName, "mp4")

	outPath := filepath.Join(mt.OutputPath, outFile)

	os.Remove(outPath)

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

	log.Info("Media convert task completed successfully: ", output)

	return nil
}
