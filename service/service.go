package service

import (
	"fmt"
	"log"
	"net"

	"github.com/caevv/simple-go-prepaid-card/data"
	"gorm.io/driver/postgres"

	"github.com/caevv/simple-go-prepaid-card/api"
	"github.com/caevv/simple-go-prepaid-card/env"
	"github.com/caevv/simple-go-prepaid-card/repository"
	"github.com/caevv/simple-go-prepaid-card/router"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func Start() {
	db, err := gorm.Open(postgres.Open(env.Settings.DBAddress), &gorm.Config{})
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open connection"))
	}

	if err := applySchemaMigration(db); err != nil {
		log.Fatal(errors.Wrap(err, "failed to apply schema migration"))
	}

	r, err := repository.New(db)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to create new repository"))
	}

	startService(router.New(r))
}

func startService(router api.PrepaidCardServer) {
	log.Print("service started")

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

func applySchemaMigration(db *gorm.DB) error {
	err := db.AutoMigrate(&data.Card{})
	if err != nil {
		return errors.Wrap(err, "failed auto migrate")
	}

	return nil
}
