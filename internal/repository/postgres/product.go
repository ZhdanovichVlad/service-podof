package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

const queryCreateProduct = `
    INSERT INTO products (type, reception_id)
    SELECT 
        $1,  -- type
        r.id as reception_id
    FROM receptions r
    WHERE r.pvz_id = $2  -- pvz_id
    AND r.deleted_at IS NULL
    ORDER BY r.created_at DESC
    LIMIT 1
    RETURNING id, type, date_time, reception_id`


    const queryDeleteLastProduct = `
    UPDATE products p
    SET deleted_at = CURRENT_TIMESTAMP
    WHERE p.id = (
        SELECT p2.id
        FROM products p2
        JOIN receptions r ON r.id = p2.reception_id
        WHERE r.pvz_id = $1
          AND p2.deleted_at IS NULL
          AND r.id = (
              SELECT r2.id
              FROM receptions r2
              WHERE r2.pvz_id = $1
              AND r2.deleted_at IS NULL
              ORDER BY r2.created_at DESC
              LIMIT 1
          )
        ORDER BY p2.created_at DESC
        LIMIT 1
    )`



func (s *Storage) CreateProduct(ctx context.Context, product *entity.Product, pvzID uuid.UUID) (*entity.Product, error) {
    var productDBInfo entity.Product
    err := s.db.QueryRow(ctx, queryCreateProduct, 
        product.Type, 
        pvzID).Scan(
            &productDBInfo.ID,
            &productDBInfo.Type,
            &productDBInfo.DateTime,
            &productDBInfo.ReceptionID)
    
    if err != nil {
        if errors.Is(err, pgx.ErrNoRows) {
            return nil, errorsx.ErrInternal
        }
        s.logger.Error("failed to create product",
            slog.String("method", "repository.CreateProduct"),
            slog.String("error", err.Error()))
        return nil, errorsx.ErrInternal
    }
    
    return &productDBInfo, nil
}

func (s *Storage) DeleteLastProduct(ctx context.Context, pvzID uuid.UUID) error {
	_, err := s.db.Exec(ctx, queryDeleteLastProduct, pvzID)	
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return errorsx.ErrProductNotFound
		}
		s.logger.Error("failed to delete last product",
			slog.String("method", "repository.DeleteLastProduct"),
			slog.String("error", err.Error()))
		return errorsx.ErrInternal
	}

	return nil
}
