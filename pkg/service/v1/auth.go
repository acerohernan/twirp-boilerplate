package servicev1

import (
	"context"
	"net/mail"

	"github.com/acerohernan/twirp-boilerplate/core"
	authv1 "github.com/acerohernan/twirp-boilerplate/core/auth/v1"
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

func (s *AuthService) CreateSession(ctx context.Context, req *authv1.CreateSessionRequest) (*authv1.CreateSessionResponse, error) {
	if req.Email == "" {
		return nil, service.ErrBadRequest
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return nil, service.ErrBadRequest
	}

	sess := &core.Session{
		Id: utils.NewGuid(utils.SessionPrefix),
	}

	err := s.storage.StoreSession(ctx, sess)

	if err != nil {
		return nil, err
	}

	return &authv1.CreateSessionResponse{
		Session: sess,
	}, nil
}

func (s *AuthService) ListSessions(ctx context.Context, req *authv1.ListSessionsRequest) (*authv1.ListSessionsResponse, error) {
	sess, err := s.storage.ListSessions(ctx)

	if err != nil {
		return nil, err
	}

	return &authv1.ListSessionsResponse{
		Sessions: sess,
	}, nil
}
