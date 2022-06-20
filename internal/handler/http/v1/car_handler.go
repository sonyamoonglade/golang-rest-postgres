package v1

import (
	"errors"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/app_errors"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/middleware"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/util"
	"net/http"
)

type Car interface {
	InitRoutes(r *myRouter.Router)
}

type carHandler struct {
	service service.Car
	logger  *myLogger.Logger
}

func NewCarHandler(service service.Car, logger *myLogger.Logger) *carHandler {
	h := &carHandler{
		logger:  logger,
		service: service,
	}
	return h
}

func (h *carHandler) InitRoutes(r *myRouter.Router) {

	u := middleware.NewExtractMiddleware(h.logger, middleware.UserId)

	r.Post("/car/create", h.createCar)
	r.Put("/car/update", h.updateCar)
	r.Delete("/car/delete", h.deleteCar)
	r.Get("/car", h.getCar)
	r.Put("/car/buy", u.Extract(h.buyCar))
}
func (h *carHandler) createCar(w http.ResponseWriter, r *http.Request) {

	var input dto.CreateCarDto

	if err := util.ReadRequestBody(r.Body, &input); err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	carId, err := h.service.CreateCar(input)
	if err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	util.JsonResponse(w, http.StatusCreated, map[string]interface{}{
		"carId": carId,
	})
	return
}

func (h *carHandler) getCar(w http.ResponseWriter, r *http.Request) {

}
func (h *carHandler) updateCar(w http.ResponseWriter, r *http.Request) {

}
func (h *carHandler) deleteCar(w http.ResponseWriter, r *http.Request) {

}
func (h *carHandler) buyCar(w http.ResponseWriter, r *http.Request) {

	var input dto.BuyCarDto

	if err := util.ReadRequestBody(r.Body, &input); err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	ok, err := h.service.BuyCar(input, r.Context())

	var appErr *app_errors.ApiError

	if err != nil {

		if errors.As(err, &appErr) {
			util.JsonResponse(w, http.StatusBadRequest, map[string]interface{}{
				"message": appErr.Message,
			})
			h.logger.PrintWithErr(appErr.Error())
			return
		}

		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	util.JsonResponse(w, http.StatusOK, map[string]interface{}{
		"ok": ok,
	})
	return

}
