package encoding

import (
	"bytes"
	"fmt"
	"os/exec"
)

type FFMpegEncoder struct {
}

var (
	SupportedExtensions = []string{"avi", "mp4", "mkv"}
)

func (FFMpegEncoder) Encode(input string, output string) (error) {
	cmd := exec.Command("ffmpeg", "-i", input, output)
	cmdOutput := &bytes.Buffer{}
	cmdErrOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	cmd.Stderr = cmdErrOutput

	err := cmd.Run()

	fmt.Println(string(cmdOutput.Bytes()))
	fmt.Println(string(cmdErrOutput.Bytes()))
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
