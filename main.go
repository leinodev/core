package main

import (
	"flag"
	"fmt"
	"github.com/leinodev/core/config"
	"github.com/leinodev/core/internal/application"
)

func main() {
	var cfgPath string
	flag.StringVar(&cfgPath, "c", "./config.yaml", "path to configuration file")
	flag.Parse()
	cfg, err := config.NewConfig(cfgPath)
	panicOnErr(err)

	app := application.New(cfg)

	if err := app.Configure(); err != nil {
		panicOnErr(fmt.Errorf("cannot configure app: %w", err))
	}

	if err := app.Run(); err != nil {
		panicOnErr(fmt.Errorf("cannot run app: %w", err))
	}
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
