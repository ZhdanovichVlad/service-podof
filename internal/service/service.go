package service

import (
	"context"
	"log/slog"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	minPasswordLength = 8  // минимальная длина пароля
	maxPasswordLength = 72 // максимальная длина из-за ограничения bcrypt
)

//go:generate mockgen -source=service.go -destination=mocks/mock_service.go -package=mocks
type repository interface {
	RoleExists(ctx context.Context, role string) (bool, error)
	UserExists(ctx context.Context, email string) (bool, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	CreatePvz(ctx context.Context, PvzRequestBodyCreate *entity.Pvz) (*entity.Pvz, error)
	PvzExists(ctx context.Context, id uuid.UUID) (bool, error)
	CityExists(ctx context.Context, city string) (bool, error)
	CreateReception(ctx context.Context, reception *entity.Reception) (*entity.Reception, error)
	LastReceptionStatus(ctx context.Context, pvzUUID uuid.UUID) (string, error)
	CloseReception(ctx context.Context, pvzUUID uuid.UUID) (*entity.Reception, error)
	DeleteLastProduct(ctx context.Context, pvzUUID uuid.UUID) error
	CreateProduct(ctx context.Context, product *entity.Product, pvzID uuid.UUID) (*entity.Product, error)
	GetPvzList(ctx context.Context, filter entity.Filter) ([]entity.PvzInfo, error)
	GetAllPvz(ctx context.Context) ([]entity.Pvz, error)
}

type tokenGenerator interface {
	GenerateToken(userUUID string, role string) (string, error)
}

type Service struct {
	repo           repository
	logger         *slog.Logger
	tokenGenerator tokenGenerator
}

func (s *Service) hashPassword(password string) (any, any) {
	panic("unimplemented")
}

func NewService(repo repository, logger *slog.Logger, generator tokenGenerator) *Service {
	return &Service{repo: repo, logger: logger, tokenGenerator: generator}
}

func (s *Service) HashPassword(password string) (string, error) {
	if password == "" {
		return "", errorsx.ErrInvalidPassword
	}
	if len(password) < minPasswordLength {
		return "", errorsx.ErrInvalidPassword
	}

	if len(password) > maxPasswordLength {
		return "", errorsx.ErrInvalidPassword
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Error("failed to hash password",
			slog.String("method", "service.hashPassword"),
			slog.String("error", err.Error()))
		return "", err
	}
	return string(hash), nil
}

func (s *Service) ComparePassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
