package service

import (
	"context"
	"net/mail"

	"github.com/acerohernan/twirp-boilerplate/core"
	"github.com/acerohernan/twirp-boilerplate/core/twirp"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
)

type AuthService struct {
	storage Storage
}

func NewAuthService(storage Storage) *AuthService {
	return &AuthService{
		storage: storage,
	}
}

func (s *AuthService) CreateSession(ctx context.Context, req *twirp.CreateSessionRequest) (*twirp.CreateSessionResponse, error) {
	if req.Email == "" {
		return nil, ErrBadRequest
	}

	if _, err := mail.ParseAddress(req.Email); err != nil {
		return nil, ErrBadRequest
	}

	sess := &core.Session{
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
