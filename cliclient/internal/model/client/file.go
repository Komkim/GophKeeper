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

type ResponseFileModel struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
	CreateAt    time.Time `json:"create_at"`
}

type RequestFileModel struct {
	Name        string    `json:"name"`
	UserID      uuid.UUID `json:"user_id"`
	CliCreation time.Time `json:"cli_creation"`
	CreateAt    time.Time `json:"create_at"`
}
type File struct {
	client *http.Client
	url    *url.URL
}

func NewFile(client *http.Client) *File {
	return &File{client: client}
}

func (f *File) SetFile(name string, userID uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	u := f.url.JoinPath(model.FileApi)

	data, err := json.Marshal(RequestFileModel{
		Name:        name,
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
	resp, err := f.client.Do(req)
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

func (f *File) GetFiles(userID uuid.UUID) ([]ResponseFileModel, error) {
	u := f.url.JoinPath(model.FileApi)

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
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var models []ResponseFileModel
	err = json.NewDecoder(resp.Body).Decode(&models)
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (f *File) GetFile(ID uuid.UUID) (*ResponseFileModel, error) {
	u := f.url.JoinPath(model.FileApi)

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
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var model ResponseFileModel
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
