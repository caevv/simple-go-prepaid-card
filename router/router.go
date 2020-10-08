package router

import (
	"context"
	"log"

	"github.com/caevv/simple-go-prepaid-card/api"
	"github.com/caevv/simple-go-prepaid-card/repository"
	"github.com/pkg/errors"
)

type Router struct {
	repository *repository.Repository
}

func New(r *repository.Repository) api.PrepaidCardServer {
	return &Router{repository: r}
}

func (router Router) TopUp(ctx context.Context, r *api.TopUpRequest) (*api.Balance, error) {
	card, err := router.repository.TopUp(r.CardID, r.Amount.Value)
	if err != nil {
		return nil, errors.Wrap(err, "could not top-up account")
	}

	log.Printf("your new balance is: %d", card.Amount)
	return &api.Balance{Amount: card.Amount}, nil
}

func (Router) Authorisation(ctx context.Context, r *api.AuthorisationRequest) (*api.AuthoriseResponse, error) {
	// TODO
	return &api.AuthoriseResponse{}, nil
}
