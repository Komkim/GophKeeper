package model

import (
	"GophKeeper/keepserver/internal/mistake"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserModel struct {
	ID          uuid.UUID `db:"id"`
	Password    string    `db:"password"`
	Login       string    `db:"login"`
	CliCreation time.Time `db:"cli_creation"`
	CreateAt    time.Time `db:"create_at"`
}

type User struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *User {
	return &User{db: db}
}

func (u *User) SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into user (login, password, cli_creation)
		values ($1, $2, $3, $4)
		returning id `
	var id uuid.UUID
	err := u.db.QueryRow(ctx, sqlStatement, login, password, cliCreation).Scan(&id)
	if err != nil {
		return nil, err
	}

	if id.ID() < 1 {
		return nil, mistake.ErrDBID
	}

	return &id, nil
}

func (u *User) GetUser(login string) (*UserModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var user UserModel
	err := u.db.QueryRow(ctx,
		`select id, login, password, cri_creation, create_at
			 from user where login = $1 and deleted = false
    		 order by cri_creation desc limit 1;`,
		login,
	).Scan(
		&user.ID,
		&user.Login,
		&user.Password,
		&user.CliCreation,
		&user.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &user, nil
}
