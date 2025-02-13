package database

import (
	"CoinTransfer/internal/model"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Функция для инициализации подключения к базе данных
func InitDB() {
	if DB == nil {
		dsn := "user=twitchis password=yourpassword dbname=twitchis host=localhost port=5432 sslmode=disable"
		var err error
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		} else {
			log.Println("Successfully connected to the database!") // Логирование успешного подключения
		}

		// Миграции
		err = DB.AutoMigrate(&model.User{})
		if err != nil {
			log.Fatal("Failed to migrate database:", err)
		} else {
			log.Println("Database migration successful!") // Логирование успешной миграции
		}
	}
}
