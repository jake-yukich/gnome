package db

func InitializeTables() error {
	createTableSQL := `
    CREATE TABLE IF NOT EXISTS test_entries (
        id SERIAL PRIMARY KEY,
        message TEXT,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := DB.Exec(createTableSQL)
	return err
}
