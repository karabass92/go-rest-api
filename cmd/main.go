package main

import (
	"log"

	todo "github.com/karabass92/go-rest-api"
	"github.com/karabass92/go-rest-api/pkg/handler"
	"github.com/karabass92/go-rest-api/pkg/repository"
	"github.com/karabass92/go-rest-api/pkg/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
