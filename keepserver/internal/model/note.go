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

type NoteModel struct {
	ID          uuid.UUID `db:"id"`
	Note        string    `db:"note"`
	UserID      uuid.UUID `db:"user_id"`
	CliCreation time.Time `db:"cli_creation"`
	CreateAt    time.Time `db:"create_at"`
}

type Note struct {
	db *pgxpool.Pool
}

func NewNote(db *pgxpool.Pool) *Note {
	return &Note{db: db}
}

func (n *Note) SetNote(note string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into note (note, user_id, cli_creation)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := n.db.QueryRow(ctx, sqlStatement, note, userId, cliCreation).Scan(&id)
	if err != nil {
		return nil, err
	}

	if id.ID() < 1 {
		return nil, mistake.ErrDBID
	}

	return &id, nil
}

func (n *Note) GetNotes(userID uuid.UUID) ([]NoteModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var notes []NoteModel
	rows, err := n.db.Query(ctx,
		`
			select id, note, user_id, cli_creation, create_at
			from note
			where user_id  = $1 and deleted = false
			order by cli_creation asc;
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
		note := NoteModel{}
		err := rows.Scan(
			&note.ID,
			&note.Note,
			&note.UserID,
			&note.CliCreation,
			&note.CreateAt,
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

func (n *Note) GetNote(ID uuid.UUID) (*NoteModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var note NoteModel
	err := n.db.QueryRow(ctx,
		`select id, note, user_id, cri_creation, create_at
			 from note where id = $1 and deleted = false
    		 order by cri_creation desc limit 1;`,
		ID,
	).Scan(
		&note.ID,
		&note.Note,
		&note.UserID,
		&note.CliCreation,
		&note.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &note, nil
}
