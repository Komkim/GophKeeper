package model

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type UserModel struct {
	ID          uuid.UUID `db:"id"`
	Password    string    `db:"password"`
	Login       string    `db:"login"`
	CliCreation time.Time `db:"cli_creation"`
}

type User struct {
	client *http.Client
}

func NewUser(client *http.Client) *User {
	return &User{client: client}
}
