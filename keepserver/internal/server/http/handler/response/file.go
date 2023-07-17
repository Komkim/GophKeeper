package response

import (
	"github.com/go-chi/render"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type File struct {
	Name        string    `json:"name"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
}

func (f File) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewFileList(files []*File) []render.Renderer {
	list := []render.Renderer{}
	for _, file := range files {
		list = append(list, file)
	}
	return list
}
