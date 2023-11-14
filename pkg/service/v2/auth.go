package servicev2

import (
	"context"

	"github.com/acerohernan/twirp-boilerplate/core"
	authv2 "github.com/acerohernan/twirp-boilerplate/core/auth/v2"
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
)

type AuthService struct {
	storage service.Storage
}

func NewAuthService(storage service.Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) CreateSession(ctx context.Context, req *authv2.CreateSessionRequest) (*authv2.CreateSessionResponse, error) {
	if req.PhoneNumber == "" {
		return nil, service.ErrBadRequest
	}

	sess := &core.Session{
		Id: utils.NewGuid(utils.SessionPrefix),
	}

	err := s.storage.StoreSession(ctx, sess)

	if err != nil {
		return nil, err
	}

	return &authv2.CreateSessionResponse{
		Session: sess,
	}, nil
}

func (s *AuthService) ListSessions(ctx context.Context, req *authv2.ListSessionsRequest) (*authv2.ListSessionsResponse, error) {
	sess, err := s.storage.ListSessions(ctx)

	if err != nil {
		return nil, err
	}

	return &authv2.ListSessionsResponse{
		Sessions: sess,
	}, nil
}
