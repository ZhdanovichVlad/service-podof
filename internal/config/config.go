package config

import (
	"errors"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseURL string
	Port string
	Host string
	GrpcPort string
}

func NewConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	databaseURL := os.Getenv("PG_DSN")
	if databaseURL == "" {
		return nil, errors.New("PG_DSN is not set")
	}
	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("PORT is not set")
	}
	host := os.Getenv("HOST")
	if host == "" {
		return nil, errors.New("HOST is not set")
	}
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		return nil, errors.New("GRPC_PORT is not set")
	}
	return &Config{
		DatabaseURL: databaseURL,
		Port: port,
		Host: host,
		GrpcPort: grpcPort,
	}, nil
}
