package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/ZhdanovichVlad/service-podof/internal/app/grpc"
	"github.com/ZhdanovichVlad/service-podof/internal/config"
	"github.com/ZhdanovichVlad/service-podof/internal/repository/postgres"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))


	cfg, err := config.NewConfig()
	if err != nil {
		logger.Error("failed to create config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	db, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	defer db.Close()
	if err != nil {
		logger.Error("failed to create db", slog.String("error", err.Error()))
		os.Exit(1)
	}
	
	err = db.Ping(context.Background())
	if err != nil {
		logger.Error("failed to ping db", slog.String("error", err.Error()))
		os.Exit(1)
	}

	repo := postgres.NewStorage(db, logger)

	svc := service.NewService(repo, logger, nil)

    grpcServer, err := app.NewGRPCServer(svc, logger, cfg)
    if err != nil {
        logger.Error("failed to create gRPC server", slog.String("error", err.Error()))
        os.Exit(1)
    }


    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    go func() {
        sigCh := make(chan os.Signal, 1)
        signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
        <-sigCh
        cancel()
    }()


    go func() {
        if err := grpcServer.Start(); err != nil {
            logger.Error("failed to start gRPC server", slog.String("error", err.Error()))
            cancel()
        }
    }()

  
    <-ctx.Done()
    grpcServer.Stop()
}