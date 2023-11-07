package service

import (
	"context"

	"github.com/acerohernan/twirp-boilerplate/core"
)

type Storage interface {
	StoreSession(ctx context.Context, session *core.Session) error
}
