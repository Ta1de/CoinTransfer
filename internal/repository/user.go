package repository

import (
	"CoinTransfer/internal/model"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, username, password, coins FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &user.Password, &user.Coins)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) FindByID(userID int) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, username, password, coins FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Username, &user.Password, &user.Coins)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(user *model.User) error {
	_, err := r.db.Exec("INSERT INTO users (username, password, coins) VALUES (?, ?, ?)", user.Username, user.Password, user.Coins)
	return err
}
