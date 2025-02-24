package repository_test

import (
	"CoinTransfer/internal/models"
	"CoinTransfer/internal/repository"
	"database/sql"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestGetItem(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	itemName := "test_item"
	expectedItem := models.Item{
		Id:    1,
		Item:  itemName,
		Price: 100,
	}

	// Ожидаем, что будет выполнен запрос SELECT с определенными параметрами
	rows := sqlmock.NewRows([]string{"id", "item", "price"}).
		AddRow(expectedItem.Id, expectedItem.Item, expectedItem.Price)
	mock.ExpectQuery("SELECT \\* FROM items WHERE item = \\$1").
		WithArgs(itemName).
		WillReturnRows(rows)

	// Вызываем тестируемую функцию
	item, err := repo.GetItem(itemName)

	// Проверяем, что ошибок нет, мок был вызван с ожидаемыми параметрами и результат соответствует ожидаемому
	assert.NoError(t, err)
	assert.Equal(t, expectedItem, item)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	userID := 1
	expectedBalance := 500

	rows := sqlmock.NewRows([]string{"coins"}).
		AddRow(expectedBalance)
	mock.ExpectQuery("SELECT coins FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	balance, err := repo.GetBalance(userID)

	assert.NoError(t, err)
	assert.Equal(t, expectedBalance, balance)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestAddToInventory(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	userID := 1
	itemName := "test_item"

	// Ожидаем, что будет выполнен запрос INSERT с определенными параметрами
	mock.ExpectExec("INSERT INTO inventory").
		WithArgs(userID, itemName, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Вызываем тестируемую функцию
	err = repo.AddToInventory(userID, itemName)

	// Проверяем, что ошибок нет и мок был вызван с ожидаемыми параметрами
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateBalance(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	senderID := 1
	amount := 100

	// Ожидаем, что будет выполнен запрос UPDATE с определенными параметрами
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE users SET coins = coins - \\$1 WHERE id = \\$2").
		WithArgs(amount, senderID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Вызываем тестируемую функцию
	err = repo.UpdateBalance(senderID, amount)

	// Проверяем, что ошибок нет и мок был вызван с ожидаемыми параметрами
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetItem_NotFound(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	itemName := "non_existent_item"

	// Ожидаем, что будет выполнен запрос SELECT, но предмет не будет найден
	mock.ExpectQuery("SELECT \\* FROM items WHERE item = \\$1").
		WithArgs(itemName).
		WillReturnError(sql.ErrNoRows)

	// Вызываем тестируемую функцию
	item, err := repo.GetItem(itemName)

	// Проверяем, что возвращается ошибка и объект item пустой
	assert.Error(t, err)
	assert.Equal(t, sql.ErrNoRows, err)
	assert.Equal(t, models.Item{}, item)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetItem_ScanError(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	itemName := "test_item"

	// Ожидаем, что будет выполнен запрос SELECT, но вернутся некорректные данные
	rows := sqlmock.NewRows([]string{"id", "item", "price"}).
		AddRow("invalid_id", itemName, 100) // "invalid_id" — строка вместо числа
	mock.ExpectQuery("SELECT \\* FROM items WHERE item = \\$1").
		WithArgs(itemName).
		WillReturnRows(rows)

	// Вызываем тестируемую функцию
	item, err := repo.GetItem(itemName)

	// Проверяем, что возвращается ошибка и объект item пустой
	assert.Error(t, err)
	assert.Equal(t, models.Item{}, item)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetItem_EmptyItemName(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	itemName := ""

	// Вызываем тестируемую функцию
	item, err := repo.GetItem(itemName)

	// Проверяем, что возвращается ошибка и объект item пустой
	assert.Error(t, err)
	assert.Equal(t, models.Item{}, item)
	assert.NoError(t, mock.ExpectationsWereMet()) // Убедимся, что запрос не был выполнен
}

func TestGetBalance_Error(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	userID := 1

	// Ожидаем, что будет выполнен запрос SELECT, но произойдет ошибка
	mock.ExpectQuery("SELECT coins FROM users WHERE id = \\$1").
		WithArgs(userID).
		WillReturnError(fmt.Errorf("database error"))

	// Вызываем тестируемую функцию
	balance, err := repo.GetBalance(userID)

	// Проверяем, что возвращается ошибка и баланс равен 0
	assert.Error(t, err)
	assert.Equal(t, 0, balance)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAddToInventory_Error(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	userID := 1
	itemName := "test_item"

	// Ожидаем, что будет выполнен запрос INSERT, но произойдет ошибка
	mock.ExpectExec("INSERT INTO inventory").
		WithArgs(userID, itemName, 1).
		WillReturnError(fmt.Errorf("database error"))

	// Вызываем тестируемую функцию
	err = repo.AddToInventory(userID, itemName)

	// Проверяем, что возвращается ошибка
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
func TestUpdateBalance_Error(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	senderID := 1
	amount := 100

	// Ожидаем, что будет выполнен запрос UPDATE, но произойдет ошибка
	mock.ExpectBegin()
	mock.ExpectExec("UPDATE users SET coins = coins - \\$1 WHERE id = \\$2").
		WithArgs(amount, senderID).
		WillReturnError(fmt.Errorf("database error"))
	mock.ExpectRollback() // Ожидаем откат транзакции

	// Вызываем тестируемую функцию
	err = repo.UpdateBalance(senderID, amount)

	// Проверяем, что возвращается ошибка
	assert.Error(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdateBalance_BeginTransactionError(t *testing.T) {
	// Создаем мок базы данных
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")
	repo := repository.NewBuyItemPostgres(sqlxDB)

	// Подготавливаем тестовые данные
	senderID := 1
	amount := 100

	// Ожидаем, что будет вызван Beginx(), но произойдет ошибка
	mock.ExpectBegin().WillReturnError(fmt.Errorf("could not begin transaction"))

	// Вызываем тестируемую функцию
	err = repo.UpdateBalance(senderID, amount)

	// Проверяем, что возвращается ошибка
	assert.Error(t, err)
	assert.EqualError(t, err, "could not begin transaction: could not begin transaction")
	assert.NoError(t, mock.ExpectationsWereMet())
}
