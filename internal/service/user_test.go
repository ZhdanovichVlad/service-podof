package service

import (
	"context"
	"log/slog"
	"testing"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/internal/service/mocks"
	"github.com/ZhdanovichVlad/service-podof/pkg/errorsx"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserService_RegisterSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}

	mockRepo.EXPECT().UserExists(gomock.Any(), user.Email).Return(false, nil)
	mockRepo.EXPECT().RoleExists(gomock.Any(), user.Role).Return(true, nil)
	mockRepo.EXPECT().CreateUser(gomock.Any(), user).Return(user, nil)


	registeredUser, err := service.Register(context.Background(), user)
	assert.NoError(t, err)
	assert.NotNil(t, registeredUser)
}

func TestUserService_RegisterUserExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}

	mockRepo.EXPECT().UserExists(gomock.Any(), user.Email).Return(true, nil)	

	registeredUser, err := service.Register(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrUserExists)
	assert.Nil(t, registeredUser)
}

func TestUserService_RegisterRoleNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)	

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}

	mockRepo.EXPECT().UserExists(gomock.Any(), user.Email).Return(false, nil)
	mockRepo.EXPECT().RoleExists(gomock.Any(), user.Role).Return(false, nil)

	registeredUser, err := service.Register(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrRoleNotFound)
	assert.Nil(t, registeredUser)
}

func TestUserService_RegisterError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}

	mockRepo.EXPECT().UserExists(gomock.Any(), user.Email).Return(false, nil)
	mockRepo.EXPECT().RoleExists(gomock.Any(), user.Role).Return(true, nil)
	mockRepo.EXPECT().CreateUser(gomock.Any(), user).Return(nil, errorsx.ErrInternal)

	registeredUser, err := service.Register(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, registeredUser)
}

func TestUserService_LoginSuccess(t *testing.T) {
  
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockrepository(ctrl)
    mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)
    logger := slog.Default()

    service := NewService(mockRepo, logger, mockTokenGenerator)

    password := "password123"
    hashedPassword, err := service.HashPassword(password)
    assert.NoError(t, err)

    user := &entity.User{
        Email:    "test@example.com",
        Password: hashedPassword,  
        Role:     "user",
    }

    expectedToken := "jwt.token.here"

    mockRepo.EXPECT().
        GetUserByEmail(gomock.Any(), user.Email).
        Return(user, nil)


    mockTokenGenerator.EXPECT().
        GenerateToken(gomock.Any(), user.Role).
        Return(expectedToken, nil)

    
    loginRequest := &entity.User{
        Email:    user.Email,
        Password: password, 
		Role:     user.Role,
    }

    token, err := service.Login(context.Background(), loginRequest)

    assert.NoError(t, err)
    assert.NotEmpty(t, token)
    assert.Equal(t, expectedToken, token.Token)
}


func TestUserService_LoginUserNotFound(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)	

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}

	mockRepo.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(nil, errorsx.ErrUserNotFound)

	token, err := service.Login(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrUserNotFound)
	assert.Nil(t, token)
}

func TestUserService_LoginInvalidPassword(t *testing.T) {
	ctrl := gomock.NewController(t)	
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)	

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}	

	mockRepo.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(user, nil)

	token, err := service.Login(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrInvalidPassword)
	assert.Nil(t, token)
}	

func TestUserService_LoginError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "password",
		Role: "user",
	}

	mockRepo.EXPECT().GetUserByEmail(gomock.Any(), user.Email).Return(nil, errorsx.ErrInternal)	

	token, err := service.Login(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrInternal)
	assert.Nil(t, token)
}


func TestUserService_Login_InternalError_GetUserByEmail(t *testing.T) {

    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockrepository(ctrl)
    mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)
    logger := slog.Default()

    service := NewService(mockRepo, logger, mockTokenGenerator)

    user := &entity.User{
        Email:    "test@example.com",
		Password: "password",
		Role: "user",
    }

    mockRepo.EXPECT().
        GetUserByEmail(gomock.Any(), user.Email).
        Return(nil, errorsx.ErrInternal)


    loginRequest := &entity.User{
        Email:    user.Email,
		Password: user.Password,
		Role: user.Role,
    }

    token, err := service.Login(context.Background(), loginRequest)

    assert.ErrorIs(t, err, errorsx.ErrInternal)
    assert.Nil(t, token)
  
}

func TestContextCancellation(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockRepo := mocks.NewMockrepository(ctrl)
    mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)
    logger := slog.Default()

    service := NewService(mockRepo, logger, mockTokenGenerator)

    ctx, cancel := context.WithCancel(context.Background())
    cancel()

    user := &entity.User{
        Email:    "test@example.com",
        Password: "password123",
		Role: "user",
    }

    mockRepo.EXPECT().
        GetUserByEmail(ctx, user.Email).
        Return(nil, context.Canceled)

    token, err := service.Login(ctx, user)
    assert.ErrorIs(t, err, context.Canceled)
    assert.Empty(t, token)
}

func TestUserService_LoginEmailIsEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)	

	user := &entity.User{
		Email: "",
		Password: "password",
		Role: "user",
	}

	token, err := service.Login(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrEmptyField)
	assert.Nil(t, token)
}

func TestUserService_LoginPasswordIsEmpty(t *testing.T) {	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)	

	service := NewService(mockRepo, slog.Default(), mockTokenGenerator)

	user := &entity.User{
		Email: "test@example.com",
		Password: "",
		Role: "user",
	}	

	token, err := service.Login(context.Background(), user)
	assert.ErrorIs(t, err, errorsx.ErrEmptyField)
	assert.Nil(t, token)
}	




