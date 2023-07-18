package model

import (
	"database/sql"
	"github.com/google/uuid"

	"time"
)

type DBCardModel struct {
	Number      string    `db:"number"`
	CVV         string    `db:"cvv"`
	Date        string    `db:"date"`
	UserID      uuid.UUID `db:"user_id"`
	CliCreation time.Time `db:"cli_creation"`
}

type DBCard struct {
	db *sql.DB
}

func NewDBCard(db *sql.DB) *DBCard {
	return &DBCard{db: db}
}

func (c *DBCard) SetDBCard(number, cvv, date string, userId uuid.UUID, cliCreation time.Time) error {
	sqlStatement := `
		insert into card (number, cvv, date, user_id, cli_creation)
		values ($1, $2, $3, $4, $5)
		returning id `
	var id uuid.UUID
	err := c.db.QueryRow(sqlStatement, number, cvv, date, userId, cliCreation).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (c *DBCard) GetDBCards() ([]DBCardModel, error) {

	var cards []DBCardModel
	rows, err := c.db.Query(
		`
			select id, number, cvv, date, user_id, cri_creation, create_at
			from card
			where sent = false
			order by cri_creation asc;
`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		card := DBCardModel{}
		err := rows.Scan(
			&card.Number,
			&card.CVV,
			&card.Date,
			&card.UserID,
			&card.CliCreation,
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
