package usecase

import (
	"context"

	"gateway/internal/domain/card"
	"gateway/internal/domain/card/datasource/database"
	"gateway/internal/models"
)

type UseCase interface {
	Create(ctx context.Context, card card.Request) (*models.Card, error)
	List(ctx context.Context, f *card.Filter) ([]*models.Card, error)
	Read(ctx context.Context, cardID uint64) (*models.Card, error)
	Update(ctx context.Context, card *models.Card) (*models.Card, error)
	Delete(ctx context.Context, cardID uint64) error
}

type useCase struct {
	data database.Mongodb
}

func New(datasource database.Mongodb) *useCase {
	return &useCase{
		data: datasource,
	}
}

func (u *useCase) Create(ctx context.Context, r card.Request) (*models.Card, error) {
	dto := card.ToCard(&r)
	cardFound, err := u.data.CreateRead(ctx, dto)
	if err != nil {
		return nil, err
	}

	return cardFound, nil
}

func (u *useCase) List(ctx context.Context, f *card.Filter) ([]*models.Card, error) {
	return u.data.List(ctx, f)
}

func (u *useCase) Read(ctx context.Context, cardID uint64) (*models.Card, error) {
	return u.data.Read(ctx, cardID)
}

func (u *useCase) Update(ctx context.Context, card *models.Card) (*models.Card, error) {
	err := u.data.Update(ctx, card)
	if err != nil {
		return nil, err
	}
	return u.data.Read(ctx, card.ID)
}

func (u *useCase) Delete(ctx context.Context, cardID uint64) error {
	return u.data.Delete(ctx, cardID)
}
