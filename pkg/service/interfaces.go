package service

import (
	"context"

	"github.com/acerohernan/twirp-boilerplate/rpc"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Storage
type Storage interface {
	StoreSession(ctx context.Context, session *rpc.Session) error
	ListSessions(ctx context.Context) ([]*rpc.Session, error)
}
