package service_test

import (
	"errors"
	"testing"

	"CoinTransfer/internal/services"
	"CoinTransfer/mocks"

	"github.com/stretchr/testify/assert"
)

func TestSendCoins(t *testing.T) {
	mockRepo := new(mocks.MockTransferRepository)
	transferService := services.NewTransferService(mockRepo)

	tests := []struct {
		name        string
		fromUserId  int
		toUser      string
		amount      int
		mockSetup   func()
		expectError bool
	}{
		{
			name:       "Recipient user not found",
			fromUserId: 1,
			toUser:     "unknownUser",
			amount:     50,
			mockSetup: func() {
				mockRepo.On("GetUserIDByUsername", "unknownUser").
					Return(0, errors.New("user not found"))
			},
			expectError: true,
		},
		{
			name:       "Sender balance retrieval fails",
			fromUserId: 2,
			toUser:     "pers1",
			amount:     50,
			mockSetup: func() {
				mockRepo.On("GetUserIDByUsername", "pers1").
					Return(2, nil)
				mockRepo.On("GetUserBalance", 2).
					Return(0, errors.New("database error"))
			},
			expectError: true,
		},
		{
			name:       "Insufficient balance",
			fromUserId: 3,
			toUser:     "pers2",
			amount:     100,
			mockSetup: func() {
				mockRepo.On("GetUserIDByUsername", "pers2").
					Return(2, nil)
				mockRepo.On("GetUserBalance", 3).
					Return(50, nil)
			},
			expectError: true,
		},
		{
			name:       "Balance update failure",
			fromUserId: 4,
			toUser:     "pers3",
			amount:     50,
			mockSetup: func() {
				mockRepo.On("GetUserIDByUsername", "pers3").
					Return(2, nil)
				mockRepo.On("GetUserBalance", 4).
					Return(100, nil)
				mockRepo.On("UpdateBalances", 4, 2, 50).
					Return(errors.New("update failed"))
			},
			expectError: true,
		},
		{
			name:       "Transfer save failure",
			fromUserId: 5,
			toUser:     "pers4",
			amount:     50,
			mockSetup: func() {
				mockRepo.On("GetUserIDByUsername", "pers4").
					Return(2, nil)
				mockRepo.On("GetUserBalance", 5).
					Return(100, nil)
				mockRepo.On("UpdateBalances", 5, 2, 50).
					Return(nil)
				mockRepo.On("SaveTransfer", 5, 2, 50).
					Return(errors.New("save failed"))
			},
			expectError: true,
		},
		{
			name:       "Successful transfer",
			fromUserId: 6,
			toUser:     "pers5",
			amount:     50,
			mockSetup: func() {
				mockRepo.On("GetUserIDByUsername", "pers5").
					Return(2, nil)
				mockRepo.On("GetUserBalance", 6).
					Return(100, nil)
				mockRepo.On("UpdateBalances", 6, 2, 50).
					Return(nil)
				mockRepo.On("SaveTransfer", 6, 2, 50).
					Return(nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			err := transferService.SendCoins(tt.fromUserId, tt.toUser, tt.amount)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
