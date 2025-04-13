package service

import (
	"context"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
)

func (s *Service) CreateProduct(ctx context.Context, product *entity.Product, pvzID uuid.UUID) (*entity.Product, error) {

	if pvzID == uuid.Nil {
		return nil, errorsx.ErrEmptyField
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}
	
	exists, err := s.repo.PvzExists(ctx, pvzID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errorsx.ErrPVZNotFound
	}

	lastReceptionStatus, err := s.repo.LastReceptionStatus(ctx, pvzID)
	if err != nil {
		return nil, err
	}

	if lastReceptionStatus == entity.ReceptionStatusCompleted {
		return nil, errorsx.ErrReceptionIsClosed
	}

	product, err = s.repo.CreateProduct(ctx, product, pvzID)
	if err != nil {
		return nil, err
	}

	return product, nil
}	



func (s *Service) DeleteLastProduct(ctx context.Context, pvzID uuid.UUID) (error) {

	if pvzID == uuid.Nil {
		return errorsx.ErrEmptyField
	}
	
	exists, err := s.repo.PvzExists(ctx, pvzID)
	if err != nil {
		return err
	}
	if !exists {
		return errorsx.ErrPVZNotFound
	}

	lastReceptionStatus, err := s.repo.LastReceptionStatus(ctx, pvzID)
	if err != nil {
		return err
	}
	
	if lastReceptionStatus == entity.ReceptionStatusCompleted {
		return errorsx.ErrReceptionIsClosed
	}

	err = s.repo.DeleteLastProduct(ctx, pvzID)
	if err != nil {
		return err
	}

	return nil
}	
