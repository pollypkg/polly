// Package coord handles coordination and graceful shutdown using a WaitGroup
// embedded into a context.Context
package coord

import (
	"context"
	"sync"
)

// Path of the coordination WaitGroup
const Path = "coord-wg"

// WithCancel acts like context.WithCancel, but also sets up waiting using a
// WaitGroup. Calling the returned cancel will cancel the context and wait for
// any jobs registered in the WaitGroup.
func WithCancel(ctx context.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(ctx)
	var wg sync.WaitGroup
	ctx = context.WithValue(ctx, Path, &wg)

	cancelFn := func() {
		cancel()
		wg.Wait()
	}

	return ctx, cancelFn
}

// WaitGroup extracts the WaitGroup from the context, if it exists
func WaitGroup(ctx context.Context) *sync.WaitGroup {
	ptr := ctx.Value(Path)
	if ptr == nil {
		return nil
	}

	wg, ok := ptr.(*sync.WaitGroup)
	if !ok {
		return nil
	}

	return wg
}

// Add n to the WaitGroup of the context if it exists
func Add(ctx context.Context, n int) {
	if wg := WaitGroup(ctx); wg != nil {
		wg.Add(n)
	}
}

// Done calls wg.Done() on the WaitGroup of the context if it exists
func Done(ctx context.Context) {
	if wg := WaitGroup(ctx); wg != nil {
		wg.Done()
	}
}

// Finally ensures f is run (and waited for) when the context is canceled.
func Finally(ctx context.Context, f func()) {
	Add(ctx, 1)

	go func() {
		<-ctx.Done()
		f()
		Done(ctx)
	}()
}
