package main

import (
	"log"
	"github.com/caevv/simple-go-prepaid-card/service"
	"github.com/caevv/simple-go-prepaid-card/env"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal(r)
		}
	}()

	env.Init()

	service.Start()
}
