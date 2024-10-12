package main

import (
	"apifika-core/application"
	"context"
	"os"
	"os/signal"
)

func main() {

	ctx, cancell := context.WithCancel(context.Background())

	//Create logger

	//Load config

	app, err := application.New(ctx)
	if err != nil {
		//logger.Fatal
	}

	app.Run()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill)
	<-ch
	cancell()
}
