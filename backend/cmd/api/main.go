package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/your-org/observability-tool/backend/internal/config"
	"github.com/your-org/observability-tool/backend/internal/httpapi"
	"github.com/your-org/observability-tool/backend/internal/server"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	appConfig, err := config.Load()
	if err != nil {
		logger.Error("failed to load configuration", "error", err)
		os.Exit(1)
	}

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	handler := httpapi.NewHandler(logger, appConfig.MaxRequestBytes)
	if err := server.Run(ctx, appConfig, logger, handler); err != nil {
		logger.Error("api server stopped unexpectedly", "error", err)
		os.Exit(1)
	}
}
