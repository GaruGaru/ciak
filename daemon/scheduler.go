package daemon

import (
	"github.com/GaruGaru/ciak/daemon/tasks"
	"github.com/GaruGaru/ciak/daemon/worker"
	"github.com/GaruGaru/ciak/media/discovery"
	log "github.com/sirupsen/logrus"
)

func (daemon CiakDaemon) Schedule(task worker.Task) bool {
	return daemon.WorkerPool.Schedule(task)
}

func (daemon CiakDaemon) ScheduleWithWorker(task worker.Task) bool {
	return daemon.WorkerPool.ScheduleWithWorker(task)
}

func (daemon CiakDaemon) ScheduleFullMediaConversion() error {
	mediaList, err := daemon.MediaDiscovery.Discover()

	if err != nil {
		return err
	}

	for _, media := range mediaList {
		if daemon.Encoder.CanEncode(media.Extension) {
			daemon.ScheduleMediaConversion(media)
		}
	}

	return err
}

func (daemon CiakDaemon) ScheduleMediaConversion(media discovery.Media) {
	log.Info("Scheduled ", media.Name, " for conversion")
	scheduled := daemon.Schedule(tasks.MediaConvertTask{
		Encoder:        daemon.Encoder,
		Media:          media,
		OutputPath:     daemon.Conf.OutputPath,
		DeleteOriginal: daemon.Conf.DeleteOriginal,
		OverrideOld:    false,
		OutputFormat:   "mp4",
	})

	if !scheduled {
		log.Warn("Unable to schedule media conversion task: task queue is full, try to increase queue size or/and number of workers")
	}
}

func (daemon CiakDaemon) ScheduleFileSystemMediaMonitor() {
	daemon.ScheduleWithWorker(MediaFileSystemChangesTask{
		MonitorPath: daemon.Conf.OutputPath,
		OnFileCreateFn: func() {
			log.Info("Detected changes on media folder, scheduling checks and conversions")
			daemon.ScheduleFullMediaConversion()
		},
	})
}
