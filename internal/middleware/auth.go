package migrations

import (
	"CoinTransfer/internal/model"

	"gorm.io/gorm"
)

// Функция для миграции базы данных
func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{})
}
