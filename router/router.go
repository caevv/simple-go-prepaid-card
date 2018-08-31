package router

import (
	"context"
	"github.com/caevv/simple-go-prepaid-card/api"
	"github.com/caevv/simple-go-prepaid-card/data"
	"log"
	"github.com/pkg/errors"
)

type Router struct {
	repository *data.Repository
}

func New(r *data.Repository) api.PrepaidCardServer {
	return &Router{repository: r}
}

func (router Router) TopUp(ctx context.Context, r *api.TopUpRequest) (*api.Balance, error) {
	card := data.Card{
		ID: r.CardID,
		Amount: r.Amount.Value,
	}
	newBalance, err := router.repository.TopUp(card, data.Amount{Value: r.Amount.Value})
	if err != nil {
		return nil, errors.Wrap(err, "could not top-up account")
	}

	log.Printf("your new balance is: %e", newBalance.Amount)
	return &api.Balance{Amount: newBalance.Amount}, nil
}

func (Router) Authorisation(ctx context.Context, r *api.AuthorisationRequest) (*api.AuthoriseResponse, error) {
	// TODO
	return &api.AuthoriseResponse{}, nil
}
