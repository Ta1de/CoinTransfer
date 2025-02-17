package service

import (
	"CoinTransfer/pkg/model"
	"CoinTransfer/pkg/repository"
	"crypto/sha1"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	salt       = "jh32v4hj23v2v4h2vg332h4g3hv4"
	singingKey = "efjn#JKf3%3#wegggegwe7we67W%3#23deg"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.StandardClaims
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (string, error) {
	user.Password = generatePasswordHash(user.Password)
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	})

	return token.SignedString([]byte(singingKey))
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
