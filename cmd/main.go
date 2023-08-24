package main

import (
	"github.com/milmenderov/todolist-app"
	"log"
	"todolist-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todolist_app.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occurred while running http server: %s", err.Error())
	}
}
