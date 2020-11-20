package daemon

import (
	"fmt"
	"github.com/GaruGaru/ciak/internal/daemon/tasks"
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"github.com/GaruGaru/ciak/internal/media/models"
	"github.com/GaruGaru/duty/task"
	log "github.com/sirupsen/logrus"
	"path/filepath"
)

func (daemon CiakDaemon) Task(id string) (task.ScheduledTask, error) {
	return daemon.Duty.Get(id)
}

func (daemon CiakDaemon) Schedule(task task.Task) error {
	_, err := daemon.Duty.Enqueue(task)
	return err
}

func (daemon CiakDaemon) ScheduleWithID(id string, task task.Task) error {
	_, err := daemon.Duty.EnqueueWithID(id, task)
	return err
}

func (daemon CiakDaemon) ScheduleMediaTransfer(media discovery.Media) error {

	output := filepath.Join(daemon.Conf.TransferDestination, fmt.Sprintf("%s.%s", media.Name, media.Format.Name()))

	transfer := tasks.TransferTask{
		Source:      media.FilePath,
		Destination: output,
	}

	return daemon.ScheduleWithID(media.Hash(), transfer)
}

func (daemon CiakDaemon) ScheduleFullMediaConversion() error {
	mediaList, err := daemon.MediaDiscovery.Discover()

	if err != nil {
		return err
	}

	for _, media := range mediaList {
		if daemon.Encoder.CanEncode(media.Format) {
			daemon.ScheduleMediaConversion(media)
		}
	}

	return err
}

func (daemon CiakDaemon) ScheduleMediaConversion(media discovery.Media) {
	log.Info("Scheduled ", media.Name, " for conversion")
	err := daemon.Schedule(tasks.MediaConvertTask{
		Encoder:        daemon.Encoder,
		Media:          media,
		OutputPath:     daemon.Conf.OutputPath,
		DeleteOriginal: daemon.Conf.DeleteOriginal,
		OverrideOld:    false,
		OutputFormat:   models.MediaFormatMp4,
	})

	if err != nil {
		log.Warn("Unable to schedule media conversion task: task queue is full, try to increase queue size or/and number of workers")
	}
}

func (daemon CiakDaemon) ScheduleFileSystemMediaMonitor() {
	err := daemon.Schedule(MediaFileSystemChangesTask{
		MonitorPath: daemon.Conf.OutputPath,
		OnFileCreateFn: func() {
			log.Info("Detected changes on media folder, scheduling checks and conversions")
			err := daemon.ScheduleFullMediaConversion()

			if err != nil {
				log.Warn("Unable to schedule media conversion task: task queue is full, try to increase queue size or/and number of workers")
			}
		},
	})

	if err != nil {
		log.Warn("Unable to schedule media monitor task: task queue is full, try to increase queue size or/and number of workers")
	}
}
