package context

import (
	"context"
	"github.com/leinodev/core/internal/core/context/auth"
	"github.com/leinodev/core/internal/core/context/request"
	"github.com/leinodev/core/pkg/utils"
	"runtime"
	"time"

	"github.com/google/uuid"
)

type Context interface {
	context.Context

	// Returns unique and constant id for this context
	GetRayID() string

	// Returns timestamp when this context been created
	CreatedAt() time.Time

	// Returns creation frame
	GetCreationTrace() runtime.Frame

	// Get request context
	GetRequestContext() *request.Context
	// Set request context
	SetRequestContext(c *request.Context)

	// Get auth context
	GetAuthContext() *auth.Context
	// Set auth context
	SetAuthContext(c *auth.Context)
}

type contextImpl struct {
	context.Context
	rayID       string
	createdFrom runtime.Frame
	createdAt   time.Time

	RequestContext *request.Context
	AuthContext    *auth.Context
}

func CreateContext(parent context.Context) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	ctx := &contextImpl{
		Context:     parent,
		rayID:       uuid.NewString(),
		createdFrom: utils.GetFrame(1),
		createdAt:   time.Now(),
	}
	return ctx
}

func (h *contextImpl) GetRayID() string {
	return h.rayID
}

func (h *contextImpl) CreatedAt() time.Time {
	return h.createdAt
}

func (h *contextImpl) GetCreationTrace() runtime.Frame {
	return h.createdFrom
}

func (h *contextImpl) GetRequestContext() *request.Context {
	return h.RequestContext
}
func (h *contextImpl) SetRequestContext(c *request.Context) {
	h.RequestContext = c
}

func (h *contextImpl) GetAuthContext() *auth.Context {
	return h.AuthContext
}
func (h *contextImpl) SetAuthContext(c *auth.Context) {
	h.AuthContext = c
}
