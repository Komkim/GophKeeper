package service

import (
	"GophKeeper/keepserver/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type NoteService struct {
	Repo model.NoteRepo
}

func NewNoteService(db *pgxpool.Pool) NoteService {
	return NoteService{model.NewNote(db)}
}

func (n *NoteService) SetNote(note string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	return n.Repo.SetNote(note, userId, cliCreation)
}

func (n *NoteService) GetNotes(userID uuid.UUID) ([]model.NoteModel, error) {
	return n.Repo.GetNotes(userID)
}

func (n *NoteService) GetNote(ID uuid.UUID) (*model.NoteModel, error) {
	return n.Repo.GetNote(ID)
}
