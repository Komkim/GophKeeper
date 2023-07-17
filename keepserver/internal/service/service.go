package service

import (
	"GophKeeper/keepserver/internal/model"
	"github.com/google/uuid"
	"time"
)

type Service interface {
	User
	File
	Card
	Note
}

type User interface {
	SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error)
	GetUser(login string) (*model.UserModel, error)
}

type File interface {
	SetFile(name string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetFiles(userID uuid.UUID) ([]model.FileModel, error)
	GetFile(ID uuid.UUID) (*model.FileModel, error)
}

type Card interface {
	SetCard(number, cvv, date string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetCards(userID uuid.UUID) ([]model.CardModel, error)
	GetCard(ID uuid.UUID) (*model.CardModel, error)
}

type Note interface {
	SetNote(note string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetNotes(userID uuid.UUID) ([]model.NoteModel, error)
	GetNote(ID uuid.UUID) (*model.NoteModel, error)
}
