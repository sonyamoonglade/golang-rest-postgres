package main

import (
	"fmt"
	postgres "github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/config"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/server"
	"github.com/spf13/viper"
	"log"
)

func main() {

	router := myRouter.NewRouter()

	if err := config.InitConfig(); err != nil {
		panic(fmt.Errorf("cfg fatal. %s", err.Error()))
	}

	dbConfig := &config.DbConfig{
		Host:         viper.GetString("host"),
		Port:         viper.GetString("port"),
		User:         viper.GetString("user"),
		DatabaseName: viper.GetString("database_name"),
		Driver:       viper.GetString("driver"),
		Password:     viper.GetString("password"),
	}

	db, err := postgres.GetDbInstance(dbConfig)
	if err != nil {
		panic(fmt.Errorf("error occured while connecting to database. %s", err.Error()))
	}

	storages := storage.NewStorages(db)
	services := service.NewServices(storages)
	handlers := server.NewHandlers(router, services)

	srv, err := server.NewServer(handlers)
	if err != nil {
		log.Fatalf("error occured creating new server with handlers - %v", handlers)
	}

	port := viper.GetInt("server_port")
	if err := srv.StartListeningOn(port); err != nil {
		log.Fatalf("error occured running server on port %d \n", port)
	}

}
