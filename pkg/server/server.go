package server

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/acerohernan/twirp-boilerplate/pkg/config"
	"github.com/acerohernan/twirp-boilerplate/pkg/config/logger"
	servicev1 "github.com/acerohernan/twirp-boilerplate/pkg/service/v1"
	servicev2 "github.com/acerohernan/twirp-boilerplate/pkg/service/v2"
	twirpv1 "github.com/acerohernan/twirp-boilerplate/rpc/twirp/v1"
	twirpv2 "github.com/acerohernan/twirp-boilerplate/rpc/twirp/v2"
	"github.com/rs/cors"
	"github.com/urfave/negroni/v3"
)

type Server struct {
	conf          *config.Config
	authServiceV1 *servicev1.AuthService
	authServiceV2 *servicev2.AuthService
	httpServer    *http.Server
	running       atomic.Bool
	doneChan      chan struct{}
	closedChan    chan struct{}
}

func NewServer(conf *config.Config, authServiceV1 *servicev1.AuthService,
	authServiceV2 *servicev2.AuthService) *Server {
	server := &Server{
		conf:          conf,
		authServiceV1: authServiceV1,
		authServiceV2: authServiceV2,
		doneChan:      make(chan struct{}),
		closedChan:    make(chan struct{}),
	}

	authServerV1 := twirpv1.NewAuthServiceServer(authServiceV1)
	authServerV2 := twirpv2.NewAuthServiceServer(authServiceV2)

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

	mux.Handle(authServerV1.PathPrefix(), authServerV1)
	mux.Handle(authServerV2.PathPrefix(), authServerV2)

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
	logger.Infow("starting server...")

	if s.running.Load() {
		return errors.New("already running")
	}

	listener, err := net.Listen("tcp", fmt.Sprint(":", s.conf.Port))

	if err != nil {
		return err
	}

	go s.httpServer.Serve(listener)

	s.running.Store(true)

	logger.Infow("server started successfully")

	<-s.doneChan

	return nil
}

func (s *Server) Stop(force bool) {
	if !s.running.Swap(false) {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	_ = s.httpServer.Shutdown(ctx)

	close(s.doneChan)
}
