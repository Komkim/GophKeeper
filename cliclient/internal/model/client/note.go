package model

import (
	"bytes"
	"cliclient/internal/model"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"time"
)

type RequestNoteModel struct {
	Note        string    `json:"note"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
	CreateAt    time.Time `json:"create_at"`
}

type ResponseNoteModel struct {
	ID          uuid.UUID `json:"id"`
	Note        string    `json:"note"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
	CreateAt    time.Time `json:"create_at"`
}

type Note struct {
	client *http.Client
	url    *url.URL
}

func NewNote(client *http.Client) *Note {
	return &Note{client: client}
}

func (n *Note) SetNote(note string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	u := n.url.JoinPath(model.NoteApi)

	data, err := json.Marshal(RequestNoteModel{
		Note:        note,
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
	resp, err := n.client.Do(req)
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

func (n *Note) GetNotes(userID uuid.UUID) ([]ResponseNoteModel, error) {
	u := n.url.JoinPath(model.NoteApi)

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
	resp, err := n.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var models []ResponseNoteModel
	err = json.NewDecoder(resp.Body).Decode(&models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (n *Note) GetNote(ID uuid.UUID) (*ResponseNoteModel, error) {
	u := n.url.JoinPath(model.NoteApi)

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
	resp, err := n.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var model ResponseNoteModel
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
