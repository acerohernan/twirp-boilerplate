package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acerohernan/twirp-boilerplate/pkg/config"
	"github.com/acerohernan/twirp-boilerplate/pkg/config/logger"
	"github.com/acerohernan/twirp-boilerplate/pkg/server"
)

func main() {
	err := runServer()

	if err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	conf := config.NewConfig()

	err := logger.InitLogger(conf.Logger)

	if err != nil {
		return err
	}

	server, err := server.InitializeServer(conf)

	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		logger.Infow("exist requested, shutting down...", "signal", sig)
		server.Stop(false)
	}()

	return server.Start()
}
