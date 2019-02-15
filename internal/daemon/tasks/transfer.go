package tasks

import (
	"fmt"
	"github.com/GaruGaru/ciak/pkg/files"
	log "github.com/sirupsen/logrus"
)

type TransferTask struct {
	Source      string
	Destination string
}

func (t TransferTask) Type() string {
	return "media-transfer-task"
}

func (t TransferTask) Run() error {
	log.Infof("Copying %s to %s", t.Source, t.Destination)

	err := files.CopyFile(t.Source, t.Destination)

	if err != nil {
		return err
	}

	fmt.Println("File copied successfully")

	return nil
}
