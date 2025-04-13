package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"github.com/ZhdanovichVlad/service-podof/internal/config"
	"github.com/ZhdanovichVlad/service-podof/internal/metrics"
	"github.com/ZhdanovichVlad/service-podof/internal/controller/htttp/middleware"
	"github.com/ZhdanovichVlad/service-podof/internal/repository/postgres"
	"github.com/ZhdanovichVlad/service-podof/internal/app/router"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/ZhdanovichVlad/service-podof/pkg/jwttoken"
	"github.com/ZhdanovichVlad/service-podof/internal/controller/htttp/handler"


	"github.com/jackc/pgx/v5/pgxpool"
)
func main() {

	config, err := config.NewConfig()
	if err != nil {	
		log.Fatal(err)
	}

	db, err := pgxpool.New(context.Background(), config.DatabaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	repo := postgres.NewStorage(db, logger)

	tokenGenerator := jwttoken.NewJwtTokenGenerator()
	service := service.NewService(repo, logger, tokenGenerator)
	handler := handler.NewHandler(service, logger)

	metrics := metrics.NewMetrics()
	middleware := middleware.NewMiddleware(tokenGenerator, metrics)


	addr := config.Host + ":" + config.Port
	router := router.NewRouter(handler, middleware, logger, addr)
	router.SetupRoutes()

	router.Run()

}

