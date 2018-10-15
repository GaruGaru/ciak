package server

import (
	"github.com/GaruGaru/ciak/server/auth"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"html/template"
	"net/http"
)

type LoginPage struct {
	Title string
}

var store = sessions.NewCookieStore([]byte("test"))

func (s CiakServer) LoginApiHandler(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	authUser, err := s.Authenticator.Authenticate(username, password)
	if err == nil {
		s.createSession(w, r, authUser)
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", 302)
	}

}

func (s CiakServer) createSession(w http.ResponseWriter, r *http.Request, user auth.User) {
	session, err := store.Get(r, "user")

	if err != nil {
		logrus.Warn("Error creating the session: ", err)
		return
	}

	session.Values["username"] = user.Name
	store.Save(r, w, session)
}

func (s CiakServer) LoginPageHandler(w http.ResponseWriter, r *http.Request) {

	pageTemplate, err := template.ParseFiles("static/base.html", "static/login.html")

	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	pageTemplate.Execute(w, LoginPage{
		Title: "Login",
	})

}
