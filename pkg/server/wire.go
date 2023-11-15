//go:build wireinject
// +build wireinject

package server

import (
	"github.com/acerohernan/twirp-boilerplate/pkg/config"
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	servicev1 "github.com/acerohernan/twirp-boilerplate/pkg/service/v1"
	servicev2 "github.com/acerohernan/twirp-boilerplate/pkg/service/v2"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
	"github.com/google/wire"
)

func InitializeServer(conf *config.Config) (*Server, error) {
	wire.Build(
		createSessionStorage,
		utils.NewProtoValidator,
		servicev1.NewAuthService,
		servicev2.NewAuthService,
		NewServer,
	)
	return &Server{}, nil
}

func createSessionStorage() service.Storage {
	return service.NewLocalStorage()
}
