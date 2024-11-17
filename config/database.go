package config

import (
	"database/sql"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3", "store.db")
	if err != nil {
		panic(err)
	}

	schemaSQL := `CREATE TABLE IF NOT EXISTS Tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title VARCHAR(25),
	    description VARCHAR(255),
		duedate TIMESTAMP,
		overdue BOOLEAN,
		completed BOOLEAN
	);`

	_, err = db.Exec(schemaSQL)
	if err != nil {
		panic(err)
	}

	slog.Info("Conncted with sqllite")
	return db
}
