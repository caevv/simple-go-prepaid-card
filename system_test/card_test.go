package system

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/caevv/simple-go-prepaid-card/api"
	"github.com/caevv/simple-go-prepaid-card/data"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/pkg/errors"
	"github.com/smartystreets/assertions"
	"google.golang.org/grpc"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	card       data.Card
	svc        api.PrepaidCardClient
	newBalance *api.Balance
	db         *gorm.DB
)

var opts = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opts)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opts.Paths = flag.Args()

	// godog v0.10.0 (latest)
	status := godog.TestSuite{
		Name:                 "godogs",
		TestSuiteInitializer: InitializeTestSuite,
		ScenarioInitializer:  InitializeScenario,
		Options:              &opts,
	}.Run()

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func InitializeTestSuite(ctx *godog.TestSuiteContext) {
	ctx.BeforeSuite(func() {
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
	})
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.BeforeScenario(func(*godog.Scenario) {
		if _, err := db.DB().Exec("TRUNCATE cards CASCADE"); err != nil {
			log.Fatal(err)
		}
	})

	ctx.Step(`^I have a card "([^"]*)" with balance of "([^"]*)"$`, iHaveACardWithBalanceOf)
	ctx.Step(`^I top-up for an amount of "([^"]*)"$`, iTopupForAnAmountOf)
	ctx.Step(`^my card "([^"]*)" should have a balance of "([^"]*)"$`, myCardShouldHaveABalanceOf)
}

func iHaveACardWithBalanceOf(cardID, balanceAmount string) error {
	value, err := strconv.ParseInt(balanceAmount, 10, 64)
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
	value, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to convert balance")
	}

	newBalance, err = svc.TopUp(
		context.Background(),
		&api.TopUpRequest{CardID: card.ID, Amount: &api.Amount{Value: value}},
	)
	if err != nil {
		return err
	}

	return nil
}

func myCardShouldHaveABalanceOf(cardId string, balance string) error {
	value, err := strconv.ParseInt(balance, 10, 64)
	if err != nil {
		return errors.Wrap(err, "failed to convert balance")
	}

	if ok, message := assertions.So(newBalance.Amount, assertions.ShouldEqual, value); !ok {
		return errors.New(message)
	}

	var card data.Card
	err = db.First(&card, cardId).Error
	if err != nil {
		return err
	}

	if ok, message := assertions.So(card.Amount, assertions.ShouldEqual, value); !ok {
		return errors.New(message)
	}

	return nil
}
