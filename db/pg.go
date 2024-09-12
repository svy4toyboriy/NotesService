package db

import (
	"SportsTestWork/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

// Dial создает новое подключение к БД
func Dial(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.PgUser, cfg.PgPassword, cfg.PgAddr, cfg.PgDb)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Создание таблицы, если не была создана вручную заранее
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS notes (
		id SERIAL PRIMARY KEY,
		content TEXT NOT NULL
	);`

	_, err = db.Exec(createTableQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to create table: %v", err)
	} else {
		fmt.Println("Table 'notes' is alive.")
	}
	return db, nil
}
