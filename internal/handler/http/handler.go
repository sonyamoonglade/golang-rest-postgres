package handler

import (
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/v1"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
)

type Handler struct {
	*myRouter.Router
	v1.Car
	v1.User
}

func NewHandlers(router *myRouter.Router, services *service.Service, logger *myLogger.Logger) *Handler {
	car := v1.NewCarHandler(services.Car, logger)
	car.InitRoutes(router)

	user := v1.NewUserHandler(services.User, logger)
	user.InitRoutes(router)

	garage := v1.NewGarageHandler(services.Garage, logger)
	garage.InitRoutes(router)

	router.LogAllRoutes()

	return &Handler{
		Router: router,
		Car:    car,
		User:   user,
	}
}
