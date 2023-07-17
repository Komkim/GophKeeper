package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateNote, downCreateNote)
}

func upCreateNote(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table note
			(
				id               uuid default gen_random_uuid() not null primary key,
				note	         varchar(400) not null,
				user_id			 uuid REFERENCES users (id),
				cli_creation	 timestamp,
				deleted			 boolean default false,
				create_at        timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateNote(tx *sql.Tx) error {
	_, err := tx.Exec("drop table note")
	return err
}
