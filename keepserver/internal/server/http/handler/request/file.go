package request

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type File struct {
	Name        string    `json:"name"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
}

func (f *File) Bind(r *http.Request) error {
	return nil
}
