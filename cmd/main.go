package main

import (
	"log"

	_go "devcortes.com/go"
	"devcortes.com/go/pkg/handler"
	"devcortes.com/go/repository"
	"devcortes.com/go/service"
)

func main() {
	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHanler(services)

	srv := new(_go.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running server %s", err.Error())
	}
}
