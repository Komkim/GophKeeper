package request

import (
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Note struct {
	Note        string    `json:"note"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
}

func (n *Note) Bind(r *http.Request) error {
	return nil
}
