package servicev1_test

import (
	"context"
	"testing"

	"github.com/acerohernan/twirp-boilerplate/pkg/service/servicefakes"
	servicev1 "github.com/acerohernan/twirp-boilerplate/pkg/service/v1"
	"github.com/acerohernan/twirp-boilerplate/rpc/twirp/v1"
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
	svc := servicev1.NewAuthService(storage)
	return &TestAuthService{
		AuthService: *svc,
		storage:     storage,
	}

}

type TestAuthService struct {
	servicev1.AuthService
	storage *servicefakes.FakeStorage
}