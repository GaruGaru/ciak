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
	panic("TransferTask")
}

func (t TransferTask) Run() error {
	log.Infof("Copying %s to %s", t.Source, t.Destination)

	sourceHash, err := files.HashDir(t.Source)

	if err != nil {
		return nil
	}

	err = files.CopyDirectory(t.Source, t.Destination)

	if err != nil {
		return nil
	}

	destinationHash, err := files.HashDir(t.Destination)

	if err != nil {
		return nil
	}

	if sourceHash != destinationHash {
		return fmt.Errorf("copy successfully but hashes don't match %s != %s", sourceHash, destinationHash)
	}

	return nil
}
