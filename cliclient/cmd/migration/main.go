package main

import (
	"database/sql"
	"github.com/pressly/goose/v3"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "client.db")

	if err != nil {
		panic(err)
	}

	err = goose.Up(db, "storage/sqlite3/migrations/")
	if err != nil {
		panic(err)
	}
}
