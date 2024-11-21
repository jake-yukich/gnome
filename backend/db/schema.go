// Package db schema.go handles database migrations and schema management.
//
// This file:
// 1. Embeds SQL migration files from the migrations/ directory
// 2. Executes migrations in order when InitializeTables() is called
// 3. Creates the tables that will later be populated by the ETL pipeline
//
// Related files:
// - migrations/*.sql: Contains the actual SQL schema definitions
// - db.go: Calls InitializeTables() after connection
// - models/variant.go: Contains Go structs matching these tables
package db

import (
	"embed"
	"fmt"
)

//go:embed migrations
var migrations embed.FS

// Read and execute migration files in the migrations/ directory
func InitializeTables() error {
	migrationFiles, err := migrations.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("failed to read migrations: %w", err)
	}

	for _, file := range migrationFiles {
		if !file.IsDir() && len(file.Name()) > 3 && file.Name()[len(file.Name())-6:] == "up.sql" {
			content, err := migrations.ReadFile("migrations/" + file.Name())
			if err != nil {
				return fmt.Errorf("failed to read migration %s: %w", file.Name(), err)
			}

			if _, err := DB.Exec(string(content)); err != nil {
				return fmt.Errorf("failed to execute migration %s: %w", file.Name(), err)
			}
		}
	}

	return nil
}

// Initial test table from testing Go connection to Postgres:
// func InitializeTables() error {
// 	createTableSQL := `
//     CREATE TABLE IF NOT EXISTS test_entries (
//         id SERIAL PRIMARY KEY,
//         message TEXT,
//         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
//     );`

// 	_, err := DB.Exec(createTableSQL)
// 	return err
// }
