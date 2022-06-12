package handler

import (
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
)

type Controller interface {
	initRoutes(r *myRouter.Router)
}

type CarController struct {
}

func (c *CarController) InitRoutes(r *myRouter.Router) {

	r.POST("/createCar", c.createCar)
	r.GET("/getCar", c.getCar)

}
