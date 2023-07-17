package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateUsers, downCreateUsers)
}

func upCreateUsers(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table users
			(				
				login	         varchar(40) not null,
				password  varchar(500) not null,
				cli_creation	 timestamp,
				deleted			 boolean default false,
				create_at        timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateUsers(tx *sql.Tx) error {
	_, err := tx.Exec("drop table users")
	return err
}
