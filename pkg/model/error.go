package model

type Error struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
