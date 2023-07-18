package service

import (
	"GophKeeper/keepserver/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type UserService struct {
	Repo model.UserRepo
}

func NewUserService(db *pgxpool.Pool) User {
	return &UserService{model.NewUser(db)}
}

func (u *UserService) SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error) {
	return u.Repo.SetUser(login, password, cliCreation)
}

func (u *UserService) GetUser(login string) (*model.UserModel, error) {
	return u.Repo.GetUser(login)
}
