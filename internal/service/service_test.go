package service_test

import (
	"log/slog"
	"strings"
	"testing"

	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/ZhdanovichVlad/service-podof/internal/service/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestService_HashPassword(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)
	logger := slog.Default()

	service := service.NewService(mockRepo, logger, mockTokenGenerator)

	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "valid password",
			password: "validPassword123",
			wantErr:  false,
		},
		{
			name:     "empty password",
			password: "",
			wantErr:  true,
		},
		{
			name:     "very long password",
			password: strings.Repeat("a", 73),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		
			hashedPassword, err := service.HashPassword(tt.password)

		
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, hashedPassword)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, hashedPassword)
			assert.NotEqual(t, tt.password, hashedPassword)

	
			isValid := service.ComparePassword(tt.password, hashedPassword)
			assert.True(t, isValid)
		})
	}
}

func TestService_ComparePassword(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)
	logger := slog.Default()

	service := service.NewService(mockRepo, logger, mockTokenGenerator)

	tests := []struct {
		name           string
		password       string
		hashedPassword string
		want           bool
	}{
		{
			name:     "valid password",
			password: "correctPassword",
		
			hashedPassword: "$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12",
			want:           true,
		},
		{
			name:     "invalid password",
			password: "wrongPassword",

			hashedPassword: "$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12",
			want:           false,
		},
		{
			name:           "empty password",
			password:       "",
			hashedPassword: "$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12",
			want:           false,
		},
		{
			name:           "invalid hash format",
			password:       "somePassword",
			hashedPassword: "invalid_hash",
			want:           false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		
			if tt.name == "valid password" {
				var err error
				tt.hashedPassword, err = service.HashPassword(tt.password)
				assert.NoError(t, err)
			}
	
			got := service.ComparePassword(tt.password, tt.hashedPassword)


			assert.Equal(t, tt.want, got)
		})
	}
}


func TestService_PasswordHashingAndComparing(t *testing.T) {

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockrepository(ctrl)
	mockTokenGenerator := mocks.NewMocktokenGenerator(ctrl)
	logger := slog.Default()

	service := service.NewService(mockRepo, logger, mockTokenGenerator)

	password := "mySecurePassword123"


	hash1, err := service.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash1)


	isValid := service.ComparePassword(password, hash1)
	assert.True(t, isValid)


	isValid = service.ComparePassword("wrongPassword", hash1)
	assert.False(t, isValid)

	
	hash2, err := service.HashPassword(password)
	assert.NoError(t, err)
	assert.NotEqual(t, hash1, hash2)


	isValid = service.ComparePassword(password, hash2)
	assert.True(t, isValid)
}
