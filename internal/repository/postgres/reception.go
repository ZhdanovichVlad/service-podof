package postgres

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const (
	queryCreateReception        = `INSERT INTO receptions (pvz_id, status)  VALUES ($1, $2) 
	                               RETURNING id, date_time, pvz_id, status`
    queryGetLastReceptionStatus = `
								   SELECT status 
								   FROM receptions 
								   WHERE deleted_at IS NULL 
								   AND pvz_id = $1
								   ORDER BY date_time DESC 
								   LIMIT 1`
	queryCloseLastReception     = `UPDATE receptions SET status = $2, date_time = $3 
	                               WHERE pvz_id = $1 AND created_at = (
		                           SELECT created_at FROM receptions 
		                           WHERE pvz_id = $1 
		                           ORDER BY created_at DESC 
		                           LIMIT 1 ) AND deleted_at IS NULL
	                               RETURNING id, date_time, pvz_id, status`
)

func (s *Storage) CreateReception(ctx context.Context, reception *entity.Reception) (*entity.Reception, error) {
	var receptionDBInfo entity.Reception
	err := s.db.QueryRow(ctx, 
		                 queryCreateReception, 
						 reception.PvzID, 
						 reception.Status).
				 	     Scan(&receptionDBInfo.Id, 
						 	  &receptionDBInfo.DateTime, 
							  &receptionDBInfo.PvzID, 
						      &receptionDBInfo.Status)
	if err != nil {
		s.logger.Error("failed to create reception",
		slog.String("method", "repository.CreateReception"),
		slog.String("error", err.Error()))
		return nil, errorsx.ErrInternal
	}
	return &receptionDBInfo, nil
}

func (s *Storage) LastReceptionStatus(ctx context.Context, pvzUUID uuid.UUID) (string, error) {
	var status string
	err := s.db.QueryRow(ctx, queryGetLastReceptionStatus, pvzUUID).Scan(&status)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errorsx.ErrReceptionNotFound
		}
		s.logger.Error("failed to get last reception status",
		slog.String("method", "repository.LastReceptionStatus"),
		slog.String("error", err.Error()))
		return "", errorsx.ErrInternal
	}

	return status, nil
}

func (s *Storage) CloseReception(ctx context.Context, pvzUUID uuid.UUID) (*entity.Reception, error) {
	var receptionDBInfo entity.Reception
	err := s.db.QueryRow(ctx,
		                queryCloseLastReception,
						pvzUUID,
						entity.ReceptionStatusCompleted,
						time.Now()).
						Scan(&receptionDBInfo.Id, 
						     &receptionDBInfo.DateTime,
					      	 &receptionDBInfo.PvzID,
						     &receptionDBInfo.Status)
	if err != nil {
		s.logger.Error("failed to close reception",
		slog.String("method", "repository.CloseReception"),
		slog.String("error", err.Error()))
		return nil, errorsx.ErrInternal
	}
	return &receptionDBInfo, nil
}