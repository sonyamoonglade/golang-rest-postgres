package v1

import (
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/util"
	"net/http"
)

type Car interface {
	InitRoutes(r *myRouter.Router)
}

type carHandler struct {
	service service.Car
}

func NewCarHandler(service service.Car) *carHandler {
	h := &carHandler{
		service: service,
	}
	return h
}

func (h *carHandler) InitRoutes(r *myRouter.Router) {
	r.Post("/car/create", h.CreateCar)
	r.Put("/car/update", h.UpdateCar)
	r.Delete("/car/delete", h.DeleteCar)
	r.Get("/car/", h.GetCar)
}
func (h *carHandler) CreateCar(w http.ResponseWriter, r *http.Request) {

	var input dto.CreateCarDto

	if err := util.ReadRequestBody(r.Body, &input); err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		return
	}

	carId, err := h.service.CreateCar(input)
	if err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
	}

	util.JsonResponse(w, http.StatusOK, map[string]interface{}{
		"id": carId,
	})
	return

}
func (h *carHandler) GetCar(w http.ResponseWriter, r *http.Request) {

}
func (h *carHandler) UpdateCar(w http.ResponseWriter, r *http.Request) {

}
func (h *carHandler) DeleteCar(w http.ResponseWriter, r *http.Request) {

}
func (h *carHandler) BuyCar(w http.ResponseWriter, r *http.Request) {

}
