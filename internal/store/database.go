package store

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func Open() (*sql.DB, error) {
	// later we will move to an .env file or a secret manager.
	// for now, because its a sample project, we are leaving it here
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable")

	if err != nil {
		return nil, fmt.Errorf("db: open %w", err)
	}

	// enhance db connection using this:
	// db.SetMaxOpenConns(), db.SetMaxIdleConns(), and db.SetConnMaxIdleTime()

	fmt.Println("Connected to database...")

	return db, nil
}
