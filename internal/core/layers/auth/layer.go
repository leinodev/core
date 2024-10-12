package auth

import (
	"github.com/leinodev/core/internal/core/context"
	"github.com/leinodev/core/internal/core/context/request"
)

type Layer struct {
}

func (l *Layer) Execute(ctx context.Context) error {
	// TODO not implemented yet
	switch ctx.GetRequestContext().GetTransportType() {
	case request.HTTP_TRANSPORT_TYPE:
		return l.checkHttpAuth(ctx)
	}
	return nil
}

func (l *Layer) checkHttpAuth(ctx context.Context) error {
	// TODO not implemented yet
	return nil
}
