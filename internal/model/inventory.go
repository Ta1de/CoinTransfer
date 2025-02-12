package model

type InventoryItem struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type Inventory struct {
	UserID int             `json:"user_id"`
	Items  []InventoryItem `json:"items"`
}
