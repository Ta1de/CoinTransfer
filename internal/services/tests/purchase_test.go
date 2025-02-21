package service_test

import (
	"errors"
	"testing"

	"CoinTransfer/internal/models"
	"CoinTransfer/internal/services"
	"CoinTransfer/mocks"

	"github.com/stretchr/testify/assert"
)

func TestBuyItemByName(t *testing.T) {
	mockRepo := new(mocks.MockBuyItemRepository)
	buyItemService := services.NewBuyItemService(mockRepo)

	tests := []struct {
		name        string
		itemName    string
		userID      int
		mockSetup   func()
		expectError bool
	}{
		{
			name:     "Item not found",
			itemName: "NonExistentItem",
			userID:   1,
			mockSetup: func() {
				mockRepo.On("GetItem", "NonExistentItem").
					Return(models.Item{}, errors.New("item not found"))
			},
			expectError: true,
		},
		{
			name:     "Insufficient balance",
			itemName: "ExpensiveItem",
			userID:   2,
			mockSetup: func() {
				mockRepo.On("GetItem", "ExpensiveItem").
					Return(models.Item{Item: "ExpensiveItem", Price: 500}, nil)
				mockRepo.On("GetBalance", 2).
					Return(100, nil)
			},
			expectError: true,
		},
		{
			name:     "Successful purchase",
			itemName: "Sword",
			userID:   3,
			mockSetup: func() {
				mockRepo.On("GetItem", "Sword").
					Return(models.Item{Item: "Sword", Price: 150}, nil)
				mockRepo.On("GetBalance", 3).
					Return(200, nil)
				mockRepo.On("UpdateBalance", 3, 150).
					Return(nil)
				mockRepo.On("AddToInventory", 3, "Sword").
					Return(nil)
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mockSetup()

			_, err := buyItemService.BuyItemByName(tt.itemName, tt.userID)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}
