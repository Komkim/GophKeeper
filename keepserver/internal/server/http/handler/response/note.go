package response

import (
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Note struct {
	Note        string    `json:"note"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
}

func (n Note) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewNoteList(notes []*Note) []render.Renderer {
	list := []render.Renderer{}
	for _, note := range notes {
		list = append(list, note)
	}
	return list
}
