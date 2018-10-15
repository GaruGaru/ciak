package daemon

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	"github.com/GaruGaru/ciak/worker"
	log "github.com/sirupsen/logrus"
)

type CiakDaemon struct {
	Conf           config.CiakDaemonConfig
	WorkerPool     worker.WorkerPool
	MediaDiscovery discovery.MediaDiscovery
	Encoder        encoding.MediaEncoder
}

func NewCiakDaemon(conf config.CiakDaemonConfig, MediaDiscovery discovery.MediaDiscovery, encoder encoding.MediaEncoder) CiakDaemon {
	return CiakDaemon{
		Conf:           conf,
		WorkerPool:     worker.NewWorkerPool(conf.Workers, conf.QueueSize),
		MediaDiscovery: MediaDiscovery,
		Encoder:        encoder,
	}
}

func (daemon CiakDaemon) initialize() error {
	if daemon.Conf.AutoConvertMedia {
		err := daemon.ScheduleFullMediaConversion()
		if err != nil {
			log.Warn("unable to schedule auto media conversion: ", err)
			return err
		}
		daemon.ScheduleFileSystemMediaMonitor()
	}
	return nil
}

func (daemon CiakDaemon) Start() error {
	log.WithFields(log.Fields{
		"workers":    daemon.Conf.Workers,
		"queue_size": daemon.Conf.QueueSize,
	}).Info("Ciak daemon started")
	err := daemon.initialize()
	daemon.WorkerPool.Start()
	return err
}

func (daemon CiakDaemon) Stop() {
	daemon.WorkerPool.Stop()
}

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
	scheduled := daemon.Schedule(MediaConvertTask{
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
