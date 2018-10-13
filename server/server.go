package server

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/server/middlewares"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CiakServer struct {
	Config         config.CiakServerConfig
	MediaDiscovery discovery.MediaDiscovery
}

func NewCiakServer(conf config.CiakServerConfig, discovery discovery.MediaDiscovery) CiakServer {
	return CiakServer{
		Config:         conf,
		MediaDiscovery: discovery,
	}
}

func (s CiakServer) Run() error {
	log.WithFields(log.Fields{
		"bind": s.Config.ServerBinding,
	}).Info("Ciak server started")
	router := mux.NewRouter()
	s.initRouting(router)
	return http.ListenAndServe(s.Config.ServerBinding, router)
}

func (s CiakServer) initRouting(router *mux.Router) {
	router.HandleFunc("/probe", ProbeHandler)
	router.HandleFunc("/media", s.MediaListHandler)
	router.HandleFunc("/", s.MediaListHandler)
	router.HandleFunc("/media/{media}", s.MediaStreamingHandler)

	router.Use(middlewares.LoggingMiddleware)
}
