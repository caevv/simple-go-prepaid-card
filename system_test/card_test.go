package system

import (
	"github.com/DATA-DOG/godog"
	"github.com/caevv/simple-go-prepaid-card/data"
	"github.com/caevv/simple-go-prepaid-card/api"
	"github.com/smartystreets/assertions"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"strconv"
	"context"
)

var (
	card       data.Card
	svc        api.PrepaidCardClient
	newBalance *api.Balance
)

func init() {
	apiConn, err := grpc.Dial("localhost:8100", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	svc = api.NewPrepaidCardClient(apiConn)
}

func iTopupForAnAmountOf(amount string) error {
	value, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		return errors.Wrap(err, "failed to convert balance")
	}

	newBalance, err = svc.TopUp(
		context.Background(),
		&api.TopUpRequest{CardID: "123", Amount: &api.Amount{float32(value)}},
	)
	if err != nil {
		return err
	}

	return nil
}

func iShouldHaveABalanceOf(balance string) error {
	value, err := strconv.ParseFloat(balance, 32)
	if err != nil {
		return errors.Wrap(err, "failed to convert balance")
	}

	if ok, message := assertions.So(newBalance.Amount, assertions.ShouldEqual, float32(value)); !ok {
		return errors.New(message)
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I top-up for an amount of "([^"]*)"$`, iTopupForAnAmountOf)
	s.Step(`^I should have a balance of "([^"]*)"$`, iShouldHaveABalanceOf)
}
