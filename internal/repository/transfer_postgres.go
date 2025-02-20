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

func (r *TransferPostgres) GetUserBalance(userID int) (int, error) {
	var balance int
	query := "SELECT coins FROM users WHERE id = $1"
	err := r.db.Get(&balance, query, userID)
	if err != nil {
		return 0, fmt.Errorf("could not get balance for user %d: %w", userID, err)
	}
	return balance, nil
}

func (r *TransferPostgres) UpdateBalances(senderID, receiverID, amount int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", amount, senderID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update sender's balance: %w", err)
	}

	_, err = tx.Exec("UPDATE users SET coins = coins + $1 WHERE id = $2", amount, receiverID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update receiver's balance: %w", err)
	}

	return tx.Commit()
}

func (r *TransferPostgres) SaveTransfer(fromUserId, toUserId, amount int) error {
	_, err := r.db.Exec(
		"INSERT INTO coinHistory (from_user, to_user, amount) VALUES ($1, $2, $3)",
		fromUserId, toUserId, amount,
	)
	if err != nil {
		return fmt.Errorf("failed to insert transfer: %v", err)
	}
	return nil
}
