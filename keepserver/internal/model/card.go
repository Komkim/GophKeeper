package model

import (
	"GophKeeper/keepserver/internal/mistake"
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type CardModel struct {
	ID          uuid.UUID `db:"id"`
	Number      string    `db:"number"`
	CVV         string    `db:"cvv"`
	Date        string    `db:"date"`
	UserID      uuid.UUID `db:"user_id"`
	CliCreation time.Time `db:"cli_creation"`
	CreateAt    time.Time `db:"create_at"`
}

type Card struct {
	db *pgxpool.Pool
}

func NewCard(db *pgxpool.Pool) *Card {
	return &Card{db: db}
}

func (c *Card) SetCard(number, cvv, date string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into card (number, cvv, date, user_id, cli_creation)
		values ($1, $2, $3, $4, $5)
		returning id `
	var id uuid.UUID
	err := c.db.QueryRow(ctx, sqlStatement, number, cvv, date, userId, cliCreation).Scan(&id)
	if err != nil {
		return nil, err
	}

	if id.ID() < 1 {
		return nil, mistake.ErrDBID
	}

	return &id, nil
}

func (c *Card) GetCards(userID uuid.UUID) ([]CardModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var cards []CardModel
	rows, err := c.db.Query(ctx,
		`
			select id, number, cvv, date, user_id, cri_creation, create_at
			from card
			where user_id  = $1 and deleted = false
			order by cri_creation asc;
`,
		userID,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		card := CardModel{}
		err := rows.Scan(
			&card.ID,
			&card.Number,
			&card.CVV,
			&card.Date,
			&card.UserID,
			&card.CliCreation,
			&card.CreateAt,
		)
		if err != nil {
			return nil, err
		}
		cards = append(cards, card)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return cards, nil
}

func (c *Card) GetCard(ID uuid.UUID) (*CardModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var card CardModel
	err := c.db.QueryRow(ctx,
		`select id, number, cvv, date, user_id, cri_creation, create_at
			 from card where id = $1 and deleted = false
    		 order by cri_creation desc limit 1;`,
		ID,
	).Scan(
		&card.ID,
		&card.Number,
		&card.CVV,
		&card.Date,
		&card.UserID,
		&card.CliCreation,
		&card.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &card, nil
}
