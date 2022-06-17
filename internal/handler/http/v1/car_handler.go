package v1

import (
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"net/http"
)

type Car interface {
}

type CarHandler struct {
	service service.Car
}

func NewCarHandler(service service.Car) *CarHandler {
	h := &CarHandler{
		service: service,
	}
	return h
}

func (h *CarHandler) InitRoutes(r *myRouter.Router) {

	r.Post("/car/create", h.CreateCar)
	r.Put("/car/update", h.UpdateCar)
	r.Delete("/car/delete", h.DeleteCar)
	r.Get("/car", h.GetCar)
}
func (h *CarHandler) CreateCar(w http.ResponseWriter, r *http.Request) {

}
func (h *CarHandler) GetCar(w http.ResponseWriter, r *http.Request) {

}
func (h *CarHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {

}
func (h *CarHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {

}
