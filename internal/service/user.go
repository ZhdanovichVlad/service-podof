package service

import (
	"context"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
)

func (s *Service) Register(ctx context.Context, user *entity.User) (*entity.User, error) {
	
	if err := user.Validate(); err != nil {
		return nil, err
	}

	exists, err := s.repo.UserExists(ctx, user.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errorsx.ErrUserExists
	}

	exists, err = s.repo.RoleExists(ctx, user.Role)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errorsx.ErrRoleNotFound
	}

	hashedPassword, err := s.HashPassword(user.Password)
	if err != nil {
		return nil,err
	}

	user.Password = hashedPassword

	createdUser, err := s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, errorsx.ErrInternal
	}

	return createdUser, nil
}

func (s *Service) Login(ctx context.Context, user *entity.User) (*entity.JwtToken, error) {
	if user.Email == "" || user.Password == "" {
		return nil, errorsx.ErrEmptyField
	}
	
	userDBInfo, err := s.repo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if !s.ComparePassword(user.Password, userDBInfo.Password) {
		return nil, errorsx.ErrInvalidPassword
	}

	token, err := s.tokenGenerator.GenerateToken(userDBInfo.ID.String(), userDBInfo.Role)
	if err != nil {
		return nil, err
	}

	return &entity.JwtToken{Token: token}, nil

}
