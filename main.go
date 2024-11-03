package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/leinodev/core/application"
	"github.com/leinodev/core/config"
)

var Version string

func main() {
	launchConfig, err := config.LoadLaunchConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancelFunc := context.WithCancel(context.Background())

	app, err := application.New(ctx, launchConfig)
	if err != nil {
		log.Fatal(err)
	}

	app.Start()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// Wait for killcode and stop application
	<-sigCh
	cancelFunc()
}
