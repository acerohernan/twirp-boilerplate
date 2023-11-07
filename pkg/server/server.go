package server

import (
	"context"
	"errors"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/acerohernan/twirp-boilerplate/core/twirp"
	"github.com/acerohernan/twirp-boilerplate/pkg/service"
	"github.com/rs/cors"
	"github.com/urfave/negroni/v3"
)

type Server struct {
	authService *service.AuthService
	httpServer  *http.Server
	running     atomic.Bool
	doneChan    chan struct{}
	closedChan  chan struct{}
}

func NewServer(authService *service.AuthService) *Server {
	server := &Server{
		authService: authService,
	}

	authServer := twirp.NewAuthServiceServer(authService)

	middleares := []negroni.Handler{
		// always first
		negroni.NewRecovery(),

		// CORS is allowed, we rely on token auth
		cors.New(cors.Options{
			AllowOriginFunc: func(origin string) bool {
				return true
			},
			AllowedHeaders: []string{"*"},
			// allow preflight to be cached for a day
			MaxAge: 86400,
		}),
	}

	// setup router
	mux := http.NewServeMux()

	mux.Handle(authServer.PathPrefix(), authServer)

	// setup middlewares
	n := negroni.New()

	for _, m := range middleares {
		n.Use(m)
	}

	n.UseHandler(mux)

	server.httpServer = &http.Server{
		Handler: n,
	}

	return server
}

func (s *Server) Start() error {
	if s.running.Load() {
		return errors.New("already running")
	}

	listener, err := net.Listen("tcp", ":3001")

	if err != nil {
		return err
	}

	go s.httpServer.Serve(listener)

	s.running.Store(true)

	<-s.doneChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_ = s.httpServer.Shutdown(ctx)

	close(s.closedChan)
	return nil
}

func (s *Server) Stop(force bool) {
	if !s.running.Swap(false) {
		return
	}

	close(s.doneChan)

	// wait for fully closed
	<-s.closedChan
}
