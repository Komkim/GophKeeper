package model

import (
	client "cliclient/internal/model/client"
	db "cliclient/internal/model/db"
	"github.com/google/uuid"
	"time"
)

const (
	CardApi   = "/api/card"
	CardApiId = "/api/card/:id"
	FileApi   = "/api/file"
	FileApiId = "api/file/:id"
	NoteApi   = "/api/note"
	NoteApiId = "/api/note/:id"
	UserApi   = "/api/user"
	UserApiId = "/api/user/:id"
)

const DBTIMEOUT = 5

type CardRepo interface {
	SetCard(number, cvv, date string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetCards(userID uuid.UUID) ([]client.ResponseCardModel, error)
	GetCard(ID uuid.UUID) (*client.ResponseCardModel, error)
}

type FileRepo interface {
	SetFile(name string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetFiles(userID uuid.UUID) ([]client.ResponseFileModel, error)
	GetFile(ID uuid.UUID) (*client.ResponseFileModel, error)
}

type NoteRepo interface {
	SetNote(note string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetNotes(userID uuid.UUID) ([]client.ResponseNoteModel, error)
	GetNote(ID uuid.UUID) (*client.ResponseNoteModel, error)
}

type UserRepo interface {
	SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error)
	GetUser(ID uuid.UUID) (*client.ResponseUserModel, error)
}

type DBCardRepo interface {
	SetDBCard(number, cvv, date string, userId uuid.UUID, cliCreation time.Time) error
	GetDBCards() ([]db.DBCardModel, error)
}

type DBFileRepo interface {
	SetDBFile(name string, userId uuid.UUID, cliCreation time.Time) error
	GetDBFiles() ([]db.DBFileModel, error)
}

type DBNoteRepo interface {
	SetDBNote(note string, userId uuid.UUID, cliCreation time.Time) error
	GetDBNotes() ([]db.DBNoteModel, error)
}

type DBUserRepo interface {
	SetDBUser(login, password string, cliCreation time.Time) error
	GetDBUsers() ([]db.DBUserModel, error)
}
