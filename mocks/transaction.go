package mocks

import (
	"github.com/stretchr/testify/mock"
)

type MockTransferRepository struct {
	mock.Mock
}

func (m *MockTransferRepository) GetUserIDByUsername(username string) (int, error) {
	args := m.Called(username)
	return args.Int(0), args.Error(1)
}

func (m *MockTransferRepository) GetUserBalance(userID int) (int, error) {
	args := m.Called(userID)
	return args.Int(0), args.Error(1)
}

func (m *MockTransferRepository) UpdateBalances(fromUserID, toUserID, amount int) error {
	args := m.Called(fromUserID, toUserID, amount)
	return args.Error(0)
}

func (m *MockTransferRepository) SaveTransfer(fromUserID, toUserID, amount int) error {
	args := m.Called(fromUserID, toUserID, amount)
	return args.Error(0)
}
