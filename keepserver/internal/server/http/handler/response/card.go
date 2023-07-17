package response

import (
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type Card struct {
	Number      string    `json:"number"`
	CVV         string    `json:"cvv"`
	Date        string    `json:"date"`
	CliCreation time.Time `json:"cli_creation"`
}

func (c Card) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCardList(cards []*Card) []render.Renderer {
	list := []render.Renderer{}
	for _, card := range cards {
		list = append(list, card)
	}
	return list
}
