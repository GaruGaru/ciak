package tasks

import (
	"github.com/GaruGaru/ciak/internal/files"
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

	log.Infof("File %s copied to %s", t.Source, t.Destination)

	return nil
}
