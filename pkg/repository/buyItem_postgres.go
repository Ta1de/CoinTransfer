package repository

import (
	"CoinTransfer/pkg/model"

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
	err := r.db.Get(&item, "SELECT * FROM items WHERE name = $1", ItemName)
	if err != nil {
		return item, err
	}
	return item, nil
}
