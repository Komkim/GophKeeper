package service

import (
	"authserver/internal/model"
	"github.com/redis/go-redis/v9"
)

type AuthService struct {
	Repo model.AuthRepo
}

func NewAuthService(db *redis.Client) Auth {
	return &AuthService{model.NewAuth(db)}
}

func (a *AuthService) GetToken(tokenUuid string) (string, error) {
	return a.Repo.GetToken(tokenUuid)
}

func (a *AuthService) SetToken(tokenUuid, userID string, tokenExpiresIn *int64) error {
	return a.Repo.SetToken(tokenUuid, userID, tokenExpiresIn)
}

func (a *AuthService) DelToken(tokenUuid string) error {
	return a.Repo.DelToken(tokenUuid)
}
