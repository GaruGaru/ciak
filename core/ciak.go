package core

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/daemon"
	"github.com/GaruGaru/ciak/server"
)

type Ciak struct {
	Config config.CiakConfig
	Server server.CiakServer
	Daemon daemon.CiakDaemon
}
