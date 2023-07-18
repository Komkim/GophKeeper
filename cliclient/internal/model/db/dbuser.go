package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type DBUserModel struct {
	Password    string    `db:"password"`
	Login       string    `db:"login"`
	CliCreation time.Time `db:"cli_creation"`
}

type DBUser struct {
	db *sql.DB
}

func NewDBUser(db *sql.DB) *DBUser {
	return &DBUser{db: db}
}

func (u *DBUser) SetDBUser(login, password string, cliCreation time.Time) error {
	sqlStatement := `
		insert into user (login, password, cli_creation)
		values ($1, $2, $3, $4)
		returning id `
	var id uuid.UUID
	err := u.db.QueryRow(sqlStatement, login, password, cliCreation).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (u *DBUser) GetDBUsers() ([]DBUserModel, error) {
	var users []DBUserModel
	rows, err := u.db.Query(
		`select id, login, password, cri_creation, create_at
			 from user where sent = false
    		 order by cri_creation asc ;`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := DBUserModel{}
		err := rows.Scan(
			&user.Login,
			&user.Password,
			&user.CliCreation,
		)

		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return users, nil
}
