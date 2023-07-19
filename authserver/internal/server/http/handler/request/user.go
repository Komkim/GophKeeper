package request

import (
	"net/http"
	"time"
)

type User struct {
	Login       string    `json:"name" validate:"required"`
	Password    string    `json:"password" validate:"required,min=8"`
	CliCreation time.Time `json:"cliCreation" validate:"required"`
	//PasswordConfirm string `json:"passwordConfirm" validate:"required,min=8"`
}

func (c *User) Bind(r *http.Request) error {
	return nil
}
