package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // Импортируйте драйвер для вашей СУБД
)

func Connect(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("could not ping database: %v", err)
	}

	return db, nil
}
