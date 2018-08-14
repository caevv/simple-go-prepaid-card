package router

import (
		"context"
	"github.com/caevv/simple-go-mortgage-investing/api"
)

type Router struct {
}

func New() api.PrepaidCardServer {
	return &Router{}
}

func (Router) TopUp(ctx context.Context, in *api.Amount) (*api.Balance, error) {
	return &api.Balance{}, nil
}
