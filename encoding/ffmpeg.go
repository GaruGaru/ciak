package encoding

import (
	"bytes"
	log "github.com/sirupsen/logrus"
	"os/exec"
)

type FFMpegEncoder struct {
}

var (
	SupportedExtensions = []string{"avi", "mp4", "mkv"}
)

func (FFMpegEncoder) Encode(input string, output string) error {
	cmd := exec.Command("ffmpeg", "-i", input, "-vcodec", "copy", "-acodec", "copy", output)
	cmdOutput := &bytes.Buffer{}
	cmdErrOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdErrOutput

	err := cmd.Run()

	log.Info(string(cmdOutput.Bytes()))
	log.Error(string(cmdErrOutput.Bytes()))

	return err
}

func (FFMpegEncoder) CanEncode(extension string) bool {
	for _, ext := range SupportedExtensions {
		if ext == extension {
			return true
		}
	}
	return false
}
