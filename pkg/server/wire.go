//go:build wireinject
// +build wireinject

package server

import (
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/google/wire"
)

func InitializeServer() *Server {
	wire.Build(
		createSessionStorage,
		NewServer,
	)
	return &Server{}
}

func createSessionStorage() service.Storage {
	return service.NewLocalStorage()
}
