package repository_test

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/repository"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewAuthPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	user := models.User{
		Username: "testuser",
		Password: "testpassword",
	}

	// Ожидаем, что будет выполнен запрос INSERT с определенными параметрами
	mock.ExpectExec("INSERT INTO users").
		WithArgs(user.Username, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Вызываем тестируемую функцию
	err = repo.CreateUser(user)

	// Проверяем, что ошибок нет и мок был вызван с ожидаемыми параметрами
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUser(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewAuthPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	username := "testuser"
	password := "testpassword"
	expectedUser := models.User{
		ID: 1, // Проверяем только ID
	}

	// Ожидаем, что будет выполнен запрос SELECT с определенными параметрами
	rows := sqlmock.NewRows([]string{"id"}). // Возвращаем только поле id
							AddRow(expectedUser.ID)
	mock.ExpectQuery("SELECT id FROM users WHERE username = \\$1 AND password_hash = \\$2").
		WithArgs(username, password).
		WillReturnRows(rows)

	// Вызываем тестируемую функцию
	user, err := repo.GetUser(username, password)

	// Проверяем, что ошибок нет, мок был вызван с ожидаемыми параметрами и результат соответствует ожидаемому
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.ID, user.ID) // Проверяем только ID
	assert.NoError(t, mock.ExpectationsWereMet())
}
