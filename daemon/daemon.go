package daemon

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/daemon/worker"
	"github.com/GaruGaru/ciak/media/discovery"
	"github.com/GaruGaru/ciak/media/encoding"
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
