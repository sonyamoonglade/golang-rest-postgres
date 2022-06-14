package handler

import (
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/service"
	"net/http"
)

type Controller struct {
	User
	Router  *myRouter.Router
	Service service.Service
}

func NewController(router *myRouter.Router, services *service.Service) *Controller {

	userController := NewUserController(services.User)
	userController.InitRoutes(router)

	return &Controller{
		Router: router,
		User:   userController,
	}
}

type User interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
	InitRoutes(r *myRouter.Router)
}
