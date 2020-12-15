package server

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func ProbeHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("OK")); err != nil {
		logrus.Error(err.Error())
	}
}
