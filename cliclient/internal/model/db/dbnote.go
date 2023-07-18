package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type DBNoteModel struct {
	Note        string    `db:"note"`
	UserID      uuid.UUID `db:"user_id"`
	CliCreation time.Time `db:"cli_creation"`
}

type DBNote struct {
	db *sql.DB
}

func NewDBNote(db *sql.DB) *DBNote {
	return &DBNote{db: db}
}

func (n *DBNote) SetDBNote(note string, userId uuid.UUID, cliCreation time.Time) error {
	sqlStatement := `
		insert into note (note, user_id, cli_creation)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := n.db.QueryRow(sqlStatement, note, userId, cliCreation).Scan(&id)
	if err != nil {
		return err
	}
	return nil
}

func (n *DBNote) GetDBNotes() ([]DBNoteModel, error) {
	var notes []DBNoteModel
	rows, err := n.db.Query(
		`
			select id, note, user_id, cli_creation, create_at
			from note
			where sent = false
			order by cli_creation asc;
`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		note := DBNoteModel{}
		err := rows.Scan(
			&note.Note,
			&note.UserID,
			&note.CliCreation,
		)
		if err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return notes, nil
}
