package model

import (
	"github.com/google/uuid"
	"time"
)

const DBTIMEOUT = 5

type FileRepo interface {
	SetFile(name string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetFiles(userID uuid.UUID) ([]FileModel, error)
	GetFile(ID uuid.UUID) (*FileModel, error)
}

type NoteRepo interface {
	SetNote(note string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetNotes(userID uuid.UUID) ([]NoteModel, error)
	GetNote(ID uuid.UUID) (*NoteModel, error)
}

type CardRepo interface {
	SetCard(number, cvv, date string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetCards(userID uuid.UUID) ([]CardModel, error)
	GetCard(ID uuid.UUID) (*CardModel, error)
}

type UserRepo interface {
	SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error)
	GetUser(login string) (*UserModel, error)
}
