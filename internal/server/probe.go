package server

import "net/http"

func ProbeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
