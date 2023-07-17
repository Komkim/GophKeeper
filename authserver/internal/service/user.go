package service

import (
	"authserver/internal/model"
	"github.com/redis/go-redis/v9"
)

type UserService struct {
	Repo model.UserRepo
}

func NewUserService(db *redis.Client) *UserService {
	return
}
