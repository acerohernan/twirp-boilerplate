package service

import (
	"context"

	"github.com/acerohernan/twirp-boilerplate/core"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Storage
type Storage interface {
	StoreSession(ctx context.Context, session *core.Session) error
	ListSessions(ctx context.Context) ([]*core.Session, error)
}
