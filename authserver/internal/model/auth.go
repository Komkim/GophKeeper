package model

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type AuthModel struct {
}

type Auth struct {
	db *redis.Client
}

func NewAuth(db *redis.Client) *Auth {
	return &Auth{db: db}
}

func (a *Auth) GetToken(tokenUuid string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	return a.db.Get(ctx, tokenUuid).Result()
}

func (a *Auth) SetToken(tokenUuid, userID string, tokenExpiresIn *int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	now := time.Now()
	return a.db.Set(ctx, tokenUuid, userID, time.Unix(*tokenExpiresIn, 0).Sub(now)).Err()
}

func (a *Auth) DelToken(tokenUuid string) error {
	ctx, cancel := context.WithTimeout(context.Background(), DBTIMEOUT*time.Second)
	defer cancel()

	return a.db.Del(ctx, tokenUuid).Err()
}
