package migrations

import (
	"database/sql"
	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upCreateCard, downCreateCard)
}

func upCreateCard(tx *sql.Tx) error {
	_, err := tx.Exec(`
			create table card
			(				
				number	         varchar(40) not null,
				cvv  			 varchar(40) not null,
				date			 varchar(40) not null,
				user_id			 uuid not null,
				cli_creation	 timestamp,
				sent			 boolean default false,
				create_at        timestamp with time zone default current_timestamp
			);
	`)
	return err
}

func downCreateCard(tx *sql.Tx) error {
	_, err := tx.Exec("drop table card")
	return err
}
