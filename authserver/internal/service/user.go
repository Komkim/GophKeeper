package service

import (
	"authserver/internal/model"
	"net/http"
	"time"
)

type UserService struct {
	Repo model.UserRepo
}

func NewUserService(client *http.Client) User {
	return &UserService{model.NewUser(client)}
}

func (u *UserService) CreateUser(login, password string, liCreation time.Time) *http.Response {
	return nil
}
func (u *UserService) GetUserByLogin(login string) (*model.UserModel, error) {
	return nil, nil
}
func (u *UserService) GetUserById(userID string) (*model.UserModel, error) {
	return nil, nil
}
