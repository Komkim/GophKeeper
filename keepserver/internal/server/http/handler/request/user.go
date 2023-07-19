package request

import (
	"net/http"
	"time"
)

type User struct {
	Password    string    `json:"password"`
	Login       string    `json:"login"`
	CliCreation time.Time `json:"cli_creation"`
}

func (u *User) Bind(r *http.Request) error {
	return nil
}
