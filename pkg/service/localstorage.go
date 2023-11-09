package service

import (
	"context"
	"sync"

	"github.com/acerohernan/twirp-boilerplate/rpc"
)

type LocalStorage struct {
	sessions map[string]*rpc.Session
	mu       sync.RWMutex
}

func NewLocalStorage() *LocalStorage {
	return &LocalStorage{
		sessions: make(map[string]*rpc.Session),
	}
}

func (s *LocalStorage) StoreSession(_ context.Context, session *rpc.Session) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sessions[session.Id] = session

	return nil
}

func (s *LocalStorage) ListSessions(ctx context.Context) ([]*rpc.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sessions := make([]*rpc.Session, 0)

	for _, s := range s.sessions {
		sessions = append(sessions, s)
	}

	return sessions, nil
}
