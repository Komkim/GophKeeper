package service

import (
	"authserver/internal/model"
	"github.com/redis/go-redis/v9"
	"net/http"
	"time"
)

type Service struct {
	Auth
	User
}

func NewService(db *redis.Client, client *http.Client) *Service {
	return &Service{
		Auth: NewAuthService(db),
		User: NewUserService(client),
	}
}

type Auth interface {
	GetToken(tokenUuid string) (string, error)
	SetToken(tokenUuid, userID string, tokenExpiresIn *int64) error
	DelToken(tokenUuid string) error
}

type User interface {
	CreateUser(login, password string, liCreation time.Time) *http.Response
	GetUserByLogin(login string) (*model.UserModel, error)
	GetUserById(userID string) (*model.UserModel, error)
}
