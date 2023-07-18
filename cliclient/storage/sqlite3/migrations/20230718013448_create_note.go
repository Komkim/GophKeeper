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
				note	         varchar(400) not null,
				user_id			 uuid not null,
				cli_creation	 timestamp,
				sent			 boolean default false,
				create_at        timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateNote(tx *sql.Tx) error {
	_, err := tx.Exec("drop table note")
	return err
}
