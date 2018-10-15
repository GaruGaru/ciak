package auth

import (
	"fmt"
	"os"
)

type User struct {
	Name string
}

type Authenticator interface {
	Authenticate(username string, password string) (User, error)
}

type NoOpAuthenticator struct{}

func (a NoOpAuthenticator) Authenticate(username string, password string) (User, error) {
	return User{Name: username}, nil
}

type EnvAuthenticator struct{}

func (a EnvAuthenticator) Authenticate(username string, password string) (User, error) {
	envUser := os.Getenv("CIAK_USERNAME")
	envPassword := os.Getenv("CIAK_PASSWORD")

	if username == envUser && password == envPassword {
		return User{Name: username}, nil
	} else {
		return User{}, fmt.Errorf("login error")
	}
}
