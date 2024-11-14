package config

import (
	"database/sql"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "test"
)

func DatabaseConnection() *sql.DB {
	db, err := sql.Open("sqlite3_extended", "./taks.db")
	if err != nil {
		panic(err)
	}

	return db
}

//docker run --name habr-pg-13.3 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=test -d postgres:13.3
