package service_test

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/services"
	"CoinTransfer/mocks"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateUser_Success(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	password := "password123"
	hashedPassword := services.GeneratePasswordHash(password)

	user := models.User{
		Username: "testuser",
		Password: hashedPassword,
	}

	mockRepo.On("CreateUser", user).Return(nil)
	mockRepo.On("GetUser", user.Username, user.Password).Return(models.User{ID: 1, Username: user.Username}, nil)

	token, err := authService.CreateUser(models.User{
		Username: "testuser",
		Password: "password123",
	})

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_TokenCreationFailure(t *testing.T) {

	mockRepo := new(mocks.MockAuthorizationRepo)

	authService := services.NewAuthService(mockRepo)

	password := "password123"
	hashedPassword := services.GeneratePasswordHash(password)

	user := models.User{
		Username: "testuser",
		Password: hashedPassword,
	}

	mockRepo.On("CreateUser", user).Return(nil)

	mockRepo.On("GetUser", user.Username, hashedPassword).Return(models.User{ID: 1, Username: user.Username}, fmt.Errorf("failed to create user"))

	token, err := authService.CreateUser(models.User{
		Username: "testuser",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Equal(t, "failed to create user", err.Error())
	assert.Empty(t, token)

	mockRepo.AssertExpectations(t)
}

func TestCreateUser_Failure_CreateUser(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	password := "password123"
	hashedPassword := services.GeneratePasswordHash(password)

	user := models.User{
		Username: "testuser",
		Password: hashedPassword,
	}

	mockRepo.On("CreateUser", user).Return(fmt.Errorf("failed to create user"))

	token, err := authService.CreateUser(models.User{
		Username: "testuser",
		Password: "password123",
	})

	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestCreateToken_Success(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	username := "testuser"
	password := "password123"
	user := models.User{ID: 1, Username: username}

	mockRepo.On("GetUser", username, password).Return(user, nil)

	token, err := authService.CreateToken(username, password)

	assert.NoError(t, err)
	assert.NotEmpty(t, token)
	mockRepo.AssertExpectations(t)
}

func TestCreateToken_Failure(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	username := "testuser"
	password := "wrongpassword"

	mockRepo.On("GetUser", username, password).Return(models.User{}, fmt.Errorf("user not found"))

	token, err := authService.CreateToken(username, password)

	assert.Error(t, err)
	assert.Empty(t, token)
	mockRepo.AssertExpectations(t)
}
