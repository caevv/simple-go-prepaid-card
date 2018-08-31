package main

import (
		"github.com/caevv/simple-go-prepaid-card/service"
)

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		log.Fatal(r)
	// 	}
	// }()

	service.Start()
}
