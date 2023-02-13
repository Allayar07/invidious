package main

import (
	"invidious/internal/handler"
	"invidious/internal/service"
	"log"
)

func main() {
	newService := service.NewService()
	newHandler := handler.NewHandler(newService)

	app := newHandler.InitRoutes()

	if err := app.Listen(":8000"); err != nil {
		log.Fatalln(err)
	}

}
