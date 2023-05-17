package app

import (
	"context"
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/xid"
	"golang.org/x/sync/errgroup"
)

// App global app
type App struct {
	opts   options
	ctx    context.Context
	cancel func()
}

// New create a app globally
func New(opts ...Option) *App {
	o := options{id: xid.New().String()}
	for _, opt := range opts {
		opt(&o)
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
	}
}

// Run start app
func (a *App) Run() error {
	eg, ctx := errgroup.WithContext(a.ctx)

	// start server
	for _, srv := range a.opts.servers {
		srv := srv
		eg.Go(func() error {
			// wait for stop signal
			<-ctx.Done()
			return srv.Stop(ctx)
		})
		eg.Go(func() error {
			return srv.Start(ctx)
		})
	}

	// watch signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	eg.Go(func() error {
		defer log.Println("signal defer")
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case s := <-quit:
				log.Printf("Server receive a quit signal: %s", s.String())
				if err := a.Stop(); err != nil {
					log.Printf("failed to stop app, err: %v", err)
					return err
				}
			}
		}
	})
	if err := eg.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}

	return nil
}

// Stop stops the application gracefully.
func (a *App) Stop() error {
	if a.cancel != nil {
		a.cancel()
	}
	return nil
}
