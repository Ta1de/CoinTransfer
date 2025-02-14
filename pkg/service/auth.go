package service

import (
	"CoinTransfer/pkg/model"
	"CoinTransfer/pkg/repository"
	"crypto/sha1"
	"fmt"
)

const salt = "jh32v4hj23v2v4h2vg332h4g3hv4"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
