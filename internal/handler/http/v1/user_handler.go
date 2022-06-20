package v1

import (
	"errors"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/app_errors"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/domain/service"
	"github.com/sonyamoonglade/golang-rest-postgres/internal/handler/http/dto"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/util"
	"net/http"
	"strconv"
)

type User interface {
	InitRoutes(r *myRouter.Router)
}

type userHandler struct {
	service service.User
	logger  *myLogger.Logger
}

func NewUserHandler(service service.User, logger *myLogger.Logger) *userHandler {
	return &userHandler{service: service, logger: logger}
}

func (h *userHandler) InitRoutes(r *myRouter.Router) {
	r.Post("/user/register", h.register)
	r.Get("/user/", h.getById)
}
func (h *userHandler) register(w http.ResponseWriter, r *http.Request) {
	var input dto.CreateUserDto

	if err := util.ReadRequestBody(r.Body, &input); err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	userId, err := h.service.Create(input)
	if err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	util.JsonResponse(w, http.StatusCreated, map[string]interface{}{
		"id": userId,
	})
	return

}
func (h *userHandler) getById(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()

	if hasId := q.Has("id"); hasId == false {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "provide id to find a user",
		})
		return
	}

	userId, err := strconv.ParseInt(q.Get("id"), 10, 64)
	if err != nil {
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "incorrect id format",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}

	user, err := h.service.GetById(userId)
	if err != nil {
		var appErr *app_errors.ApiError
		if errors.As(err, &appErr) {
			util.JsonResponse(w, http.StatusNotFound, map[string]interface{}{
				"message": appErr.Message,
			})
			h.logger.PrintWithErr(appErr.Unwrap().Error())
			return
		}
		util.JsonResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		h.logger.PrintWithErr(err.Error())
		return
	}
	util.JsonResponse(w, http.StatusOK, map[string]interface{}{
		"user": user,
	})
	return

}
