package main

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/daemon"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	"github.com/GaruGaru/ciak/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.CiakConfig{
		PortBinding: ":8082",
		MediaPath:   "/home/garu/Downloads/",
	}

	mediaDiscovery := discovery.FileSystemMediaDiscovery{BasePath: conf.MediaPath}

	ciakServer := server.CiakServer{
		Config:         conf,
		MediaDiscovery: mediaDiscovery,
	}

	err := ScheduleMediaConversions(mediaDiscovery)

	err = ciakServer.Run()
	log.Fatal(err)
}

func ScheduleMediaConversions(mediaDiscovery discovery.FileSystemMediaDiscovery) error {
	workerPool := daemon.NewWorkerPool(2, 1000)
	mediaList, err := mediaDiscovery.Discover()
	if err != nil {
		panic(err)
	}
	encoder := encoding.FFMpegEncoder{}
	for _, media := range mediaList {
		log.Info("Scheduled ", media.Name, " for conversion")
		workerPool.Schedule(daemon.MediaConvertTask{
			Encoder:    encoder,
			Media:      media,
			OutputPath: "/tmp/",
		})
	}
	workerPool.Start()
	return err
}
