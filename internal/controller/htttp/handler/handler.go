package handler

import (
	"context"
	"log/slog"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/google/uuid"
)
//go:generate mockgen -source=handler.go -destination=mocks/mock_handler.go -package=mocks
type api interface {
	DummyLogin(ctx context.Context, dummyLogin *entity.DummyLogin) (*entity.JwtToken, error)
	Register(ctx context.Context, user *entity.User) (*entity.User, error)
	Login(ctx context.Context, user *entity.User) (*entity.JwtToken, error)
	CreatePvz(ctx context.Context, pvz *entity.Pvz) (*entity.Pvz, error)	
	CreateReception(ctx context.Context, reception *entity.Reception) (*entity.Reception, error)
	CloseReception(ctx context.Context, pvzUUID uuid.UUID) (*entity.Reception, error)
	CreateProduct(ctx context.Context, product *entity.Product, pvzID uuid.UUID) (*entity.Product, error)
	DeleteLastProduct(ctx context.Context, pvzID uuid.UUID) (error)
	GetPvzList(ctx context.Context, filter entity.Filter) ([]entity.PvzInfo, error)
}

type Handler struct {
	service api
	logger  *slog.Logger
}

func NewHandler(api api, logger *slog.Logger) *Handler {
	return &Handler{api, logger}
}
