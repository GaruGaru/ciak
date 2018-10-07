package server

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/gorilla/mux"
	"net/http"
)

func NewCiakServer(config config.CiakConfig) CiakServer {
	return CiakServer{
		Config: config,
		MediaDiscovery: discovery.FileSystemMediaDiscovery{
			BasePath: config.MediaPath,
		},
	}
}

type CiakServer struct {
	Config         config.CiakConfig
	MediaDiscovery discovery.MediaDiscovery
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
