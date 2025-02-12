package service

import (
	"CoinTransfer/internal/model"
	"CoinTransfer/internal/repository"
	"errors"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Authenticate(username, password string) (*model.User, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		// Если пользователь не найден, создаем нового
		user = &model.User{
			Username: username,
			Password: password,
			Coins:    0, // Начальное количество монет
		}
		if err := s.userRepo.Create(user); err != nil {
			return nil, err
		}
	} else if user.Password != password {
		return nil, errors.New("invalid password")
	}

	return user, nil
}

func (s *AuthService) GetUserByID(userID int) (*model.User, error) {
	return s.userRepo.FindByID(userID)
}
