package handler

import (
	"encoding/json"
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/entities"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myRouter"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/service"
	"io"
	"log"
	"net/http"
)

type UserController struct {
	Service service.User
}

func NewUserController(service service.User) *UserController {
	return &UserController{Service: service}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var input entities.User

	body, err := io.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(500)
		log.Fatalf("Couldnt parse the body. %s", err.Error())
	}

	json.Unmarshal(body, &input)

	userId, err := c.Service.CreateUser(input)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("unexpected server error, sorry"))
		fmt.Println(err)
		return
	}

	response, _ := json.Marshal(userId)

	w.WriteHeader(201)
	w.Write(response)
	return
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {

}

func (c *UserController) InitRoutes(r *myRouter.Router) {

	r.POST("/register", c.CreateUser)
	r.GET("/getUser", c.GetUser)

}

