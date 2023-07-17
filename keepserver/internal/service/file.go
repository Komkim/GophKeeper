package service

import (
	"GophKeeper/keepserver/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type FileService struct {
	Repo model.FileRepo
}

func NewFileService(db *pgxpool.Pool) FileService {
	return FileService{model.NewFile(db)}
}

func (f *FileService) SetFile(name string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	return f.Repo.SetFile(name, userId, cliCreation)
}

func (f *FileService) GetFiles(userID uuid.UUID) ([]model.FileModel, error) {
	return f.Repo.GetFiles(userID)
}

func (f *FileService) GetFile(ID uuid.UUID) (*model.FileModel, error) {
	return f.Repo.GetFile(ID)
}
