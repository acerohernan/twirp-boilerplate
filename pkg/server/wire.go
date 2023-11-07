//go:build wireinject
// +build wireinject

package server

import (
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/google/wire"
)

func InitializeServer() (*Server, error) {
	wire.Build(
		createSessionStorage,
		service.NewAuthService,
		NewServer,
	)
	return &Server{}, nil
}

func createSessionStorage() service.Storage {
	return service.NewLocalStorage()
}
