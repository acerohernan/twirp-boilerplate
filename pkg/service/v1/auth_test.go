package servicev1_test

import (
	"context"
	"testing"

	authv1 "github.com/acerohernan/twirp-boilerplate/core/auth/v1"
	"github.com/acerohernan/twirp-boilerplate/pkg/service/servicefakes"
	servicev1 "github.com/acerohernan/twirp-boilerplate/pkg/service/v1"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {
	t.Run("should throw an error if the email is invalid", func(t *testing.T) {
		svc := newTestAuthService()
		req := &authv1.CreateSessionRequest{
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
	val, _ := utils.NewProtoValidator()
	svc := servicev1.NewAuthService(storage, val)
	return &TestAuthService{
		AuthService: *svc,
		storage:     storage,
	}

}

type TestAuthService struct {
	servicev1.AuthService
	storage *servicefakes.FakeStorage
}
