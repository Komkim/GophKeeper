package response

import (
	"net/http"
	"time"
)

type User struct {
	Password    string    `json:"password"`
	Login       string    `json:"login"`
	CliCreation time.Time `json:"cli_creation"`
}

func (u User) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

var UserSuccessfully = &Response{HTTPStatusCode: 202, StatusText: "the user has been successfully added."}
