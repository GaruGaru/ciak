package tasks

import (
	"fmt"
	"github.com/GaruGaru/ciak/pkg/files"
	log "github.com/sirupsen/logrus"
	"time"
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


	sourceHash, err := files.HashFile(t.Source)

	if err != nil {
		return err
	}

	err = files.CopyFile(t.Source, t.Destination)

	if err != nil {
		return err
	}

	destinationHash, err := files.HashDir(t.Destination)

	if err != nil {
		return err
	}

	if sourceHash != destinationHash {
		return fmt.Errorf("copy successfully but hashes don't match %s != %s", sourceHash, destinationHash)
	}

	fmt.Println("File copied successfully")

	return nil
}
