package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type DBFileModel struct {
	Name        string    `db:"name"`
	UserID      uuid.UUID `db:"user_id"`
	CliCreation time.Time `db:"cli_creation"`
}

type DBFile struct {
	db *sql.DB
}

func NewDBFile(db *sql.DB) *DBFile {
	return &DBFile{db: db}
}

func (f *DBFile) SetDBFile(name string, userId uuid.UUID, cliCreation time.Time) error {

	sqlStatement := `
		insert into file (name, user_id, cli_creation)
		values ($1, $2, $3)
		returning id `
	var id uuid.UUID
	err := f.db.QueryRow(sqlStatement, name, userId, cliCreation).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

func (f *DBFile) GetDBFiles() ([]DBFileModel, error) {
	var files []DBFileModel
	rows, err := f.db.Query(
		`
			select id, name, user_id, cli_creation, create_at
			from file
			where sent = false
			order by cli_creation asc;
`,
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		file := DBFileModel{}
		err := rows.Scan(
			&file.Name,
			&file.UserID,
			&file.CliCreation,
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
