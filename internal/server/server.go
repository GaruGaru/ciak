package server

import (
	"github.com/GaruGaru/ciak/internal/config"
	"github.com/GaruGaru/ciak/internal/daemon"
	"github.com/GaruGaru/ciak/internal/media/details"
	"github.com/GaruGaru/ciak/internal/media/discovery"
	"github.com/GaruGaru/ciak/internal/server/auth"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	serverVersion = "0.0.2"
)

type CiakServer struct {
	Config           config.CiakServerConfig
	MediaDiscovery   discovery.MediaDiscovery
	Authenticator    auth.Authenticator
	Daemon           daemon.CiakDaemon
	DetailsRetriever *details.Controller
}

func NewCiakServer(
	conf config.CiakServerConfig,
	discovery discovery.MediaDiscovery,
	authenticator auth.Authenticator,
	daemon daemon.CiakDaemon,
	DetailsRetriever *details.Controller,
) CiakServer {
	return CiakServer{
		Config:           conf,
		MediaDiscovery:   discovery,
		Authenticator:    authenticator,
		Daemon:           daemon,
		DetailsRetriever: DetailsRetriever,
	}
}

func (s CiakServer) Run() error {
	log.WithFields(log.Fields{
		"bind":    s.Config.ServerBinding,
		"version": serverVersion,
	}).Info("Ciak server started")

	router := mux.NewRouter()
	s.initRouting(router)
	return http.ListenAndServe(s.Config.ServerBinding, router)
}

func (s CiakServer) initRouting(router *mux.Router) {
	router.HandleFunc("/probe", ProbeHandler)
	router.HandleFunc("/", s.MediaListHandler)
	router.HandleFunc("/media/{media}", s.MediaStreamingHandler)
	router.HandleFunc("/login", s.LoginPageHandler)
	router.HandleFunc("/api/login", s.LoginApiHandler)
	router.HandleFunc("/api/media/transfer", s.MediaTransferApi)
	router.Use(LoggingMiddleware)
	router.Use(s.SessionAuthMiddleware)
}
