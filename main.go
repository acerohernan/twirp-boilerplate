package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/acerohernan/twirp-boilerplate/pkg/server"
)

func main() {
	err := runServer()

	if err != nil {
		log.Fatal(err)
	}
}

func runServer() error {
	server, err := server.InitializeServer()

	if err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigChan
		fmt.Printf("exist requested, shutting down. signal: %v", sig)
		server.Stop(false)
	}()

	return server.Start()
}
