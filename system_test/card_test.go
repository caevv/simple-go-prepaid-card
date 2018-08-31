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
	"github.com/jinzhu/gorm"
	"log"
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	card       data.Card
	svc        api.PrepaidCardClient
	newBalance *api.Balance
	db         *gorm.DB
)

func init() {
	apiConn, err := grpc.Dial("localhost:8110", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	svc = api.NewPrepaidCardClient(apiConn)

	db, err = gorm.Open("postgres", fmt.Sprintf(
		"postgres://%v:%v@%v:%v/%v?sslmode=%s",
		"postgres",
		"postgres",
		"localhost",
		"5462",
		"postgres",
		"disable",
	))

	if err != nil {
		log.Fatal(err)
	}
}

func iHaveACardWithBalanceOf(cardID, balanceAmount string) error {
	value, err := strconv.ParseFloat(balanceAmount, 32)
	if err != nil {
		return errors.Wrap(err, "failed to convert balance")
	}

	err = db.Exec(`INSERT INTO cards(id, amount, blocked_amount) VALUES (?, ?, ?)`, cardID, value, 0).Error
	if err != nil {
		return err
	}

	return nil
}

func iTopupForAnAmountOf(amount string) error {
	value, err := strconv.ParseFloat(amount, 32)
	if err != nil {
		return errors.Wrap(err, "failed to convert balance")
	}

	newBalance, err = svc.TopUp(
		context.Background(),
		&api.TopUpRequest{CardID: card.ID, Amount: &api.Amount{float32(value)}},
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
	s.Step(`^I have a card "([^"]*)" with balance of "([^"]*)"$`, iHaveACardWithBalanceOf)
	s.Step(`^I top-up for an amount of "([^"]*)"$`, iTopupForAnAmountOf)
	s.Step(`^I should have a balance of "([^"]*)"$`, iShouldHaveABalanceOf)

	s.BeforeScenario(cleanDB)
}

func cleanDB(i interface{}) {
	db.DB().Exec("TRUNCATE cards CASCADE")
}
