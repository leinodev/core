package application

import (
	"apifika-core/core/corecontext"
	"apifika-core/core/corecontext/request"
	"context"
	"net/http"
	"time"
)

type Application struct {
	ctx context.Context
}

func New(ctx context.Context) (*Application, error) {
	return &Application{
		ctx: ctx,
	}, nil
}

func (h *Application) Run() {

}

func (h *Application) httpHandler(w http.ResponseWriter, r *http.Request) {
	parent, cancell := context.WithTimeout(h.ctx, time.Minute)
	defer cancell()
	ctx := corecontext.CreateContext(parent)
	ctx.SetRequestContext(&request.Context{
		HttpContext: &request.HttpContext{
			R: r,
			W: w,
		},
	})
	//next(ctx)

}
