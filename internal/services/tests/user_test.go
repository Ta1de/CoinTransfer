package service_test

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/services"
	"CoinTransfer/mocks"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

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

func TestParseToken_Success(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	userId := 1
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})
	signedToken, _ := token.SignedString([]byte(os.Getenv("singingKey")))

	parsedUserId, err := authService.ParseToken(signedToken)

	assert.NoError(t, err)
	assert.Equal(t, userId, parsedUserId)
}

func TestParseToken_Failure_InvalidToken(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	invalidToken := "invalidToken"

	parsedUserId, err := authService.ParseToken(invalidToken)

	assert.Error(t, err)
	assert.Equal(t, 0, parsedUserId)
}

func TestParseToken_Failure_ExpiredToken(t *testing.T) {
	mockRepo := new(mocks.MockAuthorizationRepo)
	authService := services.NewAuthService(mockRepo)

	userId := 1
	expiredToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: userId,
	})
	signedToken, _ := expiredToken.SignedString([]byte(os.Getenv("singingKey")))

	parsedUserId, err := authService.ParseToken(signedToken)

	assert.Error(t, err)
	assert.Equal(t, 0, parsedUserId)
}
