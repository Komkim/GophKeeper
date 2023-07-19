package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateFile, downCreateFile)
}

func upCreateFile(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table file
			(
				id               uuid default gen_random_uuid() not null primary key,
				name	         varchar(40) not null,
				user_id			 uuid REFERENCES users (id),
				cli_creation	 timestamp,
				deleted			 boolean default false,
				create_at        timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateFile(tx *sql.Tx) error {
	_, err := tx.Exec("drop table file")
	return err
}
