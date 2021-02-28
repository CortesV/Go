package main

import (
	_go "devcortes.com/go"
	"devcortes.com/go/pkg/handler"
	"log"
)

func main() {
	srv := new(_go.Server)
	handlers := new(handler.Handler)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server %s", err.Error())
	}
}
