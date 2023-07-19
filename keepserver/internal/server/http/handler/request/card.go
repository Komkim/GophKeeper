package request

import (
	"net/http"
	"time"
)

type Card struct {
	Number      string    `json:"number"`
	CVV         string    `json:"cvv"`
	Date        string    `json:"date"`
	CliCreation time.Time `json:"cli_creation"`
}

func (c *Card) Bind(r *http.Request) error {
	return nil
}
