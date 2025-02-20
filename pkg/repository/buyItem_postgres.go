package repository

import (
	"CoinTransfer/pkg/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type BuyItemPostgres struct {
	db *sqlx.DB
}

func NewBuyItemPostgres(db *sqlx.DB) *BuyItemPostgres {
	return &BuyItemPostgres{db: db}
}

func (r *BuyItemPostgres) GetItem(ItemName string) (model.Item, error) {
	var item model.Item
	err := r.db.Get(&item, "SELECT * FROM items WHERE item = $1", ItemName)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (r *BuyItemPostgres) GetBalance(userID int) (int, error) {
	var balance int
	query := "SELECT coins FROM users WHERE id = $1"
	err := r.db.Get(&balance, query, userID)
	if err != nil {
		return 0, fmt.Errorf("could not get balance for user %d: %w", userID, err)
	}
	return balance, nil
}

func (r *BuyItemPostgres) AddToInventory(userID int, itemName string) error {
	query := `
	INSERT INTO inventory (user_id, item, quantity) 
	VALUES ($1, $2, $3) 
	ON CONFLICT (user_id, item) 
	DO UPDATE SET quantity = inventory.quantity + EXCLUDED.quantity`

	_, err := r.db.Exec(query, userID, itemName, 1)
	if err != nil {
		return fmt.Errorf("failed to add item to inventory: %w", err)
	}
	return nil
}

func (r *BuyItemPostgres) UpdateBalance(senderID, amount int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("could not begin transaction: %w", err)
	}

	_, err = tx.Exec("UPDATE users SET coins = coins - $1 WHERE id = $2", amount, senderID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("could not update sender's balance: %w", err)
	}

	return tx.Commit()
}
