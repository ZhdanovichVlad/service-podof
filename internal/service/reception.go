package service

import (
	"context"
	"errors"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
)


func (s *Service) CreateReception(ctx context.Context, reception *entity.Reception) (*entity.Reception, error) {
	
	if err := reception.Validate(); err != nil {
		return nil, err
	}
	
	exists, err := s.repo.PvzExists(ctx, reception.PvzID)
	if err != nil {
		return nil, err
	} 

	if !exists {
		return nil, errorsx.ErrPVZNotFound
	}

	lastReceptionStatus, err := s.repo.LastReceptionStatus(ctx, reception.PvzID)
	if err != nil {
		if !errors.Is(err, errorsx.ErrReceptionNotFound) {
			return nil, err
		}
	}

	if lastReceptionStatus == entity.ReceptionStatusInProgress {
		return nil, errorsx.ErrReceptionIsNotClosed
	}


	reception.Status = entity.ReceptionStatusInProgress

	reception, err = s.repo.CreateReception(ctx, reception)
	if err != nil {
		return nil, err
	}

	return reception, nil
}


func (s *Service) CloseReception(ctx context.Context, pvzUUID uuid.UUID) (*entity.Reception, error) {
	
	exists, err := s.repo.PvzExists(ctx, pvzUUID)
	if err != nil {
		return nil, err
	} 

	if !exists {
		return nil, errorsx.ErrPVZNotFound
	}


	lastReceptionStatus, err := s.repo.LastReceptionStatus(ctx, pvzUUID)
	if err != nil {
		return nil, err
	}

	if lastReceptionStatus == entity.ReceptionStatusCompleted {
		return nil, errorsx.ErrReceptionIsClosed
	}
	
	reception, err := s.repo.CloseReception(ctx, pvzUUID)
	if err != nil {
		return nil, err
	}

	return reception, nil
}




