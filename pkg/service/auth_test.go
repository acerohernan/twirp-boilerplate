package service_test

import (
	"context"
	"testing"

	"github.com/acerohernan/twirp-boilerplate/core/twirp"
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/acerohernan/twirp-boilerplate/pkg/service/servicefakes"
	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {
	t.Run("should throw an error if the email is invalid", func(t *testing.T) {
		svc := newTestAuthService()
		req := &twirp.CreateSessionRequest{
			Email: "invalid_email",
		}
		res, err := svc.CreateSession(context.Background(), req)
		require.Error(t, err)
		require.Nil(t, res)
		require.Equal(t, svc.storage.StoreSessionCallCount(), 0)
	})
}

func newTestAuthService() *TestAuthService {
	storage := &servicefakes.FakeStorage{}
	svc := service.NewAuthService(storage)
	return &TestAuthService{
		AuthService: *svc,
		storage:     storage,
	}

}

type TestAuthService struct {
	service.AuthService
	storage *servicefakes.FakeStorage
}
