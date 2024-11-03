package application

import (
	"context"

	"github.com/leinodev/core/config"
)

type Application struct {
	ctx context.Context
	cfg *config.LaunchConfig
}

func New(ctx context.Context, cfg *config.LaunchConfig) (*Application, error) {
	h := &Application{
		ctx: ctx,
		cfg: cfg,
	}
	return h, nil
}

func (h *Application) Start() {

}
