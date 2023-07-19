package model

const DBTIMEOUT = 5

type AuthRepo interface {
	GetToken(tokenUuid string) (string, error)
	SetToken(tokenUuid, userID string, tokenExpiresIn *int64) error
	DelToken(tokenUuid string) error
}

type UserRepo interface {
}
