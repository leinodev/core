package layers

import "github.com/leinodev/core/internal/core/context"

// Layer is interface that must be implemented by any layer in the core
// to provide an opportunity to switch the chain of calls.
// Layer1 -Execute(ctx)-> Layer2 -Execute(ctx)-> Layer3
// Layer1 -Execute(ctx)-> Layer3 -Execute(ctx)-> Layer2
type Layer interface {
	Execute(ctx context.Context) error
}
