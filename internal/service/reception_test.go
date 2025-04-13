package service_test

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/ZhdanovichVlad/service-podof/internal/service/mocks"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


func TestCreateReception_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		PvzID: uuid.New(),
		DateTime: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), reception.PvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), reception.PvzID).Return(entity.ReceptionStatusCompleted, nil)
	mockRepo.EXPECT().CreateReception(gomock.Any(), reception).Return(reception, nil)

	reception, err := service.CreateReception(context.Background(), reception)
	assert.NoError(t, err)
	assert.NotNil(t, reception)
}

func TestCreateReception_PvzNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)	

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		PvzID: uuid.New(),
		DateTime: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), reception.PvzID).Return(false, errorsx.ErrPVZNotFound)

	reception, err := service.CreateReception(context.Background(), reception)
	assert.ErrorIs(t, err, errorsx.ErrPVZNotFound)
	assert.Nil(t, reception)
}

func TestCreateReception_PvzNotFound_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)	

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		PvzID: uuid.New(),
		DateTime: time.Now(),
	}	

	mockRepo.EXPECT().PvzExists(gomock.Any(), reception.PvzID).Return(false, errorsx.ErrInternal)

	reception, err := service.CreateReception(context.Background(), reception)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, reception)
}

func TestCreateReception_LastReceptionStatusInProgress(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		PvzID: uuid.New(),
		DateTime: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), reception.PvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), reception.PvzID).Return(entity.ReceptionStatusInProgress, nil)

	reception, err := service.CreateReception(context.Background(), reception)
	assert.ErrorIs(t, err, errorsx.ErrReceptionIsNotClosed)
	assert.Nil(t, reception)
}

func TestCreateReception_LastReceptionStatusInProgress_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		PvzID: uuid.New(),
		DateTime: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), reception.PvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), reception.PvzID).Return("", errorsx.ErrInternal)

	reception, err := service.CreateReception(context.Background(), reception)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, reception)
}

func TestCreateReception_CreateReceptionError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		PvzID: uuid.New(),
		DateTime: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), reception.PvzID).Return(true, nil)
	mockRepo.EXPECT().LastReceptionStatus(gomock.Any(), reception.PvzID).Return(entity.ReceptionStatusCompleted, nil)
	mockRepo.EXPECT().CreateReception(gomock.Any(), reception).Return(nil, errorsx.ErrInternal)

	reception, err := service.CreateReception(context.Background(), reception)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, reception)
}	


func TestCreateReception_PvzIDIsNil(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	reception := &entity.Reception{
		Status: entity.ReceptionStatusInProgress,
		DateTime: time.Now(),
	}


	reception, err := service.CreateReception(context.Background(), reception)
	assert.ErrorIs(t, err, errorsx.ErrEmptyField)
	assert.Nil(t, reception)
}	



