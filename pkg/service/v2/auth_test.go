package servicev2_test

import (
	"context"
	"testing"

	authv2 "github.com/acerohernan/twirp-boilerplate/core/auth/v2"
	"github.com/acerohernan/twirp-boilerplate/pkg/service/servicefakes"
	servicev2 "github.com/acerohernan/twirp-boilerplate/pkg/service/v2"
	"github.com/acerohernan/twirp-boilerplate/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestCreateSession(t *testing.T) {
	t.Run("should throw an error if the phone is invalid", func(t *testing.T) {
		svc := newTestAuthService()
		req := &authv2.CreateSessionRequest{
			PhoneNumber: "invalid_number",
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
	svc := servicev2.NewAuthService(storage, val)
	return &TestAuthService{
		AuthService: *svc,
		storage:     storage,
	}

}

type TestAuthService struct {
	servicev2.AuthService
	storage *servicefakes.FakeStorage
}
