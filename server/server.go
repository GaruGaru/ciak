package server

import (
	"github.com/GaruGaru/ciak/config"
	"github.com/GaruGaru/ciak/discovery"
	"github.com/GaruGaru/ciak/server/auth"
	"github.com/GaruGaru/ciak/server/common"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type CiakServer struct {
	Config         config.CiakServerConfig
	MediaDiscovery discovery.MediaDiscovery
	Authenticator  auth.Authenticator
}

func NewCiakServer(conf config.CiakServerConfig, discovery discovery.MediaDiscovery) CiakServer {
	var authenticator auth.Authenticator = auth.NoOpAuthenticator{}

	if conf.AuthenticationEnabled {
		authenticator = auth.EnvAuthenticator{}
	}

	return CiakServer{
		Config:         conf,
		MediaDiscovery: discovery,
		Authenticator:  authenticator,
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
	router.HandleFunc("/", s.MediaListHandler)
	router.HandleFunc("/media/{media}", s.MediaStreamingHandler)
	router.HandleFunc("/login", s.LoginPageHandler)
	router.HandleFunc("/api/login", s.LoginApiHandler)
	router.Use(common.LoggingMiddleware)
	router.Use(s.SessionAuthMiddleware)
}
