package service_test

import (
	"context"
	"log/slog"
	"testing"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/ZhdanovichVlad/service-podof/internal/service/mocks"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestDummyLogin_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	moderator := &entity.DummyLogin{
		Role: "moderator",
	}

	userUUID := "00000000-0000-0000-0000-000000000000"

	mockRepo.EXPECT().RoleExists(gomock.Any(), moderator.Role).Return(true, nil)
	mockTokenGenerator.EXPECT().GenerateToken(userUUID, moderator.Role).Return("token", nil)

	token, err := service.DummyLogin(context.Background(), moderator)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	assert.Equal(t, token.Token, "token")
}

func TestDummyLogin_RoleNotExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	role := &entity.DummyLogin{
		Role: "test",
	}

	mockRepo.EXPECT().RoleExists(gomock.Any(), role.Role).Return(false, nil)

	token, err := service.DummyLogin(context.Background(), role)
	assert.ErrorIs(t, err, errorsx.ErrRoleNotFound)
	assert.Empty(t, token)
}

func TestDummyLogin_RoleExists_InternalError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	role := &entity.DummyLogin{
		Role: "test",
	}

	mockRepo.EXPECT().RoleExists(gomock.Any(), role.Role).Return(false, errorsx.ErrInternal)

	token, err := service.DummyLogin(context.Background(), role)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Empty(t, token)
}

func TestDummyLogin_GenerateTokenError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := service.NewService(mockRepo, slog.Default(), mockTokenGenerator)

	role := &entity.DummyLogin{
		Role: "test",
	}
	userUUID := "00000000-0000-0000-0000-000000000000"

	mockRepo.EXPECT().RoleExists(gomock.Any(), role.Role).Return(true, nil)
	mockTokenGenerator.EXPECT().GenerateToken(userUUID, role.Role).Return("", errorsx.ErrInternal)

	token, err := service.DummyLogin(context.Background(), role)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Empty(t, token)
}
