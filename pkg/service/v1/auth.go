package servicev1

import (
	"context"

	"github.com/acerohernan/twirp-boilerplate/core"
	authv1 "github.com/acerohernan/twirp-boilerplate/core/auth/v1"
	"github.com/acerohernan/twirp-boilerplate/pkg/config/logger"
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
)

type AuthService struct {
	storage service.Storage
	val     *utils.ProtoValidator
}

func NewAuthService(storage service.Storage, val *utils.ProtoValidator) *AuthService {
	return &AuthService{
		storage: storage,
		val:     val,
	}
}

func (s *AuthService) CreateSession(ctx context.Context, req *authv1.CreateSessionRequest) (*authv1.CreateSessionResponse, error) {
	err := s.val.Validate(req)

	if err != nil {
		logger.Errorw("proto validation error", err)
		return nil, err
	}

	sess := &core.Session{
		Id: utils.NewGuid(utils.SessionPrefix),
	}

	err = s.storage.StoreSession(ctx, sess)

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
