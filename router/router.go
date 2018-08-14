package router

import (
	"context"
	"github.com/caevv/simple-go-prepaid-card/api"
	"github.com/caevv/simple-go-prepaid-card/data"
	"log"
)

type Router struct {
}

func New() api.PrepaidCardServer {
	return &Router{}
}

func (Router) TopUp(ctx context.Context, r *api.TopUpRequest) (*api.Balance, error) {
	newBalance := data.TopUp(r.CardID, data.Amount{Value: r.Amount.Value})

	log.Printf("your new balance is: %e", newBalance)
	return &api.Balance{Amount: newBalance.Value}, nil
}
