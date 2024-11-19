package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(databaseURL string) error {
	var err error
	DB, err = sql.Open("postgres", databaseURL)
	if err != nil {
		return err
	}

	return DB.Ping()
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}