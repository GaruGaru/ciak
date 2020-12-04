package encoding

import (
	"bytes"
	"github.com/GaruGaru/ciak/internal/media/models"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type FFMpegEncoder struct{}

var (
	ConvertibleFormats = []models.MediaFormat{
		models.MediaFormatAvi,
		models.MediaFormatMkv,
	}
)

func FFMpeg() FFMpegEncoder {
	return FFMpegEncoder{}
}

func (FFMpegEncoder) Encode(input string, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "copy", "-acodec", "copy", output)
	cmdOutput := &bytes.Buffer{}
	cmdErrOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdErrOutput

	err := cmd.Run()

	if err != nil {
		log.Info(cmdOutput.String())
		log.Error(cmdErrOutput.String())
	}

	return err
}

func (FFMpegEncoder) CanEncode(format models.MediaFormat) bool {
	for _, convertibleFormat := range ConvertibleFormats {
		if convertibleFormat == format {
			return true
		}
	}
	return false
}
