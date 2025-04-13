package postgres

import (
	"log/slog"


	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	db     *pgxpool.Pool
	logger *slog.Logger
}

func NewStorage(db *pgxpool.Pool, logger *slog.Logger) *Storage {
	return &Storage{db: db, logger: logger}
}


