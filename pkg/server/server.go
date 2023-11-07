package server

import "github.com/acerohernan/twirp-boilerplate/pkg/service"

type Server struct {
	storage service.Storage
}

func NewServer(storage service.Storage) *Server {
	return &Server{
		storage: storage,
	}
}

func (s *Server) Start() {}

func (s *Server) Stop(force bool) {}
