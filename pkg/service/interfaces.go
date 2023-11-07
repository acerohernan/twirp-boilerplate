package service

import (
	"context"

	"github.com/acerohernan/twirp-boilerplate/pkg/auth"
)

type Storage interface {
	StoreSession(ctx context.Context, session *auth.Session) error
}
