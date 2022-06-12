package main

import (
	"encoding/json"
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/handler"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/repository"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/service"
	"io"
	"log"
	"net/http"
)

func main() {

	router := myRouter.NewRouter()

	repositories := repository.NewRepository()
	services := service.CreateService(repositories)
	controller := handler.CreateController(services)

	controller.InitRoutes(router)

	server, err := todo.NewServer(router)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	port := 5000
	if err := server.StartListeningOn(port); err != nil {
		log.Fatalf("error occured running server on port %d \n", port)
	}

}

func createCarHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(201)

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Fatalf("could read the body...")
		return
	}

	var data Car

	json.Unmarshal(body, &data)

	fmt.Println(data)

	response, _ := json.Marshal(data)

	w.Write(response)

}

func getCarHandler(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(200)

	w.Write([]byte("this is your car"))
	return
}

type Car struct {
	Model string `json:"model"`
	Year  int    `json:"year"`
}

func newCar(model string, year int) *Car {
	c := Car{
		Model: model,
		Year:  year,
	}
	return &c
}
