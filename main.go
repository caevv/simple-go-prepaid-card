package main

import (
	"log"
	"github.com/caevv/simple-go-mortgage-investing/service"
	"github.com/caevv/simple-go-mortgage-investing/env"
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
