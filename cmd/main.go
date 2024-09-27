package main

import (
	"log"

	"github.com/Fazom/todo-appLearnRestAPI.git"
	"github.com/Fazom/todo-appLearnRestAPI.git/package/handler"
	"github.com/Fazom/todo-appLearnRestAPI.git/package/repository"
	"github.com/Fazom/todo-appLearnRestAPI.git/package/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Ошибка инициализации конфигурации: %v", err)
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
