package main

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	"github.com/GaruGaru/ciak/server"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.CiakConfig{
		PortBinding: ":8082",
		MediaPath:   "/tmp",
	}

	log.WithFields(log.Fields{
		"port":  conf.PortBinding,
		"media": conf.MediaPath,
	}).Info("Ciak server")

	mediaDiscovery := discovery.FileSystemMediaDiscovery{BasePath: conf.MediaPath}

	mediaEncoder := encoding.FFMpegEncoder{}

	ciak := server.NewCiakServer(conf, mediaDiscovery, mediaEncoder)

	err := ciak.Run()
	log.Fatal(err)
}
