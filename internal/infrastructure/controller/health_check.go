package controller

import (
	"net/http"

	"github.com/go-chi/render"
)

// HealthCheck is the system monitoring tool.
type HealthCheck struct{}

// NewHealthCheck returns a new HealthCheck implementation.
func NewHealthCheck() HealthCheck {
	return HealthCheck{}
}

// heartbeat is a handler function that checks the heartbeat of the API.
func (h HealthCheck) heartbeat(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, basicMessage{
		Message: http.StatusText(http.StatusOK),
	})
}
