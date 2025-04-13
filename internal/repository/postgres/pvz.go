package postgres

import (
	"context"
	"database/sql"
	"log/slog"
	"time"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
)

const (
    queryGetAllPvz = `SELECT id, registration_date, city FROM pvzs`
	queryPvzExists = `SELECT EXISTS(SELECT 1 FROM pvzs WHERE id = $1 AND deleted_at IS NULL)`
	queryCreatePvz = `INSERT INTO pvzs (id, registration_date, city) VALUES ($1, $2, $3)
	           				RETURNING id, registration_date, city`
	queryCheckCityExists = `SELECT EXISTS(SELECT 1 FROM cities WHERE name = $1 AND deleted_at IS NULL)`
    queryGetPvzList = `
    SELECT 
        p.id,
        p.registration_date,
        p.city,
        r.id as reception_id,
        r.date_time as reception_date,
        r.pvz_id as reception_pvz_id,
        r.status as reception_status,
        array_remove(array_agg(CASE WHEN pr.deleted_at IS NULL THEN pr.id END ORDER BY pr.date_time DESC), NULL) as product_ids,
        array_remove(array_agg(CASE WHEN pr.deleted_at IS NULL THEN pr.type END ORDER BY pr.date_time DESC), NULL) as product_types,
        array_remove(array_agg(CASE WHEN pr.deleted_at IS NULL THEN pr.date_time END ORDER BY pr.date_time DESC), NULL) as product_dates
    FROM pvzs p
    LEFT JOIN receptions r ON r.pvz_id = p.id AND r.deleted_at IS NULL
        AND ($1::timestamp IS NULL OR r.date_time >= $1)
        AND ($2::timestamp IS NULL OR r.date_time <= $2)
    LEFT JOIN products pr ON pr.reception_id = r.id
    WHERE p.deleted_at IS NULL
    GROUP BY 
        p.id, 
        p.registration_date, 
        p.city,
        r.id,
        r.date_time,
        r.status
    ORDER BY r.date_time DESC NULLS LAST
    LIMIT $3 OFFSET $4
`
)


func (s *Storage) GetAllPvz(ctx context.Context) ([]entity.Pvz, error) {
	rows, err := s.db.Query(ctx, queryGetAllPvz)
	if err != nil {
		s.logger.Error("failed to get all pvz",
			slog.String("method", "repository.GetAllPvz"),
			slog.String("error", err.Error()))
		return nil, errorsx.ErrInternal
	}
	defer rows.Close()

	pvzList := []entity.Pvz{}
	for rows.Next() {
		var pvz entity.Pvz
		err := rows.Scan(&pvz.Id, &pvz.RegistrationDate, &pvz.City)
		if err != nil {
			s.logger.Error("failed to scan pvz",
				slog.String("method", "repository.GetAllPvz"),
				slog.String("error", err.Error()))
			return nil, errorsx.ErrInternal
		}
		pvzList = append(pvzList, pvz)
	}

	return pvzList, nil
}

func (s *Storage) PvzExists(ctx context.Context, id uuid.UUID) (bool, error) {
	var exists bool
	err := s.db.QueryRow(ctx, queryPvzExists, id).Scan(&exists)
	if err != nil {
		s.logger.Error("failed to check if pvz exists",
			slog.String("method", "repository.PvzExists"),
			slog.String("error", err.Error()))
		return false, errorsx.ErrInternal
	}

	return exists, nil
}

func (s *Storage) CreatePvz(ctx context.Context, pvz *entity.Pvz) (*entity.Pvz, error) {
	var pvzDBInfo entity.Pvz
	err := s.db.QueryRow(ctx,
		queryCreatePvz,
		pvz.Id,
		pvz.RegistrationDate,
		pvz.City).
		Scan(&pvzDBInfo.Id,
			&pvzDBInfo.RegistrationDate,
			&pvzDBInfo.City)
	if err != nil {
		s.logger.Error("failed to create pvz",
			slog.String("method", "repository.CreatePvz"),
			slog.String("error", err.Error()))
		return nil, errorsx.ErrInternal
	}
	return &pvzDBInfo, nil
}

func (s *Storage) CityExists(ctx context.Context, city string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(ctx, queryCheckCityExists, city).Scan(&exists)
	if err != nil {
		s.logger.Error("failed to check if city exists",
			slog.String("method", "repository.CityExists"),
			slog.String("error", err.Error()))
		return false, errorsx.ErrInternal
	}

	return exists, nil
}

func (s *Storage) GetPvzList(ctx context.Context, filter entity.Filter) ([]entity.PvzInfo, error) {
    rows, err := s.db.Query(ctx, queryGetPvzList,
        filter.StartDate,
        filter.EndDate,
        filter.Limit,
        filter.Offset)
    if err != nil {
        s.logger.Error("failed to get pvz list",
            slog.String("method", "storage.GetPvzList"),
            slog.String("error", err.Error()))
        return nil, errorsx.ErrInternal
    }
    defer rows.Close()

    // Создаем map для группировки приёмок по PVZ
    pvzMap := make(map[uuid.UUID]*entity.PvzInfo)

    for rows.Next() {
        var pvzID uuid.UUID
        var registrationDate time.Time
        var city string
        var receptionID, receptionStatus sql.NullString
        var receptionDate sql.NullTime
        var receptionPvzID sql.NullString
        var productIDs []uuid.UUID
        var productTypes []string
        var productDates []time.Time

        err := rows.Scan(
            &pvzID,
            &registrationDate,
            &city,
            &receptionID,
            &receptionDate,
            &receptionPvzID,
            &receptionStatus,
            &productIDs,
            &productTypes,
            &productDates,
        )
        if err != nil {
            s.logger.Error("failed to scan pvz list row",
                slog.String("method", "storage.GetPvzList"),
                slog.String("error", err.Error()))
            return nil, errorsx.ErrInternal
        }

        // Получаем или создаем PVZ
        info, exists := pvzMap[pvzID]
        if !exists {
            info = &entity.PvzInfo{}
            info.Pvz.Id = pvzID
            info.Pvz.RegistrationDate = registrationDate
            info.Pvz.City = city
            pvzMap[pvzID] = info
        }

        // Если есть приёмка, добавляем её и её продукты
        if receptionID.Valid {
            reception := entity.ReceptionWithProducts{
                Reception: entity.Reception{
                    Id:       uuid.MustParse(receptionID.String),
                    DateTime: receptionDate.Time,
                    PvzID: uuid.MustParse(receptionPvzID.String),
                    Status:   receptionStatus.String,
                },
            }

            // Добавляем продукты
            reception.Products = make([]entity.Product, len(productIDs))
            for i := range productIDs {
                reception.Products[i] = entity.Product{
                    ID:          productIDs[i],
                    Type:        productTypes[i],
                    DateTime:    productDates[i],
                    ReceptionID: reception.Reception.Id,
                }
            }

            info.Products = append(info.Products, reception)
        }
    }

    // Преобразуем map в slice
    result := make([]entity.PvzInfo, 0, len(pvzMap))
    for _, info := range pvzMap {
        result = append(result, *info)
    }

    return result, nil
}
