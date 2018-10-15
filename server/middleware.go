package server

import (
	"fmt"
	"github.com/GaruGaru/ciak/utils"
	"net/http"
)

var UnauthenticatedUrls = []string{
	"/login",
	"/probe",
	"/api/login",
}

func (s CiakServer) SessionAuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if !s.Config.AuthenticationEnabled || utils.StringIn(r.URL.Path, UnauthenticatedUrls) {
			next.ServeHTTP(w, r)
			return
		}

		session, err := store.Get(r, "user")

		if err != nil {
			fmt.Println("Session error ", err)
			return
		}

		if session.Values["username"] != nil {
			next.ServeHTTP(w, r)
		} else {
			http.Redirect(w, r, "/login", http.StatusFound)
		}

	})
}
