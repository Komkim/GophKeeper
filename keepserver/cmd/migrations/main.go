package main

import (
	"database/sql"

	_ "GophKeeper/keepserver/storage/postgres/migrations"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

func main() {
	pgDSN := "postgres://postgres:postgres@localhost:5432/keep?sslmode=disable"

	db, err := sql.Open("postgres", pgDSN)
	if err != nil {
		panic(err)
	}

	err = goose.Up(db, "/var")
	if err != nil {
		panic(err)
	}
}
