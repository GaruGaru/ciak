package daemon

import (
	"github.com/GaruGaru/ciak/internal/config"
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"github.com/GaruGaru/ciak/internal/media/encoding"
	"github.com/GaruGaru/duty/duty"
	"github.com/GaruGaru/duty/pool"
	"github.com/GaruGaru/duty/storage"
	log "github.com/sirupsen/logrus"
)

type CiakDaemon struct {
	Conf           config.CiakDaemonConfig
	Duty           duty.Duty
	MediaDiscovery discovery.MediaDiscovery
	Encoder        encoding.MediaEncoder
}

func New(conf config.CiakDaemonConfig, MediaDiscovery discovery.MediaDiscovery, encoder encoding.MediaEncoder) (CiakDaemon, error) {
	store, err := storage.NewBoltDBStorage(conf.Database)

	if err != nil {
		return CiakDaemon{}, err
	}

	dut := duty.New(store, duty.Options{
		Workers:   conf.Workers,
		QueueSize: conf.QueueSize,
		ResultCallback: func(result pool.ScheduledTaskResult) {
		},
	})

	return CiakDaemon{
		Conf:           conf,
		Duty:           dut,
		MediaDiscovery: MediaDiscovery,
		Encoder:        encoder,
	}, nil
}

func (daemon CiakDaemon) Start() error {
	dateFormatter := &log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.001 -0700 MST",
	}
	log.SetFormatter(dateFormatter)
	log.SetReportCaller(true)
	log.WithFields(log.Fields{
		"workers":    daemon.Conf.Workers,
		"queue_size": daemon.Conf.QueueSize,
	}).Info("Ciak daemon started")

	err := daemon.Duty.Init()

	if err != nil {
		return err
	}

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

func (daemon CiakDaemon) Stop() {
	daemon.Duty.Close()
}
