package v1

import (
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/middleware"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/util"
	"net/http"
)

type Garage interface {
	InitRoutes(r *myRouter.Router)
}

type garageHandler struct {
	logger  *myLogger.Logger
	service service.Garage
}

func NewGarageHandler(service service.Garage, logger *myLogger.Logger) *garageHandler {
	return &garageHandler{
		logger:  logger,
		service: service,
	}
}

func (h *garageHandler) InitRoutes(r *myRouter.Router) {

	m := middleware.NewExtractMiddleware(h.logger, middleware.UserId)

	r.Post("/garage/create", m.Extract(h.create))
	r.Get("/garage/", h.get)
}

func (h *garageHandler) create(w http.ResponseWriter, r *http.Request) {

	var input dto.CreateGarageDto

	if err := util.ReadRequestBody(r.Body, &input); err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}
	fmt.Println(input)
	userId := r.Context().Value(middleware.UserId).(int64)

	input.UserID = userId
	fmt.Println(input)

	garageId, err := h.service.Create(input)
	if err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	util.JsonResponse(w, http.StatusCreated, map[string]interface{}{
		"garageId": garageId,
	})
	return

}

func (h *garageHandler) get(w http.ResponseWriter, r *http.Request) {

}
