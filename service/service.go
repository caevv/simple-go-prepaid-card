package service

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"github.com/caevv/simple-go-prepaid-card/env"
	"github.com/caevv/simple-go-prepaid-card/router"
	"github.com/caevv/simple-go-prepaid-card/api"
	"log"
)

func Start() {
	startService(router.New())
}

func startService(router api.PrepaidCardServer) {
	log.Print("service started")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Settings.GRPCPort))

	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()

	api.RegisterPrepaidCardServer(grpcServer, router)

	err = grpcServer.Serve(lis)

	if err != nil {
		panic(err)
	}
}
