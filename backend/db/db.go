// Package db handles database connectivity and initialization for the gnome service.
//
// This file manages the PostgreSQL database connection lifecycle:
// 1. Establishes connection to the Docker PostgreSQL container
// 2. Verifies connection with ping
// 3. Triggers schema initialization via InitializeTables()
// 4. Provides connection cleanup
//
// Related files:
// - docker-compose.yml: Defines the PostgreSQL container
// - schema.go: Handles database migrations
// - models/variant.go: Defines structs matching table schema
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

	if err := DB.Ping(); err != nil {
		return err
	}

	// Initialize tables from schema
	if err := InitializeTables(); err != nil {
		return err
	}

	return nil
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
