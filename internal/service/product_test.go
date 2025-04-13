package service_test

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/ZhdanovichVlad/service-podof/internal/service/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateProduct_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	product := &entity.Product{
		Type      : "test",
		ReceptionID : uuid.New(),
		DateTime  : time.Now(),
	}

	pvzID := uuid.New()


	mockRepo.EXPECT().PvzExists(gomock.Any(), pvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), pvzID).Return(entity.ReceptionStatusInProgress, nil)
	mockRepo.EXPECT().CreateProduct(gomock.Any(), product, pvzID).Return(product, nil)

	product, err := service.CreateProduct(context.Background(), product, pvzID)
	assert.NoError(t, err)
	assert.NotNil(t, product)
}

func TestCreateProduct_PvzNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	product := &entity.Product{
		Type      : "test",
		ReceptionID : uuid.New(),
		DateTime  : time.Now(),
	}

	pvzID := uuid.New()


	mockRepo.EXPECT().PvzExists(gomock.Any(), pvzID).Return(false, nil)
	

	product, err := service.CreateProduct(context.Background(), product, pvzID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errorsx.ErrPVZNotFound)
	assert.Nil(t, product)
}

func TestCreateProduct_ReceptionIsCompleted(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	product := &entity.Product{
		Type      : "test",
		ReceptionID : uuid.New(),
		DateTime  : time.Now(),
	}

	pvzID := uuid.New()

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), pvzID).Return(entity.ReceptionStatusCompleted, nil)


	product, err := service.CreateProduct(context.Background(), product, pvzID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errorsx.ErrReceptionIsClosed)
	assert.Nil(t, product)
}

func TestCreateProduct_InternalPVZError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	product := &entity.Product{
		Type      : "test",
		ReceptionID : uuid.New(),
		DateTime  : time.Now(),
	}	

	pvzID := uuid.New()

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvzID).Return(false, errorsx.ErrInternal)


	product, err := service.CreateProduct(context.Background(), product, pvzID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, product)
}
func TestCreateProduct_LastReceptionStatusInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	product := &entity.Product{
		Type      : "test",
		ReceptionID : uuid.New(),
		DateTime  : time.Now(),
	}	

	pvzID := uuid.New()

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), pvzID).Return("", errorsx.ErrInternal)
	
	product, err := service.CreateProduct(context.Background(), product, pvzID)
	assert.Error(t, err)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, product)
}

func TestCreateProduct_CreateProductInternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()	

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	product := &entity.Product{
		Type      : "test",
		ReceptionID : uuid.New(),
		DateTime  : time.Now(),	
	}

	pvzID := uuid.New()

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), pvzID).Return(entity.ReceptionStatusInProgress, nil)
	mockRepo.EXPECT().CreateProduct(gomock.Any(), product, pvzID).Return(nil, errorsx.ErrInternal)

	product, err := service.CreateProduct(context.Background(), product, pvzID)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, product)
}