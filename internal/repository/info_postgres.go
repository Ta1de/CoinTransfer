package repository

import (
	"CoinTransfer/internal/models"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (r *InfoPostgres) GetInventory(UserId int) (models.InventoryItems, error) {
	var inventoryItems models.InventoryItems
	query := `SELECT item, quantity FROM inventory WHERE user_id = $1`

	rows, err := r.db.Query(query, UserId)
	if err != nil {
		return inventoryItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.InventoryItem
		if err := rows.Scan(&item.Type, &item.Quantity); err != nil {
			return inventoryItems, err
		}
		inventoryItems.Inventory = append(inventoryItems.Inventory, item)
	}

	if err := rows.Err(); err != nil {
		return inventoryItems, err
	}

	return inventoryItems, nil
}

func (r *InfoPostgres) TransferHistory(UserId int) (models.History, error) {
	var history models.History

	receivedTransactions, err := r.getReceivedTransactions(UserId)
	if err != nil {
		return history, err
	}

	sentTransactions, err := r.getSentTransactions(UserId)
	if err != nil {
		return history, err
	}

	for _, tx := range receivedTransactions {
		history.Received = append(history.Received, models.ReceivedTransaction{
			FromUser: tx.FromUser,
			Amount:   tx.Amount,
		})
	}

	for _, tx := range sentTransactions {
		history.Sent = append(history.Sent, models.SentTransaction{
			ToUser: tx.ToUser,
			Amount: tx.Amount,
		})
	}

	return history, nil
}

func (r *InfoPostgres) getReceivedTransactions(UserId int) ([]models.ReceivedTransaction, error) {
	var transactions []models.ReceivedTransaction
	query := `
		SELECT u.username, ch.amount 
		FROM coinHistory ch
		JOIN users u ON u.id = ch.from_user 
		WHERE ch.to_user = $1
	`
	rows, err := r.db.Query(query, UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tx models.ReceivedTransaction
		if err := rows.Scan(&tx.FromUser, &tx.Amount); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *InfoPostgres) getSentTransactions(UserId int) ([]models.SentTransaction, error) {
	var transactions []models.SentTransaction
	query := `
		SELECT u.username, ch.amount 
		FROM coinHistory ch
		JOIN users u ON u.id = ch.to_user 
		WHERE ch.from_user = $1
	`
	rows, err := r.db.Query(query, UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var tx models.SentTransaction
		if err := rows.Scan(&tx.ToUser, &tx.Amount); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *InfoPostgres) GetCoins(userID int) (int, error) {
	var balance int
	query := "SELECT coins FROM users WHERE id = $1"
	err := r.db.Get(&balance, query, userID)
	if err != nil {
		return 0, fmt.Errorf("could not get balance for user %d: %w", userID, err)
	}
	return balance, nil
}
