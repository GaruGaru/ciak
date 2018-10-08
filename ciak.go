package main

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	"github.com/GaruGaru/ciak/server"
	"log"
)

func main() {
	conf := config.CiakConfig{
		PortBinding: ":8082",
		MediaPath:   "/tmp/",
	}

	mediaDiscovery := discovery.FileSystemMediaDiscovery{BasePath: conf.MediaPath}

	mediaEncoder := encoding.FFMpegEncoder{}

	ciak := server.NewCiakServer(conf, mediaDiscovery, mediaEncoder)

	err := ciak.Run()
	log.Fatal(err)
}
