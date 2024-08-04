package main

import (
	"context"
	"errors"
	"os/signal"
	"syscall"

	"github.com/jessevdk/go-flags"
	"go.uber.org/zap"

	verifier "github.com/komandakycto/pow-example/internal/server"
	"github.com/komandakycto/pow-example/internal/service/hashcash"
	"github.com/komandakycto/pow-example/internal/service/quotes"
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

	quoteService, err := quotes.New(cfg.QuotesPath)
	if err != nil {
		logger.Fatal("failed to create quote service", zap.Error(err))

		return
	}

	hashcashService := hashcash.New(cfg.Difficulty)

	server := verifier.NewPOWServer(
		cfg.Port,
		logger,
		quoteService,
		hashcashService,
	)
	err = server.Start(ctx)
	if err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}

	defer server.Stop()
}
