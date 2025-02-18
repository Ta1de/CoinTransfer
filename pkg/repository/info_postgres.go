package repository

import (
	"CoinTransfer/pkg/model"

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

func (r *InfoPostgres) TransferHistory(UserId int) (int, error) {

	return 0, nil
}
