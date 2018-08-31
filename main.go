package main

import (
	"github.com/caevv/simple-go-prepaid-card/service"
	"runtime/debug"
	"log"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%s: %s", r, string(debug.Stack()))
		}
	}()
	service.Start()
}
