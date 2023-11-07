package service

import (
	"context"
	"sync"

	"github.com/acerohernan/twirp-boilerplate/pkg/auth"
)

type LocalStorage struct {
	sessions map[string]*auth.Session
	mu       sync.Mutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		sessions: make(map[string]*auth.Session),
	}
}

func (s *LocalStorage) StoreSession(_ context.Context, session *auth.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.Id] = session

	return nil
}
