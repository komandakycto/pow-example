package main

import (
	"context"
	"errors"
	"os/signal"
	"syscall"

	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	prover "github.com/komandakycto/pow-example/internal/client"
	"github.com/komandakycto/pow-example/internal/service/solver"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGHUP, syscall.SIGTERM)
	defer cancel()

	logger, _ := zap.NewProduction()
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {
			logger.Fatal("failed to sync logger", zap.Error(err))
		}
	}(logger)

	var cfg Config
	if _, err := flags.Parse(&cfg); err != nil {
		var fErr *flags.Error
		if errors.As(err, &fErr) && errors.Is(fErr.Type, flags.ErrHelp) {
			logger.Fatal("failed to parse flags")

			return
		}
	}

	client := prover.New(cfg.ServerAddr, logger, solver.New(cfg.Difficulty))
	client.Run(ctx)
}
