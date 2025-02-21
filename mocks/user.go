package mocks

import (
	"CoinTransfer/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockAuthorizationRepo struct {
	mock.Mock
}

func (m *MockAuthorizationRepo) CreateUser(user models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockAuthorizationRepo) GetUser(username, password string) (models.User, error) {
	args := m.Called(username, password)
	return args.Get(0).(models.User), args.Error(1)
}
