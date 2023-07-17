package model

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"time"
)

type RequestCardModel struct {
	Number      string    `json:"number"`
	CVV         string    `json:"cvv"`
	Date        string    `json:"date"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
}

type ResponseCardModel struct {
	ID          uuid.UUID `json:"id"`
	Number      string    `json:"number"`
	CVV         string    `json:"cvv"`
	Date        string    `json:"date"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
}

type Card struct {
	client *http.Client
	url    *url.URL
}

func NewCard(client *http.Client) *Card {
	return &Card{client: client}
}

func (c *Card) SetCard(number, cvv, date string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	u := c.url.JoinPath(CardApi)

	data, err := json.Marshal(RequestCardModel{
		Number:      number,
		CVV:         cvv,
		Date:        date,
		UserID:      userID,
		CliCreation: cliCreation,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userId uuid.UUID
	err = json.NewDecoder(resp.Body).Decode(&userId)
	if err != nil {
		return nil, err
	}

	return &userId, nil
}

func (c *Card) GetCards(userID uuid.UUID) ([]ResponseCardModel, error) {
	u := c.url.JoinPath(CardApi)

	type us struct {
		UserID uuid.UUID `json:"user_id"`
	}

	u1 := us{UserID: userID}
	data, err := json.Marshal(u1)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		http.MethodGet,
		u.String(),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var models []ResponseCardModel
	err = json.NewDecoder(resp.Body).Decode(&models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (c *Card) GetCard(ID uuid.UUID) (*ResponseCardModel, error) {
	u := c.url.JoinPath(CardApi)

	type us struct {
		ID uuid.UUID `json:"user_id"`
	}

	u1 := us{ID: ID}
	data, err := json.Marshal(u1)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(
		http.MethodGet,
		u.String(),
		bytes.NewBuffer(data),
	)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var model ResponseCardModel
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
