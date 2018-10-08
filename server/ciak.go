package server

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/encoding"
	"github.com/gorilla/mux"
	"net/http"
)

func NewCiakServer(config config.CiakConfig, mediaDiscovery discovery.MediaDiscovery, encoder encoding.MediaEncoder) CiakServer {
	return CiakServer{
		Config:         config,
		MediaDiscovery: mediaDiscovery,
		MediaEncoder:   encoder,
	}
}

type CiakServer struct {
	Config         config.CiakConfig
	MediaDiscovery discovery.MediaDiscovery
	MediaEncoder   encoding.MediaEncoder
}

func (s CiakServer) Run() error {
	router := mux.NewRouter()
	s.initRouting(router)
	return http.ListenAndServe(s.Config.PortBinding, router)
}

func (s CiakServer) initRouting(router *mux.Router) {
	router.HandleFunc("/probe", ProbeHandler)
	router.HandleFunc("/media", s.MediaListHandler)
	router.HandleFunc("/", s.MediaListHandler)
	router.HandleFunc("/media/{media}", s.MediaStreamingHandler)
}
