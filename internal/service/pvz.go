package service

import (
	"context"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
)


func (s *Service) CreatePvz(ctx context.Context, pvz *entity.Pvz) (*entity.Pvz, error) {

	exists, err := s.repo.PvzExists(ctx, pvz.Id)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, errorsx.ErrPvzExists
	}

	exists, err = s.repo.CityExists(ctx, pvz.City)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errorsx.ErrCityIsNotExists
	}


	pvzAnswer, err := s.repo.CreatePvz(ctx, pvz)
	if err != nil {
		return nil, err
	}

	return pvzAnswer, nil
}


func (s *Service) GetPvzList(ctx context.Context, filter entity.Filter) ([]entity.PvzInfo, error) {
	list, err := s.repo.GetPvzList(ctx, filter)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *Service) GetAllPvz(ctx context.Context) ([]entity.Pvz, error) {
	list, err := s.repo.GetAllPvz(ctx)
	if err != nil {
		return nil, err
	}
	return list, nil
}
