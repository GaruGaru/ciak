package server

import (
	"github.com/GaruGaru/ciak/internal/config"
	"github.com/GaruGaru/ciak/internal/daemon"
	"github.com/GaruGaru/ciak/internal/server/auth"
	"github.com/GaruGaru/ciak/pkg/omdb"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginApiSuccess(t *testing.T) {
	const password = "test_password"
	const username = "test_username"

	srv := NewCiakServer(
		config.CiakServerConfig{AuthenticationEnabled: false},
		nil,
		auth.NewStaticCredentialsApi(username, password),
		daemon.CiakDaemon{},
		omdb.NoOpClient{},
	)

	form := url.Values{}
	form.Add("username", username)
	form.Add("password", password)

	req := httptest.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.PostForm = form

	resp := httptest.NewRecorder()

	srv.LoginApiHandler(resp, req)

	require.Equal(t, http.StatusFound, resp.Code)
	val, present := resp.Header()["Set-Cookie"]
	require.True(t, present)
	require.NotEmpty(t, val)
}

func TestLoginApiFail(t *testing.T) {
	const password = "test_password"
	const username = "test_username"

	srv := NewCiakServer(
		config.CiakServerConfig{AuthenticationEnabled: false},
		nil,
		auth.NewStaticCredentialsApi(username, password),
		daemon.CiakDaemon{},
		omdb.NoOpClient{},
	)

	form := url.Values{}
	form.Add("username", username)
	form.Add("password", "incorrect"+password)

	req := httptest.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
	req.PostForm = form

	resp := httptest.NewRecorder()

	srv.LoginApiHandler(resp, req)

	require.Equal(t, http.StatusFound, resp.Code)
	_, present := resp.Header()["Set-Cookie"]
	require.False(t, present)
}
