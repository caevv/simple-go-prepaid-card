package service

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"github.com/caevv/simple-go-mortgage-investing/env"
	"github.com/caevv/simple-go-mortgage-investing/router"
	"github.com/caevv/simple-go-mortgage-investing/api"
)

func Start() {
	go startService(router.New())
}

func startService(router api.PrepaidCardServer) {
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
