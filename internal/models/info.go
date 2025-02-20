package models

type InventoryItems struct {
	Inventory []InventoryItem `json:"inventory"`
}

type InventoryItem struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type History struct {
	Received []ReceivedTransaction `json:"received"`
	Sent     []SentTransaction     `json:"sent"`
}

type ReceivedTransaction struct {
	FromUser string `json:"fromUser"`
	Amount   int    `json:"amount"`
}

type SentTransaction struct {
	ToUser string `json:"toUser"`
	Amount int    `json:"amount"`
}

type Info struct {
	Coins       int            `json:"coins"`
	Inventory   InventoryItems `json:"inventory"`
	CoinHistory History        `json:"coinhistory"`
}
