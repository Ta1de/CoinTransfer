package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TransferPostgres struct {
	db *sqlx.DB
}

func NewTransferPostgres(db *sqlx.DB) *TransferPostgres {
	return &TransferPostgres{db: db}
}

func (r *TransferPostgres) GetUserIDByUsername(username string) (int, error) {
	var userID int
	query := "SELECT id FROM users WHERE username = $1"
	err := r.db.Get(&userID, query, username)
	if err != nil {
		return 0, fmt.Errorf("could not find user with username %s: %w", username, err)
	}
	return userID, nil
}

func (r *TransferPostgres) GetUserBalance(userID int) (float64, error) {
	var balance float64
	query := "SELECT balance FROM users WHERE id = $1"
	err := r.db.Get(&balance, query, userID)
	if err != nil {
		return 0, fmt.Errorf("could not get balance for user %d: %w", userID, err)
	}
	return balance, nil
}

func (r *TransferPostgres) UpdateBalances(senderID, receiverID int, amount float64) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	// Уменьшаем баланс отправителя
	_, err = tx.Exec("UPDATE users SET balance = balance - $1 WHERE id = $2", amount, senderID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update sender's balance: %w", err)
	}

	// Увеличиваем баланс получателя
	_, err = tx.Exec("UPDATE users SET balance = balance + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update receiver's balance: %w", err)
	}

	return tx.Commit()
}

func (r *TransferPostgres) SaveTransfer(fromUserId int, toUser string, amount int) error {
	var toUserId int
	err := r.db.QueryRow("SELECT id FROM users WHERE username = $1", toUser).Scan(&toUserId)
	if err != nil {
		return fmt.Errorf("failed to find user: %v", err)
	}

	_, err = r.db.Exec(
		"INSERT INTO transfer (from_user_id, to_user_id, amount) VALUES ($1, $2, $3)",
		fromUserId, toUserId, amount,
	)
	if err != nil {
		return fmt.Errorf("failed to insert transfer: %v", err)
	}

	return nil
}
