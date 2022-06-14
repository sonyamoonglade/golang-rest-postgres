package main

import (
	"github.com/sonyamoonglade/golang-rest-postgres"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/handler"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/repository"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/service"
	"log"
)

func main() {

	router := myRouter.NewRouter()

	db, err := repository.GetDbInstance(repository.DbConfig{
		DBName:   "golang",
		Dialect:  "postgres",
		Port:     "5432",
		Username: "postgres",
		Host:     "localhost",
		Password: "admin",
	})
	if err != nil {
		log.Fatalf("error connecting to database... %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.CreateService(repositories)
	controller := handler.NewController(router, services)

	server, err := todo.NewServer(controller)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	port := 5000
	if err := server.StartListeningOn(port); err != nil {
		log.Fatalf("error occured running server on port %d \n", port)
	}

}
