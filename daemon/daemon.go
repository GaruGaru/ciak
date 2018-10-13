package daemon

import (
	"github.com/GaruGaru/ciak/config"
	log "github.com/sirupsen/logrus"
)

type CiakDaemon struct {
	Conf       config.CiakDaemonConfig
	WorkerPool WorkerPool
}

func NewCiakDaemon(conf config.CiakDaemonConfig) CiakDaemon {
	return CiakDaemon{
		Conf:       conf,
		WorkerPool: NewWorkerPool(conf.Workers, conf.QueueSize),
	}
}

func (d CiakDaemon) Start() error {
	log.WithFields(log.Fields{
		"workers":    d.Conf.Workers,
		"queue_size": d.Conf.QueueSize,
	}).Info("Ciak daemon started")
	d.WorkerPool.Start()
	return nil
}

func (d CiakDaemon) Stop() {
	d.WorkerPool.Stop()
}

func (d CiakDaemon) Schedule(task Task) {
	d.WorkerPool.Schedule(task)
}
