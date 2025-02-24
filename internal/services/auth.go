package services

import (
	"CoinTransfer/internal/middleware"
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/repository"
	"crypto/sha1"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

const tokenTTL = 12 * time.Hour

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (string, error) {
	user.Password = GeneratePasswordHash(user.Password)
	err := s.repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	token, err := s.CreateToken(user.Username, user.Password)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *AuthService) CreateToken(username, password string) (string, error) {
	user, err := s.repo.GetUser(username, password)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &middleware.TokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		UserId: user.ID,
	})

	return token.SignedString([]byte(os.Getenv("singingKey")))
}

func GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(os.Getenv("salt"))))
}
