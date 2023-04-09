package main

import (
	"log"

	"github.com/spf13/viper"

	todo "github.com/ShawaDev/RestAPI"
	"github.com/ShawaDev/RestAPI/pkg/handler"
	"github.com/ShawaDev/RestAPI/pkg/repository"
	"github.com/ShawaDev/RestAPI/pkg/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error init config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService()
	handlers := handler.NewHandler()

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error while running app: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
