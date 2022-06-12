package handler

import (
	"net/http"
)

func (c *Controller) getCar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get car"))
	return
}
func (c *Controller) createCar(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create car"))
	return
}
