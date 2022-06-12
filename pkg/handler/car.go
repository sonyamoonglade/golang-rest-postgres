package handler

import (
	"net/http"
)

func (c *CarController) getCar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get car"))
	return
}
func (c *CarController) createCar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create car"))
	return
}
