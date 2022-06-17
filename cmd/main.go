package main

import (
	"fmt"
	"github.com/joho/godotenv"
	postgres "github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/config"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/server"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {

	router := myRouter.NewRouter()

	if err := config.InitConfig(); err != nil {
		panic(fmt.Errorf("cfg fatal. %s", err.Error()))
	}

	if err := godotenv.Load(); err != nil {
		panic(fmt.Errorf("env fatal. %s", err.Error()))
	}

	dbConfig := &config.DbConfig{
		Host:         viper.GetString("db.host"),
		Port:         viper.GetString("db.port"),
		User:         viper.GetString("db.user"),
		DatabaseName: viper.GetString("db.name"),
		Driver:       viper.GetString("db.driver"),
		Password:     os.Getenv("DB_PASSWORD"),
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

	port := viper.GetInt("app.port")
	if err := srv.StartListeningOn(port); err != nil {
		log.Fatalf("error occured running server on port %d \n", port)
	}

}
