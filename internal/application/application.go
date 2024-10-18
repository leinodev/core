package application

import (
	"fmt"
	"github.com/leinodev/core/config"
	"net/http"
)

type Application struct {
	cfg *config.CoreConfig
}

func New(cfg *config.CoreConfig) *Application {
	return &Application{
		cfg: cfg,
	}
}

func (app *Application) Configure() error {
	return nil
}

func (app *Application) Run() error {

	fmt.Println(app.cfg)

	return nil
}

func (app *Application) httpHandler(w http.ResponseWriter, r *http.Request) {
	//parent, cancell := context.WithTimeout(h.ctx, time.Minute)
	//defer cancell()
	//ctx := corecontext.CreateContext(parent)
	//ctx.SetRequestContext(&corerequest.Context{
	//	HttpContext: &corerequest.HttpContext{
	//		R: r,
	//		W: w,
	//	},
	//})
	//next(ctx)

}
