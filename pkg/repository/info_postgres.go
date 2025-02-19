package repository

import (
	"CoinTransfer/pkg/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type InfoPostgres struct {
	db *sqlx.DB
}

func NewInfoPostgres(db *sqlx.DB) *InfoPostgres {
	return &InfoPostgres{db: db}
}

func (r *InfoPostgres) GetInventory(UserId int) (model.InventoryItems, error) {
	var inventoryItems model.InventoryItems
	query := `SELECT item, quantity FROM inventory WHERE user_id = $1`

	rows, err := r.db.Query(query, UserId)
	if err != nil {
		return inventoryItems, err
	}
	defer rows.Close()

	for rows.Next() {
		var item model.InventoryItem
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

func (r *InfoPostgres) TransferHistory(UserId int) (model.History, error) {
	var history model.History

	receivedTransactions, err := r.getReceivedTransactions(UserId)
	if err != nil {
		return history, err
	}

	sentTransactions, err := r.getSentTransactions(UserId)
	if err != nil {
		return history, err
	}

	for _, tx := range receivedTransactions {
		history.Received = append(history.Received, model.ReceivedTransaction{
			FromUser: tx.FromUser,
			Amount:   tx.Amount,
		})
	}

	for _, tx := range sentTransactions {
		history.Sent = append(history.Sent, model.SentTransaction{
			ToUser: tx.ToUser,
			Amount: tx.Amount,
		})
	}

	return history, nil
}

func (r *InfoPostgres) getReceivedTransactions(UserId int) ([]model.ReceivedTransaction, error) {
	var transactions []model.ReceivedTransaction
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
		var tx model.ReceivedTransaction
		if err := rows.Scan(&tx.FromUser, &tx.Amount); err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}
	return transactions, nil
}

func (r *InfoPostgres) getSentTransactions(UserId int) ([]model.SentTransaction, error) {
	var transactions []model.SentTransaction
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
		var tx model.SentTransaction
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
