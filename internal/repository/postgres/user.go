package postgres

import (
	"context"
	"errors"
	"log/slog"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/jackc/pgx/v5"
)

const (
	queryRoleExists = `SELECT EXISTS(SELECT 1 FROM roles WHERE name = $1 AND deleted_at IS NULL)`
	queryUserExists = `SELECT EXISTS(SELECT 1 FROM users WHERE email = $1 AND deleted_at IS NULL)`
	queryCreateUser = `INSERT INTO users (email, password_hash, role_name) 
    				   VALUES ($1, $2, $3) 
    				   RETURNING id, email, password_hash, role_name`

	queryGetUserByEmail = `SELECT id, email, password_hash, role_name FROM users WHERE email = $1 AND deleted_at IS NULL`
)

func (s *Storage) RoleExists(ctx context.Context, role string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(ctx, queryRoleExists, role).Scan(&exists)
	if err != nil {
		s.logger.Error("failed to check if role exists",
			slog.String("method", "repository.RoleExists"),
			slog.String("error", err.Error()))
		return false, errorsx.ErrInternal
	}

	return exists, nil
}

func (s *Storage) UserExists(ctx context.Context, email string) (bool, error) {
	var exists bool
	err := s.db.QueryRow(ctx, queryUserExists, email).Scan(&exists)
	if err != nil {
		s.logger.Error("failed to check if user exists",
			slog.String("method", "repository.UserExists"),
			slog.String("error", err.Error()))
		return false, errorsx.ErrInternal
	}

	return exists, nil
}

func (s *Storage) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	var createdUser entity.User
	err := s.db.QueryRow(ctx,
		queryCreateUser,
		user.Email,
		user.Password,
		user.Role).
		Scan(&createdUser.ID,
			&createdUser.Email,
			&createdUser.Password,
			&createdUser.Role)
	if err != nil {
		s.logger.Error("failed to create user",
			slog.String("method", "repository.CreateUser"),
			slog.String("error", err.Error()))
		return nil, errorsx.ErrInternal
	}
	return &createdUser, nil
}

func (s *Storage) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	err := s.db.QueryRow(ctx, queryGetUserByEmail, email).
	            Scan(&user.ID, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, errorsx.ErrUserNotFound
		}
		s.logger.Error("failed to get user by email",
			slog.String("method", "repository.GetUserByEmail"),
			slog.String("error", err.Error()))
		return nil, errorsx.ErrInternal
	}
	return &user, nil
}
