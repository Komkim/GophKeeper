package model

import (
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

type CardRepo interface {
	SetCard(number, cvv, date string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetCards(userID uuid.UUID) ([]ResponseCardModel, error)
	GetCard(ID uuid.UUID) (*ResponseCardModel, error)
}

type FileRepo interface {
	SetFile(name string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetFiles(userID uuid.UUID) ([]ResponseFileModel, error)
	GetFile(ID uuid.UUID) (*ResponseFileModel, error)
}

type NoteRepo interface {
	SetNote(note string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error)
	GetNotes(userID uuid.UUID) ([]ResponseNoteModel, error)
	GetNote(ID uuid.UUID) (*ResponseNoteModel, error)
}

type UserRepo interface {
	SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error)
	GetUser(ID uuid.UUID) (*ResponseUserModel, error)
}
