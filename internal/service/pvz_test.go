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

func TestGetPvzList_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	mockRepo.EXPECT().GetPvzList(gomock.Any(), gomock.Any()).Return([]entity.PvzInfo{}, nil)
	
	pvzList, err := service.GetPvzList(context.Background(), entity.Filter{})
	assert.NoError(t, err)
	assert.NotNil(t, pvzList)
	
}

func TestGetPvzList_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)
	
	mockRepo.EXPECT().GetPvzList(gomock.Any(), gomock.Any()).Return(nil, errorsx.ErrInternal)

	pvzList, err := service.GetPvzList(context.Background(), entity.Filter{})
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, pvzList)
}

func NewService(mockRepo *mocks.Mockrepository, logger *slog.Logger, mockTokenGenerator *mocks.MocktokenGenerator) any {
	panic("unimplemented")
}


func TestCreatePvz_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	pvz := &entity.Pvz{
		Id: uuid.New(),
		City: "Москва",
		RegistrationDate: time.Now(),
	}
	newPvz := &entity.Pvz{}

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvz.Id).Return(false, nil)
	mockRepo.EXPECT().CityExists(gomock.Any(), pvz.City).Return(true, nil)
	mockRepo.EXPECT().CreatePvz(gomock.Any(), pvz).Return(newPvz, nil)

	pvz, err := service.CreatePvz(context.Background(), pvz)
	assert.NoError(t, err)
	assert.NotNil(t, newPvz)
}

func TestCreatePvz_PvzExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	pvz := &entity.Pvz{
		Id: uuid.New(),
		City: "Москва",
		RegistrationDate: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvz.Id).Return(true, nil)	

	pvz, err := service.CreatePvz(context.Background(), pvz)
	assert.ErrorIs(t, err, errorsx.ErrPvzExists)
	assert.Nil(t, pvz)
}

func TestCreatePvz_PvzExists_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	pvz := &entity.Pvz{
		Id: uuid.New(),
		City: "Москва",
		RegistrationDate: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvz.Id).Return(true, errorsx.ErrInternal)	

	pvz, err := service.CreatePvz(context.Background(), pvz)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, pvz)
}

func TestCreatePvz_CityNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	pvz := &entity.Pvz{
		Id: uuid.New(),
		City: "Москва-Питер	",
		RegistrationDate: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvz.Id).Return(false, nil)
	mockRepo.EXPECT().CityExists(gomock.Any(), pvz.City).Return(false, nil)

	pvz, err := service.CreatePvz(context.Background(), pvz)
	assert.ErrorIs(t, err, errorsx.ErrCityIsNotExists)
	assert.Nil(t, pvz)
}

func TestCreatePvz_CityNotExists_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	pvz := &entity.Pvz{
		Id: uuid.New(),
		City: "Москва-Питер	",
		RegistrationDate: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvz.Id).Return(false, nil)
	mockRepo.EXPECT().CityExists(gomock.Any(), pvz.City).Return(false, errorsx.ErrInternal)

	pvz, err := service.CreatePvz(context.Background(), pvz)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, pvz)
}


func TestCreatePvz_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	pvz := &entity.Pvz{
		Id: uuid.New(),
		City: "Москва",
		RegistrationDate: time.Now(),
	}

	mockRepo.EXPECT().PvzExists(gomock.Any(), pvz.Id).Return(false, nil)
	mockRepo.EXPECT().CityExists(gomock.Any(), pvz.City).Return(true, nil)
	mockRepo.EXPECT().CreatePvz(gomock.Any(), pvz).Return(nil, errorsx.ErrInternal)

	pvz, err := service.CreatePvz(context.Background(), pvz)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, pvz)
}

