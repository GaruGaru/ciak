package main

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/server"
	"log"
)

func main() {
	conf := config.CiakConfig{
		PortBinding: ":8080",
		MediaPath:   "/tmp/",
	}
	ciak := server.NewCiakServer(conf)
	err := ciak.Run()
	log.Fatal(err)
}
