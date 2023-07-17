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

type FileModel struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	UserID      uuid.UUID `db:"user_id"`
	CliCreation time.Time `db:"cli_creation"`
	CreateAt    time.Time `db:"create_at"`
}

type File struct {
	db *pgxpool.Pool
}

func NewFile(db *pgxpool.Pool) *File {
	return &File{db: db}
}

func (f *File) SetFile(name string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	sqlStatement := `
		insert into file (name, user_id, cli_creation)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := f.db.QueryRow(ctx, sqlStatement, name, userId, cliCreation).Scan(&id)
	if err != nil {
		return nil, err
	}

	if id.ID() < 1 {
		return nil, mistake.ErrDBID
	}

	return &id, nil
}

func (f *File) GetFiles(userID uuid.UUID) ([]FileModel, error) {

	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var files []FileModel
	rows, err := f.db.Query(ctx,
		`
			select id, name, user_id, cli_creation, create_at
			from file
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
		file := FileModel{}
		err := rows.Scan(
			&file.ID,
			&file.Name,
			&file.UserID,
			&file.CliCreation,
			&file.CreateAt,
		)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (f *File) GetFile(ID uuid.UUID) (*FileModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	var file FileModel
	err := f.db.QueryRow(ctx,
		`select id, name, user_id, cri_creation, create_at
			 from file where id = $1 and deleted = false
    		 order by cri_creation desc limit 1;`,
		ID,
	).Scan(
		&file.ID,
		&file.Name,
		&file.UserID,
		&file.CliCreation,
		&file.CreateAt,
	)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &file, nil
}
