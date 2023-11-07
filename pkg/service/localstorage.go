package service

import (
	"context"
	"sync"

	"github.com/acerohernan/twirp-boilerplate/core"
)

type LocalStorage struct {
	sessions map[string]*core.Session
	mu       sync.Mutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		sessions: make(map[string]*core.Session),
	}
}

func (s *LocalStorage) StoreSession(_ context.Context, session *core.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.Id] = session

	return nil
}
