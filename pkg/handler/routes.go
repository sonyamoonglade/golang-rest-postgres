package handler

import (
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/service"
)

type Controller struct {
	Services *service.Service
}

func CreateController(services *service.Service) *Controller {
	return &Controller{Services: services}
}

func (c *Controller) InitRoutes(r *myRouter.Router) {

	r.POST("/createCar", c.createCar)
	r.GET("/getCar", c.getCar)

}
