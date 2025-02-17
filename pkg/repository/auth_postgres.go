package repository

import (
	"CoinTransfer/pkg/model"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user model.User) error {
	query := fmt.Sprintf("INSERT INTO %s (username, password_hash) VALUES ($1, $2)", usersTable)
	_, err := r.db.Exec(query, user.Username, user.Password)
	return err
}

func (r *AuthPostgres) GetUser(username, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE username = $1 AND password_hash = $2", usersTable)
	err := r.db.Get(&user, query, username, password)
	return user, err
}
