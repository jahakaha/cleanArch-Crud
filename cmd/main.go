package main

import (
	"log"

	"todoo/pkg/handler"
	"todoo/pkg/repository"
	"todoo/pkg/service"
	"todoo/server"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
