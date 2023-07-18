package service

import (
	"GophKeeper/keepserver/internal/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"time"
)

type CardService struct {
	Repo model.CardRepo
}

func NewCardService(db *pgxpool.Pool) Card {
	return &CardService{model.NewCard(db)}
}

func (c *CardService) SetCard(number, cvv, date string, userId uuid.UUID, cliCreation time.Time) (*uuid.UUID, error) {
	return c.Repo.SetCard(number, cvv, date, userId, cliCreation)
}

func (c *CardService) GetCards(userID uuid.UUID) ([]model.CardModel, error) {
	return c.Repo.GetCards(userID)
}

func (c *CardService) GetCard(ID uuid.UUID) (*model.CardModel, error) {
	return c.Repo.GetCard(ID)
}
