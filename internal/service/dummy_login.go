package service

import (
	"context"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
)

func (s *Service) DummyLogin(ctx context.Context, dummyLogin *entity.DummyLogin) (*entity.JwtToken, error) {
	
	if err := dummyLogin.Validate(); err != nil {
		return nil, err
	}
	exists, err := s.repo.RoleExists(ctx, dummyLogin.Role)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errorsx.ErrRoleNotFound
	}
	userUUID := "00000000-0000-0000-0000-000000000000"
	token, err := s.tokenGenerator.GenerateToken(userUUID, dummyLogin.Role)
	if err != nil {
		return nil, err
	}
	return &entity.JwtToken{Token: token}, nil
}
