package controllers

import (
	"net/http"
)

var (
	HealthController healthControllerInterface = &healthController{}
)

type healthControllerInterface interface {
	Health(http.ResponseWriter, *http.Request)
}

type healthController struct{}

func (c *healthController) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}
