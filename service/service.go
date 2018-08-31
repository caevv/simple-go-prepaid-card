package service

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"github.com/caevv/simple-go-prepaid-card/env"
	"github.com/caevv/simple-go-prepaid-card/router"
	"github.com/caevv/simple-go-prepaid-card/api"
	"log"
	"github.com/jinzhu/gorm"

	"github.com/mattes/migrate"
	"github.com/caevv/simple-go-prepaid-card/data"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"
)

func Start() {
	r, err := data.New()
	if err != nil {
		log.Fatal(err)
	}
	startService(router.New(r))
}

func startService(router api.PrepaidCardServer) {
	log.Print("service started")

	if err := applySchemaMigration(); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Settings.GRPCPort))

	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterPrepaidCardServer(grpcServer, router)

	err = grpcServer.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}
}

func applySchemaMigration() error {
	db, err := gorm.Open("postgres", env.Settings.DBAddress)
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}

	m, err := migrate.New(
		"file://data/migrations",
		env.Settings.DBAddress,
	)

	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}
	_, _, err = m.Version()

	return err
}
