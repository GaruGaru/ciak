package server

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		startTime := time.Now()
		next.ServeHTTP(w, r)

		log.WithFields(log.Fields{
			"duration": time.Since(startTime).Milliseconds(),
		}).Info(r.RequestURI)

	})
}
