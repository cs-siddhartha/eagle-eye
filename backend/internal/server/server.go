package server

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"

	"github.com/your-org/observability-tool/backend/internal/config"
)

func Run(ctx context.Context, config config.Config, logger *slog.Logger, handler http.Handler) error {
	httpServer := &http.Server{
		Addr:              config.Address,
		Handler:           handler,
		ReadHeaderTimeout: config.ReadHeaderTimeout,
		ReadTimeout:       config.ReadTimeout,
		WriteTimeout:      config.WriteTimeout,
		IdleTimeout:       config.IdleTimeout,
	}

	listener, err := net.Listen("tcp", config.Address)
	if err != nil {
		return err
	}

	serverErrors := make(chan error, 1)
	go func() {
		logger.Info("api server listening", "address", listener.Addr().String())
		serverErrors <- httpServer.Serve(listener)
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-ctx.Done():
		shutdownContext, cancel := context.WithTimeout(context.Background(), config.ShutdownTimeout)
		defer cancel()

		logger.Info("api server shutting down")
		if err := httpServer.Shutdown(shutdownContext); err != nil {
			return err
		}

		err := <-serverErrors
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	}
}
