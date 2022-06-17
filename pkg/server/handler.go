package server

import (
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	v1 "github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/v1"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
)

type Handler struct {
	*myRouter.Router
	v1.Car
}

func NewHandlers(router *myRouter.Router, services *service.Service) *Handler {

	car := v1.NewCarHandler(services.Car)
	car.InitRoutes(router)

	return &Handler{
		Router: router,
		Car:    car,
	}
}
