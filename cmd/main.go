package main

import (
	"invidious/internal/handler"
	"invidious/internal/model"
	"invidious/internal/repository"
	"log"
)

func main() {

	config, err := model.NewAppConfig("config/configs.yml")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := repository.NewPostgres(repository.Config{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		Username: config.DB.UserName,
		Password: config.DB.Password,
		DbName:   config.DB.DbName,
		SSLMode:  config.DB.SslMode,
	})
	if err != nil {
		log.Fatalln(err)
	}

	newRepository := repository.NewVideoRepository(db)
	//newService := service.NewService()
	newHandler := handler.NewHandler(newRepository)

	app := newHandler.InitRoutes()

	if err := app.Listen(":8000"); err != nil {
		log.Fatalln(err)
	}

}
