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
				name	         varchar(40) not null,
				user_id			 uuid not null,
				cli_creation	 timestamp,
				sent			 boolean default false,				
			);
	`)
	return err
}

func downCreateFile(tx *sql.Tx) error {
	_, err := tx.Exec("drop table file")
	return err
}
