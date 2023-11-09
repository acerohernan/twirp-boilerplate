package servicev1

import (
	"context"
	"net/mail"

	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
	"github.com/acerohernan/twirp-boilerplate/rpc"
	"github.com/acerohernan/twirp-boilerplate/rpc/twirp/v1"
)

type AuthService struct {
	storage service.Storage
}

func NewAuthService(storage service.Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) CreateSession(ctx context.Context, req *twirp.CreateSessionRequest) (*twirp.CreateSessionResponse, error) {
	if req.Email == "" {
		return nil, service.ErrBadRequest
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return nil, service.ErrBadRequest
	}

	sess := &rpc.Session{
		Id: utils.NewGuid(utils.SessionPrefix),
	}

	err := s.storage.StoreSession(ctx, sess)

	if err != nil {
		return nil, err
	}

	return &twirp.CreateSessionResponse{
		Session: sess,
	}, nil
}

func (s *AuthService) ListSessions(ctx context.Context, req *twirp.ListSessionsRequest) (*twirp.ListSessionsResponse, error) {
	sess, err := s.storage.ListSessions(ctx)

	if err != nil {
		return nil, err
	}

	return &twirp.ListSessionsResponse{
		Sessions: sess,
	}, nil
}
