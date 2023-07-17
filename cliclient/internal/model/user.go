package model

import (
	"bytes"
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"net/url"
	"time"
)

type RequestUserModel struct {
	Password    string    `json:"password"`
	Login       string    `json:"login"`
	CliCreation time.Time `json:"cli_creation"`
	CreateAt    time.Time `json:"create_at"`
}

type ResponseUserModel struct {
	ID          uuid.UUID `json:"id"`
	Password    string    `json:"password"`
	Login       string    `json:"login"`
	CliCreation time.Time `json:"cli_creation"`
	CreateAt    time.Time `json:"create_at"`
}

type User struct {
	client *http.Client
	url    *url.URL
}

func NewUser(client *http.Client) *User {
	return &User{client: client}
}

func (s *User) SetUser(login, password string, cliCreation time.Time) (*uuid.UUID, error) {
	u := s.url.JoinPath(UserApi)

	data, err := json.Marshal(RequestUserModel{
		Login:       login,
		Password:    password,
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
	resp, err := s.client.Do(req)
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

func (s *User) GetUser(ID uuid.UUID) (*ResponseUserModel, error) {
	u := s.url.JoinPath(UserApi)

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
	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var model ResponseUserModel
	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		return nil, err
	}

	return &model, nil
}
