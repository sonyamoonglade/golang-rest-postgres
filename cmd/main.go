package main

import (
	"fmt"
	"github.com/joho/godotenv"
	postgres "github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/adapters/db/storage"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/config"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/server"
	"github.com/spf13/viper"
	"io"
	"os"
)

func main() {

	logger := myLogger.NewLogger([]io.Writer{os.Stdout})
	router := myRouter.NewRouter(logger)

	if err := config.InitConfig(); err != nil {
		logger.PrintWithCrit(fmt.Sprintf("error when reading config. %s", err.Error()))
		panic(fmt.Errorf("cfg fatal. %s", err.Error()))
	}
	if err := godotenv.Load(); err != nil {
		logger.PrintWithCrit(fmt.Sprintf("error occured loading env variables. %s", err.Error()))
		panic(fmt.Errorf("env fatal. %s", err.Error()))
	}
	logger.PrintWithInf("initialized config successfully")

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
		logger.PrintWithCrit(fmt.Sprintf("error occured while connecting to database. %s", err.Error()))
		panic(fmt.Errorf("error occured while connecting to database. %s", err.Error()))
	}
	logger.PrintWithInf("db has connected successfully")

	storages := storage.NewStorages(db, logger)
	services := service.NewServices(storages)
	handlers := handler.NewHandlers(router, services, logger)
	logger.PrintWithInf("initialized dependencies")
	router.LogAllRoutes()

	srv, err := server.NewServer(handlers)
	logger.PrintWithInf("created new server")

	if err != nil {
		logger.PrintWithCrit(fmt.Sprintf("error occured creating new server with handlers - %v", handlers))
		panic("server fatal")
	}

	port := viper.GetInt("app.port")
	logger.PrintWithInf(fmt.Sprintf("listening on localhost:%d", port))

	if err := srv.StartListeningOn(port); err != nil {
		logger.PrintWithCrit(fmt.Sprintf("error occured running server on port %d", port))
	}

}
